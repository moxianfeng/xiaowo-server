package main

import (
    "log"
    "os"
    "net/http"
    "api"
    "flag"
    _ "xiaowo/api"
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

    if debug {
        log.Print("Run in debug mode");
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "10011"
    }

    http.HandleFunc("/call", call);
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
