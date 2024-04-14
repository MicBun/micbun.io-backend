package main

import (
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/repo"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	twirpHandler micbunio.TwirpServer
	redisDB      *db.Redis
	postgres     *gorm.DB
}

func NewApp(
	twirpHandler micbunio.TwirpServer,
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
	return errors.WithStack(http.ListenAndServe(":8080", cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(a.twirpHandler)))
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

	guestbookRepo := repo.NewGuestbook(postgres)
	service := micbunioserver.NewPbService(redisDB, guestbookRepo)

	return NewApp(
		micbunio.NewGuestbookServiceServer(service),
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
