package xsql2


type XSql2limit interface {
	Select() []map[string]interface{}
}


func (order *XSql2Order)LIMIT(a,z int) XSql2limit{
	if z==0{
		z = -1
	}
	order.args = append(order.args,a,z)
	order.limit = " LIMIT ?,?"
	return order
}
