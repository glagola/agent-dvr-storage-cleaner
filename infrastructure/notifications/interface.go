package notifications

type Repository interface {
	Clear(objectId int) error
}
