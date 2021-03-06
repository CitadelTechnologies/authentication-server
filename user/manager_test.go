package user

import(
    "ct-authentication-server/server"
    "ct-authentication-server/client"
    "bytes"
    "testing"
)

func TestCreateUser(t *testing.T) {
    username := "Toto"
    password := []byte("overki11P4$$W0RD")

    user := CreateUser(username, password)

    if user.Username != username {
        t.Errorf("Username was incorrect, got '%s', want '%s'", user.Username, username)
    }
    if bytes.Equal(user.Password, password) {
        t.Errorf("Password was incorrect, got plain password, want encoded")
    }
    if user.Id != 2 {
        t.Errorf("ID was incorrect, got %d, want 2", user.Id)
    }
}

func testConnect(t *testing.T) {
    username := "Toto"
    password := []byte("overki11P4$$W0RD")
    service := &client.Client{
        Name: "space_client",
        RedirectUrl: "http://local.la-citadelle.net",
    }
    user := Connect(service, username, password)

    if user.AccessToken == nil {
        t.Errorf("Access Token was not set")
    }
    if len(user.AccessToken) != 64 {
        t.Errorf("Access Token length was incorrect, got %d, want %d", len(user.AccessToken), 64)
    }
    data, err := server.App.Redis.HGetAll("user." + string(user.AccessToken)).Result()
    if err != nil {
        t.Errorf("User session storage failed, got error '%s'", err.Error())
    }
    if data["username"] != username {
        t.Errorf("Session storage was incorrect, got '%s', want '%s'", data["username"], username)
    }
}
