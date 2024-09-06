package main

import (
	"context"
	"database/sql"
	"fmt"
	"gin-postgres/controllers"
	dbCon "gin-postgres/db/sqlc"
	"gin-postgres/routes"
	"gin-postgres/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.ContactController
	ContactRoutes     routes.ContactRoutes
)

func init() {
	ctx = context.TODO()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Couldn't load config %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Couldn't load database %v", err)
	}

	db = dbCon.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	ContactController = *controllers.NewContactController(db, ctx)
	ContactRoutes = routes.NewRouteContact(ContactController)

	server = gin.Default()
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("Couldn't load config %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "The API endpoint is working fine",
		})
	})

	// load contact routes
	ContactRoutes.ContactRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL),
		})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))

}
