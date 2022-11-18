package model

type User struct {
	ID       uint64
	username string
	password string
	urls     []UrL
}
