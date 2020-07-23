package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v8"
)

// Global variable
var ctx = context.Background()
var redisClient *libredis.Client

// Struct
type PostType struct {
	Type string `json:"type"`
}

type PostMessage struct {
	GroupID string `json:"groupID"`
	Type    string `json:"type"`
	Data    struct {
		Text string `json:"text"`
	} `json:"data"`
}

type PostImage struct {
	GroupID string `json:"groupID"`
	Type    string `json:"type"`
	Data    struct {
		ImageID string `json:"imageID"`
	} `json:"data"`
}

// Util function
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readBody(c *gin.Context, v interface{}) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), v)
	check(err)
}

// Initial
func initRedis() *libredis.Client {
	option, err := libredis.ParseURL("redis://localhost:6379/0")
	check(err)
	client := libredis.NewClient(option)
	return client
}

// Controller
func webhookHanlder(c *gin.Context) {
	var postType PostType
	readBody(c, &postType) // 讀取request的body，並把struct套上，這個PostType struct是專門用來解析request type的
	switch postType.Type { // 如果是message type就以textHandler function處理，如果是image type就以imageHandler function處理
	case "message":
		textHandler(c)
	case "image":
		imageHandler(c)
	}
}

func textHandler(c *gin.Context) {
	var postMessage PostMessage
	readBody(c, &postMessage)                       // 讀取request的body，並把struct套上，這個PostMessage struct是專門用來讀取message json的
	key := postMessage.GroupID + ":search"          // 將GroupID+type
	err := redisClient.Set(ctx, key, true, 0).Err() // 設置key為GroupID+type，value為true
	check(err)
	err = redisClient.Expire(ctx, key, 24*time.Hour).Err() // 設置key-value的過期時間為24小時
	check(err)
	c.Status(204) // 回傳204 response
}

func imageHandler(c *gin.Context) {
	isThereSearchRequest := func(redisGroupResult string, err error) bool {
		if err != nil || redisGroupResult == "" {
			return false
		}
		redisGroupResultBool, err := strconv.ParseBool(redisGroupResult)
		check(err)
		return redisGroupResultBool
	}

	var postImage PostImage
	readBody(c, &postImage)                                       // 讀取request的body，並把struct套上，這個PPostImage struct是專門用來讀取image json的
	key := postImage.GroupID + ":search"                          // 將GroupID+type
	if isThereSearchRequest(redisClient.Get(ctx, key).Result()) { // 如果GroupID+type存在就回傳女優資訊，不存在就回傳「你沒有要求辨識」的error message
		c.JSON(200, gin.H{
			"AVStar": "佐倉絆",
		})
		return
	}
	c.JSON(403, gin.H{
		"error": "No search requested",
	})
}

func main() {
	r := gin.Default()
	redisClient = initRedis() // 初始化Redis

	r.POST("/api/WebhookHanlder", webhookHanlder) // 新增POST API api/WebhookHanlder
	r.Run()
}
