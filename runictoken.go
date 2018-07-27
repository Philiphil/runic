package main
/*RunicToken*/

	func makeRunicToken(s string)(rt RunicToken){
		rt.Content = s

		switch s{

			case rts_comma : rt.Symbol=rts_comma
			case rts_not_parsed : rt.Symbol=rts_not_parsed

			case rts_space: rt.Symbol=rts_space
			case rts_tabs: rt.Symbol=rts_tabs
			case rts_win_newline  : rt.Symbol=rts_win_newline
			case rts_unix_newline  : rt.Symbol=rts_unix_newline

			case rts_if: rt.Symbol=rts_if
			case rts_else: rt.Symbol=rts_else
			case rts_switch: rt.Symbol=rts_switch
			case rts_case: rt.Symbol=rts_case
			case rts_default: rt.Symbol=rts_default
			case rts_for: rt.Symbol=rts_for
			case rts_do: rt.Symbol=rts_do
			case rts_while: rt.Symbol=rts_while
			case rts_break: rt.Symbol=rts_break
			case rts_continue: rt.Symbol=rts_continue
			case rts_return: rt.Symbol=rts_return
			case rts_array: rt.Symbol=rts_array
			case rts_int: rt.Symbol=rts_int
			case rts_string: rt.Symbol=rts_string
			case rts_void: rt.Symbol=rts_void
			case rts_bool: rt.Symbol=rts_bool
			case rts_class: rt.Symbol=rts_class
			case rts_enum: rt.Symbol=rts_enum
			case rts_struct: rt.Symbol=rts_struct
			case rts_parent: rt.Symbol=rts_parent
			case rts_this: rt.Symbol=rts_this
			case rts_semicolon: rt.Symbol=rts_semicolon
			case rts_plus: rt.Symbol=rts_plus
			case rts_moins: rt.Symbol=rts_moins
			case rts_multiplier: rt.Symbol=rts_multiplier
			case rts_diviser: rt.Symbol=rts_diviser
			case rts_modulo: rt.Symbol=rts_modulo
			case rts_equal: rt.Symbol=rts_equal
			case rts_increment: rt.Symbol=rts_increment
			case rts_decrement: rt.Symbol=rts_decrement
			case rts_colon: rt.Symbol=rts_colon
			case rts_question: rt.Symbol=rts_question
			case rts_boolean_inversion: rt.Symbol=rts_boolean_inversion
			case rts_equality: rt.Symbol=rts_equality
			case rts_not_equal: rt.Symbol=rts_not_equal
			case rts_less_than: rt.Symbol=rts_less_than
			case rts_more_than: rt.Symbol=rts_more_than
			case rts_less_equal_than: rt.Symbol=rts_less_equal_than
			case rts_more_equal_than: rt.Symbol=rts_more_equal_than
			case rts_type_equality: rt.Symbol=rts_type_equality
			case rts_and: rt.Symbol=rts_and
			case rts_or: rt.Symbol=rts_or
			case rts_bit_left: rt.Symbol=rts_bit_left
			case rts_bit_right: rt.Symbol=rts_bit_right
			case rts_bit_and: rt.Symbol=rts_bit_and
			case rts_bit_or: rt.Symbol=rts_bit_or
			case rts_bit_xor: rt.Symbol=rts_bit_xor
			case rts_bit_nor: rt.Symbol=rts_bit_nor
			case rts_equal_plus: rt.Symbol=rts_equal_plus
			case rts_equal_moins: rt.Symbol=rts_equal_moins
			case rts_equal_multiplier: rt.Symbol=rts_equal_multiplier
			case rts_equal_diviser: rt.Symbol=rts_equal_diviser
			case rts_equal_modulo: rt.Symbol=rts_equal_modulo
			case rts_equal_bit_left: rt.Symbol=rts_equal_bit_left
			case rts_equal_bit_right: rt.Symbol=rts_equal_bit_right
			case rts_equal_bit_and: rt.Symbol=rts_equal_bit_and
			case rts_equal_bit_or: rt.Symbol=rts_equal_bit_or
			case rts_equal_bit_xor: rt.Symbol=rts_equal_bit_xor
			case rts_double_colon: rt.Symbol=rts_double_colon
			case rts_dot: rt.Symbol=rts_dot
			case rts_open_parenthese: rt.Symbol=rts_open_parenthese
			case rts_close_parenthese: rt.Symbol=rts_close_parenthese
			case rts_open_bracket: rt.Symbol=rts_open_bracket
			case rts_close_bracket: rt.Symbol=rts_close_bracket
			case rts_open_brace: rt.Symbol=rts_open_brace
			case rts_close_brace: rt.Symbol=rts_close_brace

			case rts_tick_1: rt.Symbol=rts_tick_1
			case rts_tick_2: rt.Symbol=rts_tick_2
			case rts_tick_3: rt.Symbol=rts_tick_3

			case rts_private: rt.Symbol=rts_private
			case rts_public: rt.Symbol=rts_public
			case rts_protected: rt.Symbol=rts_protected
			case rts_static: rt.Symbol=rts_static
			case rts_const: rt.Symbol=rts_const
			case rts_async: rt.Symbol=rts_async
			case rts_override: rt.Symbol=rts_override
			case rts_virtual: rt.Symbol=rts_virtual
			case rts_overload: rt.Symbol=rts_overload
			case rts_var: rt.Symbol=rts_var
			case rts_import: rt.Symbol=rts_import
			case rts_package: rt.Symbol=rts_package
			case rts_main: rt.Symbol=rts_main 

			case rts_tick_1c : rt.Symbol=rts_tick_1c 
			case rts_tick_2c : rt.Symbol=rts_tick_2c 
			case rts_tick_3c : rt.Symbol=rts_tick_3c 

			case rts_monocom : rt.Symbol=rts_monocom 
			case rts_open_multicom : rt.Symbol=rts_open_multicom 
			case rts_close_multicom : rt.Symbol=rts_close_multicom 


			default : rt.Symbol=rts_not_parsed
		}

		switch rt.Symbol{
			case rts_comma : rt.Type=rts_comma
			case rts_not_parsed : rt.Type=rts_not_parsed

			case rts_space: rt.Type=rtt_space
			case rts_tabs: rt.Type=rtt_space
			case rts_win_newline  : rt.Type=rtt_newline
			case rts_unix_newline  : rt.Type=rtt_newline


			case rts_if: rt.Type=rts_if
			case rts_else: rt.Type=rts_else
			case rts_switch: rt.Type=rts_switch
			case rts_case: rt.Type=rts_case
			case rts_default: rt.Type=rts_default
			case rts_for: rt.Type=rts_for
			case rts_do: rt.Type=rts_do
			case rts_while: rt.Type=rts_while
			case rts_break: rt.Type=rts_break
			case rts_continue: rt.Type=rts_continue
			case rts_return: rt.Type=rts_return

			case rts_array: rt.Type=rtt_var_type
			case rts_int: rt.Type=rtt_var_type
			case rts_string: rt.Type=rtt_var_type
			case rts_void: rt.Type=rtt_var_type
			case rts_bool: rt.Type=rtt_var_type
			case rts_var: rt.Type=rtt_var_type

			case rts_class: rt.Type=rts_class
			case rts_enum: rt.Type=rts_enum
			case rts_struct: rt.Type=rts_struct

			case rts_parent: rt.Type=rts_parent
			case rts_this: rt.Type=rts_this

			case rts_semicolon: rt.Type=rts_semicolon
			case rts_plus: rt.Type=rts_plus
			case rts_moins: rt.Type=rts_moins
			case rts_multiplier: rt.Type=rts_multiplier
			case rts_diviser: rt.Type=rts_diviser
			case rts_modulo: rt.Type=rts_modulo
			case rts_equal: rt.Type=rts_equal
			case rts_increment: rt.Type=rts_increment
			case rts_decrement: rt.Type=rts_decrement
			case rts_colon: rt.Type=rts_colon
			case rts_question: rt.Type=rts_question
			case rts_boolean_inversion: rt.Type=rts_boolean_inversion
			case rts_equality: rt.Type=rts_equality
			case rts_not_equal: rt.Type=rts_not_equal
			case rts_less_than: rt.Type=rts_less_than
			case rts_more_than: rt.Type=rts_more_than
			case rts_less_equal_than: rt.Type=rts_less_equal_than
			case rts_more_equal_than: rt.Type=rts_more_equal_than
			case rts_type_equality: rt.Type=rts_type_equality
			case rts_and: rt.Type=rts_and
			case rts_or: rt.Type=rts_or
			case rts_bit_left: rt.Type=rts_bit_left
			case rts_bit_right: rt.Type=rts_bit_right
			case rts_bit_and: rt.Type=rts_bit_and
			case rts_bit_or: rt.Type=rts_bit_or
			case rts_bit_xor: rt.Type=rts_bit_xor
			case rts_bit_nor: rt.Type=rts_bit_nor
			case rts_equal_plus: rt.Type=rts_equal_plus
			case rts_equal_moins: rt.Type=rts_equal_moins
			case rts_equal_multiplier: rt.Type=rts_equal_multiplier
			case rts_equal_diviser: rt.Type=rts_equal_diviser
			case rts_equal_modulo: rt.Type=rts_equal_modulo
			case rts_equal_bit_left: rt.Type=rts_equal_bit_left
			case rts_equal_bit_right: rt.Type=rts_equal_bit_right
			case rts_equal_bit_and: rt.Type=rts_equal_bit_and
			case rts_equal_bit_or: rt.Type=rts_equal_bit_or
			case rts_equal_bit_xor: rt.Type=rts_equal_bit_xor
			case rts_double_colon: rt.Type=rts_double_colon
			case rts_dot: rt.Type=rts_dot
			case rts_open_parenthese: rt.Type=rts_open_parenthese
			case rts_close_parenthese: rt.Type=rts_close_parenthese
			case rts_open_bracket: rt.Type=rts_open_bracket
			case rts_close_bracket: rt.Type=rts_close_bracket

			case rts_open_brace: rt.Type=rts_open_brace
			case rts_close_brace: rt.Type=rts_close_brace
			
			case rts_tick_1: rt.Type=rtt_tick
			case rts_tick_2: rt.Type=rtt_tick
			case rts_tick_3: rt.Type=rtt_tick
			case rts_tick_1c: rt.Type=rtt_tick
			case rts_tick_2c: rt.Type=rtt_tick
			case rts_tick_3c: rt.Type=rtt_tick

			case rts_private: rt.Type=rts_private
			case rts_public: rt.Type=rts_public
			case rts_protected: rt.Type=rts_protected
			case rts_static: rt.Type=rts_static
			case rts_const: rt.Type=rts_const
			case rts_async: rt.Type=rts_async
			case rts_override: rt.Type=rts_override
			case rts_virtual: rt.Type=rts_virtual


			case rts_overload: rt.Type=rts_overload

			
			case rts_monocom : rt.Type=rtt_commentary 
			case rts_open_multicom : rt.Type=rtt_commentary 
			case rts_close_multicom : rt.Type=rtt_commentary 

			

			default : rt.Type = rtt_not_parsed
		}


		if rt.Symbol == ""{
			rt.Symbol= rts_not_parsed
		}
		if rt.Type == ""{
			rt.Type= rtt_not_parsed
		}

		return
	}