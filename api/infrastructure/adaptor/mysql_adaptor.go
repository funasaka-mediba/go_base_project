package adaptor

import (
	"context"
	"fmt"
	"go_base_project/constant"
	"go_base_project/packages/customError"

	"github.com/jmoiron/sqlx"
)

var readDb DBAdaptor
var writeDb DBAdaptor

func ReadDb(ctx context.Context) (DBAdaptor, *customError.CustomError) {
	if readDb == nil {
		s := fmt.Sprintf(
			// "%s:%s@tcp(%s:%s)/%s",
			// env.Env().MysqlUser,
			// env.Env().MysqlPass,
			// env.Env().MysqlReadHost,
			// env.Env().MysqlPort,
			// env.Env().MysqlDbName,
			"hogehoge", // 後で消す
		)
		db, err := sqlx.Connect("mysql", s)
		if err != nil {
			return nil, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
		}
		// db.SetMaxIdleConns(env.Env().MysqlMaxIdleConns)
		// db.SetMaxOpenConns(env.Env().MysqlMaxOpenConns)
		// db.SetConnMaxLifetime(env.Env().MysqlMaxLifetime)
		readDb = &basedb{sd: db}
	}
	return readDb, nil
}

func WriteDb(ctx context.Context) (DBAdaptor, *customError.CustomError) {
	if writeDb == nil {
		s := fmt.Sprintf(
			// "%s:%s@tcp(%s:%s)/%s",
			// env.Env().MysqlUser,
			// env.Env().MysqlPass,
			// env.Env().MysqlReadHost,
			// env.Env().MysqlPort,
			// env.Env().MysqlDbName,
			"hogehoge", // 後で消す
		)
		db, err := sqlx.Connect("mysql", s)
		if err != nil {
			return nil, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
		}
		// db.SetMaxIdleConns(env.Env().MysqlMaxIdleConns)
		// db.SetMaxOpenConns(env.Env().MysqlMaxOpenConns)
		// db.SetConnMaxLifetime(env.Env().MysqlMaxLifetime)
		writeDb = &basedb{sd: db}
	}
	return writeDb, nil
}
