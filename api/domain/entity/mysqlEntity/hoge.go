package mysqlEntity

type Hoge struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
