package api

import (
    "log"
    "math/rand"
    "time"
    "fmt"

    "api"
    "xiaowo/api/model"
)

const (
    CODE_LENGTH = 6
    GET_CODE_SPACE = 60 * time.Second;
)

var (
    CODE_LIBRARY = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

type verifyCodeHandler struct {
    api.BaseRequest
}

func (self *verifyCodeHandler) genCode() string {
    var ret string;
    rand.Seed(int64(time.Now().Nanosecond()));
    for i := 0;i < CODE_LENGTH;i++ {
        r := rand.Int();
        ret += fmt.Sprintf("%d", CODE_LIBRARY[r % len(CODE_LIBRARY)]);
    }
    return ret;
}

func (self *verifyCodeHandler) Execute() {
    phoneNumber, err := self.GetString("phoneNumber");
    if nil != err {
        self.Response(err);
        return;
    }

    db, e := GetDB(self.Debug);
    if nil != e {
        self.Response(api.E_MISS_DB_CONNECTION.Apply("%s", e.Error()));
        return;
    }

    verifyCode := self.genCode();
    newVerifyCode := model.VerifyCode{PhoneNumber: phoneNumber, VerifyCode: verifyCode};

    var verifyCodeObject model.VerifyCode;
    result := db.Where(&model.VerifyCode{PhoneNumber: phoneNumber}).First(&verifyCodeObject);
    if result.RecordNotFound() {
        if e = db.Create(&newVerifyCode).Error; nil != e {
            self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
            return;
        }
    } else {
        valid := verifyCodeObject.UpdatedAt.Add(GET_CODE_SPACE);
        if valid.After(time.Now()) {
            self.Response(api.E_BAD_REQUEST.Apply("GetVerifyCode too often"));
            return;
        } else {
            newVerifyCode.ID = verifyCodeObject.ID;
            if e = db.Model(&newVerifyCode).Update(&newVerifyCode).Error; nil != e {
                self.Response(api.E_SERVER_ERROR.Apply("%s", e.Error()));
                return;
            }
        }
    }

    self.Response(api.E_OK.Replace("GetVerifyCode ok"));

    log.Printf("GetVerifyCode return %s for %s", verifyCode, phoneNumber);
}

func newVerifyCodeHandler(br api.BaseRequest) api.Request {
    return &verifyCodeHandler{BaseRequest: br}
}

func init() {
    api.RegisterExecutor("GetVerifyCode", newVerifyCodeHandler);
}
