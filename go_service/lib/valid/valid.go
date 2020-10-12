package valid

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"mime/multipart"
	"strconv"
	"strings"
)

//Validate 验证
type Validate struct {
	validation validation.Validation
}

//Rule 规则
type Rule struct {
	Name  string
	Input interface{}
	Rule  []string
}

//Check 批量验证 true 通过 false 未通过
func (v *Validate) Check(key string, obj interface{}, rules []string) (string, bool) {
	var param string
	access := true
	message := ""
	for i := range rules {
		action := rules[i]
		s := strings.Split(action, ":")
		if len(s) > 1 {
			action = s[0]
			param = s[1]
		}
		switch action {
		case "require":
			v.validation.Required(obj, key).Message("[%s]不能为空", key)
			break
		case "numeric":
			v.validation.Numeric(obj, key).Message("[%s]必须是数字", key)
			break
		case "float":
			// v.validation.Numeric(obj, key).Message("[%s]必须是数字", key)
			v.floatCheck(obj, key)
			break
		case "mobile":
			v.validation.Mobile(obj, key).Message("[%s]格式不是手机号码", key)
			break
		case "min":
			i, _ := strconv.Atoi(param)
			num := interface2Int(obj)
			v.validation.Min(num, i, key).Message("[%s]最小值为: %s,请求值为: %s", key, param, obj.(string))
			break
		case "max":
			i, _ := strconv.Atoi(param)
			num := interface2Int(obj)
			v.validation.Max(num, i, key).Message("[%s]最大值为: %s,请求值为: %s", key, param, obj.(string))
			break
		case "length":
			i, _ := strconv.Atoi(param)
			v.validation.Length(obj, i, key).Message("[%s]长度必须为: %s,请求长度为: %d", key, param, len(obj.(string)))
			break
		case "minsize":
			i, _ := strconv.Atoi(param)
			v.validation.MinSize(obj, i, key).Message("[%s]最小长度为: %s,请求长度为: %d", key, param, len(obj.(string)))
			break
		case "maxsize":
			i, _ := strconv.Atoi(param)
			v.validation.MaxSize(obj, i, key).Message("[%s]最大长度为: %s,请求长度为: %d", key, param, len(obj.(string)))
			break
		case "image":
			v.imageCheck(obj, key)
		default:
		}
		if v.validation.HasErrors() {
			access = false
			message = v.validation.Errors[0].String()
			v.validation.Clear()
			break
		}
	}
	return message, access
}

//CheckList 批量验证
func (v *Validate) CheckList(rules []Rule) (string, bool) {
	for i := range rules {
		message, ok := v.Check(rules[i].Name, rules[i].Input, rules[i].Rule)
		if !ok {
			return message, ok
		}
	}
	return "", true
}

//验证是不是图片
func (v *Validate) imageCheck(obj interface{}, name string) {
	fh, ok := obj.(*multipart.FileHeader)
	if !ok {
		v.validation.Error(fmt.Sprintf("[%s]必须是图片文件", name))
		return
	}
	header := fh.Header.Values("Content-Type")
	fmt.Println(header[0])
	if ok := len(header) > 0; !ok || header[0] != "image/png" || header[0] != "image/jpeg" {
		v.validation.Error(fmt.Sprintf("[%s]必须是图片文件(png,jpg)", name))
		return
	}
}

func interface2Int(obj interface{}) int64 {
	str, ok := obj.(string)
	if !ok {
		return 0
	}
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func interface2Float(obj interface{}) (float64, bool) {
	str, ok := obj.(string)
	if !ok {
		return 0, false
	}
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

func (v *Validate) floatCheck(obj interface{}, name string) {
	//由于post请求过来的数据全是字符串,而类型断言出来就是字符串,所以需要先将字符串转换成相应的数字类型然后进行类型断言
	_, ok := interface2Float(obj)
	if !ok {
		v.validation.Error(fmt.Sprintf("[%s]必须是整数或浮点", name))
		return
	}
}
