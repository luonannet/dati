package models

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

var (
	//UserMap key 用户名，value用户
	UserMap map[string]*User
	//Users 用户列表
	Users UserList //[]*User
	//Admin 主持人
	Admin AdminStruct
)

//User 玩家
type User struct {
	Username string
	Count    int
}

//AdminStruct 主持人
type AdminStruct struct {
	Name string `json:"name"`
	Pswd string `json:"pswd"`
}

//UserList 用户列表的别名
type UserList []*User

func init() {
	UserMap = make(map[string]*User)
	Users = UserList(make([]*User, 0))
}

func (u UserList) Less(i, j int) bool {
	return u[i].Count > u[j].Count
}
func (u UserList) Len() int {

	return len(u)
}
func (u UserList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]

}

//SaveAdminToFile 保存主持人
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
