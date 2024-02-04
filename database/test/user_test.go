package test

import (
	"blog-server/database"
	"testing"
)

func TestCreateUser(t *testing.T) {
	name, pass := "huangfl", "12345"
	database.CreateUser(name, pass)
}

func TestDeleteUser(t *testing.T) {
	name := "huangfl"
	database.DeleteUser(name)
}
