package main //main包是一个主函数，作为一个入口的标记
import "fmt"

//引入包

/*变量设置
	1:var var_name var_type
	2:var var_name=value
	多变量的声明
	var var_name1 var_name2 var_name3 var_type
	var var_name1 var_name2 var_name3 var_type=value1,value2,value3
	var var_name1 var_name2 var_name3 =value1,value2,value3

	var (
	var_name,var_type
	var_name2,var_type
)

*/

var a ="123"
var b ="bbb"
var c bool
var d=0
var e ="123"

var f ,g  ,h int =1,2,3

var  (
	y int=10
	x int=20

)

func main() {
fmt.Println(a,b,c,d)
fmt.Println(f,g,h)
fmt.Println(x+y)
}
