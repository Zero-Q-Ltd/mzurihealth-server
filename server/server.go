package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/gql/gen"
)

//	cors "github.com/rs/cors/wrapper/gin"

const defaultPort = "4242"

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
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
	logPath := dt.Format("01-2006") + ".log"
	logfile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	//Create the database connection
	db.ConnectDB("test")

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

	router.POST("/api", graphqlHandler())
	router.GET("/api", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run(":" + port)

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &gen.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
