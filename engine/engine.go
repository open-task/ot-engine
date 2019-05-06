package engine

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanocz/goginjsonrpc"
	"github.com/open-task/ot-engine/jsonrpc"
	"log"
	"net/http"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

type Config struct {
	Address string
	DSN     string
}
type OtEngineServer struct {
	Config *Config
	Engine *gin.Engine
	DB     *sql.DB
	RPC    *jsonrpc.EngineRPC
}

func New(conf *Config) (*OtEngineServer, error) {
	return &OtEngineServer{
		Config: conf,
	}, nil
}

func (e *OtEngineServer) Setup() {
	db, err := sql.Open("mysql", e.Config.DSN)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	e.DB = db

	e.Engine = gin.Default()

	// Health Check
	e.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	e.Engine.GET("healthz", e.healthz)
	e.RPC = &jsonrpc.EngineRPC{Version: "0.2.0", DB: e.DB}
	e.Engine.POST("/v1/", func(c *gin.Context) {
		goginjsonrpc.ProcessJsonRPC(c, e.RPC)
	})

	backend := e.Engine.Group("/backend/v1")
	{
		user := backend.Group("/user/:user")
		{
			user.GET("/skill", fetchAllSkill)
			user.GET("/skill/:skill", fetchSkill)
			user.POST("/skill", addSkill)
			user.PUT("/skill", updateSkill)
			user.PATCH("/skill", updateSkill)
			user.DELETE("/skill", deleteSkill)
		}
		skills := backend.Group("/skills")
		{
			skills.GET("/top",topSkills)
		}
	}
}

func (e *OtEngineServer) healthz(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	err := e.DB.PingContext(ctx)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{"message": fmt.Sprintf("db down: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (e *OtEngineServer) Serve() {
	defer e.DB.Close()

	e.Engine.Run(e.Config.Address)
}
