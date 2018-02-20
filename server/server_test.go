package server

import(
    "fmt"
    //"database/sql"
    //"github.com/gorilla/mux"
    //_ "github.com/go-sql-driver/mysql"
    //"github.com/go-redis/redis"
    "testing"
)

func TestInitialize(t *testing.T) {
    fmt.Println(App.Router.Get("/register"))
}
