package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	"mime/multipart"
	"net/http"
)

type Delivery struct {
	uc domain.PostUsecase
}

type CreateData struct {
	Data  *domain.PostStore     `form:"data" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (d *Delivery) Create(c *gin.Context) {
	var form CreateData

	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := d.uc.Store(c.Request.Context(), form.Data, form.Image); err != nil {
		// TODO handle error
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (d *Delivery) Get(c *gin.Context) {
	var req domain.PostGet

	if err := c.BindJSON(&req); err != nil {
		return
	}

	posts, err := d.uc.Get(c.Request.Context(), &req)
	if err != nil {
		// TODO handle error
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}
