package xsql2

func (order *XSql2Order)OrderByASC(obj... *XSqlParam) *XSql2Order {


	for _,o := range obj {
		order.orderbys = append(order.orderbys, o.Target.GetName() + "." + o.Name + " asc")
	}

	return order
}

func (order *XSql2Order)OrderByDESC(obj... *XSqlParam) *XSql2Order {
	for _,o := range obj {
		order.orderbys = append(order.orderbys, o.Target.GetName() + "." + o.Name + " desc")
	}
	return order
}

func (order *XSql2Order)getOrderString() string {
	return splicOrder(" order by "," ","," , order.orderbys)
}