# squid-tech
squid tech hackathon participation

# How to run
```
go mod tidy
go run github.com/steebchen/prisma-client-go db push --schema ./prisma
go cmd/server/main.go
```

# Run with docker
```
docker-compose build api
docker-compose build web
docker-compose up web
```
