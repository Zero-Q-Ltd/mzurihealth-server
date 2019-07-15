package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	ginlogrus "github.com/Bose/go-gin-logrus"
	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	logger "github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/gql/gen"
)

//	cors "github.com/rs/cors/wrapper/gin"

const defaultPort = "4242"

var logFile *os.File

func main() {
	//We init the logger first because it's usued universally

	// setup tracing...
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}

	tracer, reporter, closer, err := ginopentracing.InitTracing(
		fmt.Sprintf("go-gin-logrus-example::%s", hostName), // service name for the traces
		"localhost:5775",                        // where to send the spans
		ginopentracing.WithEnableInfoLog(false)) // WithEnableLogInfo(false) will not log info on every span sent... if set to true it will log and they won't be aggregated
	if err != nil {
		panic("unable to init tracing")
	}
	defer closer.Close()
	defer reporter.Close()
	opentracing.SetGlobalTracer(tracer)

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
	r := gin.New() // don't use the Default(), since it comes with a logger
	gin.SetMode(gin.DebugMode)

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	c := cors.Default()
	r.Use(c)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.DisableConsoleColor()

	p := ginopentracing.OpenTracer([]byte("api-request-"))
	r.Use(p)

	r.Use(gin.Recovery()) // add Recovery middleware
	useBanner := true
	useUTC := true

	r.Use(ginlogrus.WithTracing(logrus.StandardLogger(),
		useBanner,
		time.RFC3339,
		useUTC,
		"requestID",
		[]byte("uber-trace-id"), // where jaeger might have put the trace id
		[]byte("RequestID"),     // where the trace ID might already be populated in the headers
		ginlogrus.WithAggregateLogging(true)))

	r.POST("/api", graphqlHandler())
	r.GET("/api", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}

func initGin(logge *logger.Logger) {

	router := gin.New()
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

	router.POST("/api", graphqlHandler())
	router.GET("/api", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run(":" + port)
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	h := handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &gen.Resolver{}}))

	return func(c *gin.Context) {
		ginlogrus.SetCtxLoggerHeader(c, "new-header-index-name", "graphql")

		logger := ginlogrus.GetCtxLogger(c) // will get a logger with the aggregate Logger set if it's enabled - handy if you've already set fields for the request
		logger.Info("this will be aggregated into one write with the access log and will show up when the request is completed")

		// add some new fields to the existing logger
		logger = ginlogrus.SetCtxLogger(c, logger.WithFields(logrus.Fields{"comment": "this is an aggregated log entry with initial comment field"}))
		logger.Debug("aggregated entry with new comment field")

		// replace existing logger fields with new ones (notice it's logrus.WithFields())
		logger = ginlogrus.SetCtxLogger(c, logrus.WithFields(logrus.Fields{"new-comment": "this is an aggregated log entry with reset comment field"}))
		logger.Error("aggregated error entry with new-comment field")

		logrus.Info("this will NOT be aggregated and will be logged immediately")
		span := newSpanFromContext(c, "sleep-span")
		defer span.Finish() // this will get logged because tracing was setup with ginopentracing.WithEnableInfoLog(true)

		go func() {
			// need a NewBuffer for aggregate logging of this goroutine (since the req will be done long before this thing finishes)
			// it will inherit header info from the existing request
			buff := ginlogrus.NewBuffer(logger)
			time.Sleep(1 * time.Second)
			logger.Info("Hi from a goroutine completing after the request")
			fmt.Printf(buff.String())
		}()
		c.JSON(200, "Hello world!")
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func newSpanFromContext(c *gin.Context, operationName string) opentracing.Span {
	parentSpan, _ := c.Get("tracing-context")
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: c.Request.Method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: c.Request.URL.Path},
		opentracing.Tag{Key: "current-goroutines", Value: runtime.NumGoroutine()},
	}

	if parentSpan != nil {
		options = append(options, opentracing.ChildOf(parentSpan.(opentracing.Span).Context()))
	}

	return opentracing.StartSpan(operationName, options...)
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/api")

	return func(c *gin.Context) {
		ginlogrus.SetCtxLoggerHeader(c, "new-header-index-name", "this is how you set new header level data")

		logger := ginlogrus.GetCtxLogger(c) // will get a logger with the aggregate Logger set if it's enabled - handy if you've already set fields for the request
		logger.Info("this will be aggregated into one write with the access log and will show up when the request is completed")

		// add some new fields to the existing logger
		logger = ginlogrus.SetCtxLogger(c, logger.WithFields(logrus.Fields{"comment": "this is an aggregated log entry with initial comment field"}))
		logger.Debug("aggregated entry with new comment field")

		// replace existing logger fields with new ones (notice it's logrus.WithFields())
		logger = ginlogrus.SetCtxLogger(c, logrus.WithFields(logrus.Fields{"new-comment": "this is an aggregated log entry with reset comment field"}))
		logger.Error("aggregated error entry with new-comment field")

		logrus.Info("this will NOT be aggregated and will be logged immediately")
		span := newSpanFromContext(c, "sleep-span")
		defer span.Finish() // this will get logged because tracing was setup with ginopentracing.WithEnableInfoLog(true)

		go func() {
			// need a NewBuffer for aggregate logging of this goroutine (since the req will be done long before this thing finishes)
			// it will inherit header info from the existing request
			buff := ginlogrus.NewBuffer(logger)
			time.Sleep(1 * time.Second)
			logger.Info("Hi from a goroutine completing after the request")
			fmt.Printf(buff.String())
		}()
		c.JSON(200, "Hello world!")
		h.ServeHTTP(c.Writer, c.Request)
	}
}
