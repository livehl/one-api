package model

import (
	"github.com/songquanpeng/one-api/common/helper"
)

type Contact struct {
	Id      int    `json:"id"`
	Phone   string `json:"phone" gorm:"index"`
	Created int64  `json:"created" gorm:"bigint"`
	Name    string `json:"name"`
	Info    string `json:"info"`
}

func AddContact(phone string, name string, info string) (err error) {
	contact := &Contact{
		Phone:   phone,
		Name:    name,
		Info:    info,
		Created: helper.GetTimestamp(),
	}
	err = DB.Create(contact).Error
	return err
}

func GetAllContacts(startIdx int, num int, order string) (contacts []*Contact, err error) {
	query := DB.Limit(num).Offset(startIdx)

	switch order {
	case "name":
		query = query.Order("name desc")
	case "phone":
		query = query.Order("phone desc")
	case "created":
		query = query.Order("created desc")
	default:
		query = query.Order("id desc")
	}
	err = query.Find(&contacts).Error
	return contacts, err
}
