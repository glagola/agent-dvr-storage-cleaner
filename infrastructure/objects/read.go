package objects

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type repository struct {
	pathToFile string
}

var _ ReadRepository = (*repository)(nil)

type xmlObject struct {
	ID        int    `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Directory string `xml:"directory,attr"`
}

type xmlObjectsConfig struct {
	Cameras     []xmlObject `xml:"cameras>camera"`
	Microphones []xmlObject `xml:"microphones>microphone"`
}

func (self *repository) decode() (*xmlObjectsConfig, error) {
	var res *xmlObjectsConfig

	f, err := os.Open(self.pathToFile)

	if err != nil {
		return nil, fmt.Errorf("Failed to read objects from %s, reason: %w", self.pathToFile, err)
	}

	defer f.Close()

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	err = decoder.Decode(&res)

	if err != nil {
		return nil, fmt.Errorf("Failed to read objects from %s, reason: %w", self.pathToFile, err)
	}

	return res, nil
}

func (self *xmlObject) toMediaSource() MediaSource {
	return MediaSource{
		Id:        self.ID,
		Name:      self.Name,
		Directory: self.Directory,
	}
}

func (self *repository) listMediaSources(getter func(config *xmlObjectsConfig) []xmlObject) ([]MediaSource, error) {
	rawObjects, err := self.decode()

	if err != nil {
		return nil, err
	}

	rawMediaSource := getter(rawObjects)

	res := make([]MediaSource, 0, len(rawMediaSource))

	for _, rawObj := range rawMediaSource {
		res = append(res, rawObj.toMediaSource())
	}

	return res, nil
}

func (self *repository) Cameras() ([]MediaSource, error) {
	res, err := self.listMediaSources(func(config *xmlObjectsConfig) []xmlObject { return config.Cameras })

	if err != nil {
		return nil, fmt.Errorf("Failed to get cameras, reason: %w", err)
	}

	return res, nil
}

func (self *repository) Microphones() ([]MediaSource, error) {
	res, err := self.listMediaSources(func(config *xmlObjectsConfig) []xmlObject { return config.Microphones })

	if err != nil {
		return nil, fmt.Errorf("Failed to get microphones, reason: %w", err)
	}

	return res, nil
}

func NewReadRepository(pathToXmlFile string) ReadRepository {
	return &repository{
		pathToFile: pathToXmlFile,
	}
}
