module dragon-compiler

go 1.17

require (
	expression_parser v0.0.0-00010101000000-000000000000
	lexer v0.0.0-00010101000000-000000000000
)

replace lexer => ./lexer

replace simple_parser => ./parser

replace expression_parser => ./expression_parser
