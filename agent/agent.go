package agent

import (
	"finalTaskLMS/agent/internal/handlers"
	"finalTaskLMS/agent/types"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"sync"
)

var once sync.Once

type Server struct {
	R *gin.Engine
	A *types.Agent
}

func NewAgentServer() *Server {
	var s Server
	once.Do(func() {
		s = Server{
			A: &types.Agent{},
		}
	})
	return &s
}

func (s *Server) InitializeAgent() {
	res, err := strconv.Atoi(os.Getenv("LIMIT_OF_GOROUTINES"))
	if err != nil {
		fmt.Println("err : can't parse LIMIT_OF_GOROUTINES (env)")
	}
	s.A.LimitOfGoroutines = res
	s.A.CountOfGoroutines = 0
}

func (s *Server) ConfigureRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	s.R = r
}

func (s *Server) RunServer() {
	go handlers.CycleTask(s.A)

	if err := s.R.Run(":8081"); err != nil {
		log.Fatal("qwekrqwkerkopqwkeopr[k")
	}

}
