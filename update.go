package xsql2

import "fmt"

func (order *XSql2Order)Set(obj *XSqlParam,value string) *XSql2Order {

	if len(order.tables) > 1 {
		order.sets = append(order.sets,obj.Target.GetName() + "." + obj.Name + "='" + value + "'")
	} else {
		order.sets = append(order.sets,obj.Name + "='" + value + "'")
	}
	return order
}

func (order *XSql2Order)Update() bool {

	if len(order.sets) <= 0 || len(order.tables) <=0 {
		fmt.Println("数据库执行出错，len(order.sets) <= 0 || len(order.tables) <=0：",len(order.sets) ,len(order.tables) )
		return false
	}

	sqlOrder := "update "

	for index,table := range order.tables {
		sqlOrder += table.GetName()

		if index != len(order.tables) - 1 {
			sqlOrder += " , "
		}
	}

	sqlOrder += " set "

	for index,set := range order.sets {
		fmt.Println("set:",set)

		sqlOrder += set

		if index != len(order.sets) - 1 {
			sqlOrder += " , "
		}
	}

	sqlOrder += order.getWhereString()
	/*stmt := order.xsql2.stmts[sqlOrder]
	if stmt != nil {
		rows,err := stmt.Query()
		checkErr(err)
	}*/
	order.executeNoResult(sqlOrder)
	return true
}
