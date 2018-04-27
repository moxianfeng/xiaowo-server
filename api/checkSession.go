package api

import (
    "api"
    "time"
)

type checkSessionHandler struct {
    api.BaseRequest
}

type checkSessionResponse struct {
    api.Response
    SessionID string `json:"sessionID"`;
    SessionExpireTo time.Time `json:"sessionExpireTo"`;
}

func init() {
    api.RegisterExecutor("CheckSession", func (br api.BaseRequest) api.Request {
        return &checkSessionHandler{BaseRequest: br}
    });
}

func (self *checkSessionHandler) Execute() {
    sessionID, err := self.GetString("sessionID");
    if nil != err {
        self.Response(err);
        return;
    }

    user, err := checkSession(sessionID, self.Debug);
    if nil != err {
        self.Response(err);
        return;
    }

    res := &checkSessionResponse {
        Response: api.Response{ErrCode: 0, ErrMsg: "CheckSession: ok"},
        SessionID: sessionID,
        SessionExpireTo: user.UpdatedAt.Add(VALID_SESSION),
    }

    self.Response(res);
}

