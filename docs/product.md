# Product API Spec

## Add Product API

Endpoint :  POST /api/product/create

Request Body :

```json
{
  "name": "Apple 12",
  "price": 10000000,
  "description": "A phone",
  "category": "phone"
}
```

Response Body Success :

```json
{
  "message": "success",
  "data": {
    "id": 4,
    "name": "Apple 12",
    "description": "",
    "price": 10000000,
    "category": "phone"
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to create product"
}
```

## Get Products API

Endpoint :  GET /api/product

Request Body :

```json
{}
```

Response Body Success :

```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "Apple 13 Pro",
      "description": "",
      "price": 15000000,
      "category": "phone",
      "inventory": {
        "id": 1,
        "product_id": 1,
        "quantity": 25,
        "location": "Karawang"
      }
    },
    {
      "id": 2,
      "name": "Apple 15 Pro",
      "description": "",
      "price": 23000000,
      "category": "phone",
      "inventory": {
        "id": 2,
        "product_id": 2,
        "quantity": 10,
        "location": "Karawang"
      }
    },
    {
      "id": 3,
      "name": "Apple 15",
      "description": "",
      "price": 18000000,
      "category": "phone",
      "inventory": {
        "id": 3,
        "product_id": 3,
        "quantity": 3,
        "location": "Karawang"
      }
    },
    {
      "id": 4,
      "name": "Apple 12",
      "description": "",
      "price": 10000000,
      "category": "phone",
      "inventory": {
        "id": 4,
        "product_id": 4,
        "quantity": 10,
        "location": "Karawang"
      }
    }
  ]
}
```

Response Body Error :

```json
{
  "errors" : "Failed to get products"
}
```

## Get Product Details API

Endpoint :  GET /api/product/:id

Request Body :

```json
{}
```

Response Body Success :

```json
{
  "message": "success",
  "data": {
    "id": 2,
    "name": "Apple 15 Pro",
    "description": "",
    "price": 23000000,
    "category": "phone",
    "inventory": {
      "id": 2,
      "product_id": 2,
      "quantity": 10,
      "location": "Karawang"
    }
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to get product"
}
```

## Update Product API

Endpoint :  PUT /api/product/:id

Request Body :

```json
{
  "name": "Apple 13 Pro",
  "price": 15000000,
  "category": "electronic"
}
```

Response Body Success :

```json
{
  "message": "success",
  "data": {
    "id": 1,
    "name": "Apple 13 Pro",
    "description": "",
    "price": 15000000,
    "category": "electronic",
    "inventory": {
      "id": 1,
      "product_id": 1,
      "quantity": 25,
      "location": "Karawang"
    }
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to update product"
}
```

## Get Inventory API

Endpoint :  GET /api/product/inventory/:id

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
    "product_id": 1,
    "quantity": 25,
    "location": "Karawang",
    "product": {
      "id": 1,
      "name": "Apple 13 Pro",
      "description": "",
      "price": 15000000,
      "category": "electronic"
    }
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to get inventory"
}
```

## Update Inventory API

Endpoint :  PUT /api/product/inventory

Request Body :

```json
{
  "product_id": 4,
  "quantity": 10,
  "location": "Karawang"
}
```

Response Body Success :

```json
{
  "message": "success",
  "data": {
    "id": 4,
    "product_id": 4,
    "quantity": 10,
    "location": "Karawang",
    "product": {
      "id": 4,
      "name": "Apple 12",
      "description": "",
      "price": 10000000,
      "category": "phone"
    }
  }
}
```

Response Body Error :

```json
{
  "errors" : "Failed to update inventory"
}
```

## Create New Order API

Endpoint :  POST /api/order

Request Body :

```json
{
  "user_id": 1,
  "order_items": [
    {
      "product_id": 1,
      "quantity": 7
    },
    {
      "product_id": 2,
      "quantity": 10
    },
    {
      "product_id": 3,
      "quantity": 2
    }
  ]
}
```

Response Body Success :

```json
{
  "message": "success"
}
```

Response Body Error :

```json
{
  "errors" : "Failed to create order"
}
```

## Get Order API

Endpoint :  GET /api/order/:id

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
}
```

Response Body Error :

```json
{
  "errors" : "Failed to get order"
}
```