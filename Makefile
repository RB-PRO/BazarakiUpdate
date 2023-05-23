all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/BazarakiUpdate.git

pull:
	git pull git@github.com:RB-PRO/BazarakiUpdate.git

push-car:
	export GOARCH=amd64
	export GOOS=linux
	export CGO_ENABLED=0
	go env GOOS GOARCH CGO_ENABLED
	go build -o main ./cmd/main/main.go
	scp main bd root root@185.154.192.111:go/