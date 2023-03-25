# Go AVIA CRUD
Golang implementation of Avia managment JSON API service

## Startup

### Startup on host machine
1. Rename [suggested .env.example](.env.example) file to .env
2. Configure .env file under your environment
3. Start application using following command in terminal:
`go build | .\go-avia-crud.exe -migratedb -seeddb`

### Startup via Docker
1. Rename [suggested .env.example](.env.example) file to .env
2. **REMOVE** `LISTEN_PORT` variable
3. Add `DOCKER_HOST_APP_PORT` and `DOCKER_HOST_DB_PORT` for youself. These variables are describe ports through which you may access to containers' applications (see [docker compose file](docker-compose.yml))
4. Start application in container using following command
`docker-compose --env-file ./.env up -d`
- If you want fully restart application, removing containers and images, paste following command

  `docker-compose --env-file ./.env down; docker container prune; docker image rm avia_app:latest; docker volume rm avia_db; docker-compose --env-file ./.env up -d`

## Command-line options
* `-migratedb` - execute [structure.sql](dbo/structure.sql) script to initialize database structure
* `-seeddb` - execute [seeder.sql](dbo/seeder.sql) to fill database by default data