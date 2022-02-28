package main

import (
	"fmt"
	"reflect"
)

/* type order struct {
	orderId    int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
} */

// Example 1 program test
/* func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.orderId, o.customerId)
	return i
} */

// Reflection: TypeOf, ValueOf, Kind
/* func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()

	fmt.Println("Type", t)
	fmt.Println("Value", v)
	fmt.Println("Kind", k)
} */

// Reflection: NumField and Field
/* func createQuery(q interface{}) {
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Struct {
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
} */

// Reflection: Complete Program
type order struct {
	orderId    int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

/* func addQueryFields(q interface{}) string {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		tableName := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s(", tableName)
		value := reflect.ValueOf(q)
		for i := 0; i < value.NumField(); i++ {
			name := value.Field(i)
			if i == 0 {
				query = fmt.Sprintf("%s%s", query, name)
			} else {
				query = fmt.Sprintf("%s, %s", query, name)
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return query
	}
	fmt.Println("unsupported format")
	return ""
} */
func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		// query := addQueryFields(q)
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Format is not supported")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	/* i := 10
	fmt.Printf("%d %T\n", i, i) */

	/* //Example with struct
	o := order{
		orderId:    1234,
		customerId: 567,
	}
	// fmt.Println(o)
	fmt.Println(createQuery(o)) */

	/* o := order{
		orderId:    12434,
		customerId: 4590,
	}
	createQuery(o) */

	// Reflection: Int and String
	/* a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type: %T value: %v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type: %T value:%v\n", y, y) */
	o := order{
		orderId:    12434,
		customerId: 4590,
	}
	createQuery(o)
	e := employee{
		name:    "Andere",
		id:      567,
		address: "Kisumu",
		salary:  90000,
		country: "Kenya",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}
