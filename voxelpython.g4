grammar voxelpython;

program
    : statement* EOF
    ;

statement
    : varDecl
    | ifStatement
    | command
    | NEWLINE
    ;

varDecl
    : 'var' IDENTIFIER '=' expr NEWLINE
    ;

ifStatement
    : 'if' expr ':' block
      ('elif' expr ':' block)*
      ('else' ':' block)?
    ;

block
    : INDENT statement+ DEDENT
    ;

command
    : 'cmd' '.' IDENTIFIER (argument)* NEWLINE
    ;

argument
    : STRING
    | NUMBER
    | IDENTIFIER
    ;

expr
    : expr op=('*'|'/') expr      # MulDiv
    | expr op=('+'|'-') expr      # AddSub
    | '(' expr ')'                # Parens
    | NUMBER                     # Number
    | IDENTIFIER                 # Var
    | 'true'                     # TrueLiteral
    | 'false'                    # FalseLiteral
    ;

IDENTIFIER
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;

NUMBER
    : [0-9]+ ('.' [0-9]+)?
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'
    ;

NEWLINE
    : ('\r'? '\n')+
    ;

INDENT
    :   // handled by lexer or parser helper
    ;

DEDENT
    :   // handled by lexer or parser helper
    ;

WS
    : [ \t]+ -> skip
    ;

// Comments
COMMENT
    : '#' ~[\r\n]* -> skip
    ;
