package controllers

import (
	"dati/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 玩家操作API
type UserController struct {
	beego.Controller
	Username string
}

func (u *UserController) Prepare() {
	fmt.Println(" ........... Prepare...........")
	currentUser := u.GetSession("username")
	if strings.Contains(u.Ctx.Request.URL.Path, "/login") {
		return
	}
	if currentUser == nil || currentUser == "" {
		var response models.ResponseData
		response.Status = 500
		response.Msg = "请先登录"
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
	u.Username, _ = currentUser.(string)
}

// @Title answerQuestion
// @Description 普通玩家回答问题
// @Param	questID		query 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Param	answer		query 	string	true		"答案选项，比如 A ，B  C ，不分大小写"
// @Success 200 成功
// @Failure 500 失败
// @router /answer [post]
func (u *UserController) Post() {
	var response models.ResponseData
	var thisAnswer string
	var err error
	var questid int
	questid, err = u.GetInt("questID")
	if err != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "问题的id 不是数字" + err.Error()
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
	if questid != models.CurrentQuestion {
		response.Status = 500
		if models.CurrentQuestion == -1 {
			response.Msg = "哥，还没开始做题."
		} else {

			response.Msg = "兄弟，现在不是在做这道题."
		}
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
	thisAnswer = u.GetString("answer")
	if thisAnswer == "" {
		response.Status = 500
		response.Msg = "answer不合法"
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
	inputanswer := strings.ToLower(thisAnswer)
	switch inputanswer {
	case "a":
	case "b":
	case "c":
	default:
		response.Status = 500
		response.Msg = "只能abc"
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
	temp2 := *(models.AllUserAnswer[u.Username])
	temp2[questid] = thisAnswer
	response.Status = 200
	response.Msg = "录入成功"
	response.Data = &thisAnswer
	u.Data["json"] = &response
	u.ServeJSON()
}

// @Title getMyAllAnswer
// @Description 获取自己的所有答案
// @Success 200 成功
// @Failure 500 失败
// @router /getMyAllAnswer [get]
func (u *UserController) GetAll() {
	var response models.ResponseData
	response.Status = 200
	response.Msg = "我的答案"
	response.Data = models.AllUserAnswer[u.Username]

	u.Data["json"] = &response
	u.ServeJSON()
}

// @Title GetAnswer
// @Description 获取自己某一题的答案
// @Param	id		path 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Success 200 成功
// @Failure 500 失败
// @router /answer/:id [get]
func (u *UserController) Get() {
	var response models.ResponseData

	questidstr := u.Ctx.Input.Param(":id")
	questid, questerr := strconv.Atoi(questidstr)
	if questerr != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "id 不合法" + questerr.Error()
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}

	response.Status = 200
	response.Msg = "获取成功"
	myAllanswer := *(models.AllUserAnswer[u.Username])
	response.Data = myAllanswer[questid]
	u.Data["json"] = &response
	u.ServeJSON()
}

// @Title Login
// @Description 普通玩家登录
// @Param	name		query 	string	true		"登录名"
// @Success 200 成功
// @Failure 500 失败
// @router /login [post]
func (u *UserController) Login() {
	var response models.ResponseData
	name := u.GetString("name")
	if name == "" {
		response.Status = 500
		response.Msg = "name不合法"
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}

	currentUserSession := u.GetSession("username")
	if currentUserSession != nil {
		if currentUser, ok := currentUserSession.(string); ok {
			if currentUser != name {
				response.Status = 500
				response.Msg = "请不要中途换名"
				u.Data["json"] = &response
				u.ServeJSON()
				return
			}
		}
	} else {
		if models.AllUserAnswer[name] != nil {
			response.Status = 500
			response.Msg = "此用户名已经被占用。"
			u.Data["json"] = &response
			u.ServeJSON()
			return
		}

	}
	// myAnswer := (models.AllUserAnswer[currentUser])
	// if myAnswer == nil {
	temp := make(map[int]string)
	models.AllUserAnswer[name] = &temp
	//}
	u.SetSession("username", name)
	response.Status = 200
	response.Msg = "登录成功"
	response.Data = models.AllUserAnswer[name]
	u.Data["json"] = &response
	u.ServeJSON()
}
