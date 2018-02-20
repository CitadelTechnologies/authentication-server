package user

import(
    "time"
)

type(
    User struct {
        Id uint `json:"id"`
        Username string `json:"username"`
        Password []byte `json:"-"`
        AccessToken []byte `json:"access_token"`
        RefreshToken []byte `json:"refresh_token"`
        CreatedAt time.Time `json:"created_at"`
        LastConnectedAt time.Time `json:"last_connected_at"`
    }
)
