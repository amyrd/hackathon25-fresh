package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"golang.org/x/oauth2/instagram"
)

// In-memory store for MVP (use a real database later if needed)
var posts = []map[string]interface{}{}

const (
	username = "myUser"
	pass     = "myPass"
)

func main() {
	// this will create "router" thing, listens for HTTP reqs
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// bind the JSON payload we got from front to the struct
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// confirm (this is not secure)
		if credentials.Username == username && credentials.Password == pass {

			// we can make and return a StatusOK
			c.JSON(http.StatusOK, gin.H{
				"message": "login successful",
				"token":   "dummy-token",
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials, dawg they are hard coded"})
		}
	})

	// Endpoint to get all posts
	router.GET("/posts", func(c *gin.Context) {

		c.JSON(http.StatusOK, posts)
	})

	// Endpoint to create a new post
	router.POST("/posts", func(c *gin.Context) {
		var newPost map[string]interface{}
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, newPost)
		c.JSON(http.StatusCreated, newPost)
	})

	// Endpoint to fetch analytics data (dummy data for now)
	router.GET("/analytics", func(c *gin.Context) {
		analytics := gin.H{
			"totalPosts": len(posts),
			// Add more analytics fields as needed
		}
		c.JSON(http.StatusOK, analytics)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
