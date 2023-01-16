## Cafe Reservation
Mini project which focus for development application cafe shop, which is the application is used for customer can table reservation and order menus in the cafe shop.
Another feature on application is can manage menus, customer and can see transactions which is running.

## 🛠️ Installation Steps

### Installation project

clone project
``` bson
git clone https://github.com/bimaagung/cafe-reservation.git
```

running app
```bson 
make run
```

<br>

## 📁 Project Structure

```
        +---.vscode
        +---bin
        +---database
        |   \---migration
        +---menu
        |   +---controller
        |   +---domain
        |   +---mocks
        |   +---repository
        |   |   +---minio
        |   |   +---postgres
        |   |   \---redis
        |   +---usecase
        |   \---validation
        +---middleware
        |   \---authorization
        +---pkg
        |   +---dotenv
        |   +---minio
        |   +---postgres
        |   \---redis
        +---user
        |   +---controller
        |   +---domain
        |   +---mocks
        |   +---repository
        |   \---usecase
        \---utils
            +---exception
            +---response
            \---token_manager
```

## 💻 Built with

- Golang
- Fiber
- Postgres
- Gorm
- JWT Auth
- Testify
- APM ELK


## Testing
- Test Coverage
```bson
    go test -v ./... -coverprofile profile.out
```
```bson
    go tool cover -func profile.out // per function
```
```bson
    go tool cover -html profile.out // view html
```

