package main
import "fmt"

func main(){

  //1) var <variable_name> <type>
  var x int
  x=5 //assign x with value 5
  fmt.Println("x:", x) //prints x

  //2) var <variable_name> <type> = <value>
  var y int = 100
  fmt.Println("y:", y)// prints y

  //3) var <variable_name> = <value>
  var z = 55
  fmt.Println("z:", z)

  //4) var <variable_name1>, <variable_name2>  = <value1>, <value2>
  var i, j = 10, "shawal"
  fmt.Println("i and j :", i, j)
  fmt.Println("i :", i)
  fmt.Println("j :", j)


  //5) <variable_name> := <value>
  m := 30
  fmt.Println(m)
  m := 40
  fmt.Println(m)
}
