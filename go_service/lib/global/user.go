package global

//UserModel 用户
type UserModel struct {
	UID      int    `gorose:"id"`
	Mobile   string `gorose:"mobile"`
	Nickname string `gorose:"nickname"`
	Token    string `gorose:"token"`
	Expire   int    `gorose:"expired_at"`
}

//TableName 用户model的表名
func (u *UserModel) TableName() string {
	return "user"
}
