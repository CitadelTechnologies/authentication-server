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
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
)
