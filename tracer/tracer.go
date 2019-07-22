package tracer

import (
	"fmt"
	"os"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func InitTracing() gin.HandlerFunc {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}
	// initialize the global singleton for tracing...
	tracer, reporter, closer, err := ginopentracing.InitTracing(fmt.Sprintf("mzurhihealth::%s", hostName),
		"localhost:5775",
		ginopentracing.WithEnableInfoLog(true))
	if err != nil {
		panic("unable to init tracing")
	}
	fmt.Print(closer, reporter)
	// defer closer.Close()
	// defer reporter.Close()

	opentracing.SetGlobalTracer(tracer)

	// create the middleware
	p := ginopentracing.OpenTracer([]byte("api-request-"))
	return p
}
