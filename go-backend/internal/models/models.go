package models

type Post struct {
	Id       string `json:"id" gorm:"primaryKey" validate:"uuid"`
	Text     string `json:"text" validate:"required,min=5,max=255"`
	AuthorId string `json:"authorId" gorm:"column:author_id;not null;index"`
	ImgSrc   string `json:"imgSrc" gorm:"column:img_src"`
	Likes    []Like `json:"-" gorm:"foreignKey:PostId"`
	PostDate string `json:"postDate" gorm:"column:post_date"`
}

type User struct {
	Id        string `json:"id" gorm:"primaryKey" validate:"required,uuid"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Username  string `json:"username" gorm:"unique" validate:"required,min=3,max=16"`
	Password  string `json:"password" validate:"required,min=6"`
	Avatar    string `json:"avatar"`
	Followers int    `json:"followers" validate:"gte=0"`
	Likes     []Like `json:"-" gorm:"foreignKey:UserId"`
}

type PostRequest struct {
	Text     string `json:"text" validate:"required,min=5,max=255"`
	ImgSrc   string `json:"imgSrc" validate:"required,url"`
	AuthorId string `json:"authorId" validate:"required,uuid"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Avatar   string `json:"avatar"`
}

type Like struct {
	Id     string `json:"id" gorm:"primaryKey" validate:"uuid"`
	UserId string `json:"userId" gorm:"not null;index" gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	PostId string `json:"postId" gorm:"not null;index" gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
