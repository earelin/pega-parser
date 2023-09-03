package config

import "github.com/earelin/pega/tools/infoelectoral/pkg/repository"

type Config struct {
	FilePath         string
	RepositoryConfig repository.Config
}
