package xsql2

import "bytes"




type XSql2Join interface {
	On(obj *XSqlParam, op string, v interface{}) XSql2On
	OnParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2On
}
type XSql2On interface {
	On(obj *XSqlParam, op string, v interface{}) XSql2On
	OnParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2On
	LeftJoin(obj XSqlObject) XSql2Join
	RightJoin(obj XSqlObject) XSql2Join
	InnerJoin(obj XSqlObject) XSql2Join
	Field(obj ...*XSqlParam) XSql2Field
	//Where(obj *XSqlParam, op string, v interface{}) XSql2Where
	//WhereParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2Where
	//OrderByDESC(obj... *XSqlParam) XSql2OrderBy
	//OrderByASC(obj... *XSqlParam) XSql2OrderBy
	//Select() []map[string]interface{}
}

func (order *XSql2Order) LeftJoin(obj XSqlObject) XSql2Join {

	if len(order.tables) == 0 {
		return order
	}
	i := len(order.tables) - 1
	order.join = append(order.join, &XsqlJoin{pos: i, Target: obj, joinsql: " LEFT JOIN "})
	return order

}
func (order *XSql2Order) RightJoin(obj XSqlObject) XSql2Join {

	if len(order.tables) == 0 {
		return order
	}
	i := len(order.tables) - 1
	order.join = append(order.join, &XsqlJoin{pos: i, Target: obj, joinsql: " RIGHT JOIN "})
	return order

}
func (order *XSql2Order) InnerJoin(obj XSqlObject) XSql2Join {

	if len(order.tables) == 0 {
		return order
	}
	i := len(order.tables) - 1
	order.join = append(order.join, &XsqlJoin{pos: i, Target: obj, joinsql: " INNER  JOIN "})
	return order
}

func (order *XSql2Order) On(obj *XSqlParam, op string, v interface{}) XSql2On {
	if len(order.join) == 0 {
		return order
	}
	i := len(order.join) - 1
	order.join[i].conditions = append(order.join[i].conditions, &condition{param: 0, type_: 0, value: obj.Target.GetName() + "." + obj.Name + op + "?"})
	order.args = append(order.args, v)
	return order
}
func (order *XSql2Order) OnParam(obj *XSqlParam, op string, obj2 *XSqlParam) XSql2On {
	if len(order.join) == 0 {
		return order
	}
	i := len(order.join) - 1
	order.join[i].conditions = append(order.join[i].conditions, &condition{param: 1, type_: 0, value: obj.Target.GetName() + "." + obj.Name + op + obj2.Target.GetName() + "." + obj2.Name})
	return order
}

func (order *XSql2Order) getJoinString(index int) string {
	var sqlOrder bytes.Buffer
	sqlOrder.Grow(4096)
	sqlOrder.WriteString(order.join[index].joinsql)
	sqlOrder.WriteString(order.join[index].Target.GetName())
	sqlOrder.WriteString(" ON ")
	for index2, _ := range order.join[index].conditions {
		//sqlOrder += condition
		if index2 != 0 {

			if order.join[index].conditions[index2].type_ == 0 {
				sqlOrder.WriteString(" AND ")
			} else {
				sqlOrder.WriteString(" OR ")
			}
		}
		sqlOrder.WriteString(order.join[index].conditions[index2].value)
	}
	return sqlOrder.String()
}
