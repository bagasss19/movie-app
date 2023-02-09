package repository

import "gorm.io/gorm"

const (
	packagePrefix = "repository.postgres.movie"
)

type DB struct {
	repoPrefix string
	db         *gorm.DB
}
