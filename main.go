package main

import (
	"fmt"
	//GOPATH
	//     |_src
	//          |_the-way-to-go
	//                         |_GoroutinesAndChannels(文件夹)  需要注意的是 文件夹的名字和这个文件夹下所有go source文件的第一行都是package GoroutinesAndChannels
	// 在引用包时以 GOPATH/src目录为相对根目录
	//在同一个package里，多个文件被go编译器看作是一个文件，因此，这多个文件中不能出现相同的全局变量和函数，一个例外是init函数；
	//而同一个package的不同文件可以直接引用相互之间的数据。
	"encoding/json"

	pb "the-way-to-go/GoProtobuffer"
	a "the-way-to-go/GoroutinesAndChannels" //这里使用包别名  一个包可以由包含多个文件，因此文件名是无所谓的
	"the-way-to-go/JsonMarshal"
)

func main() {
	fmt.Printf("-----------using channels --------------\n")
	b := a.A(3) //虽然引用了package,但是如果不使用package名做限定是无法直接使用A(type int)的
	b.Main()
	fmt.Printf("\n-----------Marchal&Unmarchal json-----------\n")
	jsonInstance := &JsonMarshal.StatusDetails{
		Kind: "kind",
		Name: "name",
		Causes: []JsonMarshal.StatusCause{{
			Type:  "duplicate",
			Field: "",
		}},
	}
	j, err := json.Marshal(jsonInstance)
	if err != nil {
		fmt.Print("json convert fail!")
	} else {
		fmt.Print(string(j)) //j 是[]byte类型，可以直接转换成string
	}
	fmt.Printf("\n-------------------using protobuffer---------------\n")
	pb.Main()

}
