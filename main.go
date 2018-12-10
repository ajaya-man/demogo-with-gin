package main

import  (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	
	//Set the router as the default one provided by Gin
	router := gin.Default()
	
	//Process the templats at the start so that they don't have to be rloaded
	//from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use and inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the context to render a template
		c.HTML (
			//Set the HTTP status to 200 (OK)
			http.StatusOK ,
			// USe the index.html template
			"index.html",
			//Pass the data that the page uses ( in this casr, 'title')
			gin.H{
				"title" : "Home Page",
			},
		) 

	})

	// Start serving the application
	router.Run()


}