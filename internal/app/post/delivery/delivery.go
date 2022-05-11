package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	"github.com/paul-ss/pgram-backend/internal/app/post/usecase"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

func Register(r domain.GinRouter) {
	d := &Delivery{
		uc: usecase.NewUsecase(),
	}

	r.POST("/post", d.Create)
	r.POST("/post/get", d.Get)
}

type Delivery struct {
	uc domain.PostUsecase
}

type CreateData struct {
	Data  *domain.PostStoreUC   `form:"data" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (d *Delivery) Create(c *gin.Context) {
	var form CreateData

	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post, err := d.uc.Store(c.Request.Context(), form.Data, form.Image)
	if err != nil {
		// TODO handle error
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (d *Delivery) Get(c *gin.Context) {
	var req domain.PostGetUC

	if err := c.BindJSON(&req); err != nil {
		log.Error(err.Error())
		return
	}

	posts, err := d.uc.Get(c.Request.Context(), &req)
	if err != nil {
		// TODO handle error
		log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}
