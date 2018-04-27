package model

import (
    "github.com/jinzhu/gorm"
)

type VerifyCode struct {
    gorm.Model;
    PhoneNumber string `gorm:"not null;column:phone_number;type:varchar(64);index:idx_phone_number"`
    VerifyCode string `gorm:"not null;column:verify_code;type:varchar(64);"`
}

func (self VerifyCode) TableName() string {
    return "t_verify_code";
}

func init() {
    RegisterObject(&VerifyCode{});
}
