package sigma

import (
	"io/ioutil"

	"github.com/bradleyjkemp/sigma-go"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type PolicyCompiler[R any] struct {
	// Operations
	ops source.Operations[R]

	// Transformer
	transformer *source.Transformer

	// Compiled rule objects
	rules []policy.Rule[R]

	// Intermediate rule and rule config objects parsed by the Sigma parser
	sigmaRules  []sigma.Rule
	sigmaConfig []sigma.Config
}

// NewPolicyCompiler constructs a new compiler instance.
func NewPolicyCompiler[R any](ops source.Operations[R]) policy.PolicyCompiler[R] {
	pc := new(PolicyCompiler[R])
	pc.ops = ops
	pc.transformer = source.NewTransformer()
	pc.rules = make([]policy.Rule[R], 0)
	return pc
}

// Compile parses and interprets an input policy defined in path.
func (pc *PolicyCompiler[R]) compile(rulePaths []string, configPaths []string) error {
	// Read Sigma rules
	for _, path := range rulePaths {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		rule, err := sigma.ParseRule(contents)
		if err != nil {
			// ignore parsing errors
			continue
		}
		pc.sigmaRules = append(pc.sigmaRules, rule)
	}

	// Read Sigma configs
	for _, path := range configPaths {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		config, err := sigma.ParseConfig(contents)
		if err != nil {
			// ignore parsing errors
			continue
		}
		pc.sigmaConfig = append(pc.sigmaConfig, config)
	}

	// Translate the sigma rules into criterion objects
	for _, rule := range pc.sigmaRules {
		for _, conditions := range rule.Detection.Conditions {
			pc.visitSearchExpression(conditions.Search, rule.Detection.Searches)
		}
	}

	// errFound := false

	// if errFound {
	// 	return errors.New("errors found during compilation of policies. check logs for detail")
	// }

	return nil
}

// Compile parses a set of input policies defined in paths.
func (pc *PolicyCompiler[R]) Compile(paths ...string) ([]policy.Rule[R], []policy.Filter[R], error) {
	// Perhaps pass in the config object to expose separate paths for rules and config? recursive reading.
	pc.compile(nil, nil)
	return nil, nil, nil
}

func (pc *PolicyCompiler[R]) visitSearchExpression(condition sigma.SearchExpr, searches map[string]sigma.Search) policy.Criterion[R] {
	switch c := condition.(type) {
	case sigma.SearchIdentifier:
		// base case
		return policy.False[R]()

	case sigma.And:
		var andPreds []policy.Criterion[R]
		for _, expr := range c {
			andPreds = append(andPreds, pc.visitSearchExpression(expr, searches))
		}
		return policy.All(andPreds)

	case sigma.Or:
		var orPreds []policy.Criterion[R]
		for _, expr := range c {
			orPreds = append(orPreds, pc.visitSearchExpression(expr, searches))
		}
		return policy.Any(orPreds)

	case sigma.Not:
		return pc.visitSearchExpression(c, searches).Not()

	case sigma.OneOfThem:
		var orPreds []policy.Criterion[R]
		for _, search := range searches {
			orPreds = append(orPreds, pc.visitSearch(search))
		}
		return policy.Any(orPreds)
	case sigma.OneOfPattern:
	case sigma.AllOfThem:
	case sigma.AllOfPattern:
		break
	}
	return policy.True[R]()
}

func (pc *PolicyCompiler[R]) visitSearch(search sigma.Search) policy.Criterion[R] {
	var matcherPreds []policy.Criterion[R]
	for _, eventMatcher := range search.EventMatchers {
		for _, fieldMatcher := range eventMatcher {
			var fieldPreds policy.Criterion[R]
			allValuesMustMatch := false
			var transformers []source.TransformerFlags
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
						tPreds = append(tPreds, pc.visitTerm(comparators, fieldMatcher.Field, pc.transformer.TransformToString([]byte(value), t)))
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
	// TODO: check if Any is the appropriate predicate here
	return policy.Any(matcherPreds)
}

func (pc *PolicyCompiler[R]) visitTerm(ops []FieldModifier, attr string, value string) policy.Criterion[R] {
	var opPreds []policy.Criterion[R]
	if len(ops) == 0 {
		opPreds = append(opPreds, pc.ops.Eq(attr, value))
	} else {
		for _, op := range ops {
			switch op {
			case Contains:
				opPreds = append(opPreds, pc.ops.Contains(attr, value))
			case StartsWith:
				opPreds = append(opPreds, pc.ops.StartsWith(attr, value))
			case EndsWith:
				opPreds = append(opPreds, pc.ops.StartsWith(attr, value))
			case RegExp:
				opPreds = append(opPreds, pc.ops.StartsWith(attr, value))
			default:
				logger.Error.Printf("Unsupported operator %s", op)
			}
		}
	}
	return policy.All(opPreds)
}

// 	for name := range rule.Detection.Searches {
// 		// it's not possible for this call to error because the search expression parser won't allow this to contain invalid expressions
// 		matchesPattern, _ := path.Match(s.Pattern, name)
// 		if !matchesPattern {
// 			continue
// 		}
// 		if rule.visitSearchExpression(sigma.SearchIdentifier{Name: name}, searchResults) {
// 			return true
// 		}
// 	}
// 	return false

//
// 	for name := range rule.Detection.Searches {
// 		if !rule.visitSearchExpression(sigma.SearchIdentifier{Name: name}, searchResults) {
// 			return false
// 		}
// 	}
// 	return true

//
// 	for name := range rule.Detection.Searches {
// 		// it's not possible for this call to error because the search expression parser won't allow this to contain invalid expressions
// 		matchesPattern, _ := path.Match(s.Pattern, name)
// 		if !matchesPattern {
// 			continue
// 		}
// 		if !rule.visitSearchExpression(sigma.SearchIdentifier{Name: name}, searchResults) {
// 			return false
// 		}
// 	}
