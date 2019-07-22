package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"

	"github.com/kisinga/mzurihealth/config"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/gql/gen"
	"github.com/kisinga/mzurihealth/models"
	"github.com/kisinga/mzurihealth/tracer"
)

//	cors "github.com/rs/cors/wrapper/gin"

const defaultPort = "4242"

var hospital *models.Hospital

func main() {

	//Create the first context
	ctx := context.Background()

	//Create a connection to db
	dberr := db.ConnectDB(ctx, "test")
	if dberr != nil {
		initError("ConnectDB", dberr)
	}
	//Dont fotget to close connection to db
	defer db.CloseSession(ctx)

	//Read the config first
	hosp, configerr := config.ReadFile()
	hospital = &hosp
	// fmt.Print(hosp)

	if configerr != nil {
		fmt.Println("Error reading file")
		clicommands(ctx)
	}

	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	//
	p := tracer.InitTracing()
	// tell gin to use the middleware
	r.Use(p)

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	c := cors.Default()

	// r.Use(auth.Middleware())

	r.Use(c)

	//This can vary according to client, in case they have something else running
	//on port 4242
	port := os.Getenv("MZURIHEALTH PORT")
	if port == "" {
		port = defaultPort
	}

	r.Use(gin.Recovery()) // add Recovery middleware

	r.POST("/api", graphqlHandler())
	r.GET("/api", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}

func clicommands(ctx context.Context, cmd ...string) {
	fmt.Println("Checking cli options...")
	newindicator := flag.Bool("new", false, "specify name to create a new hospital")
	name := flag.String("name", "", "specify name to create a new hospital")
	flag.Parse()

	if *newindicator {
		hospital.Name = *name
		res, err := db.InserDocument(ctx, "", "hospitals", hospital)
		if err != nil {
			initError("Create Hospital", err)
		}
		str, ok := res.InsertedID.(primitive.ObjectID)
		if ok {
			fmt.Printf("ID is: %q\n", str.Hex())
		} else {
			fmt.Printf("value is not a string\n")
		}
		objID, _ := primitive.ObjectIDFromHex(str.Hex())
		result := db.QueryDocument(ctx, "", "hospitals", bson.D{{"_id", objID}})
		decodeerr := result.Decode(hospital)
		if decodeerr != nil {
			initError("Decode Hospital", decodeerr)
		}
		config.Create(*hospital)
		return
	} else {
		initError("Cli commands", nil)
	}
}

func initError(function string, e error) {
	fmt.Println(function + " Init Error:")
	fmt.Print(e)
	panic("Init Error")
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
