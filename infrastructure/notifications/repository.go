package notifications

import (
	"database/sql"
	"fmt"
)

type repository struct {
	connection *sql.DB
}

func NewRepository(connection *sql.DB) Repository {
	return &repository{
		connection: connection,
	}
}

func (repo repository) Clear(objectId int) error {
	tables := []string{
		"AgentFile",
		"AgentAlert",
	}

	for _, table := range tables {
		// TODO wrap in transaction
		sql := fmt.Sprintf("DELETE FROM %s WHERE ObjectId = ?", table)

		if _, err := repo.connection.Exec(sql, objectId); err != nil {
			return fmt.Errorf("failed to clear %s table for ObjectId = %d: %w", table, objectId, err)
		}
	}

	return nil
}
