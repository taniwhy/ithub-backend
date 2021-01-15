package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

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
	noteDatastore := datastore.NewNoteDatastore(dbConn)
	tagDatastore := datastore.NewTagDatastore(dbConn)
	noteTagDatastore := datastore.NewnNoteTagDatastore(dbConn)
	followDatastore := datastore.NewFollowDatastore(dbConn)
	commentDatastore := datastore.NewCommentDatastore(dbConn)

	userService := service.NewUserService(userDatastore)

	authHandler := handler.NewAuthHandler(userDatastore, userService)
	userHandler := handler.NewUserHandler(userDatastore, followDatastore, noteDatastore)
	noteHandler := handler.NewNoteHandler(noteDatastore, userDatastore, tagDatastore, noteTagDatastore)
	tagHandler := handler.NewTagHandler(tagDatastore)
	followHandler := handler.NewFollowHandler(followDatastore, userDatastore)
	commentHandler := handler.NewCommentHandler(commentDatastore, userDatastore)

	dbConn.LogMode(true)

	r := gin.Default()

	store := cookie.NewStore([]byte(configs.SecretKey))

	if os.Getenv("GO_ENV") != "dev" {
		store.Options(sessions.Options{SameSite: http.SameSiteNoneMode, Secure: true})
	}

	r.Use(sessions.Sessions("_session", store))
	r.Use(cors.Write())

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	notes := v1.Group("/notes")
	tags := v1.Group("/tags")
	follows := v1.Group("/follows")
	comments := notes.Group("/:id/comments")
	static := r.Group("static")
	images := static.Group("/images")
	{
		auth.GET("/reflesh", authHandler.Reflesh)
		auth.POST("/google/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)

		v1.GET("/me", userHandler.GetMe)
		users.GET("/:name", userHandler.GetByName)
		users.PUT("", userHandler.Update)
		users.DELETE("", userHandler.Delete)

		users.GET("/:name/notes", noteHandler.GetListByID)
		notes.GET("/:id", noteHandler.GetByID)
		notes.POST("", noteHandler.Create)

		comments.GET("", commentHandler.GetByNoteID)
		comments.POST("", commentHandler.Create)

		users.GET("/:name/follows", followHandler.GetFollows)
		users.GET("/:name/followers", followHandler.GetFollowers)
		follows.GET("", followHandler.Create)
		follows.GET("/delete", followHandler.Delete)

		tags.GET("", tagHandler.GetList)
		tags.POST("", tagHandler.Create)

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
