# go 语言的类型系统

## 主要内容
* 声明新的用户定义的类型
* 使用方法，为类型增加新的行为
* 了解何时使用指针；何时使用值
* 通过接口实现多态
* 通过组合来扩展和改变类型
* 公开或者未公开的标识符

## 用户自定义类型

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