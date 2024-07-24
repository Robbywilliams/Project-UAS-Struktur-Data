// File path: Project_StrukturData/model/model.go

package model

import (
	"Project_StrukturData/database"
	"Project_StrukturData/node"
)

// InsertMember menambahkan anggota baru ke dalam linked list di database.
func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := &node.TableMember{
		Data: node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	var tempLL *node.TableMember
	tempLL = &database.DBmember
	if database.DBmember.Next == nil {
		database.DBmember.Next = newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = newDataMember
	}
}

// ReadAllMember mengembalikan semua anggota dari linked list di database.
func ReadAllMember() *node.TableMember {
	if database.DBmember.Next == nil {
		return nil
	}
	return database.DBmember.Next
}

// DeleteMember menghapus anggota dari linked list berdasarkan id.
func DeleteMember(id int) bool {
	var tempLL *node.TableMember
	tempLL = &database.DBmember
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}
}

// SearchMember mencari anggota berdasarkan id di linked list.
func SearchMember(id int) *node.TableMember {
	var tempLL *node.TableMember
	tempLL = database.DBmember.Next
	if database.DBmember.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				return tempLL
			}
			tempLL = tempLL.Next
		}
	}
	return nil
}

// UpdateMember memperbarui informasi anggota yang ada.
func UpdateMember(member node.Member) bool {
	addr := SearchMember(member.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Nama = member.Nama
		addr.Data.Alamat = member.Alamat
		addr.Data.Point = member.Point
		return true
	}
}
