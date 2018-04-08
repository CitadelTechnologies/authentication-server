package user

import(
    "ct-authentication-server/client"
    "ct-authentication-server/exception"
    "ct-authentication-server/server"
    "net/http"
    "html/template"
    "strconv"
)

func RegisterAction(w http.ResponseWriter, r *http.Request) {
    defer server.CatchException(w)

    data := server.DecodeJsonRequest(r)
    user := CreateUser(data["username"], []byte(data["password"]))

    server.SendJsonResponse(w, 201, user)
}

func LoginFormAction(w http.ResponseWriter, r *http.Request) {
    defer server.CatchException(w)

    clientId := r.URL.Query().Get("clientId")
    if clientId == "" {
        panic(exception.New(400, "ID parameter is missing"))
    }
    id, _ := strconv.ParseUint(clientId, 10, 16)
    service := client.GetClient(uint(id))
    t, err := template.ParseFiles(server.App.RootPath + "/templates/login.html")
    if err != nil {
        panic(exception.New(500, "Login Form template could not be parsed"))
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
    defer server.CatchException(w)

    data := server.DecodeJsonRequest(r)
    serviceId, _ := strconv.ParseUint(data["service"], 10, 16)
    service := client.GetClient(uint(serviceId))
    user := Connect(service, data["username"], []byte(data["password"]))

    if r.Host == server.App.Origin {
        w.Header().Set("Access-Control-Allow-Origin", client.GetAllowedDomains(service))
        w.Header().Set("Location", string(service.RedirectUrl) + "?access_token=" + string(user.AccessToken))
        w.WriteHeader(302)
        return
    }
    server.SendJsonResponse(w, 200, user)
}
