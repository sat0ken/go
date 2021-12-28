package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetPingRequest struct {
}

func NewPingRequest() GetPingRequest {
	return GetPingRequest{}
}

func sleep() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("error is occur : %s\n", time.Now())
	}
	return
}

func (p *GetPingRequest) GetPing(c *gin.Context) {
	//TODO implement me
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	go sleep()
}
