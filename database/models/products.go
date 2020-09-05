package models

// Product ...
type Product struct {
	ID   int    `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name string `json:"name" gorm:"column:name;not null;varchar(50)"`
	Type string `json:"type" gorm:"column:type;not null;default:'general';varchar(25)"`
}

// TableName ...
func (p *Product) TableName() string {
	return "product"
}
