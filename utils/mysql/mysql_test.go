package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

var Db *sqlx.DB

func Init() {
	database, err := sqlx.Open("mysql", "root:leiyadi@tcp(81.70.205.15:3306)/gotool")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func TestMYSQLT(t *testing.T) {
	Init()
	// 成功插入
	r, err := Db.Exec("insert into test(id,name)values(?, ?)", 2, "test_gaga")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
