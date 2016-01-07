package LocalCmd

import (
	"bytes"
	"log"
	"os/exec"
)

func Main() {
	//执行ls命令  exec.Command("/bin/ls", "-l", "/")
	//执行who命令 exec.Command("/bin/sh","-c","wh")
	//	cmd := exec.Command("/bin/ls", "-l", "/")
	cmd := exec.Command("python", "/home/miss/haha/message/python-sdk-v2/send_sms.py", "12345678912", "Test from program")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Cmd execute error!!info:\n", err.Error())
	}
	log.Printf(out.String())
}
