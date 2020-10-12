package response

//ProductIDEmpty 商品id为空
const ProductIDEmpty = 1

//HTTPStatusOK http状态
const HTTPStatusOK = 200

//HTTPStatusFaild http状态
const HTTPStatusFaild = 400

//Response 返回内容
type Response struct {
	Code    int
	Message string
}

var TokenInvalid = Response{400, "无效用户token/token未放入header"}
var UserInvalid = Response{400, "用户不存在"}
