package storage

import (
	"news-mono/cmd/pkg/client/postgresql"
	"news-mono/cmd/pkg/logging"
)

type storage struct {
	client postgresql.Client
	logger *logging.Logger
}
