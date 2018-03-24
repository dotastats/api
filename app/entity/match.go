package entity

type matchEntity struct{}
type MatchEntity interface {
	GetList() error
}

func NewMatchEntity() MatchEntity {
	return &matchEntity{}
}

func (t *matchEntity) GetList() error {
	return nil
}
