package xsql2


func (order *XSql2Order)IN(obj *XSqlParam, v ...interface{})XSql2Where{

	in := " IN ( "
	for i,_:= range v{
		in += "?"
		if i != len(v)-1{
			in += ","
		}
	}
	in += ") "
	if len(order.tables) > 1 || len(order.join) > 0 {
		order.conditions = append(order.conditions, &condition{param: 0, type_: 0, value: obj.Target.GetName() + "." + obj.Name +in})
		order.args = append(order.args, v...)
	} else {
		order.conditions = append(order.conditions, &condition{param: 0, type_: 0, value: obj.Name+in})
		order.args = append(order.args, v...)
	}

	return order
}
