package main

import(
    "os"
    "ct-authentication-server/client"
    "ct-authentication-server/server"
    "ct-authentication-server/user"
)

func main() {
    server.App = server.Server{}
    server.App.Initialize(
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_DBNAME"),
        os.Getenv("REDIS_HOST"),
        os.Getenv("REDIS_PORT"),
        os.Getenv("REDIS_PASSWORD"),
        os.Getenv("GOPATH") + "/src/ct-authentication-server",
        os.Getenv("SSO_ORIGIN"),
    )
    initializeRouter(&server.App)
    server.App.Run()
}

func initializeRouter(s *server.Server) {
    s.Router.HandleFunc("/clients", client.CreateClientAction).Methods("POST")
    s.Router.HandleFunc("/clients/{id}", client.GetClientAction).Methods("GET")
    s.Router.HandleFunc("/clients/{id}/domains", client.AddDomainAction).Methods("POST")
    s.Router.HandleFunc("/register", user.RegisterAction).Methods("POST")
    s.Router.HandleFunc("/login", user.LoginFormAction).Methods("GET")
    s.Router.HandleFunc("/login", user.LoginAction).Methods("POST")
    s.Router.HandleFunc("/users/{access_token}", user.GetUserAction).Methods("GET")
}
