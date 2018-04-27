package model

import (
    "github.com/jinzhu/gorm"
)

type Home struct {
    gorm.Model `json:"-"`;
    UserID string `gorm:"not null;column:user_id;type:varchar(64);index:idx_user_id" json:"-"`
    HomeID string `gorm:"not null;column:home_id;type:varchar(64);unique_index:idx_home_id" json:"homeID"`
    Name string `gorm:"not null;column:name;type:varchar(64);" json:"name"`
    City string `gorm:"not null;column:city;type:varchar(64);" json:"city"`
    Province string `gorm:"not null;column:province;type:varchar(64);" json:"province"`
    Country string `gorm:"not null;column:country;type:varchar(64);" json:"country"`
    BackgroundUrl string `gorm:"not null;column:background_url;type:varchar(255);" json:"backgroundUrl"`
    RoomsCount int `gorm:"not null;column:rooms_count;type:int;" json:"roomsCount"`
    DevicesCount int `gorm:"not null;column:devices_count;type:int;" json:"devicesCount"`
}

func (self Home) TableName() string {
    return "t_home";
}

func init() {
    RegisterObject(&Home{});
}
