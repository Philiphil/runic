package main

import(
	"os"
	"bytes"
	"syscall"
	"os/exec"
	"regexp"
	"path/filepath"
	"math/rand"
	"time"
	"reflect"
)

	func stringToPath(path string) string{
		abs,err := filepath.Abs(path)
		if  _, err2 := os.Stat(abs); err2 != nil || err !=nil{
		 	panic(path+" not found, path="+abs)
		}
	   return abs
	}

	func RunCommand(name string, params []string) (stdout string, stderr string, exitCode int) {
	    var outbuf, errbuf bytes.Buffer

	    cmd := exec.Command(name, params...)
	    cmd.Stdout = &outbuf
	    cmd.Stderr = &errbuf

	    err := cmd.Run()
	    stdout = outbuf.String()
	    stderr = errbuf.String()

	    if err != nil {
	        if exitError, ok := err.(*exec.ExitError); ok {
	            ws := exitError.Sys().(syscall.WaitStatus)
	            exitCode = ws.ExitStatus()
	        } else {
	            exitCode = -1
	            if stderr == "" {
	                stderr = err.Error()
	            }
	        }
	    } else {
	        ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
	        exitCode = ws.ExitStatus()
	    }
	    return
	}

	func isAlphaNum(s byte)(b bool){
		 r, _ := regexp.Compile("([a-zA-ZÀ-ÿ|_|0-9]+)")
		return r.MatchString(string(s))
	}
	

func randSeq(n int) string {
	 rand.Seed(time.Now().UnixNano())
	 var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func HasElem(s interface{}, elem interface{}) bool {
    arrV := reflect.ValueOf(s)

    if arrV.Kind() == reflect.Slice {
        for i := 0; i < arrV.Len(); i++ {

            if arrV.Index(i).Interface() == elem {
                return true
            }
        }
    }

    return false
}