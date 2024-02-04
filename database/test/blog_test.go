package test

import (
	"blog-server/database"
	"fmt"
	"testing"
)

func TestUpdateBlog(t *testing.T) {
	blog := database.Blog{
		Id:      1,
		Title:   "双十一",
		Article: "双十一来临喜洋洋, 购物狂欢乐无边。电商盛宴蔓延芳，心愿成真喜乐颜",
	}
	if err := database.UpdateBlog(&blog); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
