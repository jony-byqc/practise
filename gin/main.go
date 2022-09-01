package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

//type reqVehicleBody struct {
//	VinCode string `json:"VinCode"`
//}

func main() {
	r := gin.Default()
	r.POST("/api/v1/barcodes", func(c *gin.Context) {
		all, _ := ioutil.ReadAll(c.Request.Body)
		req := CommonResp{}
		err := json.Unmarshal(all, &req)
		log.Println(string(all))
		//req := reqVehicleBody{}
		//err := json.Unmarshal(all, &req)
		if err != nil {
			log.Println(err)
		}
		//data := CommonResp{
		//	Code:     "200",
		//	Message:  "OK",
		//	TraceID:  "DJ2022A",
		//	ErrorMsg: "no error",
		//}
		data := CommonResp{
			Code:    "0",
			Message: "OK",
			Data: []MesVehicle{{
				AutoID:    "DJ2022A",
				Vin:       "DJ2022A12345678900000",
				Station:   "DJ2022A",
				Job:       "1",
				Configure: "DJ2022A",
			},
			},
		}

		//consumer.JSON(http.StatusOK, data)
		c.JSON(http.StatusOK, data)
	})
	r.GET("/api/v2/barcodes", func(c *gin.Context) {
		all, _ := ioutil.ReadAll(c.Request.Body)
		//req := reqVehicleBody{}
		//err := json.Unmarshal(all, &req)
		req := CommonResp{}
		err := json.Unmarshal(all, &req)
		if err != nil {
			log.Println(err)
		}
		//data := CommonResp{
		//	Code:     "200",
		//	Message:  "OK",
		//	TraceID:  "LC0",
		//	ErrorMsg: "no error",
		//}
		data := CommonResp{
			Code:    "0",
			Message: "OK",
			Data: []MesVehicle{{
				AutoID:    "DJ2022A",
				Vin:       "DJ2022A",
				Station:   "DJ2022A",
				Job:       "1",
				Configure: "DJ2022A",
			},
			},
		}
		//consumer.JSON(http.StatusOK, data)
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8080")
}

type CommonResp struct {
	Code     string       `json:"code"`
	Message  string       `json:"message"`
	Data     []MesVehicle `json:"data"`
	TraceID  string       `json:"traceID"`
	ErrorMsg string       `json:"errorMsg"`
}

type MesVehicle struct {
	AutoID    string `json:"Auto_ID"` // pin
	Vin       string `json:"Auto_VIN"`
	Station   string `json:"Station"`
	Job       string `json:"Job"`
	Configure string `json:"Configure"`
}
