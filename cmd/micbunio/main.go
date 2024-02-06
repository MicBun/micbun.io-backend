package main

import (
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog/repo"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
	pb "github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	twirpHandler pb.TwirpServer
	redisDB      *db.Redis
	postgres     *gorm.DB
}

func NewApp(
	twirpHandler pb.TwirpServer,
	redisDB *db.Redis,
	postgres *gorm.DB,
) *App {
	return &App{
		twirpHandler: twirpHandler,
		redisDB:      redisDB,
		postgres:     postgres,
	}
}

func (a *App) ServeWeb() error {
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := corsWrapper.Handler(a.twirpHandler)

	return errors.WithStack(http.ListenAndServe(":8080", handler))
}

func initApp() *App {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Server failed to load env file", err.Error())
		}
	}

	redisDB := db.NewRedisManager(db.NewRedis())
	postgres, err := db.NewPostgres()
	if err != nil {
		log.Fatalln("Server failed to connect postgres", err.Error())
	}
	blogRepo := repo.NewBlog(postgres)
	commentRepo := repo.NewComment(postgres)
	blogManager := blog.NewPbBlogServer(redisDB, blogRepo, commentRepo)
	twirpHandler := pb.NewBlogServiceServer(blogManager)

	return NewApp(
		twirpHandler,
		redisDB,
		postgres,
	)
}

func main() {
	log.Println("Server is running on port 8080")

	if err := initApp().ServeWeb(); err != nil {
		log.Fatalln("Server failed to start", err.Error())
	}
}
