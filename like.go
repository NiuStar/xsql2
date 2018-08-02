package xsql2

import "strings"

var sqlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;",
)

func (order *XSql2Order)Like(obj *XSqlParam, v string)XSql2Where{
	v = sqlEscaper.Replace(v)
	if len(order.tables) > 1 || len(order.join) > 0 {
		order.conditions = append(order.conditions, &condition{param: 1, type_: 0, value: obj.Target.GetName() + "." + obj.Name + " LIKE '" + v+"' " })
	} else {
		order.conditions = append(order.conditions, &condition{param: 1, type_: 0, value: obj.Name + " LIKE '" + v+"' "})
	}

	return order
}