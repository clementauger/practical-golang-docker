
EXPOSE ?= 8081

init:
	export GO111MODULE=on
	go mod init
	go mod tidy
reload:
	grelot
build:
	(cd .. && docker build -t pgr-consumer --target=dev -f producer/Dockerfile .)
run:
	docker run -it \
		-v $(shell pwd):/go/src/github.com/clementauger/practical-golang-docker/consumer\
	 	-p ${EXPOSE}:${EXPOSE} pgr-consumer
