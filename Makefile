#.PHONY: deps clean build
.PHONY: deps clean build

#deps:
#	go get -u ./home-state

clean: 
	rm -rf ./home-state/bin
	
build:
	GOOS=linux GOARCH=amd64 go build -o home-state/bin ./home-state
