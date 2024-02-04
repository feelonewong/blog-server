package database

import (
	"blog-server/utils"
	"fmt"
	"gorm.io/gorm"
)

type Blog struct {
	Id         int    `gorm:"column:id;primaryKey"`
	UserId     int    `gorm:"column:id;primaryKey"`
	Title      string `gorm:"column:title"`
	Article    string `gorm:"column:article"`
	UpdateTime string `gorm:"column:update_time"`
}

func (Blog) TableName() string {
	return "blog"
}

var (
	_all_blog_field = utils.GetGormFields(Blog{})
)

// id 查询博客内容
func GetBlogById(id int) *Blog {
	db := GetBlogDBConnection()
	var blog Blog
	if err := db.Select(_all_blog_field).Where("id=?", id).First(&blog).Error; err != nil {
		//	必须传blog指针
		if err != gorm.ErrRecordNotFound {
			utils.LogRus.Errorf("get content of blog %d failed: %s", id, err)
		}
		return nil
	}
	return &blog
}

// 根据作者id查询博客列表
func GetBlogByUserId(uid int) []*Blog {
	db := GetBlogDBConnection()
	var blogs []*Blog
	if err := db.Select("id, title").Where("user_id=?", uid).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			utils.LogRus.Errorf("get blog of user: %d failed: %s", uid, err)
		}
		return nil
	}
	return blogs
}

// 根据博客id更新标题和正文
func UpdateBlog(blog *Blog) error {
	if blog.Id <= 0 {
		return fmt.Errorf("could not update blog of id %d", blog.Id)
	}
	if len(blog.Article) == 0 || len(blog.Title) == 0 {
		return fmt.Errorf("could not set blog title or article to empty")
	}
	db := GetBlogDBConnection()
	return db.Model(Blog{}).Where("id=?", blog.Id).Updates(map[string]any{"title": blog.Title, "article": blog.Article}).Error
}
