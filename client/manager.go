package client

import(
    "ct-authentication-server/security"
    "ct-authentication-server/server"
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
