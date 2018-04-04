package api

import (
    "api"

    "family/api/model"
    "github.com/google/uuid"
)

type createHomeHandler struct {
    api.BaseRequest
}

type createHomeResponse struct {
    *api.Error
    Home model.Home `json:"home"`
}

func init() {
    api.RegisterExecutor("CreateHome", func (br api.BaseRequest) api.Request {
        return &createHomeHandler{BaseRequest: br}
    });
}

func (self *createHomeHandler) Execute() {
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

    name, err := self.GetString("name");
    if nil != err {
        self.Response(err);
        return;
    }

    country, err := self.GetString("country");
    if nil != err {
        country = user.Country;
    }

    province, err := self.GetString("province");
    if nil != err {
        country = user.Province;
    }

    city, err := self.GetString("city");
    if nil != err {
        country = user.City;
    }

    background, err := self.GetString("background");
    if nil != err {
        background = "";
    }

    home := model.Home {
        UserID: user.UserID,
        HomeID: uuid.New().String(),
        Name: name,
        Country: country,
        Province: province,
        City: city,
        BackgroundUrl: background,
    }

    if e = db.Create(&home).Error; nil != e {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    }

    self.Response(&createHomeResponse {
        Error: api.E_OK.Replace("GetHomes ok"),
        Home: home,
    });
}
