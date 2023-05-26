all: gen run

gen:
	@echo "Generating certificates"
	make cert -C src

clean:
	docker compose down

run:
	docker compose -f docker-compose-dev.yaml up --build
