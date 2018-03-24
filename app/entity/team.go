package entity

type teamEntity struct{}
type TeamEntity interface {
	GetInfo() error
}

func NewTeamEntity() TeamEntity {
	return &teamEntity{}
}

func (t *teamEntity) GetInfo() error {
	return nil
}
