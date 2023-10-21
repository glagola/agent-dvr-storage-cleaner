package main

import (
	"database/sql"
	"log"
	"myapp/infrastructure/media"
	"myapp/infrastructure/notifications"
	"myapp/infrastructure/objects"

	"github.com/ilyakaznacheev/cleanenv"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	PathToObjectsXML     string `env:"PATH_TO_OBJECTS_XML" env-required:"true"`
	PathToMediaFolder    string `env:"PATH_TO_MEDIA_FOLDER" env-required:"true"`
	PathToNotificationDB string `env:"PATH_TO_NOTIFICATION_DB" env-required:"true"`
}

var cfg Config

func main() {
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalln(err)
	}
	objRepo := objects.NewReadRepository(cfg.PathToObjectsXML)

	res, err := objRepo.Cameras()

	if err != nil {
		log.Fatalln(err)
	}

	mediaRepo := media.NewReadRepository(cfg.PathToMediaFolder)

	db, err := sql.Open("sqlite3", cfg.PathToNotificationDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	notificationsRepo := notifications.NewRepository(db)

	for _, cam := range res {
		log.Printf("Camera: %s", cam.Name)

		files, err := mediaRepo.Videos(cam)

		if err != nil {
			log.Fatalln(err)
		}

		for _, file := range files {
			log.Printf("    %s\n", file.PathToFile)
		}

		if err = notificationsRepo.Clear(cam.Id); err != nil {
			log.Fatal(err)
		}
	}
}
