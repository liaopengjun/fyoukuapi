package controllers

import (
	"crypto/md5"
	"encoding/hex"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64 `json:"count"`
}

//ReturnSuccess 返回成功
func ReturnSuccess(code int,msg interface{},items interface{},count int64) (json *JsonStruct){
	json = &JsonStruct{
		Code: code,
		Msg: msg,
		Items: items,
		Count: count,
	}
	return
}
// ReturnError 返回错误
func ReturnError (code int,msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{
		Code: code,
		Msg: msg,
	}
	return
}

func Md5V(password string)string  {
	h := md5.New()
	confKey, _ := beego.AppConfig.String("md5code")
	h.Write([]byte(password+confKey))
	return hex.EncodeToString(h.Sum(nil))
}

//格式化时间
func DateFormat(times int64) string {
	video_time := time.Unix(times, 0)
	return video_time.Format("2006-01-02")
}

