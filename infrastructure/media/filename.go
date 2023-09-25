package media

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type fileNameInfo struct {
	Name      string
	ObjectId  int
	CreatedAt time.Time
	orig      string
}

const (
	groupName          = 1
	groupID            = 2
	groupCreatedAt     = 3
	groupCreatedAtMili = 4
)

var fileNameRegExp *regexp.Regexp = regexp.MustCompile(`^((\d+)_(\d+-\d+-\d+_\d+-\d+-\d+)_(\d+)(_[^.]+)?)\.[^.]+$`)

func parseFileName(fileName string) (res fileNameInfo, err error) {
	matched := fileNameRegExp.FindStringSubmatch(fileName)

	if matched == nil {
		return res, fmt.Errorf("Failed to parse file name with pattern, file name = '%v'", fileName)
	}

	oid, err := strconv.Atoi(matched[groupID])
	if err != nil {
		return res, fmt.Errorf("Failed to parse ID from file name: '%v', reason: %w", fileName, err)
	}

	createdAt, err := time.Parse("2006-01-02_15-04-05.000", fmt.Sprintf("%s.%s", matched[groupCreatedAt], matched[groupCreatedAtMili]))

	if err != nil {
		return res, fmt.Errorf("Failed to parse DateTime from file name: '%v', reason: %w", fileName, err)
	}

	return fileNameInfo{
		ObjectId:  oid,
		CreatedAt: createdAt,
		Name:      matched[groupName],
		orig:      fileName,
	}, nil
}

func (self *fileNameInfo) String() string {
	return self.orig
}
