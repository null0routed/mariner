package storage 

import (
	"github.com/go-sqlx/sqlx"
)

type VaultRepository struct {
	db *sqlx.DB
}

func NewVaultRepository(db *sqlx.DB) *VaultRepository {
	return &VaultRepository{db: db}
}