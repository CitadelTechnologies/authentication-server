# Services

## Clients

### POST /clients

#### Request

*Headers*
```
Content-Type: application/json
```

*Body*
```
{
    "name": "my-service-client",
    "redirect_url": "https://your.service-redirect.url"
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
    "id": 1,
    "name": "your-service-client",
    "redirect_url": "https://your.service-redirection.url",
    "created_at": "2018-01-01 15:00:00",
    "updated_at": "2018-01-01 15:00:00"
}
```

### GET /clients/{id}

#### Request

*Parameters*
```
{id}: Your service client ID
```

#### Response 200

*Headers*
```
Content-Type: application/json
```

*Body*
```
{
    "id": 1,
    "name": "your-service-client",
    "redirect_url": "https://your.service-redirection.url",
    "created_at": "2018-01-01 15:00:00",
    "updated_at": "2018-01-01 15:00:00"
}
```

## Domains

### POST /clients/{id}/domains

#### Request

*Headers*
```
Content-Type: application/json
```

*Parameters*
```
{id}: Your service client ID
```

*Body*
```
{
    "domain": "https://your.service-redirection.domain"
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
    "name": "https://your.service-redirection.domain",
    "client": {
        "id": 1,
        "name": "your-service-client",
        "redirect_url": "https://your.service-redirection.url",
        "created_at": "2018-01-01 15:00:00",
        "updated_at": "2018-01-01 15:00:00"
    }
}
```

#### Response 400
