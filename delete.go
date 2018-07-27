package xsql2

import "fmt"

func (order *XSql2Order)Delete() string {

	if len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.tables) <=0：" ,len(order.tables) )
		return "error"
	}

	sqlOrder := "delete from "

	for index,table := range order.tables {
		sqlOrder += table.GetName()

		if index != len(order.tables) - 1 {
			sqlOrder += " , "
		}
	}

	sqlOrder += order.getWhereString()
	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	list := order.execute(sqlOrder)
	fmt.Println("list:",list)
	return "true"
}
