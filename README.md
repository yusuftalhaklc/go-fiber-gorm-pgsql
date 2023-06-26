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

**It is currently running on localhost port 8080.** [Postman Collection](https://red-shuttle-655108.postman.co/workspace/go-fiber-gorm-pgsq~0fdd766c-93af-42a4-812b-f426fd8a91e0/collection/27159195-1c71ffe2-1f00-4958-990e-6d0988fa3c4e?action=share&creator=27159195)

## API Endpoints
### Get All Users

- **Endpoint:** `/api/user`
- **Method:** `GET`
- **Description:** Get all users.

### Get Single User

- **Endpoint:** `/api/user/:id`
- **Method:** `GET`
- **Description:** Get a single user by ID.

### Login 

- **Endpoint:** `/api/user/login`
- **Method:** `POST`
- **Description:** Login.


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
    "data": [
        {
            "CreatedAt": "2023-06-26T11:46:46.46791+03:00",
            "UpdatedAt": "2023-06-26T11:46:46.46791+03:00",
            "DeletedAt": null,
            "ID": "f907c118-b2df-4920-9f04-6d13c96b355c",
            "username": "username",
            "email": "username@example.com",
            "password": "password123"
        },
        {
            "CreatedAt": "2023-06-26T11:46:52.784047+03:00",
            "UpdatedAt": "2023-06-26T11:46:52.784047+03:00",
            "DeletedAt": null,
            "ID": "2e723e10-0a12-42ee-ac46-a9e268ebb66e",
            "username": "username2",
            "email": "username2@example.com",
            "password": "password456"
        }
    ],
    "message": "Users Found",
    "status": "sucess"
}
```

### Login
- Request
```http
POST /api/user/login
```

```json
{
    "username":"username",
    "password":"password123"
}
```

- Response
```json
{
    "message": "Login successfully",
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
