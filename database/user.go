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
	_all_user_field = utils.GetGormFields(User{})
)

func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User
	if err := db.Select(_all_user_field).Where("name=?", name).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			utils.LogRus.Errorf("get password of user %s failed: %s", name, err)
		}
		return nil
	}
	return &user
}

// 创建用户
func CreateUser(name, pass string) {
	db := GetBlogDBConnection()
	pass = utils.Md5(pass)
	user := User{Name: name, PassWd: pass}
	if err := db.Create(&user).Error; err != nil {
		utils.LogRus.Errorf("create user %s failed: %s", name, err)
	} else {
		utils.LogRus.Infof("create user id %d", user, err)
	}
}

// 删除用户
func DeleteUser(name string) {
	db := GetBlogDBConnection()
	if err := db.Where("name=?", name).Delete(User{}).Error; err != nil {
		utils.LogRus.Errorf("delete user %s failed: %s", name, err)
	}
}
