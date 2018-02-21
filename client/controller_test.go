package client

import(
    "bytes"
    "ct-authentication-server/server"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestCreateClientAction(t *testing.T) {
    data, _ := json.Marshal(map[string]string{
        "name": "space_client",
    })
    req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(data))
    response := executeRequest(req)

    if response.Code != http.StatusCreated {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusCreated)
    }
    var client Client
    json.Unmarshal([]byte(response.Body.String()), &client)
    if client.Id != 1 {
        t.Errorf("ID was incorrect, got %d, want %d", client.Id, 1)
    }
    if client.Name != "space_client" {
        t.Errorf("Name was incorrect, got '%s', want '%s'", client.Name, "space_client")
    }
    if client.CreatedAt.Before(time.Now()) == false {
        t.Errorf("Creation date was incorrect, got '%s'", client.CreatedAt)
    }
    if client.UpdatedAt.Before(time.Now()) == false {
        t.Errorf("Creation date was incorrect, got '%s'", client.UpdatedAt)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    server.App.Router.ServeHTTP(rr, req)

    return rr
}
