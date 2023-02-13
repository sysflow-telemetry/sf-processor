package sigma

import (
	"os"
	"path"
	"strings"

	"github.com/bradleyjkemp/sigma-go"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type PolicyCompiler[R any] struct {
	// Operations
	ops source.Operations[R]

	// Transformer
	transformer *Transformer

	// Compiled rule objects
	rules []policy.Rule[R]

	// Intermediate rule and rule config objects parsed by the Sigma parser
	sigmaRules  []sigma.Rule
	sigmaConfig sigma.Config

	// Sigma config path
	configPath string
}

// NewPolicyCompiler constructs a new compiler instance.
func NewPolicyCompiler[R any](ops source.Operations[R], configPath string) policy.PolicyCompiler[R] {
	pc := new(PolicyCompiler[R])
	pc.ops = ops
	pc.transformer = NewTransformer()
	pc.rules = make([]policy.Rule[R], 0)
	pc.configPath = configPath
	return pc
}

// Compile parses and interprets an input policy defined in path.
func (pc *PolicyCompiler[R]) compile(rulePaths []string, configPath string) error {
	// Read Sigma rules
	for _, path := range rulePaths {
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rule, err := sigma.ParseRule(contents)
		if err != nil {
			logger.Error.Printf("Could not parse input rule ")
			continue
		}
		pc.sigmaRules = append(pc.sigmaRules, rule)
	}

	// Read Sigma config
	if p, err := os.Stat(configPath); err == nil && !p.IsDir() {
		contents, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}
		pc.sigmaConfig, err = sigma.ParseConfig(contents)
		if err != nil {
			return err
		}
	}

	// Translate the sigma rules into criterion objects
	for _, rule := range pc.sigmaRules {
		for _, conditions := range rule.Detection.Conditions {
			logger.Trace.Println("Parsing rule ", rule.ID, rule.Title)
			r := policy.Rule[R]{
				Name:      rule.ID,
				Desc:      rule.Description,
				Condition: pc.visitSearchExpression(conditions.Search, rule.Detection.Searches),
				Actions:   nil,
				Tags:      pc.getTags(rule),
				Priority:  pc.getPriority(rule),
				Prefilter: nil,
				Enabled:   true,
			}
			pc.rules = append(pc.rules, r)
		}
	}

	return nil
}

// Compile parses a set of input policies defined in paths.
func (pc *PolicyCompiler[R]) Compile(paths ...string) ([]policy.Rule[R], []policy.Filter[R], error) {
	if err := pc.compile(paths, pc.configPath); err != nil {
		return nil, nil, err
	}
	return pc.rules, nil, nil
}

func (pc *PolicyCompiler[R]) getTags(rule sigma.Rule) []policy.EnrichmentTag {
	tags := make([]policy.EnrichmentTag, len(rule.Tags))
	for i, v := range rule.Tags {
		tags[i] = v
	}
	return tags
}

func (pc *PolicyCompiler[R]) getPriority(rule sigma.Rule) policy.Priority {
	switch strings.ToLower(rule.Level) {
	case policy.Informational.String():
		return policy.Informational
	case policy.Low.String():
		return policy.Low
	case policy.Medium.String():
		return policy.Medium
	case policy.High.String():
		return policy.High
	case policy.Critical.String():
		return policy.Critical
	}
	return policy.Informational
}

func (pc *PolicyCompiler[R]) visitSearchExpression(condition sigma.SearchExpr, searches map[string]sigma.Search) policy.Criterion[R] {

	switch c := condition.(type) {

	case sigma.SearchIdentifier:
		if search, ok := searches[c.Name]; ok {
			return pc.visitSearch(search)
		}
		return policy.False[R]()

	case sigma.And:
		logger.Trace.Printf("%v", c)
		var preds []policy.Criterion[R]
		for _, expr := range c {
			preds = append(preds, pc.visitSearchExpression(expr, searches))
		}
		return policy.All(preds)

	case sigma.Or:
		var preds []policy.Criterion[R]
		for _, expr := range c {
			preds = append(preds, pc.visitSearchExpression(expr, searches))
		}
		return policy.Any(preds)

	case sigma.Not:
		return pc.visitSearchExpression(c.Expr, searches).Not()

	case sigma.OneOfThem:
		var preds []policy.Criterion[R]
		for _, search := range searches {
			preds = append(preds, pc.visitSearch(search))
		}
		return policy.Any(preds)

	case sigma.OneOfPattern:
		var preds []policy.Criterion[R]
		for name, search := range searches {
			matchesPattern, _ := path.Match(c.Pattern, name)
			if matchesPattern {
				preds = append(preds, pc.visitSearch(search))
			}
		}
		return policy.Any(preds)

	case sigma.AllOfThem:
		var preds []policy.Criterion[R]
		for _, search := range searches {
			preds = append(preds, pc.visitSearch(search))
		}
		return policy.All(preds)

	case sigma.AllOfPattern:
		var preds []policy.Criterion[R]
		for name, search := range searches {
			matchesPattern, _ := path.Match(c.Pattern, name)
			if matchesPattern {
				preds = append(preds, pc.visitSearch(search))
			}
		}
		return policy.All(preds)
	}
	return policy.False[R]()
}

func (pc *PolicyCompiler[R]) visitSearch(search sigma.Search) policy.Criterion[R] {

	if len(search.Keywords) > 0 {
		logger.Warn.Println("Keyword search is not supported. Use field selectors instead.")
		return policy.False[R]()
	}

	var matcherPreds []policy.Criterion[R]
	for _, eventMatcher := range search.EventMatchers {
		for _, fieldMatcher := range eventMatcher {
			var fieldPreds policy.Criterion[R]
			allValuesMustMatch := false
			var transformers []TransformerFlags
			var comparators []FieldModifier
			for _, modifier := range fieldMatcher.Modifiers {
				m := FieldModifier(modifier)
				if m == All {
					allValuesMustMatch = true
				}
				if m.IsTransformer() {
					transformers = append(transformers, TransformersMap[m]...)
				}
				if m.IsComparator() {
					comparators = append(comparators, m)
				}
			}
			var valuePreds []policy.Criterion[R]
			for _, value := range fieldMatcher.Values {
				if len(transformers) > 0 {
					var tPreds []policy.Criterion[R]
					for _, t := range transformers {
						values, _ := pc.transformer.Transform(value, t)
						for _, v := range values {
							tPreds = append(tPreds, pc.visitTerm(comparators, fieldMatcher.Field, v))
						}
					}
					valuePreds = append(valuePreds, policy.Any(tPreds))
				} else {
					valuePreds = append(valuePreds, pc.visitTerm(comparators, fieldMatcher.Field, value))
				}
			}
			if allValuesMustMatch {
				fieldPreds = policy.All(valuePreds)
			} else {
				fieldPreds = policy.Any(valuePreds)
			}
			matcherPreds = append(matcherPreds, fieldPreds)
		}
	}
	return policy.Any(matcherPreds)
}

func (pc *PolicyCompiler[R]) visitTerm(ops []FieldModifier, attr string, value string) policy.Criterion[R] {
	var opPreds []policy.Criterion[R]

	// apply field mappings
	if pc.sigmaConfig.FieldMappings != nil {
		if mattr, ok := pc.sigmaConfig.FieldMappings[attr]; ok {
			attr = mattr.TargetNames[0]
		}
	}

	// build predicate expression
	if len(ops) == 0 {
		opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.IEq)))
	} else {
		for _, op := range ops {
			switch op {
			case Contains:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.IContains)))
			case StartsWith:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.IStartswith)))
			case EndsWith:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.IEndswith)))
			case RegExp:
				opPreds = append(opPreds, policy.First(pc.ops.RegExp(attr, value)))
			case Lt:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.Lt)))
			case Lte:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.LEq)))
			case Gt:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.Gt)))
			case Gte:
				opPreds = append(opPreds, policy.First(pc.ops.Compare(attr, value, source.GEq)))
			default:
				logger.Error.Printf("Unsupported operator %s", op)
			}
		}
	}

	return policy.All(opPreds)
}
