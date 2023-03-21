package entity

// TableName 数据库表明自定义，默认为model的复数形式
func (User) TableName() string {
	return "user"
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	DelStatus  int    `json:"delStatus"`
	CreateTime int64  `json:"createTime"`
}
