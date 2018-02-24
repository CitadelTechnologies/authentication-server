package user

import(
    "bytes"
    "ct-authentication-server/client"
    "ct-authentication-server/server"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"
)

func TestRegisterAction(t *testing.T) {
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

func TestLoginFormAction(t *testing.T) {
    client.CreateClient("space_chat", "http://local.la-citadelle.net")

    req, _ := http.NewRequest("GET", "/login?clientId=1", bytes.NewBuffer([]byte("")))
    response := executeRequest(req)

    if response.Code != http.StatusOK {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusCreated)
    }
}

func TestLoginAction(t *testing.T) {
    redirectUrl := "http://local.la-citadelle.net"

    data, _ := json.Marshal(map[string]string{
        "username": "Foo",
        "password": "Bar",
        "service": "1",
    })
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
    response := executeRequest(req)

    if response.Code != http.StatusFound {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusFound)
    }
    location := response.Header().Get("location")
    if location == "" {
        t.Errorf("Location header was missing")
    }
    redirect := strings.Split(location, "?")
    token := strings.Split(redirect[1], "=")
    if redirect[0] != redirectUrl {
        t.Errorf("Redirection URL was incorrect, got '%s', want '%s'", redirect[0], redirectUrl)
    }
    if token[0] != "access_token" {
        t.Errorf("Access Token was missing")
    }
    if len([]byte(token[1])) != 64 {
        t.Errorf("Access Token length was incorrect, got %d bytes, want %d bytes", len([]byte(token[1])), 64)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    server.App.Router.ServeHTTP(rr, req)

    return rr
}
