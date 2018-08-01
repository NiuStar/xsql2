package xsql2

import (
	"fmt"
	"bytes"
)


type XSql2Set interface {
	Set(value... interface{}) XSql2Set
	Update() bool
	Where(obj *XSqlParam, op string, v interface{}) XSql2Where
	WhereParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2Where
}


func (order *XSql2Order)Set(value... interface{}) XSql2Set {
		order.args = append(order.args,value...)
	return order
}

func (order *XSql2Order)Update() bool {

	if len(order.args) <= 0 || len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.sets) <= 0 || len(order.tables) <=0：",len(order.sets) ,len(order.tables) )
		return false
	}

	var sqlOrder  bytes.Buffer
	sqlOrder.Grow(4096)
	sqlOrder.WriteString( "update ")

	for index,table := range order.tables {
		sqlOrder.WriteString( table.GetName())

		if index != len(order.tables) - 1 {
			sqlOrder.WriteString( ",")
		}
	}

	sqlOrder.WriteString( " set ")
	for i:=0;i<len(order.fields) ;i++  {
		if len(order.tables)==1{
			sqlOrder.WriteString(order.fields[i].Name)
			sqlOrder.WriteString("=?")
		}else {
			sqlOrder.WriteString(order.fields[i].Target.GetName())
			sqlOrder.WriteString(".")
			sqlOrder.WriteString(order.fields[i].Name)
			sqlOrder.WriteString("=?")
		}
		if i != len(order.fields) - 1 {
			sqlOrder.WriteString( ",")
		}

	}

	sqlOrder.WriteString( order.getWhereString())
	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	order.executeNoResult(sqlOrder.String())
	return true
}
