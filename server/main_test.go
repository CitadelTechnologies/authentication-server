package server

import(
    "os"
    "github.com/mattes/migrate"
    "github.com/mattes/migrate/database/mysql"
    _ "github.com/mattes/migrate/source/file"
    "testing"
)

var migration *migrate.Migrate

func TestMain(m *testing.M) {
    App = Server{}
    App.Initialize(
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_TEST_DBNAME"),
        os.Getenv("REDIS_HOST"),
        os.Getenv("REDIS_PORT"),
        os.Getenv("REDIS_PASSWORD"),
        os.Getenv("GOPATH") + "/src/ct-authentication-server",
    )
    initDatabase()
    code := m.Run()
    clearDatabase()
    os.Exit(code)
}

func initDatabase() {
    driver, err := mysql.WithInstance(App.DB, &mysql.Config{})
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
