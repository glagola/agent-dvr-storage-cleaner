package objects

type MediaSource struct {
	Id        int
	Name      string
	Directory string
}

type ReadRepository interface {
	Cameras() ([]MediaSource, error)
	Microphones() ([]MediaSource, error)
}
