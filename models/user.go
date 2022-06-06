package models

type User struct {
	UserID   string `db:"userId" json:"userId" gorm:"primaryKey; column:userId"`
	Email    string `db:"email" json:"email" gorm:"unique"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

func (User) TableName() string {
	return "usr"
}
