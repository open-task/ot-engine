package engine

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/kanocz/goginjsonrpc"
	"github.com/open-task/ot-engine/jsonrpc"
	"log"
	"net/http"
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
		c.String(http.StatusOK, "pong")
	})

	e.RPC = &jsonrpc.EngineRPC{Version: "0.2.0", DB: e.DB}
	e.Engine.POST("/v1/", func(c *gin.Context) {
		goginjsonrpc.ProcessJsonRPC(c, e.RPC)
	})
}

func (e *OtEngineServer) Serve() {
	defer e.DB.Close()

	e.Engine.Run(e.Config.Address)
}
