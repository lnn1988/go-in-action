# go 语言的类型系统

## 主要内容
* 声明新的用户定义的类型
* 使用方法，为类型增加新的行为
* 了解何时使用指针；何时使用值
* 通过接口实现多态
* 通过组合来扩展和改变类型
* 公开或者未公开的标识符

## 5.1用户自定义类型

```
# 在程序中自定义一个用户类型
type user struct {
    name string
    email string
    ext int
    priviledge bool
}

# 生明 user 类型的变量
var bill user
```

当一个变量被声明时，对应的值总是会被初始化。这个值要么使用指定的值，要么使用变量类型的默认值（零值）进行初始化。
对于数值类型来说，零值是；
对于字符串，零值是空字符串；
对于bool类型，零值是false。

```
# 声明一个变量并初始化所有字段
lise := user {
    name:"lisa", 
    email:"lisa@email.com",
    ext:123,
    priviledge:true
}
```

也可以按照顺序只声明对应的值，写在一行，但是顺序必须要和结构字段中声明的顺序一致。
```
lisa := user{"lisa", "lisa@email.com", 123, true}
```

声明结构体时，字段类型并不一定是内置类型，也可以使用用户自己定义的类型。
```
# 使用用户子定义类型
type admin struct {
    person user
    level string
}

# 初始化admin类型变量
fred ：= admin {
    person : user{
        name:"lisa", 
        email:"lisa@email.com", 
        ext:123,
        priviledge:true,
    },
    level:"super",
}

```

另一种声明用户类型的方法是基于一个已有类型，将其作为新类型使用。
```
# 声明类型Duration
type Duration int64


package main
type Duration int64 
fun main (){
    var dur Duration
    dur = int64(1000)
}
```
上述代码会抛出异常，因为dur是Duration类型，而int64(1000)是int64类型，尽管类型的值兼容，但是也不可以互相赋值。编译器不会对不同类型的值做隐式转换。

## 5.2方法
方法能给用户自定义类型增加新的行为。
方法也是函数，只是在声明时需要在关键字和方法名之间增加一个参数。
```
package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify()  {
	fmt.Printf("sending email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeMail(email string)  {
	u.email = email
}

func main()  {
	bill := user{"Bill", "bill@mail.com"}
	bill.notify()
	lisa := &user{"Lisa", "lisa@mail.com"}
	lisa.notify()

	bill.changeMail("bill@163.com")
	bill.notify()

	lisa.changeMail("lisa@163.com")
	lisa.notify()
}
```
方法的关键字 func 和函数名之间的参数被称为接受者，将函数和接收者的类型绑定在一起。
如果一个函数有接收者，那么这个函数就被称之为方法。
Go 语言有两种类型的接收者，分别是**值接收者**、**指针接收者**。
上面的代码中，使用值接收者声明了 notify 方法，使用指针接收者声明了 changeMail 方法。
如果使用值接收者声明方法，如 notify，调用时会使用这个值的一个副本来执行。
指向 user 的指针也可以调用 notify 函数，因为 go 语言调整了指针的值，用来支持方法接收者的定义，在使用指针调用 notify 时，go 做了如下的操作
```
# go 隐式的将指针变成了指针指向的值，这样就符合了要求
# 由于 notify 方法的接收者是值，所以传递的是 lisa 指针指向的值的副本
(*lisa).notify()
```

changeMail 方法的接收者是指针，当被调用时，这个方法会共享调用方法时接收者所指向的值。
类似的，也可以用一个值来调用接收者为指针的方法。当使用值调用接收者为指针的方法时，go 语言会进行调整，首先引用该值的地址得到一个指向该地址的指针，这个指针就会匹配方法的接收者类型，进行调用。

## 5.3类型的本质

### 5.3.1内置类型
go 语言的内置类型分别是数值类型、字符串类型和布尔类型。这些类型本质上是原始的类型。
因此，在对这些值进行增加、删除的时候，会创建新的值。把这些值传递给函数或者方法时，会传递值的副本。

### 5.3.2 引用类型
go 语言的引用类型有以下几个：
切片、映射、通道、接口和函数类型。

## 5.4接口

### 5.4.1标准库
### 5.4.2实现
### 5.4.3方法集
### 5.4.4多态


## 5.5嵌入类型


## 5.6 公开或未公开的标识符
当一个标识符的名字以小写字母开头时，这个标识符就是未公开的，即包外面的代码看不见；
当一个标识符以大写字母开头，这个标识符就是公开的，即被包外的代码看见。
