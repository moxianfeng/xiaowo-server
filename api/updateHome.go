package api

import (
    "api"

    "log"
    "xiaowo/api/model"
)

type updateHomeHandler struct {
    sessionCheckHandler
}

type updateHomeResponse struct {
    *api.Error
    Home model.Home `json:"home"`
}

func init() {
    api.RegisterExecutor("UpdateHome", func (br api.BaseRequest) api.Request {
        return &updateHomeHandler{sessionCheckHandler{BaseRequest: br}}
    });
}

func (self *updateHomeHandler) Execute() {
    _, db, success := self.PreExecute();
    if !success {
        return;
    }

    var newhome model.Home;
    err := self.GetStruct("home", &newhome);
    if nil != err {
        self.Response(err);
        return;
    }

    if len(newhome.HomeID) == 0 {
        self.Response(api.E_BAD_REQUEST.Apply("UpdateHome need homeID field"));
        return;
    }
    if len(newhome.Name) == 0 {
        self.Response(api.E_BAD_REQUEST.Apply("UpdateHome need name field"));
        return;
    }

    log.Print("UpdateHome: ", newhome);
    if e := db.Model(&model.Home{}).Where(&model.Home{HomeID: newhome.HomeID}).Update(newhome).Error; nil != e {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    }

    self.Response(&updateHomeResponse {
        Error: api.E_OK.Replace("UpdateHome ok"),
        Home: newhome,
    });
}
