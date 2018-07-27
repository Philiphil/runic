package main
import (
"io/ioutil"
//"fmt"
)


func (this *RunicLexer) initLexer(){
	this.setDefaultState()
}

func (this *RunicLexer) setDefaultState(){
	this.InTick = false
	this.InFunction = false
	this.InEnum = false
	this.InStruct = false
	this.InClass = false

	this.TickType = rtt_not_parsed
	this.CachingToken = false
	this.Cache=""

	this.InMain = false
	this.NextIsTypeDef = false

	this.InParenteses = 0
	this.InBrace = 0
	this.InBraket = 0

	this.Variables = []ParsedVariable{}
}

func (this *RunicLexer) parse(File RunicFile) (parsable_instruction []RunicInstruction){
	this.setDefaultState()
	content, _ := ioutil.ReadFile(File.OldPath)
	return this.parseLine(string(content))
}

func (this *RunicLexer) parseLine(line string) (instructions []RunicInstruction){
	var tokens []RunicToken
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {

		token,b := this.parseToken(runes,&i)
		if !b{
			continue
		}

		this.extractTypeDef(token)
		tokens=this.parseString(&token, tokens )

		if(token.Type != rtt_space && token.Type != rtt_tick && token.Type != rtt_newline){
		//	fmt.Println(token)
			tokens= append(tokens,token)
		}


		tokens,instructions = this.parseMultiline(token,tokens,instructions)
		tokens,instructions = this.parseMonoline(tokens,instructions)
	}

	if len(tokens) > 0{
		instructions = append(instructions, makeRunicInstruction(tokens))
	}

  	return
} 


func (this *RunicLexer) parseToken(line []rune, index *int) (RunicToken,bool){
		  	token_end := false
		  	line_end := false
		  	var token RunicToken

		  	if isAlphaNum(byte(line[*index])) {
		  		this.CachingToken=true
		  		this.Cache += string(line[*index])
		  	}else{
		  		token_end = true
		  	}
		  	if *index == len(line)-1{
		  		line_end = true
		  	}

		  	if line_end || token_end{	
		  		j := *index+1

		  		if !token_end && line_end{
		  			j--
		  		}

		  		token = makeRunicToken(string(line[*index:j]))
		  		if this.CachingToken{ 
		  			token = makeRunicToken(this.Cache)
		  			this.Cache = ""
		  			this.CachingToken=false
		  			*index--
		  		}else if j+1 < len(line) {
		  			rt_bb  := makeRunicToken(string(line[*index:j+1]))
		  			if rt_bb.Type != rtt_not_parsed{
		  				token = rt_bb
		  				*index++
		  			}
		  		}

		  	}
		  	return token, token_end
}

func (this *RunicLexer) extractTypeDef(token RunicToken) {
	if token.Symbol == rts_class || token.Symbol == rts_struct {
		this.NextIsTypeDef=true
	}else if this.NextIsTypeDef{
		this.TypeDef = append(this.TypeDef,token.Content)
		this.NextIsTypeDef=false
	}
}

func (this *RunicLexer) parseString(token *RunicToken, tokens []RunicToken)([]RunicToken){
	if !this.InTick && token.Type != rtt_tick{//no string to parse
		return tokens
	}
	if(!this.InTick && token.Type == rtt_tick){//tick = new delimiter
		this.TickType = token.Symbol
	}
	if token.Type == rtt_tick && this.TickType == token.Symbol{//if tiking
		this.InTick =!this.InTick

	  	if !this.InTick{//tick over
	  		//verif last token string
	  		tokens[len(tokens)-1].Content = "`"+ tokens[len(tokens)-1].Content +"`"
	  	}else{//tick just start
	  	}
	  }else if this.InTick{
	  	token.Type = rtt_string
	  	switch(token.Symbol){
	  		case rts_tick_1c: token.Content = rts_tick_1
	  		case rts_tick_2c: token.Content = rts_tick_2
	  		case rts_tick_3c: token.Content = rts_tick_3c
	  	}
		if tokens[len(tokens)-1].Type == rtt_string{//already string-ing
			token.Content = tokens[len(tokens)-1].Content + token.Content 
			tokens = tokens[0:len(tokens)-1]
		}
	} 
	return tokens
}

func (this *RunicLexer) parseMultiline(token RunicToken, tokens []RunicToken, instructions []RunicInstruction)([]RunicToken,[]RunicInstruction){
	if token.Type == rtt_newline{
		if(len(tokens)>2){
			bfr := tokens[len(tokens)-1]
			if bfr.Symbol != rts_semicolon && bfr.Symbol != rts_open_brace && bfr.Symbol != rts_close_brace && bfr.Type != rtt_newline{
				return tokens,instructions
			}
		}
		instructions = append(instructions, makeRunicInstruction(tokens))
		tokens = []RunicToken{}
	}
	return tokens,instructions
}

func (this *RunicLexer) parseMonoline(tokens []RunicToken, instructions []RunicInstruction)([]RunicToken,[]RunicInstruction){
	if(len(tokens) >= 2 && tokens[len(tokens)-1].Symbol == rts_open_brace && tokens[len(tokens)-2].Symbol == rts_close_parenthese){
		instructions = append(instructions, makeRunicInstruction(tokens))
		tokens = []RunicToken{}
	}
	return  tokens,instructions
}
