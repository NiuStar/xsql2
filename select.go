package xsql2

import "fmt"

func (order *XSql2Order)Select() []map[string]interface{} {

	if len(order.fields) <= 0 || len(order.tables) <=0 {
		return nil
	}


	sqlOrder := "select "

	for index,filed := range order.fields {
		fmt.Println("filed:",filed)

		if len(order.tables) == 1 {
			sqlOrder += filed.Name
		} else {
			sqlOrder += filed.Target.GetName() + "." + filed.Name
		}


		if index != len(order.fields) - 1 {
			sqlOrder += " , "
		}
	}


	sqlOrder += " from "

	for index,table := range order.tables {
		sqlOrder += table.GetName()

		if index != len(order.tables) - 1 {
			sqlOrder += " , "
		}
	}

	sqlOrder += order.getWhereString()
	sqlOrder += order.getOrderString()


	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	return order.execute(sqlOrder)
}

