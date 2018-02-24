package client

import(
    "bytes"
    "ct-authentication-server/server"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCreateClientAction(t *testing.T) {
    redirectUrl := "http://local.la-citadelle.net"
    data, _ := json.Marshal(map[string]string{
        "name": "space_client",
        "redirect_url": redirectUrl,
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
    if client.RedirectUrl != redirectUrl {
        t.Errorf("Redirect URL was incorrect, got '%s', want '%s'", client.RedirectUrl, redirectUrl)
    }
}

func TestAddDomainAction(t *testing.T) {
    dns := "http://local.la-citadelle.net"
    data, _ := json.Marshal(map[string]string{
        "domain": dns,
    })
    req, _ := http.NewRequest("POST", "/clients/1/domains", bytes.NewBuffer(data))
    response := executeRequest(req)

    if response.Code != http.StatusCreated {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusCreated)
    }
    var domain Domain
    json.Unmarshal([]byte(response.Body.String()), &domain)
    if domain.Client.Id != 1 {
        t.Errorf("Client ID was incorrect, got %d, want %d", domain.Client.Id, 1)
    }
    if domain.Name != dns {
        t.Errorf("Domain Name was incorrect, got '%s', want '%s'", domain.Name, dns)
    }
}

func TestGetClientAction(t *testing.T) {
    redirectUrl := "http://local.la-citadelle.net"
    req, _ := http.NewRequest("GET", "/clients/1", bytes.NewBuffer([]byte("")))
    response := executeRequest(req)

    if response.Code != http.StatusOK {
        t.Errorf("Status code was incorrect, got %d, want %d", response.Code, http.StatusOK)
    }
    var client Client
    json.Unmarshal([]byte(response.Body.String()), &client)
    if client.Id != 1 {
        t.Errorf("ID was incorrect, got %d, want %d", client.Id, 1)
    }
    if client.Name != "space_client" {
        t.Errorf("Name was incorrect, got '%s', want '%s'", client.Name, "space_client")
    }
    if client.RedirectUrl != redirectUrl {
        t.Errorf("Redirect URL was incorrect, got '%s', want '%s'", client.RedirectUrl, redirectUrl)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    server.App.Router.ServeHTTP(rr, req)

    return rr
}
