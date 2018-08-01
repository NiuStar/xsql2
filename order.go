package xsql2

type XSql2OrderBy interface {
	Select() []map[string]interface{}
}


func (order *XSql2Order)OrderByASC(obj... *XSqlParam) XSql2OrderBy {


	for _,o := range obj {
		order.orderbys = append(order.orderbys, o.Target.GetName() + "." + o.Name + " asc")
	}

	return order
}

func (order *XSql2Order)OrderByDESC(obj... *XSqlParam) XSql2OrderBy {
	for _,o := range obj {
		order.orderbys = append(order.orderbys, o.Target.GetName() + "." + o.Name + " desc")
	}
	return order
}

func (order *XSql2Order)getOrderString() string {
	if len(order.orderbys)==0{
		return ""
	}
	return splicOrder(" order by "," ","," , order.orderbys)
}