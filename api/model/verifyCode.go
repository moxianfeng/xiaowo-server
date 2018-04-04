package model

import (
    "github.com/jinzhu/gorm"
)

type VerifyCode struct {
    gorm.Model;
    UserID string `gorm:"not null;column:user_id;type:varchar(64);unique_index:idx_user_id"`
    PhoneNumber string `gorm:"not null;column:phone_number;type:varchar(64);index:idx_phone_number"`
    VerifyCode string `gorm:"not null;column:verify_code;type:varchar(64);"`
    SessionKey string `gorm:"not null;column:session_key;type:varchar(64);unique_index:idx_session_key"`
}

func (self VerifyCode) TableName() string {
    return "t_verify_code";
}

func init() {
    RegisterObject(&VerifyCode{});
}
