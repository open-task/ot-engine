package engine

import (
	"database/sql"
	"github.com/gin-gonic/gin"
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
}

func New(conf *Config) (*OtEngineServer, error) {
	return &OtEngineServer{
		Config: conf,
	}, nil
}

func (e *OtEngineServer) Setup() {

	e.Engine = gin.Default()

	// Health Check
	e.Engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	db, err := sql.Open("mysql", e.Config.DSN)
	if err != nil {
		log.Fatal(err)
	}
	e.DB = db
}

func (e *OtEngineServer) Serve() {
	defer e.DB.Close()

	e.Engine.Run(e.Config.Address)
}
