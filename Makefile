build:
	docker-compose down --remove-orphans
	docker-compose build
	docker-compose up -d
	docker exec -it shape dep ensure

restart:
	docker-compose restart
	
start:
	docker-compose down --remove-orphans
	docker-compose up -d

stop:
	docker-compose down --remove-orphans

shell:
	docker exec -it shape sh

logs:
	docker-compose logs -f --tail=100
