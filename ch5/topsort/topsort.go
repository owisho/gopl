package topsort

import "sort"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"liner algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating system"},
	"operating system":      {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

/**
匿名函数被递归调用时，需要先声明一个变量，再将匿名函数赋值给这个变量
否则会出现编译错误
*/
func topsort1(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	//visitAll := func(items []string){
	//	//....
	//  visitAll(m[item]) //compile error : undefined : visitAll
	//  //....
	//}
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topsort(m map[string][]string) []string {
	var courses []string
	learned := make(map[string]bool)
	var learn func(string)
	learn = func(course string) {
		if learned[course] {
			return
		}
		if pcourses, ok := m[course]; ok {
			for _, pcourse := range pcourses {
				learn(pcourse)
			}
		}
		courses = append(courses, course)
		learned[course] = true
	}
	keys := make([]string, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		learn(key)
	}
	return courses
}
