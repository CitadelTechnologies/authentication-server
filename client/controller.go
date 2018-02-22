package client

import(
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
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

func GetClientAction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    name := r.URL.Query().Get("name")

    if name == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(&struct{
            Message string `json:"message"`
        }{
            Message: "Name parameter is missing",
        })
        return
    }

    client, err := GetClient(name)

    w.Header().Set("Content-Type", "application/json")
    if err = json.NewEncoder(w).Encode(&client); err != nil {
        panic(err)
    }
}
