package databases

// CREATE TABLE `person` (
//   `id` int(11) NOT NULL AUTO_INCREMENT,
//   `first_name` varchar(40) NOT NULL DEFAULT '',
//   `last_name` varchar(40) NOT NULL DEFAULT '',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SQLDB ...
var SQLDB *sql.DB

func init() {
	var err error

	SQLDB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/person?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer SQLDB.Close()

	// 数据库的空闲连接和最大打开连接 https://www.jianshu.com/p/a3f63b5da74c
	SQLDB.SetMaxOpenConns(20)
	SQLDB.SetMaxIdleConns(20)

	if err := SQLDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
