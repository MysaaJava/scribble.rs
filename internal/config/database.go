package config

import (
	"database/sql"
	
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"	
)

type Word struct {
	bun.BaseModel `bun:"table:words,alias:w"`

	ID int64 `bun:"id,pk,autoincrement"`
	Str string `bun:"str,notnull"`
	WordGroup int64 `bun:"groupid,rel:belongs-to,join:authorid=id"`
}

type WordGroup struct {
	bun.BaseModel `bun:"table:wordGroups,alias:wg"`

	ID int64 `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
	ParentID *WordGroup `bun:"parentid,rel:belongs-to,join:parentid=id"`
}

func initDatabase(config *Config) (*bun.DB) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithAddr(config.Database.Addr),
		pgdriver.WithUser(config.Database.User),
		pgdriver.WithPassword(config.Database.Password),
		pgdriver.WithDatabase(config.Database.Database),
		pgdriver.WithApplicationName(config.Database.AppName),
		pgdriver.WithTimeout(config.Database.Timeout),
		pgdriver.WithDialTimeout(config.Database.DialTimeout),
		pgdriver.WithReadTimeout(config.Database.ReadTimeout),
		pgdriver.WithWriteTimeout(config.Database.WriteTimeout),
	)
	sqldb := sql.OpenDB(pgconn)
	db := bun.NewDB(sqldb, pgdialect.New())
	
	return db
}
