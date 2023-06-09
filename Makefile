all: run

run:
	go run cmd/main/main.go bd token -848128665

push:
	git push git@github.com:RB-PRO/BazarakiUpdate.git

pull:
	git pull git@github.com:RB-PRO/BazarakiUpdate.git

push-car:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH CGO_ENABLED
	go build cmd/main/main.go
	scp main bd token root root@185.154.192.111:go/