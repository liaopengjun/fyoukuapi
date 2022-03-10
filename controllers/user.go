package controllers

import (
	"fyoukuapi/models"
	beego "github.com/beego/beego/v2/server/web"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

//SaveRegister 用户注册
// @router /register/save [post]
func (this *UserController) SaveRegister() {
	var (
		mobile string
		password string
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" || password == "" {
		this.Data["json"] = ReturnError(4001,"手机号和密码必须填写")
	}

	isorno,_:= regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if isorno == false {
		this.Data["json"] = ReturnError(4001,"手机号不能为空")
	}
	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4005,"此手机已经注册")
	} else {
		err := models.UserSave(mobile,Md5V(password))
		if err !=nil {
			this.Data["json"] = ReturnError(5000,"")
		}else {
			this.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
		}
	}
	this.ServeJSON()

}
//Login 登录
// @router /login/do [post]
func (this *UserController) Login() {
	var (
		mobile string
		password string
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")
	if mobile == "" || password == "" {
		this.Data["json"] = ReturnError(4001,"手机号和密码必须填写")
	}
	isorno,_:= regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if isorno == false {
		this.Data["json"] = ReturnError(4001,"手机号不能为空")
	}
	id,name := models.IsUserLogin(mobile,Md5V(password))
	if id == 0 && name == ""{
		this.Data["json"] = ReturnError(4001,"用户名或密码不正确")
	}else{
		item := map[string]interface{}{"uid":id,"name":name}
		this.Data["json"] = ReturnSuccess(0,"登录成功",item,0)
	}
	this.ServeJSON()
}

//批量发送通知消息
// @router /send/message [*]
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")

	if uids == "" {
		this.Data["json"] = ReturnError(4001, "请填写接收人~")
		this.ServeJSON()
	}
	if content == "" {
		this.Data["json"] = ReturnError(4002, "请填写发送内容")
		this.ServeJSON()
	}
	messageId, err := models.SendMessageDo(content)
	if err == nil {
		uidConfig := strings.Split(uids, ",")
		for _, v := range uidConfig {
			userId, _ := strconv.Atoi(v)
			models.SendMessageUserMq(userId, messageId)
		}
		this.Data["json"] = ReturnSuccess(0, "发送成功~", "", 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, "发送失败，请联系客服~")
		this.ServeJSON()
	}
}