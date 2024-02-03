package database

import (
	"blog-server/utils"
	"gorm.io/gorm"
)

type User struct {
	Id     int    `gorm:"column:id;primaryKey"`
	Name   string `gorm:"column:name"`
	PassWd string `gorm:"column:password"`
}

func (User) TableName() string {
	return "userone"
}

var (
// _all_user_field = utils.GetGormFields(User{})
)

func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User
	if err := db.Select([]string{"id", "name", "password"}).Where("name=?", name).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			utils.LogRus.Errorf("get password of user %s failed: %s", name, err)
		}
		return nil
	}
	return &user
}
