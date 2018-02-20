package user

import(
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var body []byte
    var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 4096)); err != nil {
        panic(err)
    }
	if err = r.Body.Close(); err != nil {
        panic(err)
    }
    var data map[string]string
    if err = json.Unmarshal(body, &data); err != nil {
        w.WriteHeader(400)
        return
    }
    user, err := CreateUser(data["username"], []byte(data["password"]))
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    if err = json.NewEncoder(w).Encode(&user); err != nil {
        panic(err)
    }
}
