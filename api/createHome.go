package api

import (
    "api"

    "xiaowo/api/model"
    "github.com/google/uuid"
)

type createHomeHandler struct {
    sessionCheckHandler
}

type createHomeResponse struct {
    *api.Error
    Home model.Home `json:"home"`
}

func init() {
    api.RegisterExecutor("CreateHome", func (br api.BaseRequest) api.Request {
        return &createHomeHandler{sessionCheckHandler{BaseRequest: br}}
    });
}

func (self *createHomeHandler) Execute() {
    user, db, success := self.PreExecute();
    if !success {
        return;
    }

    var newhome model.Home;
    err := self.GetStruct("home", &newhome);
    if nil != err {
        self.Response(err);
        return;
    }

    if len(newhome.Name) == 0 {
        self.Response(api.E_BAD_REQUEST.Apply("CreateHome need name field"));
        return;
    }

    if len(newhome.Country) == 0 {
        newhome.Country = user.Country;
    }
    if len(newhome.Province) == 0 {
        newhome.Province = user.Province;
    }
    if len(newhome.City) == 0 {
        newhome.City = user.City;
    }

    newhome.UserID = user.UserID;
    newhome.HomeID = uuid.New().String();

    if e := db.Create(&newhome).Error; nil != e {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    }

    self.Response(&createHomeResponse {
        Error: api.E_OK.Replace("CreateHome ok"),
        Home: newhome,
    });
}
