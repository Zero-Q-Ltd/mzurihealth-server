package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path"
// 	"path/filepath"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"

// 	"github.com/99designs/gqlgen/handler"
// 	"github.com/gin-contrib/cors"

// 	"github.com/kisinga/mzurihealth/config"
// 	"github.com/kisinga/mzurihealth/db"
// 	"github.com/kisinga/mzurihealth/gql/gen"
// 	"github.com/kisinga/mzurihealth/models"
// )

// //	cors "github.com/rs/cors/wrapper/gin"

// const defaultPort = "4242"

// func main() {

// 	//Create the first context
// 	ctx := context.Background()

// 	//Create a connection to db
// 	dberr := db.ConnectDB(ctx, "test")
// 	if dberr != nil {
// 		initError("ConnectDB", dberr)
// 	}
// 	//Dont fotget to close connection to db
// 	defer db.CloseSession(ctx)

// 	//Read the config first
// 	hosp, configerr := config.ReadFile()
// 	//Only create a new hospital if a config file does not exist
// 	if configerr != nil {
// 		fmt.Println("Error reading file")
// 		clicommands(ctx)

// 		// if configerr.Error() == "no such file or directory" {
// 		// 	fmt.Println("File doesnt exist")
// 		// 	clicommands(ctx)
// 		// }
// 	}
// 	decodeerr := db.QueryDocument(ctx, "", "hospitals", bson.D{{"_id", hosp.ID}}).Decode(&hospital)
// 	if decodeerr != nil {
// 		initError("Decode Hospital After Reading from DB", decodeerr)
// 	}
// 	e := echo.New()

// 	//
// 	// p := tracer.InitTracing()
// 	// tell gin to use the middleware
// 	e.Use(p)

// 	// Add CORS middleware around every request
// 	// See https://github.com/rs/cors for full option listing
// 	c := cors.Default()

// 	// r.Use(auth.Middleware())

// 	e.Use(c)

// 	//This can vary according to client, in case they have something else running
// 	//on port 4242
// 	port := os.Getenv("MZURIHEALTH PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	e.Use(gin.Recovery()) // add Recovery middleware

// 	e.GET("/app", func(c *gin.Context) {
// 		dir, file := path.Split(c.Request.RequestURI)
// 		ext := filepath.Ext(file)
// 		if file == "" || ext == "" {
// 			c.File("./public/index.html")
// 		} else {
// 			// strings.Split(file, "?")
// 			c.File("./public" + path.Join(dir, file))
// 		}

// 	})
// 	e.POST("/api", graphqlHandler())
// 	e.GET("/api", graphqlHandler())
// 	e.GET("/", playgroundHandler())
// 	_ = e.Run(":" + port)
// }
import (
	"context"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kisinga/mzurihealth/config"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/gql/gen"
	"github.com/kisinga/mzurihealth/models"
	"github.com/kisinga/mzurihealth/tracer"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/urfave/cli"
)

const (
	banner = `MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWK00XWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkocclxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkollc:::lxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkollllc:::::lxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkollllllc:::::::lxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMWKkollllllllc:::::::::lxKWMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMWKkollllllllllc:::::::::::lxKWMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMWKkollllllllllllc:::::::::::::lxKWMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMWKkolllllllllllloxxxoc::::::::::::lxKWMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMWXkollllllllllllox0XNNXOoc::::::::::::lxKWMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMWKkollllllllllllox0XWWWWWNXOoc::::::::::::lxKWMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMWKkollllllllllllox0XWWMMMMMWWNXOoc::::::::::::lkXWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMWKkollllllllllllox0XWWWMMMMMMMMWWNXOoc::::::::::::lxKWMMMMMMMMMMMMM
MMMMMMMMMMMWKkollllllllllllox0XWWWMMMMMMMMMMMMWWNXOoc::::::::::::lxKWMMMMMMMMMMM
MMMMMMMMMWKkollllllllllllldOKXNNNNNNNNNNNNNNNNNNNXX0xl:::::::::::::lkKWMMMMMMMMM
MMMMMMMWXkolllllllllllllc:cccllllllllloollcccccccc:::;,;:::::::::::::lkXWMMMMMMM
MMMMMWXkolllllllllllllc;,,,,,;;;;,,;:ccc::;,'''''''''''',;:::::::::::::lkXWMMMMM
MMMWXkolllllllllllllc;,,,,,,,;;,;;:clllc:::;;,'''''''''''',;;::::::::::::lkXWMMM
MMMNxlllllllllllllc;,,,,;;,,;,;;:clllllc:::::;;,'''''''''''',;;::::::::::::xNMMM
MMMN0dllllllllllc:;,,,,,;;;;;;:clllllllc:::::::;;,'''''''''''',;:::::::::co0NWMM
MMMWNXOdllllllc:;,,,;;;;;,;;:clllllllllc:::::::::;;,'''''''''''',;;::::coOXWWMMM
MMMMWWNXOdllc:;,,,;;,;;;;;:clllllllllllc:::::::::::;;,'''''''''''',;:coOXNWWMMMM
MMMMMMWWNXOxdddddddddddlcclllllllllllllcc::::::::::::;;:loooooooooooxOXNWWMMMMMM
MMMMMMMMMWWNNNNNNNNNNKOdllllllllllllldkOkxl:::::::::::::okXNNNNNNNNNNWWWMMMMMMMM
MMMMMMMMMMMMMMMMMMWXOdllllllllllllldOKNNNX0xl:::::::::::::oONWMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMXxllllllllllllldOXNWWWWWWNKxl::::::::::::cxNMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMWXkollllllllllclkXWWMMMMMMWWKd:;::::::::::cxXWMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMWWX0xollllllc:;,:dKWMMMMMMN0l,'',;::::::cdOXWWMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMWWNXOdollc:;,,,,;:dKWWWN0o;'''''',;::cdOXNWWMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMWWNXOdc:;;:;:::::lxOkd:,,,,,,''',:oOXNWWMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMWWNXOdlllllllllllc:;;;;;;;;;;:okXNWWMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMWWNXOdollllllllc::;::;;;:coOXNWWMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMWWNXOxollllllc:::::::cdOXWWWMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNXOxollllc:::::cdOXWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNX0xollc:::cdOXWWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNXOxlc:cdOXWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNX0kk0XWWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWWWWWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
`
	timeoutInSeconds = "620s"
)

var hospital models.Hospital

const defaultPort = "4242"

// version set by LDFLAGS
var version string

func start(ctx context.Context, root string, port string, redirectHttps bool, logFormat string) {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	p := tracer.InitTracing()
	// tell gin to use the middleware
	r.Use(p)

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	c := cors.Default()

	// r.Use(auth.Middleware())

	r.Use(c)

	//Read the config first
	hosp, configerr := config.ReadFile()
	//Only create a new hospital if a config file does not exist
	if configerr != nil {
		fmt.Println("Error reading file")

	}

	decodeerr := db.QueryDocument(ctx, "", "hospitals", bson.D{{"_id", hosp.ID}}).Decode(&hospital)
	if decodeerr != nil {
		initError("Decode Hospital After Reading from DB", decodeerr)
	}
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: logFormat,
	// }))

	// if redirectHttps {
	// 	e.Pre(middleware.HTTPSRedirect())
	// }
	//This can vary according to client, in case they have something else running
	//on port 424
	r.Use(gin.Recovery()) // add Recovery middleware
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	r.POST("/api", graphqlHandler())
	r.GET("/api", graphqlHandler())
	r.GET("/playground", playgroundHandler())
	_ = r.Run(":" + port)
}

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

	app := cli.NewApp()
	app.Name = "Mzurihealth"
	app.Usage = "Mzurihealth server"
	app.Version = version

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the SPA server",
			Action: func(c *cli.Context) error {
				start(ctx, c.String("dir"), c.String("port"), c.Bool("https-redirect"), c.String("log-format")+"\n")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "dir, d",
					Value:  "./dist",
					Usage:  "Directory from which to serve files.",
					EnvVar: "MZURIHEALTH_DIR",
				},
				cli.StringFlag{
					Name:   "port, p",
					Value:  "4242",
					Usage:  "Server port.",
					EnvVar: "MZURIHEALTH_PORT",
				},
				cli.BoolFlag{
					Name:   "https-redirect",
					Usage:  "Use to force http to redirect to https",
					EnvVar: "MZURIHEALTH_HTTPS_REDIRECT",
				},
				cli.StringFlag{
					Name: "log-format",
					Value: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
						`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
						`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
						`"bytes_out":${bytes_out}}`,
					Usage:  "Specify the log format",
					EnvVar: "MZURIHEALTH_LOG_FORMAT",
				},
			},
		},
		{
			Name:  "new",
			Usage: "Create a new hospital",
			Action: func(c *cli.Context) error {
				createhospital(c.String("name"))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: "",
					Usage: "Name of the new hospital",
				},
			},
		},
	}

	// start()
	app.Run(os.Args)
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
func createhospital(hospitalname string) {
	/**
	*Make changnes before marshalling to take advantage of linter and avoid runtime errors
	**/
	//Read the config first
	_, configerr := config.ReadFile()
	//Only create a new hospital if a config file does not exist
	if configerr == nil {
		fmt.Println("Config File already exists")
		initError("Create Hospital", nil)
	}
	var ctx = context.Background()
	hospital.Name = hospitalname
	var newhosp bson.M
	b, _ := bson.Marshal(hospital)
	bson.Unmarshal([]byte(b), &newhosp)

	res, err := db.InserDocument(ctx, "", "hospitals", "", newhosp)
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
	var temp models.Hospital
	var decodeerr = result.Decode(&temp)
	if decodeerr != nil {
		initError("Decode Hospital", decodeerr)
	}
	hospital = temp
	createerr := config.Create(hospital)
	if createerr != nil {
		initError("Error creating File", createerr)
	}
	fmt.Printf("Success creating hospital: %q\n", hospitalname)
}

func initError(function string, e error) {
	fmt.Println(function + " Init Error:")
	fmt.Print(e)
	panic("Init Error")
}
