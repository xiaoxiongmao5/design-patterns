/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-05 17:53:22
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-05 21:41:35
 * @FilePath: /template/main.go
 */
package main

import (
	"fmt"

	"github.com/xiaoxiongmao5/design-patterns/template/template"
)

// ConcreteClass1 实现 AbstractClass 接口
type ConcreteClass1 struct{}

func (c *ConcreteClass1) Step1() {
	fmt.Println("具体类1的步骤1")
}

func (c *ConcreteClass1) Step2() {
	fmt.Println("具体类1的步骤2")
}

// ConcreteClass2 实现 AbstractClass 接口
type ConcreteClass2 struct{}

func (c *ConcreteClass2) Step1() {
	fmt.Println("具体类2的步骤1")
}

func (c *ConcreteClass2) Step2() {
	fmt.Println("具体类2的步骤2")
}

func main() {
	class1 := &ConcreteClass1{}
	class2 := &ConcreteClass2{}

	template.TemplateMethod(class1)
	template.TemplateMethod(class2)
}
