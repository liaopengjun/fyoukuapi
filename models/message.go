package models

import (
	"encoding/json"
	"fyoukuapi/services/mq"
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

//保存通知消息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	messageId, err := o.Insert(&message)
	return messageId, err
}

//保存消息接收人
func SendMessageUser(userId int, messageId int64) error {
	o := orm.NewOrm()
	var messageUser MessageUser
	messageUser.UserId = userId
	messageUser.MessageId = messageId
	messageUser.Status = 1
	messageUser.AddTime = time.Now().Unix()
	_, err := o.Insert(&messageUser)
	return err
}

//保存消息接收人
func SendMessageUserMq(userId int, messageId int64)  {
	type Data struct {
		UserId int
		MessageId int64
	}
	data := Data{
		UserId:userId,
		MessageId:messageId,
	}
	videoJson, _ := json.Marshal(data)
	mq.Publish("", "fyouku_send_message_user", string(videoJson))
}