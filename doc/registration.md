# Registration

### POST /register

#### Request

**Headers**
```
Content-Type: application/json
```

**Body**
```
{
    username: "John",
    password: "secret"
}
```

#### Response 201

**Headers**
```
Content-Type: application/json
```

**Body**
```
{
    id: 1,
    username: "John",
    access_token: null,
    refresh_token: null,
    created_at: "2018-01-01 11:00:00",
    last_connected_at: null
}
```
