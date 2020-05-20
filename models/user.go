package models

type UserModel struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
