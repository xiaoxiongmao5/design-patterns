/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-05 20:58:39
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-05 21:43:50
 */
package template

import "fmt"

// 定义模板方法的抽象结构
type AbstractClass interface {
	Step1()
	Step2()
}

// 定义模版方法
func TemplateMethod(c AbstractClass) {
	fmt.Println("模板方法")
	c.Step1()
	c.Step2()
}
