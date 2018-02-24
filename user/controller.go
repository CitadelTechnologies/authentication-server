package user

import(
    "ct-authentication-server/client"
    "ct-authentication-server/server"
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "html/template"
    "strconv"
)

func RegisterAction(w http.ResponseWriter, r *http.Request) {
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

func LoginFormAction(w http.ResponseWriter, r *http.Request) {
    clientId := r.URL.Query().Get("clientId")
    if clientId == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(&struct{
            Message string `json:"message"`
        }{
            Message: "ID parameter is missing",
        })
        return
    }
    id, _ := strconv.ParseUint(clientId, 10, 16)
    service, err := client.GetClient(uint(id))
    if err != nil {
        panic(err)
    }
    t, err := template.ParseFiles(server.App.RootPath + "/templates/login.html")
    if err != nil {
        panic(err)
    }
    t.Execute(w, struct{
        ServiceId uint
        ServiceName string
    }{
        ServiceId: service.Id,
        ServiceName: service.Name,
    })
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
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
    serviceId, _ := strconv.ParseUint(data["service"], 10, 16)
    service, err := client.GetClient(uint(serviceId))
    if err != nil {
        panic(err)
    }
    user, err := Connect(service, data["username"], []byte(data["password"]))
    if err != nil {
        panic(err)
    }
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Location", string(service.RedirectUrl) + "?access_token=" + string(user.AccessToken))
    w.WriteHeader(302)
}
