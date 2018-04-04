
package api

import (
    "api"

    "log"
)

type setCurrentHomeHandler struct {
    sessionCheckHandler
}

func init() {
    api.RegisterExecutor("SetCurrentHome", func (br api.BaseRequest) api.Request {
        return &setCurrentHomeHandler{sessionCheckHandler{BaseRequest: br}}
    });
}

func (self *setCurrentHomeHandler) Execute() {
    user, db, success := self.PreExecute();
    if !success {
        return;
    }

    homeID, err := self.GetString("homeID");
    if nil != err {
        self.Response(err);
        return;
    }

    log.Print(user, db, homeID);

    self.Response(api.E_OK.Replace("SetCurrentHome ok"));
}

