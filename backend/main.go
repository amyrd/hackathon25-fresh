package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

type AnalyticsMetric struct {
	Date         time.Time `json:"date"`
	Impressions  int       `json:"impressions"`
	Engagement   float64   `json:"engagement"`
	Followers    int       `json:"followers"`
	ProfileViews int       `json:"profile_views"`
}

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
		{
			ID:         "ig_2",
			Platform:   "instagram",
			Caption:    "i love neovim!!!",
			ImageURL:   "https://blog.vtutor.com/wp-content/uploads/2019/11/Vim-logo.jpg",
			Likes:      21,
			Comments:   12,
			Shares:     2,
			PostedAt:   time.Now().Add(-24 * time.Hour),
			ProfileURL: "https://instagram.com/yourbusiness",
			Engagement: 2.1,
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
			Caption:    "?? Just launched our new eco-friendly product line! ?? Join the green revolution and shop now! #Sustainability #Innovation",
			ImageURL:   "https://images.pexels.com/photos/4466176/pexels-photo-4466176.jpeg",
			Likes:      1245,
			Comments:   189,
			Shares:     342,
			PostedAt:   time.Now().Add(-24 * 7 * time.Hour),
			ProfileURL: "https://twitter.com/yourbusiness",
			Engagement: 14.2,
		},
		{
			ID:         "tw_2",
			Platform:   "twitter",
			Caption:    "?? Quarterly results are in! Grateful to our amazing team and customers for another record-breaking quarter! ?? #BusinessGrowth",
			ImageURL:   "https://images.pexels.com/photos/3184292/pexels-photo-3184292.jpeg",
			Likes:      892,
			Comments:   76,
			Shares:     155,
			PostedAt:   time.Now().Add(-24 * 6 * time.Hour),
			ProfileURL: "https://twitter.com/yourbusiness",
			Engagement: 9.8,
		},
		{
			ID:         "tw_3",
			Platform:   "twitter",
			Caption:    "?? Proud to partner with @CharityOrg to support local communities! 5% of all sales this month will be donated. #CorporateResponsibility",
			ImageURL:   "https://images.pexels.com/photos/6646918/pexels-photo-6646918.jpeg",
			Likes:      2345,
			Comments:   302,
			Shares:     489,
			PostedAt:   time.Now().Add(-24 * 5 * time.Hour),
			ProfileURL: "https://twitter.com/yourbusiness",
			Engagement: 18.7,
		},
	}
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/login", handleLogin)

	router.GET("/posts", func(c *gin.Context) {
		allPosts := append(instagramPosts, twitterPosts...)
		c.JSON(http.StatusOK, allPosts)
	})

	router.GET("/instagram/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, instagramPosts)
	})

	router.GET("/instagram/analytics", func(c *gin.Context) {
		c.JSON(http.StatusOK, instagramAnalytics)
	})

	router.GET("/twitter/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, twitterPosts)
	})

	router.GET("/twitter/analytics", func(c *gin.Context) {
		c.JSON(http.StatusOK, calculateTwitterAnalytics())
	})

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
				"followers": calculateTwitterFollowers(),
				"engagement": gin.H{
					"average": calculateAverageEngagement(twitterPosts),
					"trend":   calculateEngagementTrend(),
				},
				"metrics": calculateTwitterAnalytics(),
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

func calculateTwitterAnalytics() []AnalyticsMetric {
	var metrics []AnalyticsMetric
	baseDate := time.Now().Add(-24 * 7 * time.Hour)

	for i := 0; i < 7; i++ {
		date := baseDate.Add(time.Duration(i) * 24 * time.Hour)
		metrics = append(metrics, AnalyticsMetric{
			Date:         date,
			Impressions:  calculateDailyImpressions(date),
			Engagement:   calculateDailyEngagement(date),
			Followers:    calculateFollowersOnDate(date),
			ProfileViews: calculateProfileViews(date),
		})
	}

	return metrics
}

func calculateDailyImpressions(date time.Time) int {
	impressions := 0
	for _, post := range twitterPosts {
		if sameDate(post.PostedAt, date) {
			impressions += post.Likes*10 + post.Comments*20 + post.Shares*30
		}
	}
	if impressions == 0 {
		return 5000 + int(date.Unix())%10000
	}
	return impressions
}

func calculateDailyEngagement(date time.Time) float64 {
	total := 0.0
	count := 0
	for _, post := range twitterPosts {
		if sameDate(post.PostedAt, date) {
			total += post.Engagement
			count++
		}
	}
	if count == 0 {
		return 5.0 + float64(date.Day())/10
	}
	return total / float64(count)
}

func calculateFollowersOnDate(date time.Time) int {
	baseFollowers := 10000
	for _, post := range twitterPosts {
		if post.PostedAt.Before(date) {
			baseFollowers += post.Shares/10 + post.Comments/20
		}
	}
	return baseFollowers
}

func calculateProfileViews(date time.Time) int {
	return 500 + int(date.Unix())%1000
}

func calculateAverageEngagement(posts []SocialMediaPost) float64 {
	total := 0.0
	for _, post := range posts {
		total += post.Engagement
	}
	return total / float64(len(posts))
}

func calculateEngagementTrend() []float64 {
	return []float64{12.1, 13.4, 14.9, 15.2, 16.0, 17.3, 18.1}
}

func calculateTwitterFollowers() int {
	return calculateFollowersOnDate(time.Now())
}

func sameDate(a, b time.Time) bool {
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
