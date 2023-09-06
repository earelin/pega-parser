package config

import "github.com/earelin/pega/tools/inebase/pkg/repository"

type Config struct {
	FilePath         string
	DataSet          string
	RepositoryConfig repository.Config
}
