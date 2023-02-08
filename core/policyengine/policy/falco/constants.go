package falco

// Falco priority values.
const (
	FPriorityEmergency     = "emergency"
	FPriorityAlert         = "alert"
	FPriorityCritical      = "critical"
	FPriorityError         = "error"
	FPriorityWarning       = "warning"
	FPriorityNotice        = "notice"
	FPriorityInfo          = "info"
	FPriorityInformational = "informational"
	FPriorityDebug         = "debug"
)

// Exists creates a criterion for an existential predicate.

// 	utf16le: transforms value to UTF16-LE encoding, e.g. cmd > 63 00 6d 00 64 00 (only used in combination with base64 modifiers)

// utf16be: transforms value to UTF16-BE encoding, e.g. cmd > 00 63 00 6d 00 64 (only used in combination with base64 modifiers)

// wide: alias for utf16le modifier

// utf16: prepends a byte order mark and encodes UTF16, e.g. cmd > FF FE 63 00 6d 00 64 00 (only used in combination with base64 modifiers)

// windash: Add a new varia
// }
