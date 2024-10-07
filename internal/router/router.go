package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	CreateUserRouter(router, db)

	return router
}