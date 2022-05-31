# shopapi
A simple backend system for a shopping app

## Model
![image](https://user-images.githubusercontent.com/84314071/171119893-c8294090-fbf4-4b45-9556-424f6332f809.png)



## API
##### Expected request and response format
**--- User ---**<br/>
("/login", user.Login).Methods("GET")<br/>
**Sample:**<br/>
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
("/register", user.Register).Methods("POST")<br/>
**Sample:**
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
**--- Product ---**<br/>
("/products", product.GetAll).Methods("GET") => Get all products<br/>
("/inventory/{userId}", product.GetInventory).Methods("GET")=> Get cart/inventory of a user<br/>
**Sample:**<br/>
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

**--- Cart ---**<br/>
("/cart/{userId}",cart.GetCart).Methods("GET")<br/>
("/cart/{userId}",cart.AddItem).Methods("POST")<br/>
("/cart/{userId}",cart.DelItem).Methods("DELETE")<br/>
**--- Order ---**<br/>
("/order/{cartId}",order.NewOrder).Methods("POST")<br/>
