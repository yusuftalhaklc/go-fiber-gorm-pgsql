# Go Fiber Example API

Go Fiber Example API is a RESTful API built using the Fiber web framework.

## Installation

1. Clone the repository:
  ```shell
  git clone https://github.com/yusuftalhaklc/go-fiber-gorm-pgsql.git
  ```

2. Navigate to the project directory:
  ```shell
  cd go-fiber-gorm-pgsql
  ```
3. Install the dependencies:
  ```shell
  go mod tidy
  ```
4. Start the API server:
  ```shell
  go run main.go
  ```

## API Endpoints
### Get All Users

- **Endpoint:** `/api/user`
- **Method:** `GET`
- **Description:** Get all users.

### Get Single User

- **Endpoint:** `/api/user/:id`
- **Method:** `GET`
- **Description:** Get a single user by ID.

### Create User

- **Endpoint:** `/api/user`
- **Method:** `POST`
- **Description:** Create a new user.

### Update User

- **Endpoint:** `/api/user/:id`
- **Method:** `PUT`
- **Description:** Update an existing user by ID.

### Delete User

- **Endpoint:** `/api/user/:id`
- **Method:** `DELETE`
- **Description:** Delete a user by ID.

## Request Body and Response Examples

### Create User
- Request Body
```json
{
    "username":"username",
    "email": "username@example.com",
    "password":"password123"
}
```
- Response
```json
{
    "data": {
        "CreatedAt": "2023-06-26T11:33:44.674928+03:00",
        "UpdatedAt": "2023-06-26T11:33:44.674928+03:00",
        "DeletedAt": null,
        "ID": "f907c118-b2df-4920-9f04-6d13c96b355c",
        "username": "username",
        "email": "username@example.com",
        "password": "password123"
    },
    "message": "User has created",
    "status": "success"
}
```
### Get All User
- Request
```http
GET /api/users
```
- Response
 ```json
{
    "data": [
        {
            "CreatedAt": "2023-06-26T11:33:44.674928+03:00",
            "UpdatedAt": "2023-06-26T11:33:44.674928+03:00",
            "DeletedAt": null,
            "ID": "f907c118-b2df-4920-9f04-6d13c96b355c",
            "username": "username",
            "email": "username@example.com",
            "password": "password123"
        }
    ],
    "message": "Users Found",
    "status": "sucess"
}
```
### Get Single User
- Request
```http
GET /api/user/f907c118-b2df-4920-9f04-6d13c96b355c
```
- Response
```json
{
    "data": {
        "CreatedAt": "2023-06-26T11:33:44.674928+03:00",
        "UpdatedAt": "2023-06-26T11:33:44.674928+03:00",
        "DeletedAt": null,
        "ID": "f907c118-b2df-4920-9f04-6d13c96b355c",
        "username": "username",
        "email": "username@example.com",
        "password": "password123"
    },
    "message": "User Found",
    "status": "success"
}
```
### Update User
- Request
```http
PUT /api/user/f907c118-b2df-4920-9f04-6d13c96b355c
```

```json
{
    "username" : "yusuftalhaklc"
}
```

- Response
```json
{
    "data": {
        "CreatedAt": "2023-06-26T11:33:44.674928+03:00",
        "UpdatedAt": "2023-06-26T11:39:07.7508237+03:00",
        "DeletedAt": null,
        "ID": "f907c118-b2df-4920-9f04-6d13c96b355c",
        "username": "yusuftalhaklc",
        "email": "username@example.com",
        "password": "password123"
    },
    "message": "users Found",
    "status": "success"
}
```
### Delete User
- Request
```http
DELETE /api/user/f907c118-b2df-4920-9f04-6d13c96b355c
```
- Response
```json
{
    "message": "User deleted",
    "status": "success"
}
```
