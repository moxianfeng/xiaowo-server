package api

import (
    "api"
    "time"
    "xiaowo/api/model"
    "github.com/google/uuid"
)

type loginHandler struct {
    api.BaseRequest
}

type userStruct struct {
    UserID string `json:"userID"`;
    NickName string `json:"nickName"`;
    AvatarUrl string `json:"avatarUrl"`;
    Gender int `json:"gender"`;
    SessionID string `json:"sessionID"`;
    SessionExpireTo time.Time `json:"sessionExpireTo"`;
    Valid bool `json:"valid"`;
}

type loginResponse struct {
    api.Response
    User userStruct `json:"user"`
}

const (
    WEIXIN_URL = "https://api.weixin.qq.com/sns/jscode2session";
    APPID = "wx70db103dbe29e2fa";
    SECRET = "fe148097cb3d729d2ca27cfc6bbda24c";
)

func (self *loginHandler) Execute() {
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
    if nil != err {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    var verify model.VerifyCode;
    result := db.Where(&model.VerifyCode{PhoneNumber: phoneNumber, VerifyCode: verifyCode}).First(&verify);
    if result.RecordNotFound() {
        if self.Debug && phoneNumber == "11111111111" && verifyCode == "111111" {
            // debug的时候可以直接输入
        } else {
            self.Response(api.E_INVALID_VERIFY_CODE);
            return;
        }
    } else if result.Error != nil {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    }

    // verify code check right
    var user model.User;
    result = db.Where(&model.User{PhoneNumber: phoneNumber}).First(&user);
    if result.RecordNotFound() {
        // create user;
        user.PhoneNumber = phoneNumber;
        user.SessionKey = uuid.New().String();
        user.UserID = uuid.New().String();
        if e = db.Create(&user).Error; nil != e {
            self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
            return;
        }
    } else if result.Error != nil {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    } else {
        // check session;
        now := time.Now();
        valid := user.UpdatedAt.Add(VALID_SESSION);
        if !valid.After(now) {
            // update session key
            newSessionKey := uuid.New().String();
            db.Model(&user).Update(model.User{SessionKey: newSessionKey});
            user.SessionKey = newSessionKey;
        } else {
            // update session expire time
            db.Model(&user).Update(model.User{SessionKey: user.SessionKey});
        }
    }

    res := &loginResponse{
        Response: api.Response{ErrCode: 0, ErrMsg: "Login: ok"},
        User: userStruct {
            UserID : user.UserID,
            NickName : user.NickName,
            AvatarUrl : user.AvatarUrl,
            Gender : user.Gender,
            SessionID : user.SessionKey,
            SessionExpireTo: user.UpdatedAt.Add(VALID_SESSION),
            Valid : true,
        },
    };

    self.Response(res);
}

func newLoginHandler(br api.BaseRequest) api.Request {
    return &loginHandler{BaseRequest: br}
}

func init() {
    api.RegisterExecutor("Login", newLoginHandler);
}
