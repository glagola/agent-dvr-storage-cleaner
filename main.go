package main

import (
	"fmt"
	"log"
	"myapp/infrastructure/media"
	"myapp/infrastructure/objects"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PathToObjectsXML  string `env:"PATH_TO_OBJECTS_XML" env-required:"true"`
	PathToMediaFolder string `env:"PATH_TO_MEDIA_FOLDER" env-required:"true"`
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

	for _, cam := range res {
		files, err := mediaRepo.Videos(cam)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%+v\n", files)
	}
}
