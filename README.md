# go-design-pattern

> 📖 在线阅读（博客站）：**<https://mouxiaojun.github.io/design-patterns/>**

Go 设计模式学习仓库：23 个 GoF 模式的 **Go 落地示例（`code/`）+ 中文笔记（`document/`）**。

- 学习起点：mohuishou/go-design-pattern（一个经典 Go 设计模式仓库）；本仓库的代码与笔记是在其基础上**用自己的话重写**的练习稿
- 风格：`code/0N_name/`（可编译示例，package `_N_name`，只用标准库，测试可选用 testify）+ `document/名称模式.md`（中文笔记：分类 / 定义 / 实现方法 / 是否使用 / Go 落地要点）

## 模式索引

| # | 模式 | 类型 | code | document |
|---|---|---|---|---|
| 01 | 单例 Singleton | 创建型 | [code/01_singleton](code/01_singleton) | [单例模式.md](document/单例模式.md) |
| 02 | 工厂 Factory | 创建型 | [code/02_factory](code/02_factory) | [工厂模式.md](document/工厂模式.md) |
| 03 | 建造者 Builder | 创建型 | [code/03_builder](code/03_builder) | [建造者模式.md](document/建造者模式.md) |
| 04 | 原型 Prototype | 创建型 | [code/04_prototype](code/04_prototype) | [原型模式.md](document/原型模式.md) |
| 05 | 抽象工厂 Abstract Factory | 创建型 | [code/05_abstract_factory](code/05_abstract_factory) | [抽象工厂模式.md](document/抽象工厂模式.md) |
| 06 | 适配器 Adapter | 结构型 | [code/06_adapter](code/06_adapter) | [适配器模式.md](document/适配器模式.md) |
| 07 | 桥接 Bridge | 结构型 | [code/07_bridge](code/07_bridge) | [桥接模式.md](document/桥接模式.md) |
| 08 | 组合 Composite | 结构型 | [code/08_composite](code/08_composite) | [组合模式.md](document/组合模式.md) |
| 09 | 装饰器 Decorator | 结构型 | [code/09_decorator](code/09_decorator) | [装饰器模式.md](document/装饰器模式.md) |
| 10 | 外观 Facade | 结构型 | [code/10_facade](code/10_facade) | [外观模式.md](document/外观模式.md) |
| 11 | 享元 Flyweight | 结构型 | [code/11_flyweight](code/11_flyweight) | [享元模式.md](document/享元模式.md) |
| 12 | 代理 Proxy | 结构型 | [code/12_proxy](code/12_proxy) | [代理模式.md](document/代理模式.md) |
| 13 | 职责链 Chain of Responsibility | 行为型 | [code/13_chain_of_responsibility](code/13_chain_of_responsibility) | [职责链模式.md](document/职责链模式.md) |
| 14 | 命令 Command | 行为型 | [code/14_command](code/14_command) | [命令模式.md](document/命令模式.md) |
| 15 | 解释器 Interpreter | 行为型 | [code/15_interpreter](code/15_interpreter) | [解释器模式.md](document/解释器模式.md) |
| 16 | 迭代器 Iterator | 行为型 | [code/16_iterator](code/16_iterator) | [迭代器模式.md](document/迭代器模式.md) |
| 17 | 中介者 Mediator | 行为型 | [code/17_mediator](code/17_mediator) | [中介者模式.md](document/中介者模式.md) |
| 18 | 备忘录 Memento | 行为型 | [code/18_memento](code/18_memento) | [备忘录模式.md](document/备忘录模式.md) |
| 19 | 观察者 Observer | 行为型 | [code/19_observer](code/19_observer) | [观察者模式.md](document/观察者模式.md) |
| 20 | 状态 State | 行为型 | [code/20_state](code/20_state) | [状态模式.md](document/状态模式.md) |
| 21 | 策略 Strategy | 行为型 | [code/21_strategy](code/21_strategy) | [策略模式.md](document/策略模式.md) |
| 22 | 模板方法 Template Method | 行为型 | [code/22_template_method](code/22_template_method) | [模板方法模式.md](document/模板方法模式.md) |
| 23 | 访问者 Visitor | 行为型 | [code/23_visitor](code/23_visitor) | [访问者模式.md](document/访问者模式.md) |

## 验证

```bash
cd code
go build ./...
go test ./...
```
