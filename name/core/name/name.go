package name

import (
	//"free_go/name/storage"
	"free_go/name/storage/postgres"
	//"context"
)

type CoreSve struct {
	store *postgres.Storage
}

func NewCoreSve(b *postgres.Storage) *CoreSve {
	return &CoreSve{
		store: b,
	}
}