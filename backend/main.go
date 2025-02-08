package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SocialMediaPost represents a unified post structure for frontend
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

// AnalyticsMetric represents unified analytics data
type AnalyticsMetric struct {
	Date         time.Time `json:"date"`
	Impressions  int       `json:"impressions"`
	Engagement   float64   `json:"engagement"`
	Followers    int       `json:"followers"`
	ProfileViews int       `json:"profile_views"`
}

// Hardcoded data for development (replace with real API calls later)
var (
	instagramPosts = []SocialMediaPost{
		{
			ID:         "ig_1",
			Platform:   "instagram",
			Caption:    "Sunset vibes ??",
			ImageURL:   "https://i.pinimg.com/736x/ba/4d/3e/ba4d3e4714e2615ea65dcca8ed79e9aa.jpg",
			Likes:      1500,
			Comments:   45,
			Shares:     23,
			PostedAt:   time.Now().Add(-24 * time.Hour),
			ProfileURL: "https://instagram.com/yourbusiness",
			Engagement: 8.2,
		},
	}

	instagramAnalytics = []AnalyticsMetric{
		{
			Date:         time.Now().Add(-24 * time.Hour),
			Impressions:  15000,
			Engagement:   8.2,
			Followers:    4500,
			ProfileViews: 300,
		},
	}

	twitterPosts = []SocialMediaPost{
		{
			ID:         "tw_1",
			Platform:   "twitter",
			Caption:    "New product launch! ??",
			ImageURL:   "https://i.redd.it/9n2rfmgdykpb1.jpg",
			Likes:      890,
			Comments:   56,
			Shares:     142,
			PostedAt:   time.Now().Add(-12 * time.Hour),
			ProfileURL: "https://twitter.com/yourbusiness",
			Engagement: 12.4,
		},
	}
)

func main() {
	router := gin.Default()

	// Configure CORS for frontend development
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Login endpoint
	router.POST("/login", handleLogin)

	// Unified posts endpoint
	router.GET("/posts", func(c *gin.Context) {
		allPosts := append(instagramPosts, twitterPosts...)
		c.JSON(http.StatusOK, allPosts)
	})

	// Platform-specific endpoints
	router.GET("/instagram/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, instagramPosts)
	})

	router.GET("/instagram/analytics", func(c *gin.Context) {
		c.JSON(http.StatusOK, instagramAnalytics)
	})

	router.GET("/twitter/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, twitterPosts)
	})

	// Unified analytics endpoint
	router.GET("/analytics", func(c *gin.Context) {
		response := gin.H{
			"instagram": gin.H{
				"posts":     len(instagramPosts),
				"followers": 4500,
				"engagement": gin.H{
					"average": 8.2,
					"trend":   []float64{7.8, 8.0, 8.2, 8.5},
				},
			},
			"twitter": gin.H{
				"posts":     len(twitterPosts),
				"followers": 2300,
				"engagement": gin.H{
					"average": 12.4,
					"trend":   []float64{11.2, 11.8, 12.1, 12.4},
				},
			},
		}
		c.JSON(http.StatusOK, response)
	})

	router.Run(":8080")
}

func handleLogin(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hardcoded credentials for development
	if credentials.Username == "admin" && credentials.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"token":   "dummy-token",
			"user": gin.H{
				"name":  "Social Media Manager",
				"email": "manager@example.com",
			},
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}
}
