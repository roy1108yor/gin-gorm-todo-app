# Api go simple
It is a just simple REST API with Go using:
1. **Gin Framework**
2. **Gorm** 

## Installation & Run

1. **Initialize Go modules**
    ```bash
    go mod tidy
    ```

2. **Run the application**
    ```bash
    go run main.go
    ```

## DB Setup
1. Go to Config/Database.go
2. Update DBName, User, Password and Host, Port according to your database configuration

## API list

* `GET` : Get all todos
* `POST` : Create a todo
* `GET` : Get a todo
* `PUT` : Update a todo
* `DELETE` : Delete a todo

## Post Params
```
{
	"name": "Groceries shopping",
	"description": "Biscuits, Tea, Milk, Soap, Powder",
}
```
