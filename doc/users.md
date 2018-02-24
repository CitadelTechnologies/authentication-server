# Users

## Registration

### POST /register

#### Request

*Headers*
```
Content-Type: application/json
```

*Body*
```
{
    username: "John",
    password: "secret"
}
```

#### Response 201

*Headers*
```
Content-Type: application/json
```

*Body*
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

## Login

### GET /login

#### Request

*Query Parameters*
```
clientId=1 // your service client ID
```

#### Response 200

A HTML template with a connection form and the JS code to perform the POST call.

### POST /login

#### Request

*Headers*
```
Content-Type: application/json
```

*Body*
```
{
    "username": "Foo",
    "password": "Bar",
    "service": 1
}
```

#### Response 302

*Headers*
```
Location: http://my.service-redirection.url
Access-Control-Allow-Origin: http://my.service-redirection.domain
```
