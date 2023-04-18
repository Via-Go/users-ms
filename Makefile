gen_cert:
	@echo "Generating certificates"
	make cert -C src

gen: gen_cert

gen_and_run:
	make gen
	docker compose -f docker-compose-dev.yaml up --build

clean_and_run:
	docker compose down
	docker compose -f docker-compose-dev.yaml up --build

run:
	docker compose -f docker-compose-dev.yaml up --build