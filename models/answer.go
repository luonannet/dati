package models

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

//AnswerStruct 正确答案
// type AnswerStruct struct {
// 	Answer string `json:"answer"`
// }
//type StandAnswerStruct map[int]string

const (
	standAnswerPath = "./standAnswer.data"
	adminPath       = "./admin.data"
	bangdanPath     = "./bangdan.data"
)

//StandAnswer 正确答案的列表
var StandAnswer map[int]string
var AllUserAnswer map[string]*map[int]string //key 用户的编号，int题目的编号。value 是答案]
var CurrentQuestion int                      //当前在做哪道题
func init() {
	CurrentQuestion = -1
	StandAnswer = make(map[int]string)
	AllUserAnswer = make(map[string]*map[int]string)

	standAnswerData, err := ioutil.ReadFile(standAnswerPath)
	if err == nil {
		var standanswerBuffer bytes.Buffer
		standanswerBuffer.Write(standAnswerData)
		decoder := gob.NewDecoder(&standanswerBuffer)
		decoder.Decode(&StandAnswer)
	}
	//--------------------
	adminrData, adminerr := ioutil.ReadFile(adminPath)
	if adminerr == nil {
		var standanswerBuffer bytes.Buffer
		standanswerBuffer.Write(adminrData)
		decoder := gob.NewDecoder(&standanswerBuffer)
		decoder.Decode(&Admin)
	}

}

//保存标准答案
func SaveStandAnswerToFile() (err error) {
	var standanswerBuffer bytes.Buffer
	encoder := gob.NewEncoder(&standanswerBuffer)
	err = encoder.Encode(StandAnswer)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(standAnswerPath, standanswerBuffer.Bytes(), os.ModePerm)
	if err != nil {

		return
	}
	return nil
}

//保存榜单
func SaveBangdanToFile(bangdan UserList) (err error) {
	var bangdanBuffer bytes.Buffer
	encoder := gob.NewEncoder(&bangdanBuffer)
	err = encoder.Encode(bangdan)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(bangdanPath, bangdanBuffer.Bytes(), os.ModePerm)
	if err != nil {

		return
	}
	return nil
}

//获取榜单
func GetBangdanFile() (bangdan *UserList, err error) {
	bangdanData, err := ioutil.ReadFile(bangdanPath)
	if err == nil {
		var bangdanrBuffer bytes.Buffer
		bangdanrBuffer.Write(bangdanData)
		decoder := gob.NewDecoder(&bangdanrBuffer)
		bangdan = new(UserList)
		decoder.Decode(bangdan)
		return bangdan, err
	}
	return nil, err
}
