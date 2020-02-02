package main

import (
  //"bytes"
  //"encoding/gob"
  //"encoding/json"
  "fmt"
  "os"
  //"regexp"
  "io/ioutil"
  "strconv"

  "github.com/gin-gonic/gin"
  log "github.com/sirupsen/logrus"

  //"time"
  "net/http"
)

var debug bool

func set_cors_headers(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
}

// err can be nil here
func return_500(c *gin.Context, msg string, err error) {
  var full_msg string
  if err != nil {
    full_msg = msg + ": " + err.Error()
  } else {
    full_msg = msg
  }
  log.Warn(full_msg)
  if debug {
    c.String(500, full_msg)
  } else {
    c.String(500, msg)
  }
}

func proxy_route(c *gin.Context) {
  url := c.Query("url")
  if url == "" {
    c.String(400, "cors-proxy: url parameter missing")
    return
  }

  resp, err := http.Get(url)
  if err != nil {
    return_500(c, "cors-proxy: problem making request: ", err)
    return
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return_500(c, "cors-proxy: problem reading response: ", err)
    return
  }
  set_cors_headers(c)
  c.String(resp.StatusCode, string(body))
}

func main() {
  var err error

  // Disable Console Color
  // gin.DisableConsoleColor()

  _debug := os.Getenv("DEBUG")
  if _debug == "" {
    gin.SetMode(gin.ReleaseMode)
    log.SetLevel(log.WarnLevel)
    debug = false
  } else {
    log.SetLevel(log.DebugLevel)
    debug = true
  }

  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()

  // https://stackoverflow.com/questions/32443738/setting-up-route-not-found-in-gin
  router.NoRoute(proxy_route)

  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  port := os.Getenv("PORT")
  var iport int
  if port == "" {
    iport = 8099
  } else {
    iport, err = strconv.Atoi(port)
    if err != nil {
      log.Fatal(err)
    }
  }
  log.Info(fmt.Sprintf("Listening on port %d", iport))
  router.Run(fmt.Sprintf(":%d", iport))
  // router.Run(":3000") for a hard coded port

}

// TODO: stream responses. See:
// https://stackoverflow.com/questions/44825244/how-to-write-a-stream-api-using-gin-gonic-server-in-golang-tried-c-stream-didnt
