PROJECTNAME:=boxsize-service

build:
	go build

test:
	go test

docker-build :
		docker build -t $(PROJECTNAME) .

docker-run :
		docker run --publish 3000:3000 --name $(PROJECTNAME) --rm $(PROJECTNAME)

run :
	go run main.go handler.go

fmt :
	go fmt ./...

help :
	@cat "Readme.md"
