package contracts

type EventDeletionRepositoryContract interface {
	Make(id string) error
}
