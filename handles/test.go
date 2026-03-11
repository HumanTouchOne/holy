package handles

import "github.com/HumanTouchOne/holy/engine"

func Test(c *engine.Context) {
	c.String(200, "test")
}
