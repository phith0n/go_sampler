package mysql

type UserTab struct {
	BaseModel

	Username    string `json:"username" gorm:"column:username;not null;index;size:150;"`
	Email       string `json:"email" gorm:"column:email;not null;unique;size:254;"`
	Password    string `json:"-" gorm:"column:password;not null;size:128;"`
	IsSuperuser bool   `json:"is_superuser" gorm:"column:is_superuser;default:0;"`
}

func (u *UserTab) TableName() string {
	return "user"
}
