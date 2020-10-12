package chengyu

import (
	"encoding/json"
	D "go_api/lib/database"
	G "go_api/lib/global"
	"go_api/lib/response"
	T "go_api/lib/tools"
	V "go_api/lib/valid"
	"strings"

	"github.com/gin-gonic/gin"
)

var valid = V.Validate{}

//Data 返回数据结构
type Data struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Mp3    string `json:"mp3"`
	Gif    string `json:"bihua_gif"`
	Png    string `json:"bihua_png"`
	Pinyin string `json:"pinyin"`
	Search string `json:"search"`
	Zuci   string `json:"zuci"`
}

//CyData 返回成语数据结构
type CyData struct {
	ID        int64  `json:"id"`
	Chengyu   string `json:"chengyu"`
	Mp3       string `json:"mp3"`
	Pinyin    string `json:"pinyin"`
	Hanzi     string `json:"hanzi"`
	Mean      string `json:"mean"`
	HanziList []Data `json:"hanzi_list"`
}

//TableName 为结构体设置表名
func (c *Data) TableName() string {
	return "bihua"
}

//Get 登录
func Get(c *gin.Context) {
	db := D.DBT("chengyu")
	hanzi := c.Query("hanzi")
	redis := G.GetRedis()
	//如果未传值,随机从redis里取key
	if count, _ := redis.HLen("chengyu").Result(); len(hanzi) <= 0 && count > 500 {
		hanzi = redis.SRandMember("chengyu_key").Val()
	}
	res := redis.HGet("chengyu", hanzi)
	if err := res.Err(); err == nil {
		data := CyData{}
		b, _ := res.Bytes()
		json.Unmarshal(b, &data)
		data.HanziList = getzi(data.Hanzi)
		c.JSON(response.HTTPStatusOK, G.Json("获取数据成功", data))
		return
	}
	if len(hanzi) > 0 {
		db.Where("chengyu", hanzi)
	} else {
		db.OrderBy("rand()").Limit(1)
	}
	zi, err := db.First()
	if err != nil {
		c.JSON(response.HTTPStatusFaild, err.Error())
		return
	}
	if len(zi) < 1 {
		c.JSON(response.HTTPStatusOK, G.Json("暂未收录此成语", zi))
		return
	}
	hanzi = zi["chengyu"].(string)
	zi["hanzi_list"] = getzi(zi["hanzi"].(string))
	js, err := json.Marshal(zi)
	redis.HSet("chengyu", hanzi, string(js))
	redis.SAdd("chengyu_key", hanzi)
	c.JSON(response.HTTPStatusOK, G.Json("获取数据成功", zi))
}

//获取所有字的图片 参数格式天,天,向,上
func getzi(hanzi string) []Data {
	if len(hanzi) <= 0 {
		return []Data{}
	}
	zs := strings.Split(hanzi, ",")
	data := []Data{}
	for i := range zs {
		d := get(zs[i])
		if d.ID > 0 {
			data = append(data, d)
		}
	}
	return data
}

func get(hanzi string) Data {
	db := D.DBT("bihua")
	redis := G.GetRedis()
	//如果未传值,随机从redis里取key
	if count, _ := redis.HLen("bihua").Result(); len(hanzi) <= 0 && count > 500 {
		hanzi = redis.SRandMember("bihua_key").Val()
	}
	res := redis.HGet("bihua", hanzi)
	if err := res.Err(); err == nil {
		data := Data{}
		b, _ := res.Bytes()
		json.Unmarshal(b, &data)
		return data
	}
	if len(hanzi) > 0 {
		db.Where("search", T.MD5(hanzi))
	} else {
		db.OrderBy("rand()").Limit(1)
	}
	zi, err := db.First()
	if err != nil {
		return Data{}
	}
	if len(zi) < 1 {
		return Data{}
	}
	hanzi = zi["name"].(string)
	js, err := json.Marshal(zi)
	redis.HSet("bihua", hanzi, string(js))
	redis.SAdd("bihua_key", hanzi)
	data := Data{}
	data.ID = zi["id"].(int64)
	data.Name = zi["name"].(string)
	data.Gif = zi["bihua_gif"].(string)
	data.Png = zi["bihua_png"].(string)
	return data
}
