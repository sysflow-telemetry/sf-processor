grammar Sfpl;

RULE: 'rule';
FILTER: 'filter';
MACRO: 'macro';
LIST: 'list';
NAME: 'name';
ITEMS: 'items';
COND: 'condition';
DESC: 'desc' ;
ACTION: 'action';
OUTPUT: 'output';
PRIORITY: 'priority';
TAGS: 'tags';

policy
	: (prule | pfilter | pmacro | plist)+ EOF
	;

prule
	: DECL RULE DEF text DESC DEF text COND DEF expression (ACTION|OUTPUT) DEF text PRIORITY DEF SEVERITY TAGS DEF items 
	;

pfilter
	: DECL FILTER DEF ID COND DEF expression
	;

pmacro
	: DECL MACRO DEF ID COND DEF expression
	;

plist
	: DECL LIST DEF ID ITEMS DEF items 
	;
	
expression 
	: or_expression 
	;

or_expression 
	: and_expression (OR and_expression)*
	;

and_expression 
	: term (AND term)*
	;

term 
	: variable
	| NOT term
	| atom unary_operator 
	| atom binary_operator atom 
	| atom (IN|PMATCH) LPAREN (atom|items) (LISTSEP (atom|items))* RPAREN 
	| LPAREN expression RPAREN
	;

items 
	: LBRACK (atom (LISTSEP atom)*)? RBRACK
	;
	
variable
	: ID
	;		

atom 
	: ID
	| PATH
	| NUMBER
	| STRING
	| '<' /* event direction */
	| '>' /* event direction */
	;

text
	: ({!(p.GetCurrentToken().GetText() == "desc" ||
	      p.GetCurrentToken().GetText() == "condition" ||
	      p.GetCurrentToken().GetText() == "action" ||
	      p.GetCurrentToken().GetText() == "output" ||
	      p.GetCurrentToken().GetText() == "priority" ||
	      p.GetCurrentToken().GetText() == "tags")}? .)+
	;
	
binary_operator 
	: LT 
	| LE 
	| GT 
	| GE 
	| EQ 
	| NEQ 
	| CONTAINS 
	| ICONTAINS
	| STARTSWITH
	;

unary_operator 
	: EXISTS
	;

AND 
	: 'and'
	;

OR 
	: 'or'
	;

NOT 
	: 'not'
	;

LT 
	: '<'
	;

LE 
	: '<='
	;

GT 
	: '>'
	;

GE 
	: '>='
	;

EQ 
	: '='
	;

NEQ 
	: '!='
	;

IN 
	: 'in'
	;

CONTAINS 
	: 'contains'
	;

ICONTAINS 
	: 'icontains'
	;
	
STARTSWITH 
	: 'startswith'
	;
	
PMATCH
	: 'pmatch'
	;

EXISTS 
	: 'exists'
	;

LBRACK 
	: '['
	;

RBRACK 
	: ']'
	;

LPAREN 
	: '('
	;

RPAREN 
	: ')'
	;

LISTSEP 
	: ','
	;

DECL 
	: '-'
	;
	
DEF
	: ':' ((' ')* '>')? 
	;

SEVERITY
	: 'DEBUG'
	| 'INFO'
	| 'NOTICE'
	| 'WARNING'
	| 'ERROR'
	;

ID
	:  ('a'..'z' | 'A'..'Z' | '0'..'9' | '_') ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | ':'? '[' (NUMBER|PATH) (':' PATH)* ']' | '*' )*	
	;
	
NUMBER 
	: ('0'..'9')+ ('.' ('0'..'9')+)?
	;
	
PATH
	:  ('a'..'z' | 'A'..'Z' | '/' ) ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | '/' | '*' )*	
	;

STRING 
    : '"' (STRING|STRLIT) '"' 
    | '\'' (STRING|STRLIT) '\''
    | '\\"' (STRING|STRLIT) '\\"'
    | '\'\'' (STRING|STRLIT) '\'\''
    ;

fragment STRLIT 
    //: .*? 
    : ~[\r\n]*?
	;
	
fragment ESC : '\\"' | '\'\'' ;
		
WS
	: [ \t\r\n\u000C]+ -> channel(HIDDEN)
	;
	
NL
	: '\r'? '\n' -> channel(HIDDEN)
	;
	
COMMENT 
	: '#' ~[\r\n]* -> channel(HIDDEN)
	;
	
ANY : . ;
