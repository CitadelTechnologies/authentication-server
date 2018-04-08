package client

import(
    "ct-authentication-server/exception"
    "ct-authentication-server/security"
    "ct-authentication-server/server"
    "strings"
    "time"
)

func CreateClient(name, redirectUrl string) *Client {
    client := Client{
        Name: name,
        Token: security.GenerateRandomToken(32),
        Secret: security.GenerateRandomToken(64),
        RedirectUrl: redirectUrl,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    stmt, err := server.App.DB.Prepare("INSERT INTO client__clients(name, token, secret, redirect_url, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
    if err != nil {
        panic(exception.New(500, "Client creation query could not be prepared"))
    }
    res, err := stmt.Exec(client.Name, client.Token, client.Secret, client.RedirectUrl, client.CreatedAt, client.UpdatedAt)
    if err != nil {
        panic(exception.New(500, "Client could not be created"))
    }
    id, err := res.LastInsertId()
    if err != nil {
        panic(exception.New(500, "Created Client ID could not be retrieved"))
    }
    client.Id = uint(id)
    return &client
}

func AddDomainToClient(client *Client, name string) *Domain {
    domain := Domain{
        Client: client,
        Name: name,
    }
    stmt, err := server.App.DB.Prepare("INSERT INTO client__domains(name, client_id) VALUES(?, ?)")
    if err != nil {
        panic(exception.New(500, "Client domain creation query could not be prepared"))
    }
    if _, err = stmt.Exec(domain.Name, domain.Client.Id); err != nil {
        panic(exception.New(500, "Client domain could not be created"))
    }
    return &domain
}

func GetAllowedDomains(client *Client) string {
    rows, err := server.App.DB.Query("SELECT name FROM client__domains WHERE client_id = ?", client.Id)
    if err != nil {
        return ""
    }
    defer rows.Close()
    for rows.Next() {
        var domain string
        rows.Scan(&domain)
        if strings.Contains(client.RedirectUrl, domain) {
            return domain
        }
    }
    return ""
}

func GetClient(id uint) *Client {
    client := Client{
        Id: id,
    }
    err := server.App.DB.QueryRow("SELECT name, token, secret, redirect_url, created_at, updated_at FROM client__clients WHERE id = ?", id).Scan(
        &client.Name,
        &client.Token,
        &client.Secret,
        &client.RedirectUrl,
        &client.CreatedAt,
        &client.UpdatedAt,
    )
    if err != nil {
        panic(exception.New(404, "Client not found"))
    }
    return &client
}
