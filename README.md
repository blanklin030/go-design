### go语言的类和对象
> go语言通过struct结构体来模拟类
```
type Bike struck{
    color string // 首字母小写表示私有
    Name string // 首字母大写表示公有
}
func (b *Bike) Run() string {
    fmt.Println(b.Name+" "+b.color+" is running")
}
```
### go语言的三大基本特征
+ 封装
> 首字母大小写来表示公有私有权限
```
type Person struct {
    name string // 私有
}
// 公有
func(p *Person) Walk() {
    fmt.Println(p.name+ " is walking") 
}

```

+ 继承
> 使用内嵌的方式，对结构体struct进行组合
```
type Chinese struct {
    person Person // 继承Person类
    skin string
}
func(c *Chinese)GetSkin() string {
    return c.skin
}

```

+ 多态
> 使用interface来实现
```
type Human interface {
    Speak()
}

type American struct {
    name string
    Language string
}
// 通过American这个类实现Speak方法
func(a *American) Speak() {
    fmt.Println(a.name + " speak " +  a.Language)
}

type Chinese struct {
    name string
    Language string
}
// 通过Chinese这个类实现Speak方法
func(a *American) Speak() {
    fmt.Println(a.name + " speak " +  a.Language)
}
```

### 5大基本原则
+ 单一原则
+ 开闭原则
+ 里氏替换原则
+ 接口隔离原则
+ 依赖反转原则

### 设计模式
+ 分类
    + 创建型
    > 单例模式/简单工厂模式/工厂方法模式/抽象工厂模式/建造者模式/原型模式
    + 结构型
    > 代理模式/适配器模式/装饰器模式/桥接模式/组合模式/亨元模式/外观模式
    + 行为型
    > 观察者模式/模板方法模式/命令模式/状态模式/职责链模式/解释器模式/中介者模式/访问者模式/策略模式/备忘录模式/迭代器模式
+ 实现
    + 简单工厂模式
    > ${your project}/factory/static
    + 工厂方法模式
    > ${your project}/factory/complex
    + 装饰器模式
    > ${your project}/decorate
    + 策略模式
    > ${your project}/strategy

### 建议
+ 多实践
+ 不要为了使用设计模式而使用设计模式
+ 少即是多

