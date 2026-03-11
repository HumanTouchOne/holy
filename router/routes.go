package router

import (
	"github.com/HumanTouchOne/holy/engine"
	"github.com/HumanTouchOne/holy/handles"
	"github.com/HumanTouchOne/holy/hooks"
)

func LoadRouter(r engine.IRouter) {

	r.GET("/test", handles.Test)
	r.Rest("/user", handles.User{})
	r.Use(engine.PosAhead, engine.Prefix("/wx", []string{}), hooks.Authorize)
	r.SetGuard(engine.Cors(hooks.Authorize))

}
