package api

import (
    "time"

    "api"
    "family/api/model"
)

const (
    CODE_PERIOD = 70 * time.Second;
)


type bindPhoneHandler struct {
    api.BaseRequest
}

func (self *bindPhoneHandler) Execute() {
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

    phoneNumber, err := self.GetString("phoneNumber");
    if nil != err {
        self.Response(err);
        return;
    }

    verifyCode, err := self.GetString("verifyCode");
    if nil != err {
        self.Response(err);
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != e {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    // check verify code
    var code model.VerifyCode;
    result := db.Where(&model.VerifyCode{SessionKey: sessionId, UserID: user.UserID, PhoneNumber: phoneNumber, VerifyCode: verifyCode}).First(&code);
    if result.RecordNotFound() {
        self.Response(api.E_INVALID_VERIFY_CODE);
        return;
    } else if nil != result.Error {
        self.Response(api.E_SERVER_ERROR.Apply(result.Error.Error()));
        return;
    }
    if code.UpdatedAt.Add(CODE_PERIOD).Before(time.Now()) {
        self.Response(api.E_INVALID_VERIFY_CODE);
        return;
    }

    // bind phone
    user.CellPhone = phoneNumber;
    if e = db.Model(&user).Update(&user).Error; nil != e {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    }

    self.Response(api.E_OK.Replace("BindPhone ok"));
}


func init() {
    api.RegisterExecutor("BindPhone", func(br api.BaseRequest) api.Request {
        return &bindPhoneHandler{BaseRequest: br};
    });
}
