package api

import (
    "api"

    "xiaowo/api/model"
)

type getGardensHandler struct {
    api.BaseRequest
}

type getGardensResponse struct {
    *api.Error
    Gardens []model.Garden `json:"gardens"`
}

func init() {
    api.RegisterExecutor("GetGardens", func (br api.BaseRequest) api.Request {
        return &getGardensHandler{BaseRequest: br}
    });
}

func (self *getGardensHandler) Execute() {
    sessionID, err := self.GetString("sessionID");
    if nil != err {
        self.Response(err);
        return;
    }

    _, err = checkSession(sessionID, self.Debug);
    if nil != err {
        self.Response(err);
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != e {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    var gardens []model.Garden;
    // result := db.Where(&model.Garden{UserID: user.UserID}).Find(&gardens);
    result := db.Where(&model.Garden{}).Find(&gardens);
    if result.RecordNotFound() {
        // no garden setting
    } else if result.Error != nil {
        self.Response(api.E_SERVER_ERROR.Apply("%s", result.Error.Error()));
    }

    self.Response(&getGardensResponse {
        Error: api.E_OK.Replace("GetGardens ok"),
        Gardens: gardens,
    })
}

