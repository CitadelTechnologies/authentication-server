package main

import(
    "os"
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
    )
    initializeRouter(&server.App)
    server.App.Run()
}

func initializeRouter(s *server.Server) {
    s.Router.HandleFunc("/register", user.Register).Methods("POST")
}
