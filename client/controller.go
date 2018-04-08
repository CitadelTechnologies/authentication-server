package client

import(
    "ct-authentication-server/server"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

func CreateClientAction(w http.ResponseWriter, r *http.Request) {
    defer server.CatchException(w)

    data := server.DecodeJsonRequest(r)
    client := CreateClient(data["name"], data["redirect_url"])

    server.SendJsonResponse(w, 201, client)
}

func AddDomainAction(w http.ResponseWriter, r *http.Request) {
    defer server.CatchException(w)

    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 16)
    client := GetClient(uint(id))
    data := server.DecodeJsonRequest(r)
    domain := AddDomainToClient(client, data["domain"])

    server.SendJsonResponse(w, 201, domain)
}

func GetClientAction(w http.ResponseWriter, r *http.Request) {
    defer server.CatchException(w)

    clientId, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 16)
    client := GetClient(uint(clientId))

    server.SendJsonResponse(w, 200, client)
}
