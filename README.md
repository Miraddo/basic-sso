# Basic Single Sign-On (SSO)

This is a basic project to implement SSO with Go.

## Run
requirements

-   go 1.17
-   docker
    -   images
        -   Golang
        -   MariaDB
        -   Adminer
        -   Alpine

first clone the repository

```
$ git clone https://github.com/miraddo/basic-sso
```

after that go to inside the directory 

```
$ cd basic-sso
```

then run 

```
$ make tidy
$ make sso-build
$ make sso-up

# to close it press Ctrl+C
# then type
$ make sso-down
```


## List 

-   [x] Structure
-   [x] Configuration
-   [x] Database
-   [x] Implement Register Request
    -   [x] Check User Exist
    -   [x] Hash Password
    -   [x] Useful Comment
-   [ ] Redirect to new Service (Could be Service A)
-   [ ] Login with username and password
-   [ ] Send Token to Service (Could be Service A)