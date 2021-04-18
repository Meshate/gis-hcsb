package models

import (
	. "gis-hcsb"
	"log"
)

type User struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (u *User) CheckUsername() bool {
	if err := DB.Where("username = ?", u.Username).First(u).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u *User) Insert() bool {
	if err := DB.Create(u).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u *User) LoginCheck() bool {
	if err := DB.Where("username = ? and password = ?", u.Username, u.Password).First(u).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
