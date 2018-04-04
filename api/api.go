package api

import (
    _ "github.com/go-sql-driver/mysql"

    "family/api/model"
    "github.com/jinzhu/gorm"
)

const (
    MYSQL_DSN = "family:familyadmin@tcp(127.0.0.1:3306)/family?charset=utf8&parseTime=True&loc=Local";
)

var _db *gorm.DB;
var _dbDebug *gorm.DB;

func _getDB(debug bool) (*gorm.DB, error) {
    if debug {
        if _dbDebug != nil {
            return _dbDebug, nil;
        }
    } else {
        if nil != _db {
            return _db, nil;
        }
    }

    var e error;
    _db, e = gorm.Open("mysql", MYSQL_DSN);
    if nil == e {
        _dbDebug = _db.Debug();
    } else {
        return nil, e
    }

    if debug {
        return _dbDebug, nil;
    } else {
        return _db, nil;
    }
}

func GetDB(debug bool) (*gorm.DB, error) {
    db, err := _getDB(debug);
    return db, err;
}

func init() {
    db, err := GetDB(false);
    if nil == err {
        db.Debug().Set("gorm:table_options", "ENGINE=InnoDB,DEFAULT CHARACTER SET=utf8").AutoMigrate(model.GetModels()...);
    }
}
