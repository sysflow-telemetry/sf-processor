package sigma

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

// FieldModifier type.
type FieldModifier string

// Sigma field modifiers.
const (
	// conjunctive modifier
	All FieldModifier = "all"

	// comparators
	Contains   FieldModifier = "contains"
	EndsWith   FieldModifier = "endswith"
	StartsWith FieldModifier = "startswith"
	Lt         FieldModifier = "lt"
	Lte        FieldModifier = "lte"
	Gt         FieldModifier = "gt"
	Gte        FieldModifier = "gte"

	// transformers
	Base64       FieldModifier = "base64"
	Base64Offset FieldModifier = "base64Offset"
	UTF16        FieldModifier = "utf16"
	UTF16LE      FieldModifier = "utf16le"
	UTF16BE      FieldModifier = "utf16be"
	Wide         FieldModifier = "wide"
	WinDash      FieldModifier = "windash"
	RegExp       FieldModifier = "re"
	Cidr         FieldModifier = "cidr"
)

var exists = struct{}{}

var ComparatorsMap = map[FieldModifier]struct{}{
	Contains:   exists,
	EndsWith:   exists,
	StartsWith: exists,
	Lt:         exists,
	Lte:        exists,
	Gt:         exists,
	Gte:        exists,
}

var TransformersMap = map[FieldModifier][]source.TransformerFlags{
	Base64:       []source.TransformerFlags{source.Base64Flag},
	Base64Offset: []source.TransformerFlags{source.Base64Flag, source.Base64Offset1Flag, source.Base64Offset2Flag},
	UTF16:        []source.TransformerFlags{source.UTF16BEFlag.Set(source.UTF16BOMFlag)},
	UTF16LE:      []source.TransformerFlags{source.UTF16LEFlag},
	UTF16BE:      []source.TransformerFlags{source.UTF16BEFlag},
	Wide:         []source.TransformerFlags{source.UTF16LEFlag},
	WinDash:      []source.TransformerFlags{source.NoFlags, source.WinDashFlag},
	Cidr:         []source.TransformerFlags{source.CIDRFlag},
}

func (s FieldModifier) IsComparator() bool {
	_, ok := ComparatorsMap[s]
	return ok
}

func (s FieldModifier) IsTransformer() bool {
	_, ok := TransformersMap[s]
	return ok
}
