package models

type BaseModel struct {
}

type CreateToken struct {
	BaseModel

	TokenId         int    `json:"token_id" gorm:"size:500;primaryKey"`
	Token           string `json:"token" gorm:"size:500;index:idx_token;not null"`
	TokenExpiration int64  `json:"token_expiration" gorm:"index:;autoTokenExpiration:milli;not null"`
	CreatedAt       int64  `json:"created_at" gorm:"index:;autoCreatedAt:milli;not null"`
}

func (o CreateToken) TableName() string {
	return "token"
}
