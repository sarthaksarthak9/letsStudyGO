package main

import (
	"net/http"
	"proj2/controller"
	internal "proj2/internal/database"
	"proj2/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// how to create router
	router := gin.Default()
	db := internal.InitDB()

	if db == nil {
		// error while connection to the database
		return
	}

	notesServices := &services.NotesService{}
	notesServices.InitService(db)

	// c has both req and res
	// c represents the current HTTP request and response.
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from gin",
			"new var": "done for part2",
		})
	})

	// how to pass param, how to retrive req from frontend and send req
	router.GET("/me/:id/:newId", func(c *gin.Context) {

		var id = c.Param("id")
		var newId = c.Param("newId")

		c.JSON(http.StatusOK, gin.H{
			"user_ID":    id,
			"user_newID": newId,
		})
	})

	// how to post
	router.POST("/me", func(c *gin.Context) {
		// email, password
		type MeRequest struct {
			// `json: "email"` the format from which I get the response from the frontend
			// Email string `json:"email"`
			// suppose binding is required
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var meRequest MeRequest

		// Bind the var ?? In Go, you cannot directly read JSON into variables unless you:
		// Read the body, Parse JSON, Map fields manually
		// Gin’s binding system does all of this for you.

		if err := c.BindJSON(&meRequest);

		// error handling
		err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		// It: Sets the HTTP status code, Converts Go data → JSON
		// Writes the JSON to the response body,  Sets Content-Type: application/json
		// So this line ends the request by sending a JSON response back to the client.
		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	// Put use for updation. It update the complete data.
	router.PUT("/me", func(c *gin.Context) {

		type MeRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"passoword"`
		}

		var meRequest MeRequest

		if err := c.BindJSON(meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"passowrd": meRequest.Password,
		})
	})

	// Patch it use to update a small segment of the data
	router.PATCH("/me", func(c *gin.Context) {

		type MeRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var meRequest MeRequest

		if err := c.Bind(meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	// we send a ID when we use delete
	router.DELETE("/me/:id", func(c *gin.Context) {
		var id = c.Param("id")

		c.JSON(200, gin.H{
			"id":      id,
			"message": "deleted",
		})
	})

	// Register the controller

	notesController := &controller.NotesController{}
	// we never send data directly from main. It is from services either models
	notesController.InitNotesControllerRoutes(router, *notesServices)

	router.Run(":8000")
}
