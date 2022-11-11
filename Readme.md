# goauth
## This is a library to use jwt-token (web)

### How to install?
```
$ go get github.com/DUNA-E-Commmerce/goauth
```

### How to create a token?
First at all we need to instance and set the secretKey we will use to create the token and then to retrieve the encrypted data to.

```
    // Create an instance of goauth to use in your code.
    goAuth := NewGoAuth("my-secret-key")
    
    /* 
        Create a map with the data you want to add 
        on your new token
    */
    data := make(map[string]interface{})
    data[<key>] = <value>
    
    // Create the new token without expiration
    token, err := goAuth.CreateToken(data, 0)
    if err != nil {
        fmt.Errorf("Something is wrong", err);
    }

    // Example new token with expiration (1 hour for example)
    token, err := goAuth.CreateToken(data, 1 * time.Hour)
    if err != nil {
        fmt.Errorf("Something is wrong", err);
    }
    
    fmt.Println(token) // This print your new token
```

### How to decrypt a token?
First at all we need to instance and set the secretKey we will use to create the token and then to retrieve the encrypted data to.

```
    // Use the instance with the same secret key that you have created the specify token.
    goAuth := NewGoAuth("my-secret-key")
    
    claim, err := goAuth.DecryptToken(token)
    if err != nil {
        fmt.Errorf("Something is wrong", err);
    }

    print(claim[<key>]) // This will return the same map you specify before on the token creation step.
```

You can implement this library on a middleware and manage there the authentification/authorization with this library.
