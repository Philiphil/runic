package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
)

	func (this *RunicCompiler) getNewFileList()([]string){
		var s []string
		for i:=0;i < len(this.Files); i++{
		 	s = append(s,this.Files[i].NewPath)
		}
		return append(s, this.RunicHeaderPath)
	}


	func (this *RunicCompiler) getRunCommand()(string, []string){
		r := this.getNewFileList()
		r= append( []string{this.Action/*,this.Naming+" "+this.Name*/}, r...)
		return this.Compiler, r
	}

	func (this *RunicCompiler) parse(){
		for i:=0; i < len(this.Files); i++{
			this.Files[i].Content = this.Lexer.parse(this.Files[i])
		}

		this.Parser.initParser(this.Lexer)
		var errs []RunicError
		for i:=0; i < len(this.Files); i++{
			var err []RunicError
			this.Files[i].Content, err = this.Parser.parseInstruction(this.Files[i].Content)

			if (len(err) > 0){
				for _,r := range err{
					r.File = this.Files[i].OldPath
					errs = append(errs,r)
				}
			}
			
		}
		if(len(errs)>0){
			this.panic(errs)
		}
	}

	func (this *RunicCompiler) compile(){
		for i:=0; i < len(this.Files); i++{
			this.Files[i].compile()
			if len(filepath.Dir(this.Files[i].OldPath)) < len(this.RunicHeaderPath) || this.RunicHeaderPath == "" {
				this.RunicHeaderPath = filepath.Dir(this.Files[i].OldPath)
			}
		}

		this.RunicHeaderPath += string(os.PathSeparator)
		this.Path = this.RunicHeaderPath

		this.RunicHeaderPath += this.Mask + "runic.go"
		ioutil.WriteFile(this.RunicHeaderPath,[]byte(setPreprocessorInstruction("")), 0644)	
	}

	func makeRunicCompiler()(rc RunicCompiler){
	 	rc.Compiler = "go"
	 	rc.Action = "";
	 	rc.Naming =  "-o"
	 	rc.RunicHeaderPath = ""
	 	rc.Path = ""
	 	rc.Name = "a.o"
	 	rc.Mask = rc.getMask()
		return
	}

	func (this *RunicCompiler) getMask()(string){
		return randSeq(5)

	}

	func (this *RunicCompiler) panic(errs []RunicError){
		fmt.Println( len(errs), "ERRORS !")
		for _,e := range errs{
			e.print()
		}
		os.Exit(-1)
	}