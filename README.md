## Cafe Reservation
Mini project which focus for development application cafe shop, which is the application is used for customer can table reservation and order menus in the cafe shop.
Another feature on application is can manage menus, customer and can see transactions which is running

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
