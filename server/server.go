package main

import (
	"context"
	"log"
	"os"
	"time"

	ginlogrus "github.com/Bose/go-gin-logrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/gql/gen"
	"github.com/kisinga/mzurihealth/logger"
)

//	cors "github.com/rs/cors/wrapper/gin"

const defaultPort = "4242"

var logFile *os.File

func main() {
	r := gin.Default()

	p := logger.InitLogging()
	// tell gin to use the middleware
	r.Use(p)

	//Create the database connection
	ctx := context.Background()

	db.ConnectDB(ctx, "test")

	//Dont fotget to close connection to db
	defer db.CloseSession(ctx)

	dt := time.Now()
	logPath := dt.Format("01-02-2006") + ".log.json"
	logfile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}
	// use the JSON formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(logfile)

	gin.SetMode(gin.DebugMode)

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	c := cors.Default()

	// r.Use(auth.Middleware())

	r.Use(c)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.DisableConsoleColor()

	r.Use(gin.Recovery()) // add Recovery middleware

	r.POST("/api", graphqlHandler())
	r.GET("/api", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
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
		ginlogrus.SetCtxLoggerHeader(c, "playground", "this is how you set new header level data")

		// logger := ginlogrus.GetCtxLogger(c) // will get a logger with the aggregate Logger set if it's enabled - handy if you've already set fields for the request
		// logger.Info("this will be aggregated into one write with the access log and will show up when the request is completed")

		// // add some new fields to the existing logger
		// logger = ginlogrus.SetCtxLogger(c, logger.WithFields(logrus.Fields{"comment": "this is an aggregated log entry with initial comment field"}))
		// logger.Debug("aggregated entry with new comment field")

		// // replace existing logger fields with new ones (notice it's logrus.WithFields())
		// logger = ginlogrus.SetCtxLogger(c, logrus.WithFields(logrus.Fields{"new-comment": "this is an aggregated log entry with reset 	 field"}))
		// logger.Error("aggregated error entry with new-comment field")

		// logrus.Info("this will NOT be aggregated and will be logged immediately")
		// span := newSpanFromContext(c, "sleep-span")
		// defer span.Finish() // this will get logged because tracing was setup with ginopentracing.WithEnableInfoLog(true)

		// go func() {
		// 	// need a NewBuffer for aggregate logging of this goroutine (since the req will be done long before this thing finishes)
		// 	// it will inherit header info from the existing request
		// 	buff := ginlogrus.NewBuffer(logger)
		// 	time.Sleep(1 * time.Second)
		// 	logger.Info("Hi from a goroutine completing after the request")
		// 	fmt.Printf(buff.String())
		// }()
		h.ServeHTTP(c.Writer, c.Request)
	}
}
