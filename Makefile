migrate.init:
	migrate create -ext sql -dir migration/ -seq
# Example: migrate create -ext sql -dir migration/ -seq init_schema (MIGRATION NAME)
migrate.up:
	migrate -path migration/ -database "mysql://root@tcp(127.0.0.1:3306)/intikom-test?charset=utf8&parseTime=True&loc=Local" -verbose up
migrate.down:
	migrate -path migration/ -database "mysql://root@tcp(127.0.0.1:3306)/intikom-test?charset=utf8&parseTime=True&loc=Local" -verbose down
migrate.fix:
	migrate -path migration/ -database "mysql://root@tcp(127.0.0.1:3306)/intikom-test?charset=utf8&parseTime=True&loc=Local" force
# migrate -path migration/ -database "mysql://root@tcp(127.0.0.1:3306)/intikom-test?charset=utf8&parseTime=True&loc=Local" force 00002 (VERSION)

