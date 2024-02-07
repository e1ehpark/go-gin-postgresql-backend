run:
	go run main.go

install:
	go get -u github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres
	go get -u github.com/e1ehpark/go-gin-postgresql-backend/src/middlewares
	go get -u github.com/e1ehpark/go-gin-postgresql-backend/src/models
	go get -u github.com/e1ehpark/go-gin-postgresql-backend/src/routes
	go get -u github.com/e1ehpark/go-gin-postgresql-backend/src/utils

