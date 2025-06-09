package migrations

import (
	"pokemon/internal/domains/user"
)

func GetAllMigrators() []Migrator {
	return []Migrator{
		user.UserMigrator{},
	}
}
