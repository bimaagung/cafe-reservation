## Cafe Reservation
Project learning go, this concept is case simulation table resevation and order menu in cafe   

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