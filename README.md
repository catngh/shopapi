# shopapi
A simple backend system for a shopping app
## How to run
`make init` to setup container and migrate db
`make run` to start server

## Table of content
- [Model](https://github.com/BerIincat/shopapi#model)
- [API](https://github.com/BerIincat/shopapi#api)
- [Diagrams](https://github.com/BerIincat/shopapi#diagrams)

## Model
![image](https://user-images.githubusercontent.com/84314071/171119893-c8294090-fbf4-4b45-9556-424f6332f809.png)



## API
### Expected request and response format
#### User<br/>
User login: ```GET /login```
```
Request body
{
  email: "address@domain,
  password: "password"
}
Response body
{
  userId: "1",
  email: "address@domain",
  role: "buyer"
}
```

Register user: ```POST /register```
```
Request body
{
  email: "address@domain,
  password: "password"
  role: "vendor"
}
Response body
{
  userId: "2",
  email: "address@domain",
  role: "vendor"
}
```
#### Product
Get all products: ```GET /products```<br/>
Get cart/inventory of a user: ```GET /products/{userId}```

```
Response body
{
  [
    {
      productId: "1",
      name: "abc",
      price: 12.5
    },
    {
      productId: "2",
      name: "bca",
      price: 10
    },
    ...
  ]
}
```

#### Cart
Get cart info of a user: ```GET /cart/{userId}```
```
Response body
{
  cartId: "1",
  userId: "2",
  items:
  [
    {
      productId: "1",
      name: "abc",
      price: 12.5
    },
    {
      productId: "2",
      name: "bca",
      price: 10
    },
    ...
  ]
}
```

Add item to cart: ```POST /cart/{userId}```
```
Request body
{
  productId: "1"
}
```

Delete item in cart: ```DELETE /cart/{userId}```
```
Request body
{
  productId: "2",
}
```

#### Order
Create order with existing cart: ```POST /order/userid```
```
Response body
{
  orderId: "32",
  cartId: "1",
  subTotal: "2000",
  timeCreated: 2022-04-23T18:25
}
```

## Diagrams
#### Login user
<img width="412" alt="Login" src="https://user-images.githubusercontent.com/84314071/171130586-5f1a6898-3f89-46e3-aa00-a395013de447.png"><br/>

#### Register user
<img width="412" alt="Login" src="https://user-images.githubusercontent.com/84314071/171130267-31cb488a-edb9-4a53-a2dd-02dc088ac879.png"><br/>

#### Browse products
<img width="412" alt="Login" src="https://user-images.githubusercontent.com/84314071/171128621-ed4bced2-38d2-45cb-b5ba-db7033cb8a1c.jpg"><br/>

#### View inventory
<img width="412" alt="Login" src="https://user-images.githubusercontent.com/84314071/171128690-d7d0e43a-3e96-44cd-be69-34bb6efa7247.jpg"><br/>

#### Checkout
<img width="412" alt="Login" src="https://user-images.githubusercontent.com/84314071/171128769-ce038948-408e-465c-83a2-e85a54e36a01.jpg"><br/>


