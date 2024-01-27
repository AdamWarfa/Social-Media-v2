package models

type Post struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Text     string `json:"text"`
	Author   string `json:"author"`
	ImgSrc   string `json:"imgSrc" gorm:"column:imgSrc"`
	Likes    int    `json:"likes"`
	PostDate string `json:"postDate" gorm:"column:postDate"`
}

type User struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Followers int    `json:"followers"`
}
