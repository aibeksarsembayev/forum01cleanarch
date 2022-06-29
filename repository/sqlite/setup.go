package sqlite

import (
	"database/sql"
	"io/ioutil"
	"strings"
)

func Setup(db *sql.DB, path string) error {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
