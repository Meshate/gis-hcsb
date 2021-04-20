package models

import (
	"fmt"
	. "gis-hcsb"
	"log"
)

type Record struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Road     string `gorm:"column:road" json:"road"`
}

func (r *Record) GetByUsername() []Record {
	var ret []Record
	if err := DB.Where("username = ?", r.Username).Find(&ret).Error; err != nil {
		log.Println(err)
		return ret
	}
	return ret
}

func (r *Record) Insert() bool {
	if err := DB.Create(r).Error; err != nil {
		fmt.Errorf("%s\n", err)
		return false
	}
	return true
}
