package user

import(
    "bytes"
    "testing"
)

func TestCreateUser(t *testing.T) {
    username := "Toto"
    password := []byte("overki11P4$$W0RD")

    user, err := CreateUser(username, password)

    if err != nil {
        t.Errorf("User creation failed, got error '%s'", err.Error())
    }
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
