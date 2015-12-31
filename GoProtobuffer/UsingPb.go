package GoProtobuffer

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
	miUser "milliant_protobuffer"
)

func Main() {
	user := &miUser.User{
		Uid:  proto.Int32(1),
		Name: proto.String("blackbeans"),
	}
	encObj, err := proto.Marshal(user)
	if nil == err {
		fmt.Println("length:", len(encObj))
		tobj := &miUser.User{}
		e := proto.Unmarshal(encObj, tobj)
		if nil == e {
			fmt.Println(tobj.GetName())
		} else {
			log.Fatalln("decode fail ", e)
		}
	} else {
		log.Fatalln("encode fail", err)
	}
}
