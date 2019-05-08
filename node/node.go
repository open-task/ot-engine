package node

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanocz/goginjsonrpc"
	"github.com/open-task/ot-engine/jsonrpc"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/open-task/ot-engine/engine"
)

type Config struct {
	Server    string   `json:"server"`
	Contract  string   `json:"contract""`
	FromFlag  bool     `json:"from_flag"`
	FromBlock *big.Int `json:"from_block"`
	ToFlag    bool     `json:"to_flag"`
	ToBlock   *big.Int `json:"to_block"`
}

type HTTPConfig struct {
	Port string `json:port`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:port`
	User     string `json:user`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (db DatabaseConfig) DSN() (dsn string) {
	dsn = db.User + ":" + db.Password + "@"
	if db.Host != "" {
		dsn += "tcp(" + db.Host
		if db.Port != "" {
			dsn += ":" + db.Port
		}
		dsn += ")"
	}

	dsn += "/" + db.Database
	return dsn
}

type Engine struct {
	Database DatabaseConfig `json:"database"`
	Http     HTTPConfig     `json:http`
}

func (e Engine) Address() string {
	return ":" + e.Http.Port
}

type Backend struct {
	Database DatabaseConfig `json:"database"`
}
type Node struct {
	EthConfig     *Config
	EngineConfig  *Engine
	BackendConfig *Backend

	GinServer *gin.Engine

	EngineDB *sql.DB
	RPC      *jsonrpc.EngineRPC

	BackendDB *sql.DB
}

var DefaultConfig = Config{
	Server:    "http://localhost",
	Contract:  "",
	FromFlag:  false,
	FromBlock: big.NewInt(0),
	ToFlag:    false,
	ToBlock:   big.NewInt(0),
}

func (n *Node) Setup() {
	db, err := sql.Open("mysql", n.EngineConfig.Database.DSN())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	n.EngineDB = db

	db, err = sql.Open("mysql", n.BackendConfig.Database.DSN())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	n.BackendDB = db

	n.GinServer = gin.Default()
	n.RPC = &jsonrpc.EngineRPC{Version: "0.2.0", DB: n.EngineDB}

	n.GinServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	n.GinServer.GET("healthz", n.healthz)
	n.GinServer.POST("/v1/", func(c *gin.Context) {
		goginjsonrpc.ProcessJsonRPC(c, n.RPC)
	})
	n.GinServer.POST("/engine/v2/", func(c *gin.Context) {
		goginjsonrpc.ProcessJsonRPC(c, n.RPC)
	})

	backend := n.GinServer.Group("/backend/v1")
	{
		user := backend.Group("/user/:user")
		{
			user.GET("/skill", func(c *gin.Context) {
				engine.FetchUserSkills(c, n.BackendDB)
			})
			user.GET("/skill/:skill", func(c *gin.Context) {
				engine.FetchUserSkill(c, n.BackendDB)
			})
			user.POST("/skill/:skill", func(c *gin.Context) {
				engine.AddUserSkill(c, n.BackendDB)
			})
			user.PUT("/skill/:skill", func(c *gin.Context) {
				engine.UpdateUserSkill(c, n.BackendDB)
			})
			user.PATCH("/skill/:skill", func(c *gin.Context) {
				engine.UpdateUserSkill(c, n.BackendDB)
			})
			user.DELETE("/skill/:skill", func(c *gin.Context) {
				engine.DeleteUserSkill(c, n.BackendDB)
			})
		}
		skills := backend.Group("/skill")
		{
			skills.GET("/top", func(c *gin.Context) {
				engine.TopSkills(c, n.BackendDB)
			})
		}
	}
}

func (n *Node) healthz(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	err := n.EngineDB.PingContext(ctx)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{"message": fmt.Sprintf("engine db down: %v", err)})
		return
	}
	err = n.BackendDB.PingContext(ctx)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{"message": fmt.Sprintf("backend db down: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (n *Node) Serve() {
	defer n.EngineDB.Close()
	defer n.BackendDB.Close()

	n.GinServer.Run(n.EngineConfig.Address())
}
