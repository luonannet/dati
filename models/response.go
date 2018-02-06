package models

//ResponseData 服务端输出的数据
type ResponseData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
