package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
func check(s string,lookup [10]string)int{
for i:=0;i<len(lookup);i++{
if lookup[i]==s{
return i
}
}
return -1
}
func words2num(s string)string{
lookup:=[10]string{"zero","one","two","three","four","five","six","seven","eight","nine"}
arr:=strings.Split(s," ")
ans:=""
for i:=0;i<len(arr);i++{
temp:=check(strings.ToLower(arr[i]),lookup)
if temp!=-1{
ans=ans+strconv.Itoa(temp)
}else{
ans="Invalid input"
break
}
}
return ans
}

func main() {
	arr:=os.Args
	fmt.Println(words2num(arr[1]));
}
