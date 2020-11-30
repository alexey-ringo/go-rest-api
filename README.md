## HTTP REST API SERVER
    http rest api server, powered by Go

## commands
    go mod init github.com/alexey-ringo/go-rest-api

    make; ./apiserver
    migrate create -ext sql -dir migrations create_users
    migrate -path migrations -database "postgres://localhost/restapi_dev?user=default&password=secret&sslmode=disable" up
    migrate -path migrations -database "postgres://localhost/restapi_test?user=default&password=secret&sslmode=disable" up