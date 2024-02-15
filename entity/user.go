package entity

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
