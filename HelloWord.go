package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	//var str string
	str := "Hello World!"
	fmt.Println(str)
	//const 定义 常量 不可变
	const name = "aquan"
	//name = "test" cannot assign to name
	fmt.Println(name)
	log.Println(str + name)
	ages := make(map[string]int)
	ages["linday"] = 20
	ages["michael"] = 30
	fmt.Println(ages)
	for name,age := range ages {
		fmt.Println("name:",name,",age:",age)
	}
	log.Println("现在时间：", GetTime())
	address := Address{"广州"}
	person := Person{
		age:  24,
		name: "Aquan",
		Address: address,
	}
	log.Println(person.GetName())
	getNameTwo,_ := GetTimeTwo()
	log.Println("getNameTwo:", getNameTwo)
	result := make(chan int)
	go func() {
		sum:=0
		for i:=0;i<10;i++{
			sum=sum+i
		}
		result<-sum
	}()
	fmt.Println(<-result)
}

type Person struct {
	age int
	name string
	Address
}

func (p Person) GetName() string {
	return p.name
}

func GetTime() time.Time{
	return time.Now()
}

func GetTimeTwo() (time.Time,error){
	return time.Now(),nil
}

type Address struct {
	city string
}

type Stringer interface {
	String() string
}

func (p Person) String() string {
	return "name is "+p.name+",age is "+strconv.Itoa(p.age)
}