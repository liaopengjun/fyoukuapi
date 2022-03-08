package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
	redisClient "fyoukuapi/services/redis"
)

type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func init()  {
	orm.RegisterModel(new(User))
}

// IsUserMobile 验证用户是否注册
func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user,"Mobile")
	if err == orm.ErrNoRows{
		return false
	}
	return true
}

// UserSave 保存用户
func UserSave(mobile string,password string)  (err error){
	o := orm.NewOrm()
	var user User
	user.Name = ""
	user.Password = ""
	user.Mobile = mobile
	user.Password = password
	user.Status = 1
	user.AddTime = time.Now().Unix()
	_, err = o.Insert(&user)
	return err
}
//IsUserLogin 验证登录
func IsUserLogin(mobile string,password string) (uid int,name string)  {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("mobile",mobile).Filter("password",password).One(&user)
	if err == orm.ErrNoRows || err ==orm.ErrMissPK{
		return 0,""
	}
	return user.Id,user.Name
}


//根据用户ID获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	o := orm.NewOrm()
	var user UserInfo
	err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
	return user, err
}

func RedisGetUserInfo(uid int) (UserInfo,error) {
	var user UserInfo
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "user:id:"+strconv.Itoa(uid) //int转成string
	exists,err := redis.Bool(conn.Do("exists",redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &user)
	}else{
		o := orm.NewOrm()
		err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
		if err == nil {
			//保存redis
			_, err = conn.Do("hmset", redis.Args{redisKey}.AddFlat(user)...)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return user, err
}