package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/realcp1018/tinylog"
	"os"
	"test/ent"
	"time"
)

var Logger = tinylog.NewStreamLogger(tinylog.INFO)

func main() {
	dsnStr := "user:password@tcp(host:port)/<database>?parseTime=True"
	drv, err := sql.Open("mysql", dsnStr)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	// 相比于直接Open DSN获取连接，封装entgo.io/ent/dialect/sql.*DB显然更加适合生产环境，即有连接池的支持，又可以支持执行裸SQL
	client := ent.NewClient(ent.Driver(drv))
	//client, _ := ent.Open("mysql", dsnStr)
	defer client.Close()
	_ = RawSQLQueryUser(context.Background(), client)
}

func RawSQLQueryUser(ctx context.Context, client *ent.Client) error {
	q := "select * from users"
	result, err := client.QueryContext(ctx, q)
	if err != nil {
		return err
	}
	// QueryContext
	for result.Next() {

	}
	return nil
}
