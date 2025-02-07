# User API Spec

## Register User API

Endpoint :  POST /api/user/register

Request Body :

```json
{
  "full_name": "Indrawansyah",
  "username": "indra",
  "email": "indra2example.com",
  "password": "indra123"
}
```

Response Body Success :

```json
{
  "message": "success",
  "data": {
    "id": 1,
    "full_name": "Indrawansyah",
    "username": "indra",
    "email": "indra2example.com",
    "orders": null
  }
}
```

Response Body Error :

```json
{
  "errors" : "Username already registered"
}
```