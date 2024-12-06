export MYSQL_URL='mysql://root:root@tcp(localhost:3306)/social_network'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up: 
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migarte-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down