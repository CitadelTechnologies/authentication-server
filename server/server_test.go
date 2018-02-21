package server

import(
    "testing"
)

func TestInitialize(t *testing.T) {
    if App.Router == nil {
        t.Errorf("Router was not initialized")
    }
    if App.DB == nil {
        t.Errorf("Database connection was not initialized")
    }
    if App.Redis == nil {
        t.Errorf("Redis connection was not initialized")
    }
}
