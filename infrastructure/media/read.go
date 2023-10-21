package media

import (
	"fmt"
	"myapp/infrastructure/objects"
	"os"
	"path/filepath"
)

type repository struct {
	pathToMediaFolder storageFiles
}

var _ ReadRepository = (*repository)(nil)

func (self repository) Videos(source objects.MediaSource) ([]Video, error) {
	videoFilePattern := self.pathToMediaFolder.videoPattern(source.Directory)

	matchedFiles, err := filepath.Glob(videoFilePattern)

	if err != nil {
		return nil, fmt.Errorf("Failed to find video files using pattern %v, reason: %w", videoFilePattern, err)
	}

	result := make([]Video, 0, len(matchedFiles))

	for _, pathToFile := range matchedFiles {
		info, err := parseFileName(filepath.Base(pathToFile))

		if err != nil || info.ObjectId != source.Id {
			fmt.Println(err)
			continue
		}

		small, large := self.pathToMediaFolder.pathToVideoThumbnails(source.Directory, info.Name)

		result = append(result, Video{
			mediaFile: mediaFile{
				PathToFile: pathToFile,
				CreatedAt:  info.CreatedAt,
			},
			Thumbnails: thumbnail{
				PathToLarge: large,
				PathToSmall: small,
			},
		})
	}

	return result, nil
}

func nilIfNoFile(pathToFile string) *string {
	if exists := fileExists(pathToFile); exists {
		return &pathToFile
	}

	return nil
}

func fileExists(pathToFile string) bool {
	file, err := os.Open(pathToFile)
	if err != nil {
		return false
	}
	defer file.Close()

	return true
}

func NewReadRepository(pathToMediaFolder string) ReadRepository {
	return &repository{
		pathToMediaFolder: storageFiles(pathToMediaFolder),
	}
}
