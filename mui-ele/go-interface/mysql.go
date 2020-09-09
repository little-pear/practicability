package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 指定列名是否区分大小写
const ColumnNameIgnoreCase = true

//DBRow 查询结果行数据信息（仅包括数据，不包括列名或者列序号）
type DBRow []string

//SqlDB 对mysql数据库对象的封装
type SqlDB struct {
	*sql.DB
}

//SqlDBRows 查询结果集对象
type SqlDBRows struct {
	columns  map[string]int //列名集合
	dbResult []DBRow        //查询结果数据集
	index    int            //当前查询结果遍历时的“指针”
	rowSize  int            //查询结果的总行数
}

//ResNum 查询结果的总行数
func (rows *SqlDBRows) ResNum() int {
	return rows.rowSize
}

//Colmns 返回数据库列名集合
func (rows *SqlDBRows) Colmns() map[string]int {
	return rows.columns
}

//NextT 用于遍历查询结果集，将数据指向下一条结果（行数据）
func (rows *SqlDBRows) NextT() bool {
	if rows.index+1 >= rows.rowSize || rows.rowSize == 0 {
		return false
	}
	rows.index++
	return true
}

//GetField 根据列名获取该列的“值数据”
func (rows *SqlDBRows) GetField(s string) string {
	if rows.index >= rows.rowSize || rows.rowSize == 0 {
		return string("")
	}

	//判断是否列名区分大小写
	colName := s
	if ColumnNameIgnoreCase {
		colName = strings.ToLower(s)
	}

	if v, ok := rows.columns[colName]; ok {
		return rows.dbResult[rows.index][v]
	}

	return string("")
}

//IsFieldExist 是否字段存在
func (rows *SqlDBRows) IsFieldExist(colName string) bool {
	if nil != rows.columns {
		if ColumnNameIgnoreCase {
			colName = strings.ToLower(colName)
		}
		_, ok := rows.columns[colName]
		return ok
	}
	return false
}

//Select 执行查询操作，并返回结果集
func (db *SqlDB) Select(sqlCmd string) (SqlRows *SqlDBRows, Error error) {

	if len(sqlCmd) == 0 {
		return nil, errors.New("sqlCmd is nil")
	}

	resRows := &SqlDBRows{
		dbResult: make([]DBRow, 0, 10),
		columns:  make(map[string]int, 0),
		index:    -1,
		rowSize:  0,
	}
	rows, err := db.Query(sqlCmd)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//构造列名数组
	col, err := rows.Columns()
	if err != nil {
		fmt.Println("error is ", err.Error())
		return nil, err
	}
	if ColumnNameIgnoreCase {
		for k, v := range col {
			resRows.columns[strings.ToLower(v)] = k
		}
	} else {
		for k, v := range col {
			resRows.columns[v] = k
		}
	}

	values := make([]sql.RawBytes, len(col))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		var res []string
		if err != nil {
			return nil, err
		}
		for _, v := range values {
			if v != nil {
				res = append(res, string(v))
			} else {
				res = append(res, "")
			}
		}
		resRows.rowSize++
		resRows.dbResult = append(resRows.dbResult, res)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows num is null errrrrr")
		resRows.rowSize = 0
		return nil, err
	}
	return resRows, nil
}

//Insert 执行插入操作
func (db *SqlDB) Insert(sqlCmd string) (int64, int64, error) {

	if len(sqlCmd) == 0 {
		return 0, 0, errors.New("sqlcmd is nil")
	}

	res, err := db.Exec(sqlCmd)
	if err != nil {
		return 0, 0, err
	}
	if nil == res {
		return 0, 0, errors.New("insert result is nil.")
	}

	lastId, _ := res.LastInsertId()
	rowCnt, _ := res.RowsAffected()

	return lastId, rowCnt, nil
}

//Update 执行更新操作,返回值为 LastInsertId，RowsAffected，错误信息
//一般情况下，LastInsertId = 0，RowsAffected=影响的行数.
//如果更新的值与数据库中的值一致，则不计入RowsAffected，所以RowsAffected不能作为判断操作是否成功的标志。
func (db *SqlDB) Update(sqlCmd string) (int64, int64, error) {
	return db.Insert(sqlCmd)
}

//执行删除操作
func (db *SqlDB) Delete(sqlCmd string) (int64, int64, error) {
	return db.Insert(sqlCmd)
}

//TransactionDB 批量执行SQL语句（按事务执行）
func (db *SqlDB) TransactionDB(sqlCmd []string) error {

	if len(sqlCmd) == 0 {
		return errors.New("sqlCmd len is 0")
	}
	tx, err := db.Begin()

	if err != nil {
		return errors.New("begin failed")
	}

	for _, v := range sqlCmd {
		_, err := tx.Exec(v)

		if err != nil {
			tx.Rollback()
			return errors.New(v + " exec failed")
		}
	}
	tx.Commit()

	return nil
}

//OpenDB 打开DB
func OpenDB(mode string, sqlCmd string) (*SqlDB, error) {

	db, err := sql.Open(mode, sqlCmd)

	if err != nil {
		return nil, err
	}

	return &SqlDB{
		DB: db,
	}, nil

}

//SetConnsParam 设置数据库连接参数
func (db *SqlDB) SetConnsParam(maxOpenConns int, maxOpenIdleConns int) {
	if maxOpenConns != 0 {
		db.SetMaxOpenConns(maxOpenConns)
	} else {
		db.SetMaxOpenConns(1000)
	}

	if maxOpenIdleConns != 0 {
		db.SetMaxIdleConns(maxOpenIdleConns)
	} else {
		db.SetMaxIdleConns(10)
	}
}
