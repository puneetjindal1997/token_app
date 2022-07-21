package models

type Admin struct {
	BaseModel

	Id        int    `json:"id" gorm:"size:500;primaryKey"`
	FullName  string `json:"full_name" gorm:"size:500;index:idx_full_name;not null"`
	Email     string `json:"email" gorm:"size:500;index:idx_email;not null"`
	Password  string `json:"password" gorm:"size:500;index:idx_password;not null"`
	CreatedAt int64  `json:"created_at" gorm:"index:;autoCreatedAt:milli;not null"`
}

func (o Admin) TableName() string {
	return "admin"
}
