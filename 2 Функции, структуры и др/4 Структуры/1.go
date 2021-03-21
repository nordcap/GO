/*Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.
*/
package main
import "fmt"
func main()  {
	testStruct:=new(MyStruct)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}

type MyStruct struct {
	On bool
	Ammo int
	Power int
}

func (s *MyStruct) Shoot() bool {
	if s.On == false {
		return false
	}
	if s.Ammo>0 {
		s.Ammo--
		return true
	} else {
		return false
	}

}

func (s *MyStruct) RideBike() bool {
	if s.On == false {
		return false
	}
	if s.Power>0 {
		s.Power--
		return true
	} else {
		return false
	}
}