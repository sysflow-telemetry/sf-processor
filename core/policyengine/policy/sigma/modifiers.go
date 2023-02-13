package sigma

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
	CIDR         FieldModifier = "cidr"
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

var TransformersMap = map[FieldModifier][]TransformerFlags{
	Base64:       {Base64Flag},
	Base64Offset: {Base64OffsetFlag},
	UTF16:        {NoFlags},
	UTF16LE:      {NoFlags},
	UTF16BE:      {NoFlags},
	Wide:         {NoFlags},
	WinDash:      {NoFlags, WinDashFlag},
	CIDR:         {CIDRFlag},
}

func (s FieldModifier) IsComparator() bool {
	_, ok := ComparatorsMap[s]
	return ok
}

func (s FieldModifier) IsTransformer() bool {
	_, ok := TransformersMap[s]
	return ok
}
