.PHONY: dev prod

dev:
	docker-compose -f deployments/docker-compose.yml up --build

prod:
	docker-compose -f deployments/docker-compose.yml --env-file deployments/.env.production up --build