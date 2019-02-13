package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanocz/goginjsonrpc"
	"github.com/xyths/ot-engine/jsonrpc"
	"github.com/xyths/ot-engine/config"
	"flag"
	"fmt"
	"log"
	"database/sql"
)

const version string = "0.1.1"

// 实际中应该用更好的变量名
var (
	h bool
	v bool

	configFile string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")

	flag.StringVar(&configFile, "c", "config.json", "`config`: config file")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `ot-engine: OpenTask Engine, version: %s
Usage: de [-hv] [-c config]

Options:
`, version)
	flag.PrintDefaults()
}

var db1 = make(map[string]string)

func setupRouter(rpc *jsonrpc.EngineRPC) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/v1/", func(c *gin.Context) { goginjsonrpc.ProcessJsonRPC(c, rpc) })

	// Get user value
	/*	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db1[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})*/

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	/*	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))*/

	/*	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db1[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})*/

	return r
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if v {
		fmt.Println("OpenTask Engine, version:", version)
		return
	}

	cfg, err := config.LoadConfig(configFile)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database: ", cfg.DSN())

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rpc := jsonrpc.EngineRPC{Version: version, DB: db}

	r := setupRouter(&rpc)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
