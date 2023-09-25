package media

import (
	"myapp/infrastructure/objects"
	"time"
)

type thumbnail struct {
	PathToSmall string
	PathToLarge string
}

type mediaFile struct {
	PathToFile string
	CreatedAt  time.Time
}

type Video struct {
	Thumbnails thumbnail

	mediaFile
}

type Audio struct {
	mediaFile
}

type ReadRepository interface {
	Videos(source objects.MediaSource) ([]Video, error)
}
