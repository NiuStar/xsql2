package xsql2


type XSql2limit interface {
	Select() []map[string]interface{}
}

//限制
func (order *XSql2Order)LIMIT(a,z int) XSql2limit{
	if z==0{
		z = -1
	}
	order.args = append(order.args,a,z)
	order.limit = " LIMIT ?,?"
	return order
}
func (order *XSql2Order)getLimitString() string {
	if len(order.limit)==0{
		return ""
	}
	return order.limit
}