package main



func setPreprocessorInstruction(s string)(string){
	return 	`package main
	import "os"

	func program_return(i int){
		os.Exit(i)
	}

	func runicVersion()string{
		return "Runic BETA 0.0.0"
	}

	`
}
				