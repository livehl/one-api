package model

import (
	"errors"
)

const (
	ComboStatusEnabled  = 1 // don't use 0, 0 is the default value!
	ComboStatusDisabled = 2 // also don't use 0
	ComboStatusDeleted  = 3
)

// User if you add sensitive fields, don't forget to clean them in setupLogin function.
// Otherwise, the sensitive information will be saved on local storage in plain text!
type Combo struct {
	Id       int    `json:"id"`
	Name     string `json:"name" gorm:"unique;index" validate:"max=12"`
	Detail   string `json:"detail" `
	Price    int    `json:"price"`
	Status   int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	CreateAt int64  `json:"create_at" gorm:"bigint;default:0"`
}

func GetAllCombos() (combos []*Combo, err error) {
	query := DB.Where("status != ?", ComboStatusDeleted).Order("id desc")
	err = query.Find(&combos).Error
	return combos, err
}

func GetComboById(id int) (*Combo, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	combo := Combo{Id: id}
	var err error = nil
	err = DB.First(&combo, "id = ?", id).Error
	return &combo, err
}
