package models

type Post struct {
	Id       string `json:"id" gorm:"primaryKey" validate:"required,uuid"`
	Text     string `json:"text" validate:"required,min=5,max=255"`
	Author   string `json:"author" validate:"required,uuid"`
	ImgSrc   string `json:"imgSrc" gorm:"column:imgSrc"`
	Likes    int    `json:"likes" validate:"gte=0"`
	PostDate string `json:"postDate" gorm:"column:postDate"`
}

type User struct {
	Id        string `json:"id" gorm:"primaryKey" validate:"required,uuid"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=16"`
	Password  string `json:"password" validate:"required,min=6"`
	Avatar    string `json:"avatar"`
	Followers int    `json:"followers" validate:"gte=0"`
}
