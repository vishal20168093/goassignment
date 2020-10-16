package main

import (
	"fmt"
	"os"
	"strconv"
)
func maxPower(m int,n int)int{

k:=0
temp:=1
for ;m>temp; {
temp=temp*n
k++
}
return k
}
func main() {
	arr:=os.Args

	temp1,err1:=strconv.Atoi(arr[1])
	temp2,err2:=strconv.Atoi(arr[2])
	if err1!=nil{
	fmt.Println(err1);
	}
	if err2!=nil{
	fmt.Println(err2);
	}
	fmt.Println(maxPower(temp1,temp2));
}
