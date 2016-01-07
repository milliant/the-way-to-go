package network

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func StartClient() {
	con, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic("Write: " + err.Error())
	}
	defer con.Close()
	//启动一个channel，可以理解为启动一个新的线程，如果这个函数中不包括循环结构，那么这个channel将函数中的
	//代码依次执行一遍就结束了
	go readFromServer(con)
	for {
		//		buf := make([]byte, 512)
		_, errWrite := con.Write([]byte("Hi,Server I'm client"))
		if errWrite != nil {
			panic("Write:" + errWrite.Error())
			break
		}
		time.Sleep(1e9)
	}
}

func readFromServer(con net.Conn) {
	buf := make([]byte, 512)
	for {
		n, err := con.Read(buf[:])
		if err != nil {
			panic("Write :" + err.Error())
		}
		fmt.Printf("\nClient: (get message from server)" + string(buf[0:n]))
	}
}

func UseGet() {
	fmt.Printf("\n-------------------using net(http.Get)---------------\n")
	resp, errGet := http.Get("http://ugistry.ucloud.cn/v1/search")
	if errGet != nil {
		panic(errGet.Error())
		os.Exit(1)
	}
	//从response中读取数据
	content := make([]byte, 512)
	num, errRead := resp.Body.Read(content)
	if errRead != nil && errRead != io.EOF {
		panic(errRead.Error())
		os.Exit(1)
	}
	fmt.Printf("Information is: ", string(content[:num]))
}
