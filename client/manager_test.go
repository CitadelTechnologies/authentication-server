package client

import(
    "testing"
)

func TestCreateClient(t *testing.T) {
    name := "chat_client"
    redirectUrl := "http://local.la-citadelle.net"

    client := CreateClient(name, redirectUrl)

    if client.Name != name {
        t.Errorf("Name was incorrect, got '%s', want '%s'", client.Name, name)
    }
    if len(client.Token) != 64 {
        t.Errorf("Token was incorrect, got %d bytes, want %d", len(client.Token), 64)
    }
    if len(client.Secret) != 128 {
        t.Errorf("Secret was incorrect, got %d bytes, want %d", len(client.Secret), 128)
    }
    if client.RedirectUrl != redirectUrl {
        t.Errorf("Redirect URL was incorrect, got '%s', want '%s'", client.RedirectUrl, redirectUrl)
    }
    if client.Id != 2 {
        t.Errorf("ID was incorrect, got %d, want 2", client.Id)
    }
}
