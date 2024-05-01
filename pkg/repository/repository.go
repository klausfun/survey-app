package repository

type Authorization interface {
}

type Surveys interface {
}

type Users interface {
}

type Repository struct {
	Authorization
	Surveys
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}
