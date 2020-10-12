package bihua

import (
	"encoding/json"
	D "go_api/lib/database"
	G "go_api/lib/global"
	"go_api/lib/response"
	T "go_api/lib/tools"
	V "go_api/lib/valid"

	"github.com/gin-gonic/gin"
)

var valid = V.Validate{}

//Data 返回数据结构
type Data struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Mp3     string `json:"mp3"`
	Gif     string `json:"bihua_gif"`
	Png     string `json:"bihua_png"`
	Pinyin  string `json:"pinyin"`
	Search  string `json:"search"`
	Zuci    string `json:"zuci"`
	Chengyu string `json:"chengyu"`
}

//CyData 返回成语数据结构
type CyData struct {
	ID      int    `json:"id"`
	Chengyu string `json:"chengyu"`
	Mp3     string `json:"mp3"`
	Pinyin  string `json:"pinyin"`
	Hanzi   string `json:"hanzi"`
	Mean    string `json:"mean"`
}

//TableName 为结构体设置表名
func (c *CyData) TableName() string {
	return "chengyu"
}

//Get 登录
func Get(c *gin.Context) {
	db := D.DBT("bihua")
	hanzi := c.Query("hanzi")
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
		data.Chengyu = chengyu(hanzi)
		c.JSON(response.HTTPStatusOK, G.Json("获取数据成功", data))
		return
	}
	if len(hanzi) > 0 {
		db.Where("search", T.MD5(hanzi))
	} else {
		db.OrderBy("rand()").Limit(1)
	}
	zi, err := db.First()
	if err != nil {
		c.JSON(response.HTTPStatusFaild, err.Error())
		return
	}
	if len(zi) < 1 {
		c.JSON(response.HTTPStatusOK, G.Json("暂未收录此字", zi))
		return
	}
	hanzi = zi["name"].(string)
	zi["chengyu"] = chengyu(hanzi)
	js, err := json.Marshal(zi)
	redis.HSet("bihua", hanzi, string(js))
	redis.SAdd("bihua_key", hanzi)
	c.JSON(response.HTTPStatusOK, G.Json("获取数据成功", zi))
}

//获取一列成语用来展示
func chengyu(hanzi string) string {
	if len(hanzi) <= 0 {
		return ""
	}
	var data []CyData
	db := D.DB().Table(&data)
	redis := G.GetRedis()
	res := redis.HGet("chengyu_list", hanzi)
	if err := res.Err(); err == nil {
		return res.Val()
	}
	var cache bool = true //如果有字开头的成语就存到redis,如果没有就直接读数据库取包含这个字的成语
	db.Where("hanzi", "like", hanzi+"%").Select()
	if len(data) <= 0 {
		D.DB().Table(&data).Where("hanzi", "like", "%"+hanzi+"%").OrderBy("rand()").Limit(10).Select()
		cache = false
	}
	response := ""
	for i := range data {
		response += data[i].Chengyu + ","
	}
	if cache && len(response) > 0 {
		redis.HSet("chengyu_list", hanzi, response)
	}
	return response
}
