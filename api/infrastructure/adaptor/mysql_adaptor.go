package adaptor

import (
	"context"
	"fmt"
	"go_base_project/constant"
	"go_base_project/packages/customError"
	"go_base_project/packages/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var readDb DBAdaptor
var writeDb DBAdaptor

// 読み込みようにレプリカインスタンスを用意できるようになったら使用。
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

// 基本はこちらのみ使用で良い。
func WriteDb(ctx context.Context) (DBAdaptor, *customError.CustomError) {
	if writeDb == nil {
		s := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			env.Env().DbMysqlUser,
			env.Env().DbMysqlPassword,
			env.Env().DbMysqlWriteHost,
			env.Env().DbMysqlPort,
			env.Env().DbMysqlDatabase,
		)
		db, err := sqlx.Connect("mysql", s)
		if err != nil {
			return nil, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
		}
		db.SetMaxIdleConns(env.Env().DbMysqlMaxIdleConns)
		db.SetMaxOpenConns(env.Env().DbMysqlMaxOpenConns)
		db.SetConnMaxLifetime(env.Env().DbMysqlMaxLifetime)
		writeDb = &basedb{sd: db}
	}
	return writeDb, nil
}
