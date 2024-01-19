//////////////////////////////////////////////////////////
// Package: rmain
// File: redditfetch.go
// Simple command line Go application go retrive YOUR
// comments and posts and save them to a local file in JSON Format
//
// Authors: Tony Brandao m, assisted by bard.google.com
// Copyright (c) 2024 .
// This code is licensed under the MIT license.
/////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const CONFIG_FILENAME = "reddit.config.json"

type Config struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	UserAgent    string `json:"userAgent"`
}

// RedditPost represents the structure of a Reddit post.
type RedditPost struct {
	Headline string `json:"headline"`
	Text     string `json:"text"`
	Ups      int    `json:"ups"`
	Date     string `json:"date"`
}

// RedditPostResponse represents the structure of a Reddit post response.
type RedditPostResponse struct {
	Data struct {
		Title      string  `json:"title"`
		Selftext   string  `json:"selftext"`
		Ups        int     `json:"ups"`
		CreatedUTC float64 `json:"created_utc"`
	} `json:"data"`
}

// RedditCommentResponse represents the structure of a Reddit comment response.
type RedditCommentResponse struct {
	Data struct {
		Author     string  `json:"author"`
		Score      int     `json:"score"`
		Body       string  `json:"body"`
		Ups        int     `json:"ups"`
		CreatedUTC float64 `json:"created_utc"`
		Subreddit  string  `json:"subreddit"`
		CommentID  string  `json:"id"`
		PostID     string  `json:"link_id"`
	} `json:"data"`

	// Full Reddit URL derived from PostID
	FullRedditURL string `json:"full_reddit_url"`
}

// RedditCommentResponse represents the structure of a Reddit comments.
type RedditComments struct {
	CommentID string `json:"id"`
	PostID    string `json:"link_id"`
	Subreddit string `json:"subreddit"`
	Body      string `json:"body"`
	Text      string `json:"text"`
	Score     int    `json:"score"`
	Date      string `json:"created_utc"`
}

////////////////////////////////
// start of main routnine

func main() {

	banner := `
    ░█▀▄░█▀▀░█▀▄░█▀▄░▀█▀░▀█▀░░░█▀▀░█▀█░░░█▀▀░█▀▀░▀█▀░█▀▀░█░█
    ░█▀▄░█▀▀░█░█░█░█░░█░░░█░░░░█░█░█░█░░░█▀▀░█▀▀░░█░░█░░░█▀█
    ░▀░▀░▀▀▀░▀▀░░▀▀░░▀▀▀░░▀░░░░▀▀▀░▀▀▀░▀░▀░░░▀▀▀░░▀░░▀▀▀░▀░▀
	`
	fmt.Println(banner)
	fmt.Print(" Reddit Post/Comment Go.Fetch retriever >>------> \n")

	config, err := loadConfig()
	if err != nil {
		fmt.Println("Config file not found. Please enter configuration values:")
		config = promptForConfig()
		if err := saveConfig(config); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}
	}

	// Use the loaded config values here
	fmt.Println("Reading config file :", CONFIG_FILENAME)

	token, err := getAuthToken(config)
	if err != nil {
		log.Fatal(err)
	}

	posts, err := getUserPosts(token, config)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch user comments
	comments, err := getUserComments(token, config)
	if err != nil {
		log.Fatal(err)
	}

	// Update comments with full Reddit URL for links to original comments
	updateCommentURLs(comments)

	// fmt.Println("Comments: ",comments)

	var redditPosts []RedditPost

	for _, post := range posts {
		redditPosts = append(redditPosts, RedditPost{
			Headline: post.Data.Title,
			Text:     post.Data.Selftext,
			Ups:      post.Data.Ups,
			Date:     time.Unix(int64(post.Data.CreatedUTC), 0).Format(time.RFC3339),
		})
	}

	// Combine posts and comments into a single slice
	allData := []interface{}{}
	//	allData = append(allData, posts)
	allData = append(allData, comments)

	// Save all data to a JSON file
	saveToTextFile(allData, "reddit_posts_comments.json")
	// saveToTextFile(redditPosts)

} //end of main

func getAuthToken(config Config) (string, error) {
	client := &http.Client{}
	access_token_url := "https://www.reddit.com/api/v1/access_token"

	fmt.Println("Getting Access toekn form  ", access_token_url)
	req, err := http.NewRequest("POST", access_token_url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(config.ClientID, config.ClientSecret)
	req.Header.Set("User-Agent", config.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("grant_type", "password")
	q.Add("username", config.Username)
	q.Add("password", config.Password)
	req.URL.RawQuery = q.Encode()

	fmt.Println("Submitting credenitals  to ", req.URL.RawPath)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

// ...

func getUserComments(token string, config Config) ([]RedditCommentResponse, error) {
	client := &http.Client{}

	limit := 500

	fmt.Println("Requesting User Comments for  ", config.Username)
	req, err := http.NewRequest("GET", "https://oauth.reddit.com/user/"+config.Username+"/comments?sort=top&t=all&limit="+strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("User-Agent", config.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("Comments: ", resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var commentsResponse struct {
		Data struct {
			Children []RedditCommentResponse `json:"children"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &commentsResponse); err != nil {
		return nil, err
	}

	count := len(commentsResponse.Data.Children)
	fmt.Println("Number of Comments in the structure:", count)

	return commentsResponse.Data.Children, nil
}

// ...

func getUserPosts(token string, config Config) ([]RedditPostResponse, error) {
	client := &http.Client{}

	fmt.Println("Fetching Maximum User Posts for  ", config.Username)
	req, err := http.NewRequest("GET", "https://oauth.reddit.com/user/"+config.Username+"/submitted?sort=top&t=all", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("User-Agent", config.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var postsResponse struct {
		Data struct {
			Children []RedditPostResponse `json:"children"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &postsResponse); err != nil {
		return nil, err
	}

	count := len(postsResponse.Data.Children)
	fmt.Println("Number of Posts  retrieved:", count)

	return postsResponse.Data.Children, nil
}

// func saveToTextFile(posts []RedditPost) {
// 	fileData, err := json.MarshalIndent(posts, "", "  ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = ioutil.WriteFile("reddit_posts.json", fileData, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Reddit posts saved to reddit_posts.json")
// }

func saveToTextFile(data []interface{}, filename string) {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, fileData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reddit data saved to %s\n", filename)
}

func updateCommentURLs(comments []RedditCommentResponse) {
	for i := range comments {
		// example URL: https://www.reddit.com/r/the_everything_bubble/comments/19702zi/
		comments[i].FullRedditURL = "https://www.reddit.com/r/" + comments[i].Data.Subreddit + "/comments/" + comments[i].Data.PostID
	}
}

func promptForConfig() Config {
	var config Config
	fmt.Println("visit Reddit API  https://www.reddit.com/wiki/api/ then  https://old.reddit.com/prefs/apps/")
	fmt.Println("to get the neecessary access Client ID and Cleint Secret.")

	fmt.Println("Please enter the following information:")
	fmt.Print("Client ID: ")
	fmt.Scanln(&config.ClientID)
	fmt.Print("Client Secret: ")
	fmt.Scanln(&config.ClientSecret)
	fmt.Print("Username: ")
	fmt.Scanln(&config.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&config.Password)
	fmt.Print("User Agent: ")
	fmt.Scanln(&config.UserAgent)
	return config
}

func saveConfig(config Config) error {
	file, err := os.Create(CONFIG_FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func loadConfig() (Config, error) {
	file, err := os.Open(CONFIG_FILENAME)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
