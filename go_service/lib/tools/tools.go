package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"

	"strconv"
)

//MD5 md5加密,支持字符串,整数
func MD5(obj interface{}) string {
	var str string = ""
	//change to string
	switch obj.(type) {
	case string:
		str = obj.(string)
	case int:
		str = strconv.Itoa(obj.(int))
		break
	case int64:
		str = strconv.FormatInt(obj.(int64), 10)
		break
	}
	m := md5.Sum(bytes.NewBufferString(str).Bytes())
	return hex.EncodeToString(m[:])
}
