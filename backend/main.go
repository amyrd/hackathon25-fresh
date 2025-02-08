package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var twitterClient = NewTwitterClient()

type SocialMediaPost struct {
	ID         string    `json:"id"`
	Platform   string    `json:"platform"`
	Caption    string    `json:"caption"`
	ImageURL   string    `json:"image_url"`
	Likes      int       `json:"likes"`
	Comments   int       `json:"comments"`
	Shares     int       `json:"shares"`
	PostedAt   time.Time `json:"posted_at"`
	ProfileURL string    `json:"profile_url"`
	Engagement float64   `json:"engagement"`
}

func main() {
	router := gin.Default()

	// CORS setup for local development
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Posts endpoints
	router.GET("/posts", handlePosts)
	router.GET("/twitter/posts", handleTwitterPosts)

	router.Run(":8080")
}

func handlePosts(c *gin.Context) {
	twitterPosts, err := twitterClient.FetchTweets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to fetch posts",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, append(instagramPosts, twitterPosts...))
}

func handleTwitterPosts(c *gin.Context) {
	tweets, err := twitterClient.FetchTweets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Twitter API failed",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tweets)
}

// Sample Instagram data
var instagramPosts = []SocialMediaPost{
	{
		ID:         "ig_1",
		Platform:   "instagram",
		Caption:    "Sunset vibes ??",
		ImageURL:   "https://example.com/ig1.jpg",
		Likes:      1500,
		Comments:   45,
		Shares:     23,
		PostedAt:   time.Now().Add(-24 * time.Hour),
		ProfileURL: "https://instagram.com/yourbusiness",
		Engagement: 8.2,
	},
}
