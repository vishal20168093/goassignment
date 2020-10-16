package main

import (
	"fmt"
	"os"
	"strconv"
)
func num2hex(n int,swi string)string{
if swi=="-u" {
return fmt.Sprintf("%X",n)
}else if swi=="-l" {
return fmt.Sprintf("%x",n)
}else{
return "Invalid option"
}

}


func main() {
	args:=os.Args
	if(len(args)==2){
	temp,err:=strconv.Atoi(args[1])
	if(err==nil){
	fmt.Println(num2hex(temp,"-l"))
	}else{
	fmt.Println("Invalid number")
	}
	}else if(len(args)==3){
	temp,err:=strconv.Atoi(args[2])
	if(err==nil){
	fmt.Println(num2hex(temp,args[1]))
	}else{
	fmt.Println("Invalid number")
	}
	}else{
	fmt.Println("Invalid option")
	}
}
