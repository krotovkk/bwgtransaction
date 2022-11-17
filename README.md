# bwgtransaction

Запуск приложения через cmd/app/main.go

Запуск миграций через команду:

 goose -v -dir ./migrations postgres "host=localhost port=80 user=user password=password dbname=bwgtransaction sslmode=disable" up
