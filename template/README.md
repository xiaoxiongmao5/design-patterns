### 模板方法（Template Method）模式

在 Go 语言中，没有传统面向对象语言中的类继承和模板方法的概念，因此无法像传统面向对象语言那样直接使用模板方法模式。Go 语言鼓励使用 __接口（interface）和组合（composition）__ 来实现代码重用和多态性。虽然 Go 语言没有显式的模板方法，但仍然可以使用接口和组合来实现类似的模式。

* 核心思想：定义一个算法的骨架，将一些步骤的实现延迟到子类。
* 设计优点：将通用的模版方法与具体的实现分离，这样可以轻松地添加新的实现，同时确保所有实现都遵循相同的模版结构。增强代码重用和扩展性。