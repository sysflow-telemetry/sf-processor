grammar Sfpl;

RULE: 'rule';
DROP: 'drop';
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
PREFILTER: 'prefilter';
ENABLED: 'enabled';
WARNEVTTYPE: 'warn_evttypes';
SKIPUNKNOWN: 'skip-if-unknown-filter';
FAPPEND: 'append';
REQ: 'required_engine_version';

policy
	: (prule | pfilter | pmacro | plist | preq)+ EOF
	;

defs
	: (srule | sfilter | pmacro | plist | preq)* EOF
	;

prule
	: DECL RULE DEF text DESC DEF text COND DEF expression ((ACTION|OUTPUT) DEF text | PRIORITY DEF severity | TAGS DEF tags | PREFILTER DEF prefilter | ENABLED DEF enabled | WARNEVTTYPE DEF warnevttype | SKIPUNKNOWN DEF skipunknown)*
	;

srule
	: DECL RULE DEF text DESC DEF text COND DEF expression ((ACTION|OUTPUT) DEF text | PRIORITY DEF severity | TAGS DEF tags | PREFILTER DEF prefilter | ENABLED DEF enabled | WARNEVTTYPE DEF warnevttype | SKIPUNKNOWN DEF skipunknown)*
	;

pfilter
	: DECL DROP DEF ID COND DEF expression (ENABLED DEF enabled)?
	;

sfilter
	: DECL DROP DEF ID COND DEF expression (ENABLED DEF enabled)?
	;

pmacro
	: DECL MACRO DEF ID COND DEF expression (FAPPEND DEF fappend)?
	;

plist
	: DECL LIST DEF ID ITEMS DEF items 
	;

preq
	: DECL REQ DEF atom
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
	: LBRACK (atom (LISTSEP atom)*)? (LISTSEP)? RBRACK
	;

tags
	: LBRACK (atom (LISTSEP atom)*)? (LISTSEP)? RBRACK
	;

prefilter
	: items
	;

severity
	: SEVERITY
	;

enabled
	: atom
	;

warnevttype
	: atom
	;

skipunknown
	: atom
	;

fappend
	: atom
	;

variable
	: ID
	;		

atom 
	: ID
	| PATH
	| NUMBER
	| TAG
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
	      p.GetCurrentToken().GetText() == "tags" ||
		  p.GetCurrentToken().GetText() == "prefilter" ||
		  p.GetCurrentToken().GetText() == "enabled" ||
		  p.GetCurrentToken().GetText() == "warn_evttypes" ||
		  p.GetCurrentToken().GetText() == "skip-if-unknown-filter" ||
		  p.GetCurrentToken().GetText() == "append")}? .)+
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
	| ENDSWITH
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

ENDSWITH
	: 'endswith'
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
	: SFSEVERITY
	| FSEVERITY	
	;

SFSEVERITY
	: H I G H
	| M E D I U M 
	| L O W
	;

FSEVERITY
	: E M E R G E N C Y 	
	| A L E R T	
	| C R I T I C A L	
	| E R R O R	
	| W A R N I N G	
	| N O T I C E	
	| I N F O
	| I N F O R M A T I O N A L	
	| D E B U G
	;

ID
	:  ('a'..'z' | 'A'..'Z' | '0'..'9' | '_') ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | ':'? '[' (NUMBER|PATH) (':' PATH)* ']' | '*' )*	
	;
	
NUMBER 
	: ('0'..'9')+ ('.' ('0'..'9')+)?
	;
	
PATH
	:  ('a'..'z' | 'A'..'Z' | '/' | '.') ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | '/' | '*' )*	
	;

STRING 
    : '"' (STRING|STRLIT) '"' 
    | '\'' (STRING|STRLIT) '\''
    | '\\"' (STRING|STRLIT) '\\"'
    | '\'\'' (STRING|STRLIT) '\'\''
    ;

TAG
	: ID ':' ID
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

fragment A : [aA]; // match either an 'a' or 'A'
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
