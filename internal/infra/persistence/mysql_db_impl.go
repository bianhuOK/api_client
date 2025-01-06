package persistence

import (
	"database/sql"
	"sync"

	repo "github.com/bianhuOK/api_client/internal/infra/iface"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDbRepo struct {
	db *sql.DB
}

var (
	mysqlConnPool sync.Map
)

// NewMySQLDBRepository 创建一个新的 MySQL 数据库存储库。
//
// 参数：
// dsn: 数据源名称，用于连接到 MySQL 数据库。
//
// 返回值：
// repository.DBRepository: 返回一个新的 MySQL 数据库存储库实例。
// error: 如果创建存储库时发生错误，则返回错误。
func NewMySQLDBRepository(dsn string) (repo.DbRepository, error) {
	// 检查是否已存在该 DSN 的连接池
	if db, ok := mysqlConnPool.Load(dsn); ok {
		return &MysqlDbRepo{db: db.(*sql.DB)}, nil
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 设置连接池参数 (根据需要调整) todo
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60 * 60) // 5小时

	mysqlConnPool.Store(dsn, db)
	return &MysqlDbRepo{db: db}, nil
}

func (r *MysqlDbRepo) ExecuteSql(query string) ([]map[string]interface{}, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		for i, col := range columns {
			row[col] = *(values[i].(*interface{}))
		}
		results = append(results, row)
	}
	return results, nil
}
