#.PHONY: deps clean build
.PHONY: clean build

#deps:
#	go get -u ./hello-world

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
