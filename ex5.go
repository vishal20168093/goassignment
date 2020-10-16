package main
import(
	"fmt"
	"os"
	"strconv"
)
func main(){
args:=os.Args
temp1,err1:=strconv.ParseComplex(args[1],128)
temp2,err2:=strconv.ParseComplex(args[3],128)
if(err1==nil && err2==nil){

if(args[2]=="+"){
fmt.Println(temp1+temp2)
}else if(args[2]=="-"){
fmt.Println(temp1-temp2)
}else if(args[2]=="*"){
fmt.Println(temp1*temp2)
}else{
fmt.Println("please enter a valid operator")
}

}else{
fmt.Println("please enter a valid number")
}

}