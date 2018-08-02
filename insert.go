package xsql2

import (
	"fmt"
	"bytes"
)

type XSql2Insert interface {
	Insert() int64
	Add(value... interface{}) XSql2Insert

}


func (order *XSql2Order)Add(value... interface{}) XSql2Insert {
	order.values = append(order.values,value)
	return order
}

func (order *XSql2Order)Insert() int64 {

	if len(order.values) <= 0 || len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.x.sets) <= 0 || len(order.x.tables) <=0：",len(order.sets) ,len(order.tables) )
		return 0
	}
	var sqlOrder bytes.Buffer
	sqlOrder.Grow(8192)

	sqlOrder .WriteString( "insert into ")

	for index,table := range order.tables {
		sqlOrder .WriteString( table.GetName())

		if len(order.values) > 0 {
			sqlOrder .WriteString( "(")
			for index1 , _ := range order.fields {

				sqlOrder .WriteString( order.fields[index1].Name)

				if index1 != len(order.fields) - 1 {
					sqlOrder .WriteString(  " , ")
				}
			}
			sqlOrder .WriteString( ")")
		}

		if index != len(order.tables) - 1 {
			sqlOrder .WriteString( " , ")
		}
	}

	sqlOrder .WriteString( " VALUES ")

	for index,value := range order.values {
		fmt.Println("value:",value)
		sqlOrder .WriteString( "(")

		for index1 ,v := range value {

			sqlOrder .WriteString(  "?")
			order.args = append(order.args,v)
			if index1 != len(value) - 1 {
				sqlOrder .WriteString(  " , ")
			}
		}
		sqlOrder .WriteString(  ")")
		if index != len(order.values) - 1 {
			sqlOrder .WriteString( " , ")
		}
	}

	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	n := order.executeForLastInsertId(sqlOrder.String())
	return n
}
