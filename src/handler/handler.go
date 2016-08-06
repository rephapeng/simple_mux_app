package handler

import (
    "fmt"
    "net/http"
    //"encoding/json"
    //"github.com/gorilla/mux"
    //"github.com/julienschmidt/httprouter"
    "config"
)
//get configuration data
var globalConfiguration = config.Generate()


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Simple Pattern project with gorila mux")
}
