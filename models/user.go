package models

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

var (
	UserMap map[string]*User
	Users   []*User
	Admin   AdminStruct
)

func init() {
	UserMap = make(map[string]*User)
	Users = make([]*User, 0)
}

type User struct {
	Username string
	Count    int
}

//AdminStruct 管理员
type AdminStruct struct {
	Name string `json:"name"`
	Pswd string `json:"pswd"`
}

type UserList []*User

func (u UserList) Less(i, j int) bool {

	return u[i].Count > u[j].Count
}
func (u UserList) Len() int {

	return len(u)
}
func (u UserList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]

}

//保存主持人
func SaveAdminToFile() (err error) {
	var adminBuffer bytes.Buffer
	encoder := gob.NewEncoder(&adminBuffer)
	err = encoder.Encode(Admin)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(adminPath, adminBuffer.Bytes(), os.ModePerm)
	if err != nil {

		return
	}
	return nil
}
