run-dev:
	@$(MAKE) _run-dev -j 2

_run-dev: run-apiserver run-db

run-apiserver:
	scripts/run_dev.sh

run-db:
	docker-compose up
