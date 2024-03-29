package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/pankajojha/go-test-drive/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	fmt.Println(string(body))

	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", p1)
}
