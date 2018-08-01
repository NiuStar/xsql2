package xsql2

import (
	"bytes"
	"fmt"
)

func (order *XSql2Order) Select() []map[string]interface{} {
	fmt.Println(11111)
	if len(order.fields) <= 0 || len(order.tables) <= 0 {
		return nil
	}
	var sqlOrder bytes.Buffer
	sqlOrder.WriteString( "select ")

	var pos,joinIndex int
	pos = -1
	if len(order.join)>0{
		pos=order.join[0].pos
	}
	for index, _ := range order.fields {
		//fmt.Println("filed:", order.fields[index])
		if len(order.tables) == 1 &&len(order.join)==0 {
			sqlOrder.WriteString( order.fields[index].Name)
		} else {
			sqlOrder.WriteString(  order.fields[index].Target.GetName() )
			sqlOrder.WriteString( "." )
			sqlOrder.WriteString( order.fields[index].Name)
		}
		if index != len(order.fields)-1 {
			sqlOrder.WriteString(  " , ")
		}
	}
	sqlOrder.WriteString( " from ")

	for index, _ := range order.tables {
		sqlOrder.WriteString( order.tables[index].GetName())

		if pos!= -1 && index == pos && joinIndex < len(order.join){

			sqlOrder.WriteString(order.getJoinString(joinIndex))
			joinIndex++
		}
		if index != len(order.tables)-1 {
			sqlOrder.WriteString( " , ")
		}
	}

	sqlOrder.WriteString( order.getWhereString())
	sqlOrder.WriteString( order.getOrderString())

	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/

	return order.execute(sqlOrder.String())
}
