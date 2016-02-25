package storage

import (
	"github.com/adamdecaf/horizon/utils"
)

func ExecuteGzippedSQL(filepath string) (int64, error) {
	blob, err := utils.GzipDecompressFile(filepath)
	if err != nil {
		return 0, err
	}

	db, err := InitializePostgres()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(blob)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
