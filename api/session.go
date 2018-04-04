package api

import (
    "time"
    "family/api/model"
    "api"
)

const (
    VALID_SESSION = time.Second * 3600;
)

func checkSession(sessionId string, debug bool) (*model.User, *api.Error) {
    db, err := GetDB(debug);
    if nil != err {
        return nil, api.E_MISS_DB_CONNECTION.Apply("%s", err.Error());
    }

    var user model.User;
    result := db.Where(&model.User{SessionKey: sessionId}).First(&user);
    if result.RecordNotFound() {
        return nil, api.E_INVALID_SESSION;
    }

    now := time.Now();
    valid := user.UpdatedAt.Add(VALID_SESSION);
    if !valid.After(now) {
        return nil, api.E_INVALID_SESSION;
    }

    // update UpdatedAt
    db.Model(&user).Update(model.User{SessionKey: user.SessionKey});

    return &user, nil;
}
