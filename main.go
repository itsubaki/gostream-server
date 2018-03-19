package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
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

	gost.POST("/", func(c *gin.Context) {
		event, err := NewLogEvent(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w, err := gost.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	gost.GET("/", func(c *gin.Context) {
		w, err := gost.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		events := <-w.Output()
		c.JSON(200, gocep.Newest(events))
	})

	if err := gost.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
