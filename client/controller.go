package client

import(
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

func CreateClientAction(w http.ResponseWriter, r *http.Request) {
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
    client, err := CreateClient(data["name"], data["redirect_url"])
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    if err = json.NewEncoder(w).Encode(&client); err != nil {
        panic(err)
    }
}

func AddDomainAction(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
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
    id, _ := strconv.ParseUint(vars["id"], 10, 16)
    client, err := GetClient(uint(id))
    if err != nil {
        panic(err)
    }
    domain, err := AddDomainToClient(client, data["domain"])
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    if err = json.NewEncoder(w).Encode(&domain); err != nil {
        panic(err)
    }
}

func GetClientAction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    clientId, _ := strconv.ParseUint(vars["id"], 10, 16)
    client, err := GetClient(uint(clientId))

    w.Header().Set("Content-Type", "application/json")
    if err = json.NewEncoder(w).Encode(&client); err != nil {
        panic(err)
    }
}
