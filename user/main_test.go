package user

import(
    "os"
    "ct-authentication-server/server"
    "github.com/mattes/migrate"
    "github.com/mattes/migrate/database/mysql"
    _ "github.com/mattes/migrate/source/file"
    "testing"
)

var migration *migrate.Migrate

func TestMain(m *testing.M) {
    server.App = server.Server{}
    server.App.Initialize(
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_TEST_DBNAME"),
        os.Getenv("REDIS_HOST"),
        os.Getenv("REDIS_PORT"),
        os.Getenv("REDIS_PASSWORD"),
        os.Getenv("GOPATH") + "/src/ct-authentication-server",
        os.Getenv("SSO_ORIGIN"),
    )
    initializeRouter(&server.App)
    initializeDatabase()
    code := m.Run()
    clearDatabase()
    os.Exit(code)
}

func initializeRouter(s *server.Server) {
    s.Router.HandleFunc("/register", RegisterAction).Methods("POST")
    s.Router.HandleFunc("/login", LoginFormAction).Methods("GET")
    s.Router.HandleFunc("/login", LoginAction).Methods("POST")
}

func initializeDatabase() {
    driver, err := mysql.WithInstance(server.App.DB, &mysql.Config{})
    if err != nil {
        panic(err)
    }
    migration, err = migrate.NewWithDatabaseInstance(
        "file://../migrations",
        "mysql",
        driver,
    )
    if err != nil {
        panic(err)
    }
    migration.Up()
}

func clearDatabase() {
    migration.Down()
}
