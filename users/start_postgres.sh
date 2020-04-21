export DB_USERNAME=postgres
export DB_PASSWORD=test
export DB_HOST=localhost
export DB_SCHEMA=users

docker run --name test_postgis -e POSTGRES_PASSWORD=$DB_PASSWORD -e POSTGRES_DB=$DB_SCHEMA -p 5432:5432 -d postgis/postgis

