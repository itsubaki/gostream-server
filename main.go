package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
)

type LogEvent struct {
	Time    time.Time `json:"time"`
	Level   int       `json:"level"`
	Message string    `json:"message"`
}

func main() {
	c := NewConfig()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)

	// select count(*) from LogEvent(10sec) where Level > 2
	w := gocep.NewTimeWindow(10 * time.Second)
	w.SetSelector(gocep.EqualsType{Accept: LogEvent{}})
	w.SetSelector(gocep.LargerThanInt{Name: "Level", Value: 2})
	w.SetFunction(gocep.Count{As: "count(*)"})
	gost.window.Put("/log/count", w)

	gost.engine.POST("/", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		var event LogEvent
		if unerr := json.Unmarshal(b, &event); unerr != nil {
			c.JSON(400, unerr.Error())
			return
		}

		w, err := gost.window.Get("/log/count")
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, "ok")
	})

	gost.engine.GET("/", func(c *gin.Context) {
		w, err := gost.window.Get("/log/count")
		if err != nil {
			c.JSON(400, err)
			return
		}

		events := <-w.Output()
		c.JSON(200, gocep.Newest(events))
	})

	log.Println(time.Now())
	gost.Run()
}
