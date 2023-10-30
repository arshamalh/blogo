# Sample file is the default value, but shold be overriden like `make run File=data/another.json
version=0.0.1
name=kari

help: # Generate list of targets with descriptions
	@grep '^.*\:\s#.*' Makefile | sed 's/\(.*\) # \(.*\)/\1 \2/' | column -t -s ":"

build: # build the docker image with specified name and version
	docker build -t ${name}:${version} .

dev: # run the application using go run command and is used for development
	# You should first install swag using: go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init
	go run main.go serve
