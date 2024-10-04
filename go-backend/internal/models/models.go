package models

type Post struct {
	Id       string `json:"id" gorm:"primaryKey" validate:"required,uuid"`
	Text     string `json:"text" validate:"required,min=5,max=255"`
	AuthorID string `json:"authorId" gorm:"column:author_id;not null;index"` // Foreign key to User
	Author   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ImgSrc   string `json:"imgSrc" gorm:"column:imgSrc"`
	Likes    int    `json:"likes" validate:"gte=0"`
	PostDate string `json:"postDate" gorm:"column:postDate"`
}

type User struct {
	Id        string `json:"id" gorm:"primaryKey" validate:"required,uuid"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Username  string `json:"username" gorm:"unique" validate:"required,min=3,max=16"`
	Password  string `json:"password" validate:"required,min=6"`
	Avatar    string `json:"avatar"`
	Followers int    `json:"followers" validate:"gte=0"`
	Posts     []Post `gorm:"foreignKey:AuthorID"` // One-to-many relationship with Post
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

type NbaTeam struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
}

type NbaGame struct {
	Id               int     `json:"id" gorm:"primaryKey"`
	Date             string  `json:"date"`
	HomeTeam         NbaTeam `json:"home_team"`
	HomeTeamScore    int     `json:"home_team_score"`
	Period           int     `json:"period"`
	Postseason       bool    `json:"postseason"`
	Season           int     `json:"season"`
	Status           string  `json:"status"`
	Time             string  `json:"time"`
	VisitorTeam      NbaTeam `json:"visitor_team"`
	VisitorTeamScore int     `json:"visitor_team_score"`
}
