migrate-latest:

		migrate -database 'mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST})/${MYSQL_DBNAME}' -source file://migrations up

migrate-rollback:

		migrate -database 'mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST})/${MYSQL_DBNAME}' -source file://migrations down
