package models

import (
	"time"
)

type File struct{
	Id int64
	File_sha1 string `orm:"unique"`
	File_name string
	File_size int64
	File_addr string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
	Status int64 `orm:"null"`
	Device string
	Version string
}

type User struct {
	Id int64
	UserName string `orm:"unique"`
	UserPwd string
	Email string `orm:"null"`
	Phone string `orm:"null"`
	SignUpAt string `orm:"auto_now_add;type(datetime)"`
}
