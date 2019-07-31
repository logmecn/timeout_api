package main

import (
	"./utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	glog  *zap.Logger
	sugar *zap.SugaredLogger
)

func main() {
	glog = utils.InitLogger("./log/sv.log", 16, "debug")
	f, _ := os.Create("./log/gin.log")
	sugar = glog.Sugar()

	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/", AddData)
	err := r.Run("0.0.0.0:8001")
	if err != nil {
		sugar.Fatal("Http 启动失败：", err)
	}
}

type TData struct {
	Data string `json:"data"`
	Url  string `json:"url"`
	TOut uint8  `josn:"tout"`
}

func AddData(c *gin.Context) {
	var Djson TData
	if err := c.ShouldBindJSON(&Djson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ch:
			case <-time.After(time.Second * time.Duration(Djson.TOut)):
				var url string
				if len(strings.Split(Djson.Url, "?")) > 1 {
					url = Djson.Url + "&"
				} else {
					url = Djson.Url + "?"
				}
				httpGet(url + Djson.Data)
				quit <- true
				return
			}
		}
	}()

	c.JSON(http.StatusOK, gin.H{"ret": 200, "msg": "Data had send"})
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		glog.Error("Http_Get_error", zap.String("url", url))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error("Http_GetBody_error", zap.String("url", url))
	}
	glog.Debug("Http_Get_Result", zap.String("body", string(body)))
}
