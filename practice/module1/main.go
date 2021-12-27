package main

import (
	"fmt"
	"go-practice/practice/module1/slice"
)

func main() {
	//homework1.1 给定一个字符串数组
	//[“I”,“am”,“stupid”,“and”,“weak”]
	//用 for 循环遍历该数组并修改为
	//[“I”,“am”,“smart”,“and”,“strong”]
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("mySlice:", mySlice)
	mySlice = slice.ReplaceStringInSlice(mySlice)
	fmt.Println("replaced mySlice:", mySlice)

}
