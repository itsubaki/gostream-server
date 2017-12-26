package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
	uuid "github.com/satori/go.uuid"
)

func main() {
	c := NewConfig()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)

	// select count(*) from LogEvent(10sec) where Level > 2
	w := gocep.NewTimeWindow(10 * time.Second)
	w.SetSelector(gocep.EqualsType{Accept: LogEvent{}})
	w.SetSelector(gocep.LargerThanInt{Name: "Level", Value: 2})
	w.SetFunction(gocep.Count{As: "count(*)"})
	gost.Register("/", w)

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
		event.ID = uuid.NewV4().String()

		uri := c.Request.RequestURI
		w, err := gost.Window(uri)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	gost.engine.GET("/", func(c *gin.Context) {
		uri := c.Request.RequestURI
		w, err := gost.Window(uri)
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
