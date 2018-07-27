package xsql2

/*
*/
import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"github.com/NiuStar/log"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	//"sync"
	"time"
)

const LifeTime int64 = 60 * 60

type XSql2 struct {
	db        *sql.DB
	name      string
	password  string
	ip        string
	port      string
	sqlName   string
	stmts	  map[string]*sql.Stmt
}

type XSql2Order struct {
	tables 		[]XSqlObject
	fields 		[]*XSqlParam
	sets			[]string
	values			[][]string
	conditions 	[]*condition

	orderbys	[]string
	or 			[]string
	xsql2		*XSql2
}

type XSqlObject interface {
	GetName() string
}

type XSqlParam struct {
	Name string
	Type_ string
	Target XSqlObject
}

type condition struct {

	type_ int //0 AND 1 OR
	value string
}
func checkErr(err error) {
	if err != nil {
		log.Write(err)
	}
}

func InitSql(name string, password string, ip string, port string, sqlName string,charset string) *XSql2 {
	db := connectDB(name, password, ip, port, sqlName,charset)
	fmt.Println("初始化数据库成功")
	s := new(XSql2)
	s.db = db
	s.name = name
	s.password = password
	s.ip = ip
	s.port = port
	s.sqlName = sqlName
	s.stmts = make(map[string]*sql.Stmt)

	return s
}

func connectDB(name string, password string, ip string, port string, sqlName string,charset string) *sql.DB {
	db, err := sql.Open("mysql", name+":"+password+"@tcp("+ip+":"+port+")/"+sqlName+"?charset=" + charset)

	checkErr(err)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(10 * time.Minute)
	err = db.Ping()

	checkErr(err)

	return db
}

func (x *XSql2)Table(obj... XSqlObject) *XSql2Order {
	return &XSql2Order{tables:obj,xsql2:x}
}

func (order *XSql2Order)Field(obj... *XSqlParam) *XSql2Order {

	fmt.Println("obj:",obj)
	order.fields = append(order.fields,obj...)
	return order
}

func getInitValue(pval []byte) interface{} {
	result_int, ok := ParseInt(pval)
	if !ok {
		result_float, ok := ParseFloat(pval)
		if !ok {
			fmt.Println("string")
			return string(pval)
		}
		fmt.Println("float")
		return result_float
	} else {
		s := string(pval)
		a := strings.Split(s, "0")
		if strings.EqualFold(a[0], "") {
			return string(pval)
		}
		fmt.Println("int")
		return result_int
	}
}


func byte2Int(value []byte) int64 {

	result, err := strconv.ParseInt(string(value), 10, 64)
	checkErr(err)
	return result
}
func byte2Float(value []byte) float64 {

	result, err := strconv.ParseFloat(string(value), 64)
	checkErr(err)
	return result
}

func byte2String(value []byte) string {
	return string(value)
}


func ParseInt(value []byte) (int64, bool) {
	result, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		return 0, false
	}
	return result, true
}

func ParseFloat(value []byte) (float64, bool) {
	result, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return 0, false
	}
	return result, true
}

func splicOrder(prefix,subfix ,midfix string,list []string) string {

	var orderstring string = prefix
	for index,order := range list {
		orderstring += order
		if index != len(list) - 1 {
			orderstring += midfix
		}
	}
	orderstring += subfix
	return orderstring
}