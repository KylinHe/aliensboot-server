/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/2
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package http

import (
	"github.com/KylinHe/aliensboot-server/protocol"
	"encoding/json"
	"fmt"
	"github.com/KylinHe/aliensboot-core/config"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/task"
	"github.com/gin-gonic/gin"
	"net/http"
	//_ "net/http/pprof"
)

func Init(config config.HttpConfig) {
	//_ = http.ListenAndServe("0.0.0.0:6060", nil)
	if config.Address == "" {
		return
	}
	//if !aliensboot.IsDebug() {
	gin.SetMode(gin.ReleaseMode)
	//}
	router := gin.Default()

	router.Any(config.Prefix+"/:service/:handler", handleService)
	task.SafeGo(func() {
		defer func() {
			if err := recover(); err != nil {
				exception.PrintStackDetail(err)
			}
		}()
		err := router.Run(config.Address)
		if err != nil {
			log.Fatalf("start http service err : %v", err)
		}
	})
}

func Close() {

}

func registerMsg(handler string, data []byte, request *protocol.Request) error {
	if handler == "benchmark" {
		benchmark := &protocol.Benchmark{}
		err := json.Unmarshal(data, benchmark)
		game := &protocol.Request_Benchmark{Benchmark: benchmark}
		request.Game = game
		return err
	}
	return fmt.Errorf("unexpect handler %v", handler)

}

func handleService(c *gin.Context) {
	//c.String(http.StatusOK, "ok")

	c.JSON(http.StatusOK, gin.H{"errmsg": 1})
	//service := c.Param("service")
	//handler := c.Param("handler")
	//
	//body, _ := ioutil.ReadAll(c.Request.Body)
	//request := &protocol.Request{}
	//err := registerMsg(handler, body, request)
	//
	//data, err := request.Marshal()
	//
	//response, err := route.HandleUrlMessage(service, data)
	//
	//if err != nil {
	//	response = []byte(err.Error())
	//	log.Debugf("request data %v", string(response))
	//	c.String(http.StatusInternalServerError, string(response))
	//} else {
	//	c.String(http.StatusOK, string(response))
	//}
	//_, err := w.Write(response)
	//if err != nil {
	//	log.Debug(err.Error())
	//}
}
