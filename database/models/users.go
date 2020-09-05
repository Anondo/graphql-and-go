package models

// User ...
type User struct {
	ID        int    `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	FirstName string `json:"first_name" gorm:"column:first_name;not null;varchar(25)"`
	LastName  string `json:"last_name" gorm:"column:last_name;default:'';varchar(25)"`
	PhoneNo   string `json:"phone_no" gorm:"column:phone_no;not null;varchar(32)"`
}

// TableName ...
func (u *User) TableName() string {
	return "user"
}
