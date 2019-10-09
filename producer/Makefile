
EXPOSE ?= 8080

init:
	export GO111MODULE=on
	go mod init
	go mod tidy
reload:
	grelot
build:
	(cd .. && docker build -t pgr-producer --target=dev -f producer/Dockerfile .)
run:
	docker run -it \
		-v $(shell pwd):/go/src/github.com/clementauger/practical-golang-docker/producer\
	 	-p ${EXPOSE}:${EXPOSE} pgr-producer
