all: test

test:
	URL=https://api.ytctest.ru/ex/client/loadClientInfo \
	GUID=f8946569-64a6-4772-b713-e9e001144576-1670230075741 \
	go run ./cmd/main.go

battle:
	URL=https://api.ytimes.ru/ex/client/loadClientInfo \
	GUID=e06f5ed1-7c14-4cc9-84ae-044dfa14b746 \
	go run ./cmd/main.go

full: env container wait
	migrate -path ./schema -database 'postgres://postgres:Termit561908@0.0.0.0:5432/postgres?sslmode=disable' up
	
container:
	docker run --name=coffee-db -e POSTGRES_PASSWORD=Termit561908 -p 5432:5432 -d --rm postgres

wait:
	sh wait-for-postgres.sh

env:
	$(env) DB_PASSWORD=Termit561908

clear:
	docker stop coffee-db