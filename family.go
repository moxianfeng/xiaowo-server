package main

import (
    "log"
    "os"
    "net/http"
    "api"
    "flag"
    _ "family/api"
)

var (
    debug bool = true;
)

func call(w http.ResponseWriter, req *http.Request) {
    r := api.HandleObject(w, req, debug);
    r.Execute();
}

func main() {
    flag.BoolVar(&debug, "debug", false, "debug flag");
    flag.Parse();

    port := os.Getenv("PORT")
    if port == "" {
        port = "10010"
    }

    http.HandleFunc("/call", call);
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
