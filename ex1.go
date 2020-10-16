package main

import (
	"fmt"
	"os"
)
func num2words(n string)string{
ans:=""
arr:=[10]string{"zero","one","two","three","four","five","six","seven","eight","nine"}

for i:=0;i<len(n);i++ {
temp:=n[i]-'0'
if temp>=0 && temp<=9{
ans=ans+arr[temp]+" "
}else{

ans="Not a number"
break
}

}

return ans
}
func main() {
	arr:=os.Args
	for i:=1;i<len(arr);i++{
	fmt.Println(num2words(arr[i]));
	}
}
