package client

import(
    "time"
)

type(
    Client struct {
        Id uint `json:"id"`
        Name string `json:"name"`
        Token []byte `json:"-"`
        Secret []byte `json:"-"`
        RedirectUrl string `json:"redirect_url"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
    Domain struct {
        Name string `json:"name"`
        Client *Client `json:"client"`
    }
)
