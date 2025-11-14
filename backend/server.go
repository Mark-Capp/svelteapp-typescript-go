package backend

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Server is the main program
func Server() {

	// read command-line flags
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", 8080, "Server port")
	docker := flag.Bool("docker", false, "Running in docker")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("./data/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&ListItem{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&ListItemFact{})

	// prepare service, http handler and server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default()) // All origins allowed by default

	// apis
	api := router.Group("/api")
	api.GET("/items", ListItems(db))
	api.POST("/items", AddItem(db))

	api.GET("/tags", GetTags(db))
	api.POST("/tags", AddTag(db))

	// serve static files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	router.NoRoute(func(c *gin.Context) { // fallback
		c.File("./build/index.html")
	})

	var serverPath string
	if *docker {
		serverPath = "0.0.0.0:8080"
		log.Println("Server started at http://localhost:8080 ...")
	} else {
		serverPath = fmt.Sprintf("%s:%d", *host, *port)
		log.Printf("Server started at http://%s ...\n", serverPath)
	}

	server := &http.Server{
		Addr:         serverPath,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
	log.Println("Server stopped. ")
}
