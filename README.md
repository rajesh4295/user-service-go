# user-registration-service-go
- Simple user registration microservice written in Go with Postgres database.
- Uses clean architecture to accept requests from exposed API's and persist in database.
- Flow : API --> Controller --> Service --> Business --> Database

# Running the server
- Dev mode (default) `go run main.go`
- Prod mode `go run main.go --mode prod`

# Config
- All the server config is read from `/config/config.{mode}.json`
  


# List of API's

<table>
    <tr>
        <td>Method</td>
        <td>Endpoint</td>
        <td>Payload</td>
        <td>Response</td>
        <td>Description</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>/user/signup</td>
        <td>
     
```json
// New user in new org
{
    "user": {
        "name": "test5",
        "email": "test5",
        "password": "test"
    },
    "orgName": "test"
}

// New user in existing org
{
    "user": {
        "name": "test5",
        "email": "test5",
        "password": "test",
        "orgId": "b0cec438-0af6-4fa1-b1fb-db0e2f7e1bca"
    },
    "orgName": "test"
}
```
</td>
        <td>


```json
{
  "name": "test6",
  "email": "test6",
  "orgId": "b0cec438-0af6-4fa1-b1fb-db0e2f7e1bca",
  "id": "4df5f8f1-5cf0-41f8-a680-b841f17dc963",
  "createdAt": 1621062737438,
  "updatedAt": 1621062737438
}
```
</td>
        <td>New user signup</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>/user/login</td>
        <td>
     
```json
// Login with email and password
{
    "email": "test5",
    "password": "test"
}

// Login with name and password
{
    "name": "test6",
    "password": "test"
}
```
</td>
        <td>

```json
{
  "name": "test6",
  "email": "test6",
  "orgId": "b0cec438-0af6-4fa1-b1fb-db0e2f7e1bca",
  "id": "4df5f8f1-5cf0-41f8-a680-b841f17dc963",
  "createdAt": 1621062737438,
  "updatedAt": 1621062737438
}
```
</td>
        <td>User login</td>
    </tr>
    <tr>
        <td>GET</td>
        <td>/user/{id}</td>
        <td>
</td>
        <td>
        
        
```json
{
  "name": "test6",
  "email": "test6",
  "orgId": "b0cec438-0af6-4fa1-b1fb-db0e2f7e1bca",
  "id": "4df5f8f1-5cf0-41f8-a680-b841f17dc963",
  "createdAt": 1621062737438,
  "updatedAt": 1621062737438
}
```
   </td>
        <td>Get user by `id`</td>
    </tr>
    <tr>
        <td>GET</td>
        <td>/org/{id}</td>
        <td>
</td>
        <td>
    
```json
{
  "name": "test",
  "id": "b0cec438-0af6-4fa1-b1fb-db0e2f7e1bca",
  "createdAt": 1621054868431,
  "updatedAt": 1621054868431
}
```    
   </td>
        <td>Get org by `id`</td>
    </tr>
</table>


# Packages used
- `GORM` - Database ORM
- `VIPER` - Environment variables
- `MUX` - HTTP Router

# TODO
- Add server events logging (info, success, warn, error).
- Graceful server shutdown
- Add JWT authentication.
- Add logout functionality.
- Add user update and delete functionality.
- Add org update and delete functionality.

# License
MIT