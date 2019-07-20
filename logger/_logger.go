package logger

import (
	"fmt"
	"os"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

var tracer opentracing.Tracer
var reporter jaeger.Reporter

func log(data struct{}) {

}

func InitLogging() gin.HandlerFunc {
	// setup tracing...
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}

	_tracer, _reporter, closer, err := ginopentracing.InitTracing(fmt.Sprintf("go-gin-opentracing-example::%s", hostName),
		"localhost:5775",
		ginopentracing.WithEnableInfoLog(true))
	if err != nil {
		panic("unable to init tracing")
	}

	tracer = _tracer
	reporter = _reporter
	defer closer.Close()
	defer reporter.Close()
	opentracing.SetGlobalTracer(tracer)

	// create the middleware
	p := ginopentracing.OpenTracer([]byte("api-request-"))
	return p
}
