package oauth

import (
	"context"
	"crypto/rand"
	"fmt"
	"gin-template/lib/errx"
	"gin-template/lib/resp"
	"gin-template/server"
	"gin-template/server/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/big"
	"regexp"
	"strconv"
	"time"
)

type OAuth struct {
}

func init() {
	var instance = &OAuth{}
	router.RegisterRequestHandler(instance)
}

func (O OAuth) Register() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"VCodePostCheckHandler": O.VCodePostCheckHandler(),
		"VCodePostHandler":      O.VCodePostHandler(),
	}
}

type VCodeRequest struct {
	PhoneNumber string `form:"phone_number" binding:"required"`
}

func (O OAuth) VCodePostCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// query check
		VCodePostRequest := VCodeRequest{}
		if err := c.ShouldBind(&VCodePostRequest); err != nil {
			c.JSON(resp.FormatRespAutoMsg(200, errx.QueryError, nil))
			c.Abort()
			fmt.Println(err)
			return
		}
		// frequency check
		result, err := server.Instance.DB.Redis.Get(context.Background(), VCodePostRequest.PhoneNumber).Result()
		if err != nil {
			if result == "" {
				c.Set("phone_number", VCodePostRequest.PhoneNumber)
				c.Next()
				return
			}
			c.JSON(resp.FormatRespAutoMsg(200, errx.RedisError, nil))
			c.Abort()
			return // ignore error
		}
		if result != "" {
			// get ttl
			ttl, err := server.Instance.DB.Redis.TTL(context.Background(), VCodePostRequest.PhoneNumber).Result()
			if err != nil {
				c.JSON(resp.FormatRespAutoMsg(200, errx.RedisError, nil))
				c.Abort()
				return
			}
			if ttl >= time.Minute*4 {
				c.JSON(resp.FormatRespAutoMsg(200, errx.FrequencyError, nil))
				c.Abort()
				return
			} else {
				c.Set("phone_number", VCodePostRequest.PhoneNumber)
				c.Next()
			}
			return
		}
	}
}

func (O OAuth) VCodePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check phone number
		phone, _ := c.Get("phone_number")
		isMatchPhoneNumber, err := regexp.MatchString("^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$", phone.(string))
		if isMatchPhoneNumber == false {
			c.JSON(resp.FormatRespSelfMsg(200, -1, "手机号格式不正确", nil))
			return
		}
		// Random code
		max := new(big.Int).SetInt64(int64(899999))
		min := new(big.Int).SetInt64(int64(100000))
		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			logrus.WithField("server", "Handler").Errorf("Can't generate vCode: %v, %v", i, err)
			c.JSON(resp.FormatRespAutoMsg(200, -1, nil))
			return
		}
		vCode := i.Int64() + min.Int64()

		// Store code and phone number to redis

		err = server.Instance.DB.Redis.Set(context.Background(), phone.(string), strconv.FormatInt(vCode, 10), time.Minute*5).Err()
		if err != nil {
			c.JSON(resp.FormatRespAutoMsg(200, errx.RedisError, nil))
			return
		}

		// Send code to user
		// No Money to send code :-(
		// Just print the code

		// Return
		c.JSON(resp.FormatRespAutoMsg(200, 0, nil))
	}
}
