K=$(shell basename `pwd`)

init:
	docker network create "$K" --subnet 10.0.1.0/24 # for if you are using openvpn
	docker-compose build

clean:
	docker-compose down || true
	docker network rm "$K" || true
	(docker images -a | grep "$K" | awk '{print $$3}' | grep -v '^$$' | xargs docker rmi) || true
	(docker container ls -a | grep "$K" | awk '{print $$3}' | grep -v '^$$' | xargs docker container rm -v) || true

run:
	# docker-compose up --build
	docker-compose up
