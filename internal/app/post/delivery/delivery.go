package delivery

import "github.com/gin-gonic/gin"

type Delivery struct {
}

func (d *Delivery) Get(c *gin.Context) {
	c.Request.Context()
}
