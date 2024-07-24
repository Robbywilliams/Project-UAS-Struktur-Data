package main

import (
	"Project_StrukturData/model"
	"Project_StrukturData/node"
	"fmt"
)

func main() {
	model.InsertMember(1, "Robby", "Surabaya", 100)
	model.InsertMember(2, "Williams", "London", 22)
	model.InsertMember(3, "Pascal", "Jamaica", 33)

	member1 := node.Member{
		Id:     1,
		Nama:   "ITATS",
		Alamat: "Surabaya",
		Point:  100,
	}
	tes := model.UpdateMember(member1)
	fmt.Println("Return Update Member : ", tes)
	Member := model.ReadAllMember()
	if Member != nil {
		for Member != nil {
			fmt.Println(Member.Data)
			Member = Member.Next
		}
	}
}
