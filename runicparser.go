package main
import (
//"fmt"
//"io/ioutil"
)

/*RunicTokenParser*/
func (this *RunicParser) initParser(rl RunicLexer){
	this.setDefaultState()
	this.Functions = rl.Functions
	this.TypeDef = rl.TypeDef
}

func (this *RunicParser) setDefaultState(){
	this.InTick = false
	this.InFunction = false
	this.InEnum = false
	this.InStruct = false
	this.InClass = false

	this.TickType = rtt_not_parsed

	this.InMain = false
	this.InParenteses = 0
	this.InBrace = 0
	this.InBraket = 0
}


/*STEP 2*/
func (this *RunicParser) parseInstruction(parsable_instruction []RunicInstruction) (parsed_instruction []RunicInstruction, err []RunicError){
		this.setDefaultState()	

		for i,instruction := range parsable_instruction{

			s := ""
			var r RunicError
			var rs []RunicError
			instruction_done := true

			if len(instruction.Content) == 0 {
			}else if !this.InTick && !this.InClass &&!this.InFunction && !this.InEnum && !this.InStruct{
				//if not In..., newline must be declaration

				//GLOBAL
				if instruction.Content[0].Type == rtt_var_type && instruction.Content[len(instruction.Content)-1].Symbol == rts_semicolon{

				//function
				}else if (instruction.Content[0].Type == rtt_var_type || this.isTypeDef(instruction.Content[0].Content) || instruction.Content[0].Symbol == rts_open_parenthese ) && instruction.Content[len(instruction.Content)-1].Symbol == rts_open_brace{
					
					s,rs = this.parseFunctionSignature(instruction)

				}else{
					instruction_done =false
				}
			//newline must be operation
			}else if this.InFunction{
				s,r = this.parseFunctionBody(instruction)
			}else{
				instruction_done =false
			}

			if !instruction_done{
				for _,token := range instruction.Content{
					s+=" "+ token.Content
				}
			}


			instruction.CompiledContent = s
			parsed_instruction= append(parsed_instruction,instruction)
			rs = append(rs,r)
			for _,er := range rs{
				if (er.What != ""){
					er.Line = i+1;
					err = append(err,er)
				}
			}
			
		}//endfile
		return
}

func (this *RunicParser) parseFunctionSignature(instruction RunicInstruction) (s string, err []RunicError){
		this.InFunction = true

		return_type := ""
		if instruction.Content[0].Symbol == rts_open_parenthese{ //(int) fn(){
			for i, token := range instruction.Content{
				return_type += token.Content
				if token.Symbol == rts_close_parenthese {
					instruction.Content = instruction.Content[i+1:]
					break;
				}
			}		
		}else{//int fn(){
			if instruction.Content[0].Content != "void"{
				return_type = instruction.Content[0].Content
			}
			instruction.Content = instruction.Content[1:]
		}

		if instruction.Content[0].Content == "main"{
			this.InMain = true
			if return_type != "int" && return_type != "void"{
				r := RunicError{}
				r.What = "main function must return int or void not "+return_type
				err = append(err,r)
			}
			return_type = ""
		}

		s = "func "+ instruction.Content[0].Content
		instruction.Content = instruction.Content[:1]


		args := ""
		parameters := RunicInstruction{}
		for _, token := range instruction.Content{
			if token.Symbol == rts_open_parenthese{
				continue
			}else if token.Symbol == rts_close_parenthese{
				r := RunicError{}
				args,r = this.parseParameters(parameters)
				err = append(err,r)
				break;
			}
			parameters.Content = append(parameters.Content, token)
		}//endfor

		s += "("+args+")"+return_type+"{"
		return
}


func (this *RunicParser) parseParameters(instruction RunicInstruction) (str string, r RunicError){
	var params []string
	t := false

	for _, token := range instruction.Content{
		if (!t && token.Type != rtt_var_type) || (token.Type == rtt_space || token.Symbol == rts_comma){
				continue// espaces, virgules ...
			}else if !t{
				str = token.Content
			}else{
				params = append(params, token.Content +" "+str)
			}
			t = !t
	}
	str = ""
	for _, param := range params{
		str += param+","
	}

	if len(str) > 0{
		str=str[0:len(str)-1 ]
		 return
	}
	return 
}

func (this *RunicParser) parseFunctionBody(instruction RunicInstruction) (s string, r RunicError){
	instruction_done := false

	// return (in main)
	if instruction.Content[0].Symbol == rts_return && this.InMain{
		instruction_done=true
		bfr := instruction.Content[1:len(instruction.Content)-1]
		s_bfr := ""
		for _,b := range bfr{
			s_bfr += b.Content+" "
		}
		s = "program_return("+s_bfr+")"
	}

	//variable declaration
	if (instruction.Content[0].Type == rtt_var_type ||  this.isTypeDef(instruction.Content[0].Content) ){
		//type t = ...;
		if  instruction.Content[2].Symbol == rts_equal{
			instruction_done=true
			s = instruction.Content[1].Content+" "+":="
			for i,e := range instruction.Content{
				if i <= 2 {
					continue
				}
				s += e.Content
			}
		//type t;
		}else if instruction.Content[2].Symbol == rts_semicolon{
			instruction_done=true
			s ="var "+instruction.Content[1].Content+" "+instruction.Content[0].Content
		}
	}

	//detect function end
	if instruction.Content[0].Symbol == rts_close_brace && this.InBrace == 0{
		this.InEnum = false
		this.InStruct = false
		if this.InClass &&  !this.InFunction{
			this.InClass = false
		}
		this.InFunction = false
		this.InMain = false
	}

	//compute not parsed stuff as go
	if (!instruction_done){
		for _,token := range instruction.Content{
			 s+=" "+token.Content
		}
	}


	return
}


func (this *RunicParser) isTypeDef (s string)(bool){
	for _,td := range this.TypeDef{
		if td == s{
			return true
		}
	}
	return false
}