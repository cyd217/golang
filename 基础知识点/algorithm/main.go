package main

import (
	"fmt"
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
)

type Stu struct {
	Id   int    //编号
	Name string //姓名
	Age  int    //年龄
}

func main() {
	var stus []Stu
	fmt.Println("====== 初始化数据 =====")
	stus = append(stus, Stu{Id: 1, Name: "张一", Age: 10})
	stus = append(stus, Stu{Id: 2, Name: "李二", Age: 20})
	stus = append(stus, Stu{Id: 3, Name: "王三", Age: 20})
	stus = append(stus, Stu{Id: 4, Name: "刘四", Age: 30})
	stus = append(stus, Stu{Id: 5, Name: "黄五", Age: 40})
	stus = append(stus, Stu{Id: 6, Name: "周六", Age: 50})
	stus = append(stus, Stu{Id: 7, Name: "周六", Age: 60})
	fmt.Println(stus)

	fmt.Println("====== 过滤查询学生列表 =====")
	var stus2 []Stu
	linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 20 && s.Id > 0
	}).Take(100).ToSlice(&stus2) //Take 固定取几个，也可以不写
	fmt.Println(stus2)

	fmt.Println("====== 过滤查询学生列表，并排序 =====")
	var stus3 []Stu
	linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 20 && s.Id > 0
	}).Take(100).OrderByDescendingT(func(s Stu) int {
		return s.Age
	}).ToSlice(&stus3)
	fmt.Println(stus3)

	fmt.Println("====== 过滤查询学生列表,根据名字Distinct，再排序 =====")
	var stus4 []Stu
	linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 20 && s.Id > 0
	}).Take(100).OrderByDescendingT(func(s Stu) string {
		return s.Name
	}).ThenByDescendingT(func(s Stu) int {
		return s.Age
	}).DistinctByT(func(s Stu) string {
		return s.Name
	}).ToSlice(&stus4) //DistinctByT 也可以直接改为Distinct()：完全匹配
	fmt.Println(stus4)

	fmt.Println("====== 过滤查询单个学生 =====")
	s1 := linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 40 && strings.Contains(s.Name, "周")
	}).First() //也可以用Last
	fmt.Println(s1)

	fmt.Println("====== 过滤结构体并查询姓名 =====")
	var names []string
	linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 30 && s.Id > 0
	}).SelectT(func(s Stu) string {
		return s.Name
	}).ToSlice(&names)
	fmt.Println(names)

	fmt.Println("====== 函数相关 =====")
	stuQuery := linq.From(stus).WhereT(func(s Stu) bool {
		return s.Age >= 0 && s.Id > 0
	}).SelectT(func(s Stu) int {
		return s.Age
	})
	fmt.Println("平均年龄：", stuQuery.Average())
	fmt.Println("最大年龄：", stuQuery.Max())
	fmt.Println("最小年龄：", stuQuery.Min())
	fmt.Println("总年龄：", stuQuery.SumInts()) // .SumFloats()
	fmt.Println("共查询学生：", stuQuery.Count())
	fmt.Println("共查询学生(Distinct)：", stuQuery.Distinct().Count())
}
