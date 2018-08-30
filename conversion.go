package xsql2

import (
	"strconv"
	"time"
	"math/big"
)

func InterfaceToString(data []map[string]interface{}) {
	for i, _ := range data {
		for i1, _ := range data[i] {
			switch data[i][i1].(type) {
			case int:
				data[i][i1] = strconv.Itoa(data[i][i1].(int))
			case *big.Int:
				data[i][i1] = data[i][i1].(*big.Int).String()
			case int64:
				data[i][i1] = strconv.FormatInt(data[i][i1].(int64), 10)
			case float32:
				data[i][i1] = strconv.FormatFloat(float64(data[i][i1].(float32)), 'f', 6, 64)
			case float64:
				data[i][i1] = strconv.FormatFloat(data[i][i1].(float64), 'f', -1, 64)
			case bool:
				if data[i][i1].(bool) {
					data[i][i1] = "true"
				} else {
					data[i][i1] = "false"
				}
			case time.Time:
				data[i][i1] = data[i][i1].(time.Time).Format("2006-01-02 15:04:05")
			case byte:
				data[i][i1] = string(data[i][i1].(byte))
			case string:
				data[i][i1] = data[i][i1].(string)
			default:
				if data[i][i1] != nil && data[i][i1] != "" {
					data[i][i1] = data[i][i1].(string)
				} else {
					data[i][i1] = ""
				}
			}
		}

	}
}
func InterfaceToString2(data map[string]interface{}) {
		for i, _ := range data {
			switch data[i].(type) {
			case int:
				data[i] = strconv.Itoa(data[i].(int))
			case *big.Int:
				data[i] = data[i].(*big.Int).String()
			case int64:
				data[i] = strconv.FormatInt(data[i].(int64), 10)
			case float32:
				data[i] = strconv.FormatFloat(float64(data[i].(float32)), 'f', 6, 64)
			case float64:
				data[i] = strconv.FormatFloat(data[i].(float64), 'f', -1, 64)
			case bool:
				if data[i].(bool) {
					data[i] = "true"
				} else {
					data[i] = "false"
				}
			case time.Time:
				data[i] = data[i].(time.Time).Format("2006-01-02 15:04:05")
			case byte:
				data[i] = string(data[i].(byte))
			case []byte:
				data[i] = string(data[i].([]byte))
			case string:
				data[i] = data[i].(string)
			default:
				if data[i] != nil && data[i] != "" {
					data[i] = data[i].(string)
				} else {
					data[i] = ""
				}
			}
		}
}
