package xsql2

import (

	"strings"
	"fmt"
)

func (order *XSql2Order)Where(obj *XSqlParam,operator,value string) *XSql2Order {

	str := strings.ToLower(value)//统一转为小写
	badStr := "'|and|exec|execute|insert|select|delete|update|count|drop|*|chr|mid|master|truncate|" +
		"char|declare|sitename|net user|xp_cmdshell|;|or|-|,|like'|and|exec|create|" +
		"table|from|grant|use|group_concat|column_name|" +
		"information_schema.columns|table_schema|union|where|order|by|;|-|--|,|like|//|/|#|\"";//过滤掉的sql关键字，可以手动添加
	badStrs := strings.Split(badStr,"\\|")

	for _,bad := range badStrs {
		if strings.Index(str,bad) >= 0 {
			fmt.Println("数据库非法操作：",value)
			return nil
		}
	}

	if len(order.tables) > 1 {
		order.conditions = append(order.conditions,&condition{type_:0,value:obj.Target.GetName() + "." + obj.Name + operator + `'` + value + `'`})
	} else {
		order.conditions = append(order.conditions,&condition{type_:0,value:obj.Name + operator + `'` + value + `'`})
	}

	return order
}

func (order *XSql2Order)WhereParam(obj *XSqlParam,operator string,obj2 *XSqlParam) *XSql2Order {

	if len(order.tables) > 1 {
		order.conditions = append(order.conditions,&condition{type_:0,value:obj.Target.GetName() + "." + obj.Name + operator + obj2.Target.GetName() + "." + obj2.Name })
	} else {
		order.conditions = append(order.conditions,&condition{type_:0,value:obj.Name + operator + obj2.Name})
	}

	return order
}

func (order *XSql2Order)getWhereString() string {
	var sqlOrder string
	if len(order.conditions) > 0 {
		sqlOrder += " where "
		for index,con := range order.conditions {
			//sqlOrder += condition
			if index != 0 {

				if con.type_ == 0 {
					sqlOrder += " AND "
				} else {
					sqlOrder += " OR "
				}
			}
			sqlOrder += con.value
		}
		return sqlOrder
	}
	return sqlOrder
}
