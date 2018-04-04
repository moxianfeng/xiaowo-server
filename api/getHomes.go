package api

import (
    "api"

    "family/api/model"
)

type getHomesHandler struct {
    api.BaseRequest
}

type getHomesResponse struct {
    *api.Error
    Homes []model.Home `json:"homes"`
}

func init() {
    api.RegisterExecutor("GetHomes", func (br api.BaseRequest) api.Request {
        return &getHomesHandler{BaseRequest: br}
    });
}

func (self *getHomesHandler) Execute() {
    sessionId, err := self.GetString("sessionId");
    if nil != err {
        self.Response(err);
        return;
    }

    user, err := checkSession(sessionId, self.Debug);
    if nil != err {
        self.Response(err);
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != e {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    var homes []model.Home;
    result := db.Where(&model.Home{UserID: user.UserID}).Find(&homes);
    if result.RecordNotFound() {
        // no home setting
    } else if result.Error != nil {
        self.Response(api.E_SERVER_ERROR.Apply("%s", result.Error.Error()));
    }

    self.Response(&getHomesResponse {
        Error: api.E_OK.Replace("GetHomes ok"),
        Homes: homes,
    })
}

