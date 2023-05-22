package database

type bd struct {
}

func New() (*bd, error) {
	return &bd{}, nil
}
