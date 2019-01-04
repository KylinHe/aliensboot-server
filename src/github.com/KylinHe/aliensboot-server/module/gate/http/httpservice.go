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
	"github.com/KylinHe/aliensboot-core/config"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Init(config config.HttpConfig) {
	if config.Address == "" {
		return
	}
	go func() {
		router := gin.Default()
		if !log.DEBUG {
			gin.SetMode(gin.ReleaseMode)
		}
		router.Any(config.Prefix+"/:service/:handler", handleService)
		err := router.Run(config.Address)
		if err != nil {
			log.Fatalf("start http service err : %v", err)
		}
	}()
}

func Close() {

}

//添加弹幕信息
func handleService(c *gin.Context) {
	service := c.Param("service")

	//var data map[string]interface{}
	//c.ShouldBindQuery(data)
	//c.shouldbind

	body, _ := ioutil.ReadAll(c.Request.Body)
	//data, _ := json.Marshal(c.Request.Form)

	log.Debugf("request data %v", string(body))
	response, err := route.HandleUrlMessage(service, nil)
	if err != nil {
		response = []byte(err.Error())
		c.String(http.StatusInternalServerError, string(response))
	} else {
		c.String(http.StatusOK, string(response))
	}
	//_, err := w.Write(response)
	//if err != nil {
	//	log.Debug(err.Error())
	//}
}
