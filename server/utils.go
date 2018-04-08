package server

import (
    "ct-authentication-server/exception"
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
)

func DecodeJsonRequest(r *http.Request) map[string]string {
    var body []byte
    var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 4096)); err != nil {
        panic(exception.New(500, "Request body reader could not be opened"))
    }
	if err = r.Body.Close(); err != nil {
        panic(exception.New(500, "Request body reader could not be closed"))
    }
    var data map[string]string
    if err = json.Unmarshal(body, &data); err != nil {
        panic(exception.New(400, "Request body could not be read"))
    }
    return data
}

func SendJsonResponse(w http.ResponseWriter, code int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(&data); err != nil {
        panic(err)
    }
}

func CatchException(w http.ResponseWriter) {
    if r := recover(); r != nil {
        exception := r.(*exception.Exception)
        SendJsonResponse(w, exception.Code, exception)
    }
}
