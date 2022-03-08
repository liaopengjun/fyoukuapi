package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	redisM "fyoukuapi/services/redis"
	"github.com/garyburd/redigo/redis"
)

type RedisDemoController struct {
	beego.Controller
}

// @router /redis/demo [*]
func (this *RedisDemoController) Demo() {
	c := redisM.PoolConnect()
	defer c.Close()
	//_,err := c.Do("SET","username","frog")
	//if err ==nil{
	//	c.Do("expire","username",20)
	//}
	r,err := redis.String(c.Do("get","username"))
	if err ==nil{
		fmt.Println(1)
		fmt.Println(r)
		//剩余过期时间
		ttl,_ := redis.Int64(c.Do("ttl","username"))
		fmt.Println(ttl)
	}else{
		fmt.Println(2)
		fmt.Println(err)
	}
	this.Ctx.WriteString("hello world")

}