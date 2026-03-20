# USER (/users)
## GET /
### Headers
> auth
### Body
> `nil`
### Response
> [] {Id, Email, User}

## GET /:id
### Headers
> auth
### Body
> `nil`
### Response
> {Id, Email, User}

## POST /signup
### Headers
> `nil`
### Body
> {email, password, username}
### Response
> jwt

## POST /login
### Headers
> `nil`
### Body
> {email, password}
### Response
> jwt

## PATCH /:id
### Headers
> auth
### Body
> {user?}
### Response
> {Id, Email, User}

## PUT /pass/:id
### Headers
> auth
### Body
> {password}
### Response
> OK

## DELETE /:id
### Headers
> auth
### Body
> `nil`
### Response
> OK
