package main

import(
"fmt"
"os"
)

	func main(){
		var rc RunicCompiler = makeRunicCompiler()

		if len(os.Args) < 3  {
			fmt.Println( "File(s) or option(build or run) missing" )
			return 
		}


		for i :=1 ; i < len(os.Args); i++ {
			if i == 1{
				rc.Action = os.Args[i]
				continue
			}
			rc.Files = append(rc.Files, makeRunicFile(os.Args[i]))
		}

		rc.DeleteTMP = true

		rc.parse()
		rc.compile()

		cmd, params := rc.getRunCommand()
		
		fmt.Println(RunCommand(cmd,params))

	}