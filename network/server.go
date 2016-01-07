package network

import (
	"fmt"
	"net"
	"os"
)

func StartServer() {
	//每运行一次这个文件就会产生，因此在运行之前需要将这个文件删除
	filename := "/tmp/echo.sock"
	//os.Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	listen, err := net.Listen("unix", filename)
	//defer listen.Close()这会报错误Go: panic: runtime error: invalid memory address or nil pointer dereference
	//报这个错误的原因是执行了上一个语句以后，err不是nil
	//The defer only defers the function call. The field and method are accessed immediately.
	//也就是说 defer只是推迟函数的调用的时间，而这个函数的内容是会被立即访问到的。
	//这时候报了一个err,而listen可能就为空，所以，listen.Close()就会报错误。
	//ref:http://stackoverflow.com/questions/16280176/go-panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference
	if err != nil {
		panic("Write: " + err.Error())
	}
	defer listen.Close()
	for {
		connect, errCon := listen.Accept()
		if errCon != nil {
			panic("Write " + errCon.Error())
		}
		go forEachRuquest(connect)
	}
}

func forEachRuquest(con net.Conn) {
	for {
		buf := make([]byte, 512)
		num, err := con.Read(buf)
		if err != nil {
			panic("Write: " + err.Error())
		}
		data := buf[:num]
		fmt.Printf("\nServer: (recive Message from client)---" + string(data))
		_, errWrite := con.Write([]byte("Hello, Client I'm server!\n"))

		if errWrite != nil {
			panic("Write :" + errWrite.Error())
		}
	}
}
