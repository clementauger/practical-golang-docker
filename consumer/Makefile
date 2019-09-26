
EXPOSE ?= 8081

init:
	export GO111MODULE=on
	go mod init
	go mod tidy
gobuild:
	go build -o producer .
reload:
	reflex -r '\.go$$' -s -- sh -c 'go build -o consumer . && ./consumer'
build:
	(cd .. && docker build -t pgr-producer --target=dev -f producer/Dockerfile .)
run:
	docker run -it \
		-v $(shell pwd):/gow/src/github.com/clementauger/practical-golang-docker/producer\
	 	-p ${EXPOSE}:${EXPOSE} pgr-producer
