package xsql2

import (
	"bytes"
)



type XSql2Where interface {
	Where(obj *XSqlParam, op string, v interface{}) XSql2Where
	WhereParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2Where
	LL() XSql2Where // (  id=1
	RR() XSql2Where //    id=1 )
	OR() XSql2Where//   id =1 OR
	Select() []map[string]interface{}
	Update() bool
	Count()int64
	OrderByDESC(obj... *XSqlParam) XSql2OrderBy
	OrderByASC(obj... *XSqlParam) XSql2OrderBy
	LIMIT(a,z int) XSql2limit
	IN(obj *XSqlParam, v ...interface{})XSql2Where
	Like(obj *XSqlParam, v string)XSql2Where
}
//where语句
func (order *XSql2Order) Where(obj *XSqlParam, op string, v interface{}) XSql2Where {

	//str := strings.ToLower(value)//统一转为小写
	//badStr := "'|and|exec|execute|insert|select|delete|update|count|drop|*|chr|mid|master|truncate|" +
	//	"char|declare|sitename|net user|xp_cmdshell|;|or|-|,|like'|and|exec|create|" +
	//	"table|from|grant|use|group_concat|column_name|" +
	//	"information_schema.columns|table_schema|union|where|order|by|;|-|--|,|like|//|/|#|\"";//过滤掉的sql关键字，可以手动添加
	//badStrs := strings.Split(badStr,"\\|")
	//
	//for _,bad := range badStrs {
	//	if strings.Index(str,bad) >= 0 {
	//		fmt.Println("数据库非法操作：",value)
	//		return nil
	//	}
	//}

	if len(order.tables) > 1 || len(order.join) > 0 {
		//查询语句
		order.conditions = append(order.conditions, &condition{param: 0, type_: 0, value: obj.Target.GetName() + "." + obj.Name + op + "?"})
		//传入查询参数
		order.args = append(order.args, v)
	} else {
		order.conditions = append(order.conditions, &condition{param: 0, type_: 0, value: obj.Name + op + "?"})
		order.args = append(order.args, v)
	}

	return order
}

func (order *XSql2Order) WhereParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2Where {

	if len(order.tables) > 1 || len(order.join) > 0 {
		order.conditions = append(order.conditions, &condition{param: 1, type_: 0, value: obj.Target.GetName() + "." + obj.Name + op + obj2.Target.GetName() + "." + obj2.Name})
	} else {
		order.conditions = append(order.conditions, &condition{param: 1, type_: 0, value: obj.Name + op + obj2.Name})
	}

	return order
}


//改用OR，默认AND
func (order *XSql2Order) OR() XSql2Where{
	i := len(order.conditions)-1
	order.conditions[i].type_=1
	return order
}


//左括号
func (order *XSql2Order) LL() XSql2Where{
	i := len(order.conditions)-1
	order.conditions[i].brackets=1
	order.conditions[i].bracketsnum++
	return order
}
func (order *XSql2Order) LR() XSql2Where{
	i := len(order.conditions)-1
	order.conditions[i].brackets=2
	return order
}
func (order *XSql2Order) RL() XSql2Where{
	i := len(order.conditions)-1
	order.conditions[i].brackets=3
	return order
}
//右括号
func (order *XSql2Order) RR() XSql2Where{
	i := len(order.conditions)-1
	order.conditions[i].brackets=4
	order.conditions[i].bracketsnum++
	return order
}


func (order *XSql2Order) getWhereString() string {
	var sqlOrder bytes.Buffer
	sqlOrder.Grow(4096)
	if len(order.conditions) > 0 {

		sqlOrder.WriteString(" where ")
		for index, _ := range order.conditions {
			//sqlOrder += condition
			if order.conditions[index].brackets ==1{
				for i := 0;i<order.conditions[index].bracketsnum;i++{
					sqlOrder.WriteString(" ( ")
				}
			}else if order.conditions[index].brackets == 3 {
				for i := 0;i<order.conditions[index].bracketsnum;i++ {
					sqlOrder.WriteString(" ) ")
				}
			}

			//if index != 0 {
			//
			//	if order.conditions[index].type_ == 0 {
			//		sqlOrder.WriteString(" AND ")
			//	} else {
			//		sqlOrder.WriteString(" OR ")
			//	}
			//}
			if order.conditions[index].param == 0 {

			}
			sqlOrder.WriteString(order.conditions[index].value)

			if order.conditions[index].brackets ==2{
				for i := 0;i<order.conditions[index].bracketsnum;i++ {
					sqlOrder.WriteString(" ( ")
				}
			}else if order.conditions[index].brackets == 4 {
				for i := 0;i<order.conditions[index].bracketsnum;i++ {
					sqlOrder.WriteString(" ) ")
				}
			}
			if index != len(order.conditions)-1 {

				if order.conditions[index].type_ == 0 {
					sqlOrder.WriteString(" AND ")
				} else {
					sqlOrder.WriteString(" OR ")
				}
			}
		}
		return sqlOrder.String()
	}
	return sqlOrder.String()
}
