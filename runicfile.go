package main
import "io/ioutil"

	func makeRunicFile(s string)(rf RunicFile){
	 	rf.OldPath =  stringToPath(s)
	 	rf.NewPath =  rf.OldPath[0:len(rf.OldPath)-5]+"go"
		return
	}

	func (this *RunicFile) compile(){
		s:= ""
		for _,i := range this.Content{
			s += i.CompiledContent + "\n"
		}
		ioutil.WriteFile(this.NewPath, []byte(s), 0644)	
	}