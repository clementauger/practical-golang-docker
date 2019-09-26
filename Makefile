run:
	# docker-compose up --build
	docker network create localdev --subnet 10.0.1.0/24 # for if you are using openvpn
	docker-compose up
