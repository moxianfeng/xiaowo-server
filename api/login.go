package api

import (
    "api"
    "fmt"
    "net/http"
    "encoding/json"
    "family/api/model"
    "github.com/google/uuid"
)

type loginHandler struct {
    api.BaseRequest
}

type userStruct struct {
    UserId string `json:"userId"`;
    NickName string `json:"nickName"`;
    AvatarUrl string `json:"avatarUrl"`;
    Gender int `json:"gender"`;
    SessionId string `json:"sessionId"`;
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
    code, err := self.GetString("code");
    if nil != err {
        self.Response(err);
        return;
    }
    // https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
    url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", WEIXIN_URL, APPID, SECRET, code);
    resp, e := http.Get(url);
    if nil != e {
        self.Response(api.E_NETWORK.Apply("%s", e.Error()));
        return;
    }

    defer resp.Body.Close()
    decoder := json.NewDecoder(resp.Body);

    var wxResp map[string]interface{};
    e = decoder.Decode(&wxResp);
    if nil != e {
        self.Response(api.E_BAD_JSON.Apply("%s", e.Error()));
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != err {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    unionid, ok := wxResp["unionid"];
    if !ok {
        unionid = wxResp["openid"];
    }
    unionidStr := unionid.(string);

    sessionkeyStr := wxResp["session_key"].(string)

    var user model.User;
    result := db.Where(&model.User{UnionID: unionidStr}).First(&user);
    if result.RecordNotFound() {
        err = self.GetStruct("userInfo", &user);
        if nil != err {
            self.Response(err);
            return;
        }

        user.UserID = uuid.New().String();
        user.UnionID = unionidStr;
        user.CellPhone = "";
        user.SessionKey = sessionkeyStr;

        if e = db.Create(&user).Error; nil != e {
            self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
            return;
        }
    } else if result.Error != nil {
        self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
        return;
    } else {
        var updateUser model.User;
        updateUser.ID = user.ID;
        updateUser.SessionKey = sessionkeyStr;
        db.Debug().Model(&user).Update("session_key", sessionkeyStr);
    }

    res := &loginResponse{
        Response: api.Response{ErrCode: 0, ErrMsg: "Login: ok"},
        User: userStruct {
            UserId : user.UserID,
            NickName : user.NickName,
            AvatarUrl : user.AvatarUrl,
            Gender : user.Gender,
            SessionId : user.SessionKey,
            Valid : true,
        },
    };

    // check CellPhone
    if len(user.CellPhone) == 0 {
        res.User.Valid = false;
    }
    self.Response(res);
}

func newLoginHandler(br api.BaseRequest) api.Request {
    return &loginHandler{BaseRequest: br}
}

func init() {
    api.RegisterExecutor("Login", newLoginHandler);
}
