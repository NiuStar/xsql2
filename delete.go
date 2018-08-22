package xsql2

import (
	"fmt"
	"bytes"
)

func (order *XSql2Order)Delete() string {

	if len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.tables) <=0：" ,len(order.tables) )
		return "error"
	}
	var sqlOrder bytes.Buffer
	sqlOrder.Grow(4096)
	sqlOrder.WriteString( "delete from ")

	for index,table := range order.tables {
		sqlOrder.WriteString( table.GetName())

		if index != len(order.tables) - 1 {
			sqlOrder.WriteString(  " , ")
		}
	}

	sqlOrder.WriteString( order.getWhereString())
	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	order.executeNoResult(sqlOrder.String())
	return "true"
}
