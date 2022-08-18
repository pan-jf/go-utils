# t2t

```markdown
t2t = type to type
```

t2t 用于把任意的源数据转移到任意的目标数据中，不同类型之间亦可以转移；t2t 利用了 Golang reflect 包提供运行时类型转换功能

* 由于 Golang 结构体静态定义 tag ，导致运行时才确定的数据较难处理，t2t拟补了这一缺陷；

```gotemplate
  type Object struct {
      Kind int64
  }
  
  type Car struct {
      Kind int64
      Doors int64
      Wheels int64
  }
  
  type Apple struct {
      Kind int64
      Color int64
  }
  
  var unknown interface{}
  
  var object struct
  t2t.Transfer(unknown, &object)
  
  if object.Kind == KindCar {
      var car Car
      t2t.Transfer(unknown, &car)
  } else if object.Kine == KindApple {
      var apple Apple
      t2t.Transfer(unknown, &apple)
  }
```

* Golang 属于强类型语言，在不同类型的数据当中转移需要编写额外的代码，增加代码复杂度并降低了代码可读性，t2t 在某些场合下能够直接进行转移，如： 

```gotemplate
type ErrorMessage string

var errorMessageList []ErrorMessage
var stringList []string

// 这种情况下 t2t 能够直接转移
t2t.Transfer(errorMessageList, &stringList) 
t2t.Transfer(stringList, &errorMessageList)

```

## 特点

* 兼容性强：t2t能够自定义使用tag来进行数据描述，无缝兼容json、bson、yaml等标签
* 性能高：在某些需要marshal后再unmarshal才能进行转换的场景下，t2t能够直接通过内存复制实现功能
* 弱类型兼容：浮点整型互转、字符串整型互转等场景下，使用传统的marshal/unmarshal需要编写繁琐代码，t2t只需把弱类型选项打开即可实现
* 无代码入侵：最简化使用t2t时，只需要调用t2t.Transfer即可，无需另外编写其他代码

## 类型转移支持

* bool -> bool
* string -> number
* string -> float
* string -> interface
* number -> float
* number -> string
* slice -> array
* slice -> struct
* slice -> interface
* array -> slice
* array -> struct
* array -> interface
* map -> struct
* struct -> map
* struct -> slice
* struct -> array