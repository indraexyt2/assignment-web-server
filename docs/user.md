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
  "errors" : "Failed to create user"
}
```

## Login User API

Endpoint :  POST /api/user/login

Request Body :

```json
{
  "username": "indra",
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
  "errors" : "Invalid username or password"
}
```

## Get User API

Endpoint :  GET /api/user/:id

Request Body :

```json
{}
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
    "orders": [
      {
        "id": 1,
        "user_id": 1,
        "order_items": [
          {
            "id": 1,
            "order_id": 1,
            "product_id": 1,
            "quantity": 7,
            "product": {
              "id": 1,
              "name": "Apple 13 Pro",
              "description": "",
              "price": 15000000,
              "category": "electronic"
            }
          },
          {
            "id": 2,
            "order_id": 1,
            "product_id": 2,
            "quantity": 10,
            "product": {
              "id": 2,
              "name": "Apple 15 Pro",
              "description": "",
              "price": 23000000,
              "category": "phone"
            }
          },
          {
            "id": 3,
            "order_id": 1,
            "product_id": 3,
            "quantity": 2,
            "product": {
              "id": 3,
              "name": "Apple 15",
              "description": "",
              "price": 18000000,
              "category": "phone"
            }
          }
        ]
      }
    ]
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to get user"
}
```