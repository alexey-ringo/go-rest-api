## HTTP REST API SERVER
    http rest api server, powered by Go

## commands
    go mod init github.com/alexey-ringo/go-rest-api

    go get github.com/BurntSushi/toml
    go get -u github.com/gorilla/mux
    go get github.com/stretchr/testify
    go get github.com/lib/pq
    go get github.com/go-ozzo/ozzo-validation

    make; ./apiserver
    migrate create -ext sql -dir migrations create_users
    migrate -path migrations -database "postgres://localhost/restapi_dev?user=default&password=secret&sslmode=disable" up
    migrate -path migrations -database "postgres://localhost/restapi_test?user=default&password=secret&sslmode=disable" up

    Вывод в консоль структуры
    fmt.Printf("%+v\n", srtuct)
    fmt.Printf("%T - %v\n", struct, struct)
    и строки
	fmt.Println("123")