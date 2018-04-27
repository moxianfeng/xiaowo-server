package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model;
    UserID string `gorm:"not null;column:user_id;type:varchar(64);unique_index:idx_user_id"`
    City string `gorm:"not null;column:city;type:varchar(64);" json:"city"`
    Province string `gorm:"not null;column:province;type:varchar(64);" json:"province"`
    Country string `gorm:"not null;column:country;type:varchar(64);" json:"country"`
    AvatarUrl string `gorm:"not null;column:avatar_url;type:varchar(255);" json:"avatarUrl"`
    NickName string `gorm:"not null;column:nick_name;type:varchar(64);" json:"nickName"`
    Gender int `gorm:"not null;column:gender;type:int;" json:"gender"`
    Language string `gorm:"not null;column:language;type:varchar(64);" json:"language"`
    PhoneNumber string `gorm:"not null;column:phone_number;type:varchar(20);index:idx_phone_number"`
    SessionKey string `gorm:"not null;column:session_key;type:varchar(64);unique_index:idx_session_key"`
}

func (self User) TableName() string {
    return "t_user";
}

func init() {
    RegisterObject(&User{});
}

