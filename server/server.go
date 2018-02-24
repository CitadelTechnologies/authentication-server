package server

import(
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-redis/redis"
    "net/http"
    "log"
)

type Server struct {
    Redis *redis.Client
    Router *mux.Router
    DB *sql.DB
    RootPath string
}

var App Server

func (s *Server) Initialize(dbHost, dbUser, dbPassword, dbName, redisHost, redisPort, redisPassword, rootPath string) {
    var err error
    if s.DB, err = sql.Open("mysql", dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?parseTime=true"); err != nil {
        panic(err)
    }
    s.Redis = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,  // use default DB
	})
    s.Router = mux.NewRouter()
    s.RootPath = rootPath
}

func (s *Server) Run() {
    log.Fatal(http.ListenAndServe(":80", s.Router))
}
