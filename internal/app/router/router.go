package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/configs"
	"github.com/taniwhy/ithub-backend/internal/app/datastore"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/service"
	"github.com/taniwhy/ithub-backend/internal/app/handler"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/cors"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
)

// Init : Init関数は依存性の注入とURLパスルーティングを行います
func Init(dbConn *gorm.DB) *gin.Engine {

	userDatastore := datastore.NewUserDatastore(dbConn)
	tagDatastore := datastore.NewTagDatastore(dbConn)
	userService := service.NewUserService(userDatastore)
	authHandler := handler.NewAuthHandler(userDatastore, userService)
	userHandler := handler.NewUserHandler(userDatastore)
	tagHandler := handler.NewTagHandler(tagDatastore)

	dbConn.LogMode(true)

	r := gin.Default()
	store := cookie.NewStore([]byte(configs.SecretKey))
	r.Use(sessions.Sessions("_session", store))
	r.Use(cors.Write())

	var test string

	type body struct {
		Mark string `json:"mark" validate:"required"`
	}

	r.GET("", func(c *gin.Context) {
		c.JSON(200, test)
	})

	r.POST("", func(c *gin.Context) {
		body := body{}
		_ = c.Bind(&body)
		test = body.Mark
		c.JSON(200, gin.H{
			"message": test,
		})
	})

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	tags := v1.Group("/tags")
	static := r.Group("static")
	images := static.Group("/images")
	{
		auth.POST("/google/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)

		v1.GET("/me", userHandler.GetMe)
		users.GET("/:name", userHandler.GetByName)
		users.PUT("/", userHandler.Update)
		users.DELETE("/", userHandler.Delete)

		tags.GET("/", tagHandler.GetList)
		tags.POST("/", tagHandler.Create)

		images.POST("/upload", func(c *gin.Context) {
			file, _ := c.FormFile("image")
			src, _ := file.Open()
			defer src.Close()

			req, _ := http.NewRequest("POST", "https://api.imgur.com/3/image", src)
			req.Header.Add("Authorization", "Client-ID "+"d399235f8835bde")
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			client := new(http.Client)
			resp, _ := client.Do(req)
			body, _ := ioutil.ReadAll(resp.Body)

			imgurRes := model.ImgurResponse{}
			if err := json.Unmarshal(body, &imgurRes); err != nil {
				response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
				return
			}

			c.JSON(200, gin.H{"link": imgurRes.Data.Link})
		})
	}
	return r
}
