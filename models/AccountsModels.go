package models

type UserLogin struct {
	BaseModel

	Id       int    `json:"id" gorm:"size:500;primaryKey"`
	Token    string `json:"token" gorm:"size:500;index:idx_token;not null"`
	UserName string `json:"user_name" gorm:"size:500;index:idx_user_name;not null"`
	Email    string `json:"email" gorm:"size:500;index:idx_email;not null"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

func (o UserLogin) TableName() string {
	return "user"
}
