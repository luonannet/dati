package controllers

import (
	"dati/models"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//AdminController 主持人操作的 API
type AdminController struct {
	beego.Controller
}

func (o *AdminController) Prepare() {
	if strings.Contains(o.Ctx.Request.URL.Path, "/login") || strings.Contains(o.Ctx.Request.URL.Path, "/createAdmin") {
		return
	}
	currentUser := o.GetSession("admin")
	if currentUser == nil || currentUser == "" {
		var response models.ResponseData
		response.Status = 400
		response.Msg = "请先登录"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
}

// @Title login
// @Description 主持人登录
// @Param	name		query 	string	true		"登录名"
// @Param	pswd		query 	string	true		"密码"
// @Success 200 成功
// @Success 400 {models.ResponseData}
// @Failure 500 失败
// @router /login [post]
func (o *AdminController) Login() {
	var response models.ResponseData
	name := o.GetString("name")
	if name == "" {
		response.Status = 500
		response.Msg = "name不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	pswd := o.GetString("pswd")
	if pswd == "" {
		response.Status = 500
		response.Msg = "pswd不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	if models.Admin.Name != name || models.Admin.Pswd != pswd {
		response.Status = 500
		response.Msg = "账户不对"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	o.SetSession("admin", name)
	response.Status = 200
	response.Msg = "登录成功"
	response.Data = &models.Admin
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title CreateAdmin
// @Description 创建主持人。重复创建会覆盖
// @Param	name		query 	string	true		"登录名"
// @Param	pswd		query 	string	true		"密码"
// @Success 200 成功
// @Failure 500 失败
// @router /createAdmin [post]
func (o *AdminController) CreateAdmin() {
	var response models.ResponseData
	models.Admin.Name = o.GetString("name")
	if models.Admin.Name == "" {
		response.Status = 500
		response.Msg = "name不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	models.Admin.Pswd = o.GetString("pswd")
	if models.Admin.Pswd == "" {
		response.Status = 500
		response.Msg = "pswd不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}

	err := models.SaveAdminToFile()
	if err != nil {
		response.Status = 500
		response.Msg = "models.SaveAdminToFile err :" + err.Error()
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Status = 200
	response.Msg = "admin初始化成功"
	response.Data = &models.Admin
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title InputAnswer
// @Description 主持人录入正确答案
// @Param	questID		query 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Param	answer		query 	string	true		"答案选项，比如 A ，B  C ，不分大小写"
// @Success 200 成功
// @Failure 500 失败
// @router /inputAnswer [post]
func (o *AdminController) Post() {

	var response models.ResponseData
	var thisAnswer string
	var err error
	var questid int
	questid, err = o.GetInt("questID")
	if err != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "问题的id 不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	thisAnswer = o.GetString("answer")
	if thisAnswer == "" {
		response.Status = 500
		response.Msg = "answer不合法"
		o.Data["json"] = &response
		o.ServeJSON()
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
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	models.StandAnswer[questid] = thisAnswer
	err = models.SaveStandAnswerToFile()
	if err != nil {
		response.Status = 500
		response.Msg = "SaveStandAnswerToFile err:" + err.Error()
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Status = 200
	response.Msg = "录入成功"
	response.Data = &thisAnswer
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title GetSingleAnswer
// @Description 获取一个问题的答案
// @Param	id		path 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Success 200 成功
// @Failure 500 失败
// @router /question/:id [get]
func (o *AdminController) Get() {
	questidStr := o.Ctx.Input.Param(":id")
	var response models.ResponseData
	questid, questerr := strconv.Atoi(questidStr)
	if questerr != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "id不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	if models.StandAnswer[questid] == "" {
		response.Status = 500
		response.Msg = "这题没答案"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Status = 200
	response.Msg = questidStr + "题的答案"

	response.Data = models.StandAnswer[questid]
	o.Data["json"] = &response

	o.ServeJSON()
}

// @Title GetAllAnswer
// @Description 获取所有问题的答案
// @Success 200 成功
// @Failure 500 失败
// @router /answer/getAll [get]
func (o *AdminController) GetAllAnswer() {
	var response models.ResponseData
	response.Status = 200
	response.Msg = "所有问题的答案"

	response.Data = models.StandAnswer
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title DeleteAnswer
// @Description 删除这道题的答案
// @Param	id		path 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Success 200 成功
// @Failure 500 失败
// @router /answer/:id [delete]
func (o *AdminController) Delete() {
	questidStr := o.Ctx.Input.Param(":id")
	var response models.ResponseData
	questid, questerr := strconv.Atoi(questidStr)
	if questerr != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "id不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	delete(models.StandAnswer, questid)
	err := models.SaveStandAnswerToFile()
	if err != nil {
		response.Status = 500
		response.Msg = "SaveStandAnswerToFile err:" + err.Error()
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Status = 200
	response.Msg = "所有问题的答案"
	response.Data = models.StandAnswer
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title CreateBangdan
// @Description 生成排名榜单
// @Success 200 成功
// @Failure 500 失败
// @router /createBangdan [get]
func (o *AdminController) CreateBangdan() {
	var response models.ResponseData
	response.Status = 200
	response.Msg = "榜单排行"
	var result models.UserList
	models.AllUserAnswer.Range(func(username, answerMap interface{}) bool {
		currentUser := new(models.User)
		currentUser.Username, _ = username.(string)
		tempAnswer, _ := (answerMap.(models.StandAnswerStruct))
		for myQuestid, myanswer := range tempAnswer {
			if strings.EqualFold(myanswer, models.StandAnswer[myQuestid]) {
				currentUser.Count++
			}
		}
		fmt.Println(username, currentUser.Count)
		result = append(result, currentUser)
		return true
	})
	// for username, answerMap := range models.AllUserAnswer {
	// 	currentUser := new(models.User)
	// 	currentUser.Username = username
	// 	tempAnswer := (map[int]string)(answerMap)

	// 	for myQuestid, myanswer := range tempAnswer {
	// 		if strings.EqualFold(myanswer, models.StandAnswer[myQuestid]) {
	// 			currentUser.Count++
	// 		}
	// 	}

	// 	result = append(result, currentUser)
	// }
	sort.Sort(result)
	err := models.SaveBangdanToFile(result)
	if err != nil {
		response.Status = 500
		response.Msg = "SaveBangdanToFile: " + err.Error()
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Data = result
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title GetBangdan
// @Description 获取已经生成的榜单，没有榜单的话就再去生成一次
// @Success 200 成功
// @Failure 500 失败
// @router /getBangdan [get]
func (o *AdminController) GetBangdan() {
	var response models.ResponseData
	response.Status = 200
	response.Msg = "榜单排行"
	bangdan, err := models.GetBangdanFile()
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			o.CreateBangdan()
			return
		default:
			response.Status = 500
			response.Msg = "GetBangdan: " + err.Error()
			o.Data["json"] = &response
			o.ServeJSON()
			return
		}

	}
	response.Data = bangdan
	o.Data["json"] = &response
	o.ServeJSON()
}

// @Title SetQuestionID
// @Description 设置当前是哪道题
// @Param	questID		query 	int	true		"题号， 数字，第一题就写1  第二题就写2"
// @Success 200 成功
// @Failure 500 失败
// @router /setQuestionID [post]
func (o *AdminController) SetQuestionID() {
	var response models.ResponseData
	var err error
	var questid int
	questid, err = o.GetInt("questID")
	if err != nil || questid <= 0 {
		response.Status = 500
		response.Msg = "questID 不合法"
		o.Data["json"] = &response
		o.ServeJSON()
		return
	}
	response.Status = 200
	response.Msg = "设置成功"
	models.CurrentQuestion = questid
	response.Data = models.CurrentQuestion
	o.Data["json"] = &response
	o.ServeJSON()
}
