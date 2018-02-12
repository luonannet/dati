package controllers

import (
	"dati/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 玩家操作API
type UserController struct {
	beego.Controller
}

func (u *UserController) Prepare() {
	currentUser := u.GetSession("username")
	if strings.Contains(u.Ctx.Request.URL.Path, "/login") {
		return
	}
	if currentUser == nil || currentUser == "" {
		var response models.ResponseData
		response.Status = 400
		response.Msg = "请先登录"
		u.Data["json"] = &response
		u.ServeJSON()
		return
	}
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
		response.Msg = "questID 异常：" + err.Error()
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
		response.Data = models.CurrentQuestion
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
	currentUser := u.GetSession("username")
	username, _ := currentUser.(string)
	var myAllAnswer models.StandAnswerStruct
	temp2, _ := models.AllUserAnswer.Load(username)
	if temp2 == nil {
		myAllAnswer = make(models.StandAnswerStruct)
	} else {
		myAllAnswer, _ = temp2.(models.StandAnswerStruct)
	}
	myAllAnswer[questid] = thisAnswer
	models.AllUserAnswer.Store(username, myAllAnswer)
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
	currentUser := u.GetSession("username")
	username, _ := currentUser.(string)
	response.Msg = "我的答案"
	response.Data, _ = models.AllUserAnswer.Load(username)

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
	currentUser := u.GetSession("username")
	username, _ := currentUser.(string)
	myAllanswer, _ := models.AllUserAnswer.Load(username)
	myAllAnswer, _ := myAllanswer.(map[int]string)
	response.Data = (myAllAnswer)[questid]
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
	currentUser := u.GetSession("username")
	username, _ := currentUser.(string)
	if username != "" {
		if username != name {
			response.Status = 500
			response.Msg = "请不要中途换名"
			response.Data = username
			u.Data["json"] = &response
			u.ServeJSON()
			return
		}
	} else {
		if _, ok := models.AllUserAnswer.Load(name); ok {
			response.Status = 500
			response.Msg = "此用户名已经被占用。"
			response.Data = name
			u.Data["json"] = &response
			u.ServeJSON()
			return
		}
		username = name

	}
	if _, ok := models.AllUserAnswer.Load(username); ok == false {
		models.AllUserAnswer.Store(username, make(models.StandAnswerStruct))
	}

	u.SetSession("username", name)
	response.Status = 200
	response.Msg = "登录成功"
	response.Data, _ = models.AllUserAnswer.Load(username)
	u.Data["json"] = &response
	u.ServeJSON()
}

// @Title Logout
// @Description 玩家中途退出，这个只在开发测试的时候才有用
// @Success 200 成功
// @Failure 500 失败
// @router /logout [get]
func (u *UserController) Logout() {
	var response models.ResponseData
	currentUser := u.GetSession("username")
	username, _ := currentUser.(string)
	//	models.AllUserAnswer.Delete(username)
	response.Status = 200
	response.Msg = "退出登录"
	u.DelSession("username")
	response.Data = username
	u.Data["json"] = &response
	u.ServeJSON()
}
