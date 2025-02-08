package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	TwitterBearerToken = "AAAAAAAAAAAAAAAAAAAAAFB7ywEAAAAAKtRzVnjv%2F%2Bf0JzgroNrgb2ynrjQ%3D91EFsgJTgck55oLumJ2bJOzjBKfPDgkpBgx3SByksTX1iIaML3"
	TwitterAccountID   = "1888305777993252864"
)

type TwitterClient struct {
	bearerToken string
}

func NewTwitterClient() *TwitterClient {
	return &TwitterClient{
		bearerToken: TwitterBearerToken,
	}
}

type TwitterResponse struct {
	Data []struct {
		ID            string `json:"id"`
		Text          string `json:"text"`
		CreatedAt     string `json:"created_at,omitempty"`
		PublicMetrics struct {
			RetweetCount int `json:"retweet_count,omitempty"`
			ReplyCount   int `json:"reply_count,omitempty"`
			LikeCount    int `json:"like_count,omitempty"`
			QuoteCount   int `json:"quote_count,omitempty"`
		} `json:"public_metrics,omitempty"`
	} `json:"data"`
	Errors []struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"errors"`
}

func (tc *TwitterClient) FetchTweets() ([]SocialMediaPost, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?tweet.fields=created_at,public_metrics", TwitterAccountID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tc.bearerToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewReader(body)) // Reset for decoding

	// Debug logging
	fmt.Printf("Twitter API Response (%d): %s\n", resp.StatusCode, string(body))

	var result TwitterResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode error: %v\nBody: %s", err, string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error %d: %+v", resp.StatusCode, result.Errors)
	}

	return transformTwitterPosts(result.Data), nil
}

func transformTwitterPosts(tweets []struct {
	ID            string `json:"id"`
	Text          string `json:"text"`
	CreatedAt     string `json:"created_at,omitempty"`
	PublicMetrics struct {
		RetweetCount int `json:"retweet_count,omitempty"`
		ReplyCount   int `json:"reply_count,omitempty"`
		LikeCount    int `json:"like_count,omitempty"`
		QuoteCount   int `json:"quote_count,omitempty"`
	} `json:"public_metrics,omitempty"`
}) []SocialMediaPost {
	var posts []SocialMediaPost

	for _, tweet := range tweets {
		var createdTime time.Time
		if tweet.CreatedAt != "" {
			if t, err := time.Parse(time.RFC3339, tweet.CreatedAt); err == nil {
				createdTime = t
			}
		}

		posts = append(posts, SocialMediaPost{
			ID:         tweet.ID,
			Platform:   "twitter",
			Caption:    tweet.Text,
			Likes:      tweet.PublicMetrics.LikeCount,
			Comments:   tweet.PublicMetrics.ReplyCount,
			Shares:     tweet.PublicMetrics.RetweetCount,
			PostedAt:   createdTime,
			ProfileURL: fmt.Sprintf("https://twitter.com/user/status/%s", tweet.ID),
			Engagement: calculateEngagement(tweet.PublicMetrics),
		})
	}
	return posts
}

func calculateEngagement(metrics struct {
	RetweetCount int `json:"retweet_count,omitempty"`
	ReplyCount   int `json:"reply_count,omitempty"`
	LikeCount    int `json:"like_count,omitempty"`
	QuoteCount   int `json:"quote_count,omitempty"`
}) float64 {
	return float64(
		metrics.LikeCount+
			metrics.RetweetCount*2+
			metrics.ReplyCount*3+
			metrics.QuoteCount*2,
	) / 100
}
