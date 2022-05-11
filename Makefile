# note: call scripts from /scripts

.PHONY: dbup test dbclean

dbup:
	docker run   -d --rm -p 5432:5432 --name postgres-local \
	-e POSTGRES_USER=api_user -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=api_db \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v "$(shell pwd)/pgdata/local":/var/lib/postgresql/data/pgdata -v "$(shell pwd)/configs/sql/initdb":/docker-entrypoint-initdb.d \
	postgres \
	&& docker ps

dbclean:
	docker stop postgres-local; sudo rm -r pgdata/

test:
	docker run -d --rm -p 5430:5432 --name postgres-test \
    -e POSTGRES_USER=tst -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=tst \
	-v "$(shell pwd)/configs/sql/initdb":/docker-entrypoint-initdb.d \
    postgres && sleep 5 && go test ./... ;  \
 	docker stop postgres-test;


fill:
	docker run -it --rm --name postgres-fill \
    -e POSTGRES_USER=tst -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=tst \
	-v "$(shell pwd)/configs/sql/initdb":/docker-entrypoint-initdb.d \
    postgres