package main

import (
	"github.com/HumanTouchOne/holy/components/log"
	"github.com/HumanTouchOne/holy/engine"
	"github.com/HumanTouchOne/holy/router"
)

func main() {
	e := engine.New(func(eng *engine.Engine) {
		eng.SetAddr(":3003")
	})
	log.InitLogger([]string{"./logs/output.log"}, []string{"./logs/error.log"})
	router.LoadRouter(e)
	if err := e.Run(); err != nil {
		panic(err)
	}
}
