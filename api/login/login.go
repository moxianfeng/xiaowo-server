package login

import (
    "net/http"
    "encoding/json"
    "log"
    "family/api/util"
    "io"
)

func login(w http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        util.HttpErrorResponse(w, http.StatusBadRequest, "Expect post method");
        return;
    }

    decoder := json.NewDecoder(req.Body);
    var m map[string]interface{};
    err := decoder.Decode(&m);
    if nil != err {
        log.Print(err);
        util.HttpErrorResponse(w, http.StatusBadRequest, "Invalid json request");
        return;
    }

    io.WriteString(w, "OK");
    log.Print(m);
}

func init() {
    http.HandleFunc("/api/Login", login);
}
