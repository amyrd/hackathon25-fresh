package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// In-memory store for MVP (use a real database later if needed)
var posts = []map[string]interface{}{}

// Hard-coded Instagram credentials (for MVP only)
const (
	instagramUsername = "myInstagramUser"
	instagramPassword = "myPassword123"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	// Login endpoint
	router.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if credentials.Username == instagramUsername && credentials.Password == instagramPassword {
			c.JSON(http.StatusOK, gin.H{
				"message": "login successful",
				"token":   "dummy-token",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		}
	})

	// Endpoint to get all posts (for your own posts)
	router.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, posts)
	})

	// Endpoint to create a new post (for adding posts manually)
	router.POST("/posts", func(c *gin.Context) {
		var newPost map[string]interface{}
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, newPost)
		c.JSON(http.StatusCreated, newPost)
	})

	// Dummy analytics endpoint
	router.GET("/analytics", func(c *gin.Context) {
		analytics := gin.H{
			"totalPosts": len(posts),
		}
		c.JSON(http.StatusOK, analytics)
	})

	// New endpoint for fetching Instagram posts (dummy data for now)
	router.GET("/instagram-posts", func(c *gin.Context) {
		dummyInstagramPosts := []map[string]interface{}{
			{
				"id":      "1",
				"image":   "https://example.com/image1.jpg",
				"caption": "My first Instagram post",
			},
			{
				"id":      "2",
				"image":   "https://example.com/image2.jpg",
				"caption": "Another day, another post!",
			},
		}
		c.JSON(http.StatusOK, dummyInstagramPosts)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
