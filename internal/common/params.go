package common

import (
	"database/sql"

	"go.uber.org/dig"
)

// Params ...
type Params struct {
	// dig.Inを宣言することでパラメータオブジェクトになる
	dig.In

	DB *sql.DB
}
