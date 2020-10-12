package global

import (
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

//MakeData 数据
type MakeData map[string]interface{}

var conf *goconfig.ConfigFile

//User 当前用户信息
var User UserModel

//Json json格式化
func Json(message string, data interface{}) gin.H {
	if data == nil {
		data = []string{}
	}
	return gin.H{"message": message, "data": data}
}

func initConfig() {
	var err error
	conf, err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Println(err.Error())
	}
}

//Config 配置读取
func Config(section string, key string) string {
	if conf == nil {
		initConfig()
	}
	sec, err := conf.GetSection(section)
	if err != nil {
		log.Panicln(err.Error())
	}
	return sec[key]
}

//GetRedis 从连接池中获取一个redis连接
func GetRedis() *redis.Client {
	maxActive, _ := strconv.Atoi(Config("redis", "maxActive"))
	db, _ := strconv.Atoi(Config("redis", "db"))
	return redis.NewClient(&redis.Options{
		Addr:     Config("redis", "address"),
		Password: "",
		DB:       db,
		PoolSize: maxActive,
	})
}
