package node

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kanocz/goginjsonrpc"
	"github.com/open-task/ot-engine/engine"
	"github.com/open-task/ot-engine/jsonrpc"
	"log"
	"math/big"
	"net/http"
	"time"
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

	dsn += "/" + db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
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

	BackendDB *gorm.DB
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

	gormDB, err := gorm.Open("mysql", n.BackendConfig.Database.DSN())
	if err != nil {
		log.Fatal(err)
	}
	gormDB.DB().SetMaxIdleConns(5)
	gormDB.DB().SetMaxOpenConns(10)
	gormDB.DB().SetConnMaxLifetime(time.Hour)
	n.BackendDB = gormDB

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
		backend.GET("/list_skills", func(c *gin.Context) {
			engine.ListSkills(c, n.BackendDB, n.EngineDB)
		})
		backend.GET("/get_users", func(c *gin.Context) {
			engine.GetUsers(c, n.BackendDB, n.EngineDB)
		})
		backend.POST("/add_skill", func(c *gin.Context) {
			engine.AddSkill(c, n.BackendDB, n.EngineDB)
		})
		user := backend.Group("/user/:address")
		{
			user.GET("/skill", func(c *gin.Context) {
				engine.FetchUserSkills(c, n.BackendDB)
			})
			user.POST("/skill", func(c *gin.Context) {
				engine.AddUserSkill(c, n.BackendDB)
			})
			user.GET("/skill/:id", func(c *gin.Context) {
				engine.FetchUserSkill(c, n.BackendDB)
			})
			user.PUT("/skill/:skill", func(c *gin.Context) {
				engine.UpdateUserSkill(c, n.BackendDB)
			})
			user.PATCH("/skill/:skill", func(c *gin.Context) {
				engine.UpdateUserSkill(c, n.BackendDB)
			})
			user.POST("/skill/:skill", func(c *gin.Context) {
				engine.UpdateUserSkill(c, n.BackendDB)
			})
			user.DELETE("/skill/:id", func(c *gin.Context) {
				engine.DeleteUserSkill(c, n.BackendDB)
			})

			user.GET("/info", func(c *gin.Context) {
				engine.FetchUserInfo(c, n.BackendDB)
			})
			user.POST("/info", func(c *gin.Context) {
				engine.UpdateUserInfo(c, n.BackendDB)
			})
			//info := user.Group("/info")
			//{
			//	//info.GET("/info", func(c *gin.Context) {
			//	//	engine.FetchUserInfo(c, n.BackendDB)
			//	//})
			//}
			user.GET("/mission", func(c *gin.Context) {
				engine.FetchUserMissions(c, n.BackendDB)
			})
			//mission := user.Group("/mission")
			//{
			//	//info.GET("/info", func(c *gin.Context) {
			//	//	engine.FetchUserInfo(c, n.BackendDB)
			//	//})
			//}
		}
		backend.GET("/skill", func(c *gin.Context) {
			engine.FetchSkills(c, n.BackendDB)
		})
		skill := backend.Group("/skill")
		{
			skill.POST("/update_skill", func(c *gin.Context) {
				engine.UpdateSkills(c, n.BackendDB)
			})
			skill.POST("/get_skill", func(c *gin.Context) {
				engine.GetSkills(c, n.BackendDB)
			})
			skill.POST("/del_skill", func(c *gin.Context) {
				engine.DeleteSkills(c, n.BackendDB)
			})
			skill.GET("/:id/user", func(c *gin.Context) {

				engine.FetchSkillProviders(c, n.BackendDB)
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
	err = n.BackendDB.DB().PingContext(ctx)
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
