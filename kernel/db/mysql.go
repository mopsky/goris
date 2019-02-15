package db

/**
 *	数据库模型
 *  Created By YangXB 2019.1.15
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goris/conf/yaml"
	"github.com/goris/utils"
	"github.com/kataras/iris/core/errors"
	"strconv"
	"strings"
)

// 条件结构
type Where struct {
	Column string
	Symbol string
	Value  interface{}
}

type Model struct {
	sTable   string
	sFields  string
	sJoin    string
	sWhere   string        // where 条件
	sLimit   string        // limit 条件
	bindArgs []interface{} // 绑定变量参数
	sOrder   string
	sGroup   string
	sHaving  string
	sSQL     string //最后封装形成的SQL语句

	db         *sql.DB //数据库连接
	tx         *sql.Tx //事务管理器
	dbFieldsKV map[string]string
	dbFields   []string
	dbRows     *sql.Rows
	err        error //错误类
}

func (m *Model) SetTable(sTable string) *Model {
	m.sTable = sTable
	return m
}

func (m *Model) Open(sTable string) {
	Init(m, sTable)
}

//初始化数据库
func Init(m *Model, sTable string) {
	if sTable == "" {
		m.db = nil
		m.err = errors.New("表名不能为空")
		return
	}

	m.sTable = sTable
	// 初始化数据库 "pfws:pfws2016@tcp(192.168.1.231:3306)/pfws?charset=utf8"
	m.db, m.err = sql.Open("mysql", yaml.DB_SOURCE)
	if m.err != nil {
		return
	}

	//初始化字段
	m.sFields = "*"

	//UTF8
	//m.db.Query("SET NAMES UTF8")
	//m.db.SetMaxOpenConns(300)

	//获取表列
	m.dbRows, m.err = m.db.Query("SHOW COLUMNS FROM " + m.sTable)
	if m.dbRows != nil {
		defer m.dbRows.Close()
	}
	if m.err != nil {
		return
	}

	var sField, sType string
	var sNull, sKey, sDefault, sExtra sql.NullString
	m.dbFieldsKV = make(map[string]string)
	for m.dbRows.Next() {
		if m.err = m.dbRows.Scan(&sField, &sType, &sNull, &sKey, &sDefault, &sExtra); m.err != nil {
			return
		}
		m.dbFieldsKV[sField] = sType
		m.dbFields = append(m.dbFields, sField)
	}
}

//校验数据库连接
func CheckConn(m *Model) bool {
	if m.db == nil {
		m.err = errors.New("数据库未初始化")
		return false
	}

	return true
}

func M(sTable string) *Model {
	m := new(Model)
	Init(m, sTable)
	return m
}

/**封装SQL*/
func (m *Model) makeSQL(bFind bool) bool {
	// 已经生成
	//if m.sSQL != "" {
	//	return true
	//}

	if m.sTable == "" {
		m.err = errors.New("表名不能为空")
		return false
	}

	m.sSQL = "SELECT " + m.sFields + " FROM " + m.sTable + m.sJoin + m.sWhere + m.sGroup + m.sHaving + m.sOrder

	if bFind {
		m.sSQL += " LIMIT 1"
	} else {
		m.sSQL += m.sLimit
	}

	return true
}

func (m *Model) _clear() {
	m.sWhere = ""
	m.sGroup = ""
	m.sOrder = ""
	m.sLimit = ""
	m.sHaving = ""

	m.bindArgs = nil
}

func (m *Model) _where(v ...string) {
	if m.sWhere == "" {
		m.sWhere += " WHERE "
	} else {
		m.sWhere += " AND "
	}

	if len(v) == 1 {
		m.sWhere += v[0]
	} else {
		m.sWhere += v[0] + " " + v[1] + " ?"
	}
}

// 组装where条件
func (m *Model) Where(args ...interface{}) *Model {
	for _, v := range args {
		switch at := v.(type) {
		case string:
			m._where(at)
		case Where:
			m._where(at.Column, at.Symbol)
			m.bindArgs = append(m.bindArgs, at.Value)
		case []Where:
			for _, v := range at {
				m._where(v.Column, v.Symbol)
				m.bindArgs = append(m.bindArgs, v.Value)
			}
		default:
			m.err = errors.New("不支持的Where类型")
		}
	}

	return m
}

// 字段
func (m *Model) Fields(sFields string) *Model {
	m.sFields = sFields
	return m
}

// 限制
func (m *Model) Limit(limit ...int) *Model {
	if len(limit) == 0 {
		return m
	} else if len(limit) == 1 {
		m.sLimit = " LIMIT " + strconv.Itoa(limit[0])
	} else {
		m.sLimit = " LIMIT " + strconv.Itoa(limit[0]) + "," + strconv.Itoa(limit[1])
	}

	return m
}

// 排序
func (m *Model) Order(sOrder string) *Model {
	m.sOrder = " ORDER BY " + strings.Trim(sOrder, " ")
	return m
}

// 分组
func (m *Model) Group(sGroup string) *Model {
	m.sGroup = " GROUP BY " + strings.Trim(sGroup, " ")
	return m
}

// 分组筛选
func (m *Model) Having(sHaving string) *Model {
	m.sHaving = " HAVING" + strings.Trim(sHaving, " ")
	return m
}

// 内部查询
func (m *Model) _fetchData(bSQL, bFind bool) (map[int]map[string]string, error) {
	defer m._clear()

	if !CheckConn(m) {
		return nil, m.err
	}

	//生成SQL
	if bSQL && !m.makeSQL(bFind) {
		return nil, m.err
	}

	fmt.Println(m.sSQL)
	m.dbRows, m.err = m.db.Query(m.sSQL, m.bindArgs...)
	if m.dbRows != nil {
		defer m.dbRows.Close()
	}
	if m.err != nil {
		return nil, m.err
	}

	//读出查询出的列字段名
	cols, _ := m.dbRows.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	//遍历数据
	i := 0
	results := make(map[int]map[string]string)
	for m.dbRows.Next() { //循环，让游标往下推
		if m.err = m.dbRows.Scan(scans...); m.err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return nil, m.err
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}

	if len(results) > 0 {
		return results, nil
	} else {
		return nil, nil
	}
}

func (m *Model) Select() (map[int]map[string]string, error) {
	defer m._clear()
	return m._fetchData(true, false)
}

// 单条查询
func (m *Model) Find() (map[string]string, error) {
	defer m._clear()
	result, err := m._fetchData(true, true)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result[0], nil
	} else {
		return nil, nil
	}
}

// 获取字段
func (m *Model) GetField(sFields string) (map[string]map[string]string, error) {
	defer m._clear()
	if sFields != "" {
		m.sFields = sFields
	}

	//获取首字段
	var sFirstColumn string
	if m.sFields == "*" || m.sFields == "" {
		sFirstColumn = m.dbFields[0]
	} else {
		splitColumn := strings.Split(m.sFields, ",")
		if len(splitColumn) < 2 {
			m.err = errors.New("GetField至少需要两个字段")
			return nil, m.err
		}

		sFirstColumn = splitColumn[0]
	}

	result := make(map[string]map[string]string)
	res, err := m._fetchData(true, false)
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		key := v[sFirstColumn]
		delete(v, sFirstColumn)
		result[key] = v
	}

	return result, nil
}

// 返回记录条数
func (m *Model) Count() (int, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	m.sFields = "COUNT(0) AS CNT"

	res, err := m._fetchData(true, true)
	if err != nil {
		return -1, err
	}

	iCount, err := strconv.Atoi(res[0]["CNT"])
	if err != nil {
		return -1, err
	}

	return iCount, nil
}

// 返回最小值
func (m *Model) Min(sField string) (interface{}, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if sField == "" {
		m.err = errors.New("Min字段不能为空")
		return nil, m.err
	}

	m.sFields = "MIN(" + sField + ") AS MIN_VALUE"

	res, err := m._fetchData(true, true)
	if err != nil {
		return -1, err
	}

	return res[0]["MIN_VALUE"], nil
}

// 返回最大值
func (m *Model) Max(sField string) (interface{}, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if sField == "" {
		m.err = errors.New("Max字段不能为空")
		return nil, m.err
	}

	m.sFields = "MAX(" + sField + ") AS MAX_VALUE"

	res, err := m._fetchData(true, true)
	if err != nil {
		return -1, err
	}

	return res[0]["MAX_VALUE"], nil
}

// 返回平均值
func (m *Model) Avg(sField string) (interface{}, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if sField == "" {
		m.err = errors.New("Avg字段不能为空")
		return nil, m.err
	}

	m.sFields = "AVG(" + sField + ") AS AVG_VALUE"

	res, err := m._fetchData(true, true)
	if err != nil {
		return -1, err
	}

	return res[0]["AVG_VALUE"], nil
}

// 关联
func (m *Model) Join(sJoin string) *Model {
	m.sJoin += " " + strings.Trim(sJoin, " ")
	return m
}

// 单条新增
func (m *Model) Add(data map[string]string) (int64, error) {
	if !CheckConn(m) {
		return -1, m.err
	}

	if len(data) < 1 {
		m.err = errors.New("请传入需要新增的字段数据")
		return -1, m.err
	}

	sColumns, sValues, i := "", "", 0
	var excuteArgs []interface{}
	for k, v := range data {
		sColumns += k
		sValues += "?"
		if i++; i != len(data) {
			sColumns += ", "
			sValues += ", "
		}
		excuteArgs = append(excuteArgs, v)
	}

	m.sSQL = "INSERT INTO " + m.sTable + "(" + sColumns + ") VALUES (" + sValues + ")"
	res, err := m.db.Exec(m.sSQL, excuteArgs...)
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取新增的主见
	lastID, err := res.LastInsertId()
	if err != nil {
		m.err = err
		return -1, err
	}

	return lastID, nil
}

// 批量新增
func (m *Model) AddAll(datas []map[string]string) (int64, error) {
	if !CheckConn(m) {
		return -1, m.err
	}

	if len(datas) < 1 {
		m.err = errors.New("请传入需要新增的字段数据")
		return -1, m.err
	}

	// 为保证key的顺序
	aColumns, sColumns, sTemp, sValues, i := []string{}, "", "", "", 0
	for k, _ := range datas[0] {
		aColumns = append(aColumns, k)
		sColumns += k
		sTemp += "?"
		if i++; i != len(datas[0]) {
			sColumns += ", "
			sTemp += ", "
		}
	}

	var excuteArgs []interface{}
	for k, v := range datas {
		for _, value := range aColumns {
			excuteArgs = append(excuteArgs, v[value])
		}

		sValues += "(" + sTemp + ")"
		if k+1 < len(datas) {
			sValues += ", "
		}
	}

	m.sSQL = "INSERT INTO " + m.sTable + "(" + sColumns + ") VALUES " + sValues
	res, err := m.db.Exec(m.sSQL, excuteArgs...)
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取新增的主键
	lastID, err := res.LastInsertId()
	if err != nil {
		m.err = err
		return -1, err
	}

	return lastID, nil
}

// 保存数据
func (m *Model) Save(data map[string]string) (int64, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if m.sWhere == "" {
		m.err = errors.New("没有Where条件，不允许更新")
		return -1, m.err
	}

	if len(data) < 1 {
		m.err = errors.New("请传入需要更新的字段数据")
		return -1, m.err
	}

	var aColumns []string
	for k, _ := range data {
		aColumns = append(aColumns, k)
	}

	var excuteArgs []interface{}
	m.sSQL = "UPDATE " + m.sTable + " SET "
	for k, v := range aColumns {
		excuteArgs = append(excuteArgs, data[v])
		m.sSQL += v + " = ?"
		if k+1 < len(data) {
			m.sSQL += ", "
		}
	}

	m.sSQL += m.sWhere
	res, err := m.db.Exec(m.sSQL, excuteArgs...)
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取影响的记录数
	rows, err := res.RowsAffected()
	if err != nil {
		m.err = err
		return -1, err
	}

	return rows, nil
}

// 步进数据
func (m *Model) SetInc(sField string, iInc int64, args ...interface{}) (int64, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if args != nil {
		if args[0].(bool) && m.sWhere == "" {
			m.err = errors.New("没有Where条件，不允许更新")
			return -1, m.err
		}
	}

	var symbol string
	if iInc > 0 {
		symbol = " + ?"
	} else {
		symbol = " - ?"
	}

	m.sSQL = "UPDATE " + m.sTable + " SET " + sField + " = " + sField + symbol + m.sWhere
	res, err := m.db.Exec(m.sSQL, utils.AbsInt64(iInc))
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取影响的记录数
	rows, err := res.RowsAffected()
	if err != nil {
		m.err = err
		return -1, err
	}

	return rows, nil
}

// 删除数据
// args[0] bool true的时候没有where条件也可以删除
func (m *Model) Delete(args ...interface{}) (int64, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if args != nil {
		if args[0].(bool) && m.sWhere == "" {
			m.err = errors.New("没有Where条件，不允许删除")
			return -1, m.err
		}
	}

	m.sSQL = "DELETE FROM " + m.sTable + m.sWhere
	res, err := m.db.Exec(m.sSQL, m.bindArgs...)
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取影响的记录数
	rows, err := res.RowsAffected()
	if err != nil {
		m.err = err
		return -1, err
	}

	return rows, nil
}

// 数据查询
func (m *Model) Query(sSQL string, args ...interface{}) (map[int]map[string]string, error) {
	defer m._clear()

	if !CheckConn(m) {
		return nil, m.err
	}

	if sSQL == "" {
		m.err = errors.New("查询SQL不能为空")
		return nil, m.err
	}

	m.sSQL = sSQL
	m.bindArgs = args
	return m._fetchData(false, false)
}

// 数据执行
func (m *Model) Execute(sSQL string, args ...interface{}) (int64, error) {
	defer m._clear()

	if !CheckConn(m) {
		return -1, m.err
	}

	if sSQL == "" {
		m.err = errors.New("执行SQL不能为空")
		return -1, m.err
	}

	m.sSQL = sSQL
	res, err := m.db.Exec(m.sSQL, args)
	if err != nil {
		m.err = err
		return -1, err
	}

	//获取影响的记录数
	rows, err := res.RowsAffected()
	if err != nil {
		m.err = err
		return -1, err
	}

	return rows, nil
}

// 关闭数据库
func (m *Model) Close() {
	if !CheckConn(m) {
		return
	}

	m.db.Close()
}

// 获取错误信息
func (m *Model) GetError() error {
	return m.err
}

// 获取执行语句s
func (m *Model) GetSQL() string {
	return m.sSQL
}

// 获取列信息
func (m *Model) GetColumns() map[string]string {
	return m.dbFieldsKV
}

// 开启事务
func (m *Model) StartTrans() {
	if !CheckConn(m) {
		return
	}

	m.tx, m.err = m.db.Begin()
}

// 提交事务
func (m *Model) Commit() {
	if !CheckConn(m) {
		return
	}

	m.err = m.tx.Commit()
}

// 回滚事务
func (m *Model) Rollback() {
	if !CheckConn(m) {
		return
	}

	m.err = m.tx.Rollback()
}
