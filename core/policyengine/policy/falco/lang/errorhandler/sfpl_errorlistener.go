package errorhandler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// SfplSyntaxError stores syntax error information during
// policy parsing
type SfplSyntaxError struct {
	line, column int
	msg          string
}

// Error returns a formatted string representing the syntax error
func (s *SfplSyntaxError) Error() string {
	return fmt.Sprintf("line: %d  column: %d %s", s.line, s.column, s.msg)
}

// SfplErrorListener monitors errors during the policy parsing process
// and stores them in an error list
type SfplErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

// SyntaxError is called by the antlr lexer and parser when it encounters and error
func (l *SfplErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, &SfplSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
