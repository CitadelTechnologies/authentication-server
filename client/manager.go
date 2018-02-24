package client

import(
    "ct-authentication-server/security"
    "ct-authentication-server/server"
    "strings"
    "time"
)

func CreateClient(name, redirectUrl string) (*Client, error) {
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
        return nil, err
    }
    res, err := stmt.Exec(client.Name, client.Token, client.Secret, client.RedirectUrl, client.CreatedAt, client.UpdatedAt)
    if err != nil {
        return nil, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return nil, err
    }
    client.Id = uint(id)
    return &client, nil
}

func AddDomainToClient(client *Client, name string) (*Domain, error) {
    domain := Domain{
        Client: client,
        Name: name,
    }
    stmt, err := server.App.DB.Prepare("INSERT INTO client__domains(name, client_id) VALUES(?, ?)")
    if err != nil {
        return nil, err
    }
    if _, err = stmt.Exec(domain.Name, domain.Client.Id); err != nil {
        return nil, err
    }
    return &domain, nil
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

func GetClient(id uint) (*Client, error) {
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
        return nil, err
    }
    return &client, nil
}
