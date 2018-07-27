package xsql2

import "fmt"
func (order *XSql2Order)Add(value... string) *XSql2Order {
	order.values = append(order.values,value)
	return order
}

func (order *XSql2Order)Insert() string {

	if len(order.values) <= 0 || len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.sets) <= 0 || len(order.tables) <=0：",len(order.sets) ,len(order.tables) )
		return "error"
	}

	sqlOrder := "insert into "

	for index,table := range order.tables {
		sqlOrder += table.GetName()

		if len(order.values) > 0 {
			sqlOrder += "("
			for index1 , field := range order.fields {

				sqlOrder += field.Name

				if index1 != len(order.fields) - 1 {
					sqlOrder += " , "
				}
			}
			sqlOrder += ")"
		}

		if index != len(order.tables) - 1 {
			sqlOrder += " , "
		}
	}

	sqlOrder += " VALUES "

	for index,value := range order.values {
		fmt.Println("value:",value)
		sqlOrder += "("

		for index1 , v := range value {

			sqlOrder += "'" + v + "'"

			if index1 != len(value) - 1 {
				sqlOrder += " , "
			}
		}
		sqlOrder += ")"
		if index != len(order.values) - 1 {
			sqlOrder += " , "
		}
	}

	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	list := order.execute(sqlOrder)
	fmt.Println("list:",list)
	return "true"
}
