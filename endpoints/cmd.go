package endpoints

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Cmd struct {
	Cmd      int     `json:"cmd"`
	Sequence [][]int `json:"sequence"`
}

var cmdQueue = make([]Cmd, 0)
var cmdQueueMutex sync.Mutex

func GetCmd(c *gin.Context) {
	if len(cmdQueue) == 0 {
		c.Status(http.StatusOK)
		return
	}
	cmd := popCmd()
	c.JSON(http.StatusOK, cmd)
	return
}

func PostCmd(c *gin.Context) {
	var cmd Cmd
	if c.BindJSON(&cmd) != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	pushCmd(cmd)
	c.Status(http.StatusOK)
	return
}

func pushCmd(cmd Cmd) {
	cmdQueueMutex.Lock()
	defer cmdQueueMutex.Unlock()
	cmdQueue = append(cmdQueue, cmd)
}

func popCmd() Cmd {
	cmdQueueMutex.Lock()
	defer cmdQueueMutex.Unlock()
	var cmd Cmd
	cmd, cmdQueue = cmdQueue[0], cmdQueue[1:]
	return cmd
}
