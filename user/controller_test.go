package user

import(
    "bytes"
    "ct-authentication-server/server"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestRegister(t *testing.T) {
    data, _ := json.Marshal(map[string]string{
        "username": "Foo",
        "password": "Bar",
    })
    req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(data))
    response := executeRequest(req)

    if response.Code != http.StatusCreated {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusCreated)
    }
    var user User
    json.Unmarshal([]byte(response.Body.String()), &user)
    if user.Id != 1 {
        t.Errorf("ID was incorrect, got %d, want %d", user.Id, 1)
    }
    if user.Username != "Foo" {
        t.Errorf("Username was incorrect, got '%s', want '%s'", user.Username, "Foo")
    }
    if user.CreatedAt.Before(time.Now()) == false {
        t.Errorf("Creation date was incorrect, got '%s'", user.CreatedAt)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    server.App.Router.ServeHTTP(rr, req)

    return rr
}
