
package api

import (
    "api"
    "xiaowo/api/model"
    "github.com/jinzhu/gorm"
)

type sessionCheckHandler struct {
    api.BaseRequest
}

func (self *sessionCheckHandler) PreExecute() (user *model.User, db *gorm.DB, success bool) {
    success = false;

    sessionID, err := self.GetString("sessionID");
    if nil != err {
        self.Response(err);
        return;
    }

    user, err = checkSession(sessionID, self.Debug);
    if nil != err {
        self.Response(err);
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != e {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    return user, db, true;
}

