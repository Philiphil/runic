package main
import "fmt"

	type RunicCompiler struct{
		Files []RunicFile
		
		Compiler string
		Action string

		Naming string
		Name string

		Path string
		RunicHeaderPath string

		DeleteTMP bool

		Mask string
		Parser RunicParser
		Lexer RunicLexer
	}

	type RunicFile struct{
		OldPath string
		NewPath string
		Content []RunicInstruction
	} 

	type RunicToken struct{
		Symbol RunicTokenSymbol
		Type RunicTokenType
		Content string
		CompiledContent string
	}

	type RunicParser struct{
		InTick bool//
		TickType RunicTokenSymbol

		InFunction bool//
		InEnum bool//
		InClass bool//
		InStruct bool//

		InMain bool
		PackageName string

		InParenteses int
		InBrace int
		InBraket int

		Variables []ParsedVariable //variable in file
		Functions []ParsedFunction//
		TypeDef []string//
	}

	type RunicLexer struct{
		InTick bool//
		TickType RunicTokenSymbol

		InFunction bool//
		InEnum bool//
		InClass bool//
		InStruct bool//

		InMain bool
		PackageName string

		Cache string
		CachingToken bool

		InParenteses int
		InBrace int
		InBraket int

		Variables []ParsedVariable //variable in file
		Functions []ParsedFunction//
		TypeDef []string//
		NextIsTypeDef bool
	}

	type ParsedVariable struct{
		Name string
		Type string
	}

	type ParsedFunction struct{
		Name string
		Vars []ParsedVariable
	}

	type RunicInstruction struct{
		Content []RunicToken
		CompiledContent string
	}
	type RunicTokenSymbol string
	type RunicTokenType string

	const(
		rtt_not_parsed="rtt_not_parsed"
		rtt_var_type ="rtt_var_type"

		rtt_string="rtt_string"

		rtt_tick ="rtt_tick"

		rtt_sugar="rtt_sugar"

		rtt_space ="rtt_space"
		rtt_newline="rtt_newline"

		rtt_operator="rtt_operator"
		rtt_logical_operator="rtt_logical_operator"

		rtt_method_prefix="rtt_method_prefix"
		rtt_property_prefix="rtt_property_prefix"

		rtt_commentary="rtt_commentary"

		)

	const(
		rts_not_parsed = ""
		rts_space = " "
		rts_tabs = "\t"
		rts_win_newline="\r\n"
		rts_unix_newline="\n"


		//controle
			//fallthough, try catch ?
			rts_if = "if"
			rts_else = "else"
			rts_switch = "switch"
			rts_case = "case"
			rts_default = "default"
			rts_for = "for"
			rts_do = "do"
			rts_while = "while"
			rts_break = "break"
			rts_continue = "continue"
			rts_return = "return"
			rts_async = "async"
			rts_import = "import"
			rts_package = "package"
			rts_main = "main"

		//type
			rts_var = "var"
			rts_array = "array"
			rts_int = "int"
			rts_string = "string"
			rts_void = "void"
			rts_bool = "bool"

			rts_class = "class"
			rts_enum = "enum"
			rts_struct = "struct"
			rts_const = "const" 
				//objs
				rts_parent = "parent"
				rts_this = "this"
				rts_private = "private"
				rts_public = "public"
				rts_protected = "protected"
				rts_static = "static"
				rts_override ="override"
				rts_virtual = "virtual"
				rts_overload = "overload"
			

		//operators
			//math
			rts_plus = "+"
			rts_moins = "-"
			rts_multiplier = "*"
			rts_diviser = "/"
			rts_modulo = "%"
			rts_equal = "="
			rts_increment = "++"
			rts_decrement = "--"
			//logical
			rts_colon = ":"
			rts_question = "?"
			rts_boolean_inversion = "!"
			rts_equality = "=="
			rts_not_equal = "!="
			rts_less_than = "<"
			rts_more_than = ">"
			rts_less_equal_than = "<="
			rts_more_equal_than = ">="
			rts_type_equality = "is"
			rts_and = "&&"
			rts_or = "||"
			//bitwise
			rts_bit_left = "<<"
			rts_bit_right = ">>"
			rts_bit_and = "&"
			rts_bit_or = "|"
			rts_bit_xor = "^"
			rts_bit_nor = "~"
			//equal and 
			rts_equal_plus = "+="
			rts_equal_moins = "-="
			rts_equal_multiplier = "*="
			rts_equal_diviser = "/="
			rts_equal_modulo = "%="
			rts_equal_bit_left = "<<="
			rts_equal_bit_right = ">>="
			rts_equal_bit_and = "&="
			rts_equal_bit_or = "|="
			rts_equal_bit_xor = "^="

			rts_monocom = "//"
			rts_open_multicom = "/*"
			rts_close_multicom = "*/"

		//sugar
			rts_double_colon = "::"
			rts_dot = "."
			rts_semicolon = ";"
			rts_comma = ","

			rts_open_parenthese = "("
			rts_close_parenthese = ")"

			rts_open_bracket = "["
			rts_close_bracket = "]"

			rts_open_brace= "{"
			rts_close_brace = "}"

			rts_tick_1 = "\""
			rts_tick_2 = "'"
			rts_tick_3 = "`"

			rts_tick_1c = "\\\""
			rts_tick_2c = "\\'"
			rts_tick_3c = "\\`"
		)


	type RunicError struct {
		Line int
		Char int
		File string
		What string
	}

	func (e RunicError) print() {
		fmt.Println("FILE",e.File,"LINE",e.Line,":",e.What)
	}
