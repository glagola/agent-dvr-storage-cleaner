package media

import (
	"fmt"
	"path"
)

const (
	videoExt = ".mkv"
	thumbExt = ".jpg"
)

type storageFiles string

func (self storageFiles) videoPattern(directory string) string {
	return path.Join(self.pathToVideoFolder(directory), "*"+videoExt)
}

func (self storageFiles) pathToVideoFolder(directory string) string {
	return path.Join(string(self), "video", directory)
}

func (self storageFiles) pathToVideoThumbs(directory string) string {
	return path.Join(self.pathToVideoFolder(directory), "thumbs")
}

func (self storageFiles) pathToVideoThumbnails(directory string, fileName string) (small string, large string) {
	thumbs := self.pathToVideoThumbs(directory)

	small = path.Join(thumbs, fileName+thumbExt)
	large = path.Join(thumbs, fmt.Sprintf("%s_large%s", fileName, thumbExt))

	return
}
