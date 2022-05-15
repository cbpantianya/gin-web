package gateway

import (
	"gin-template/lib/errx"
	"gin-template/lib/resp"
	"gin-template/server/database"
	"github.com/gin-gonic/gin"
	"regexp"
	"time"
)

// Gateway is the main entry point for the program.

func BaseCheckInGateway(c *gin.Context) {
	// Check if the user is legal.
	// If not, redirect to the 404 page.
	// If yes, continue.
	c.Next()

}

func AuthCheckInGateway(c *gin.Context) {
	// Check if the user have permission to the model.
	// If not, redirect to the 403 page.
	// If yes, continue.
	token := c.GetHeader("Authorization")
	// Match the token with the regex.
	match, err := regexp.MatchString("^Bearer (\\d|[A-Z]|[a-z]){32}$", token)
	if !match || err != nil {
		c.JSON(resp.FormatRespAutoMsg(403, errx.GatewayDeny, nil))
		c.Abort()
		return
	}
	// If the token is valid, continue.
	// Match the token in the database.
	var transferStruct database.AccessToken
	err = database.Instance.MySql.Model(&database.AccessToken{}).Where("access_token = ? and expired_at > ?  ", token[7:], time.Now()).First(&transferStruct).Error
	if err != nil {
		c.JSON(resp.FormatRespAutoMsg(403, errx.GatewayDeny, nil))
		c.Abort()
		return
	}
	c.Set("user_id", transferStruct.UserID)
	c.Next()

}

func SpecialCheckInGateway(c *gin.Context) {
	// Check if the user have permission to the model.
	// If not, redirect to the 403 page.
	// If yes, continue.
	c.Next()

}

func HardCheckInGateway(c *gin.Context) {
	// Check if the user have permission to the model.
	// If not, redirect to the 403 page.
	// If yes, continue.
	c.Next()

}
