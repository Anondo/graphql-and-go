package repos

import (
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/models"
	"github.com/jinzhu/gorm"
)

// UserRepo ...
type UserRepo struct {
	db *conn.DB
}

// NewUserRepo ...
func NewUserRepo(db *conn.DB) User {
	return &UserRepo{
		db: db,
	}
}

// GetUsers ...
func (u *UserRepo) GetUsers() ([]models.User, error) {

	users := []models.User{}

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// GetUser ...
func (u *UserRepo) GetUser(userID int) (*models.User, error) {
	user := models.User{}

	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}
