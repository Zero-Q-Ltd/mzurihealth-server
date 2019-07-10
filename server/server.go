package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kisinga/mzurihealth/gql/gen"
)

//	cors "github.com/rs/cors/wrapper/gin"

const defaultPort = "4242"

func main() {
	router := gin.Default()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	c := cors.Default()
	router.Use(c)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gin.DisableConsoleColor()

	//set up log file
	dt := time.Now()
	logPath := dt.Format("01-02-2006") + ".log"
	logfile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	// set request logging
	gin.DefaultWriter = logfile

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	gql := gin.WrapH(handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &gen.Resolver{}})))
	router.GET("/", gin.WrapH(handler.Playground("GraphQL playground", "/api")))
	router.GET("/api", gql)
	router.POST("/api", gql)
	router.Run(":" + port)

}
