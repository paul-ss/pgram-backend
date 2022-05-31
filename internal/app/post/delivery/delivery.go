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

	r.POST("/post", d.CreatePost)
	r.POST("/post/get", d.GetFeed)
}

type Delivery struct {
	uc domain.PostUsecase
}

type CreateData struct {
	Data  *domain.PostCreate    `form:"data" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// CreatePost
// @Summary      create new post
// @Description  create new post
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        postCreate  body  domain.PostCreate  true  "CreatePost post"
// @Success      200  {object}  domain.Post
// @Failure      400  {object}  domain.ErrorBase
// @Failure      404  {object}  domain.ErrorBase
// @Failure      500  {object}  domain.ErrorBase
// @Router       /post [post]
func (d *Delivery) CreatePost(c *gin.Context) {
	var form CreateData

	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post, err := d.uc.Create(c.Request.Context(), form.Data, form.Image)
	if err != nil {
		// TODO handle error
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (d *Delivery) GetFeed(c *gin.Context) {
	var req domain.FeedGet

	if err := c.BindJSON(&req); err != nil {
		log.Error(err.Error())
		return
	}

	posts, err := d.uc.GetFeed(c.Request.Context(), &req)
	if err != nil {
		// TODO handle error
		log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}
