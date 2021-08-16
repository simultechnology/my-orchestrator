
.PHONY: run_postgres
run_postgres:
	docker run -it -p 5439:5432 --name postgres -e POSTGRES_USER=cube -e POSTGRES_PASSWORD=secret postgres:13