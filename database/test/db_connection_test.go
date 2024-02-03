package test

import (
	"blog-server/database"
	"fmt"
	"testing"
)

func TestGetBlogDBConnection(t *testing.T) {
	const C = 100
	for i := 0; i < C; i++ {
		database.GetBlogDBConnection()
	}
}

func TestGetUserByName(t *testing.T) {
	res := database.GetUserByName("user1")
	fmt.Print(res)
}
