package ping

import (
	"fmt"
	"gin-template/lib/resp"
	"gin-template/server/router"
	"github.com/gin-gonic/gin"
)

type Ping struct {
}

func (p *Ping) Register() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"PingGetHandler":      p.GetHandler(),
		"AfterPingGetHandler": p.AfterGetHandler,
	}
}

func init() {
	var instance = &Ping{}
	router.RegisterRequestHandler(instance)
}

// GetHandler ping handler
func (p *Ping) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("GetHandler")
		userID, _ := c.Get("user_id")
		respStr := fmt.Sprintf("pong: %s", userID)
		c.JSON(resp.FormatRespSelfMsg(200, 0, "success", respStr))
		c.Next()
		return
	}
}

func (p *Ping) AfterGetHandler(c *gin.Context) {
	// Do something after the handler

}
