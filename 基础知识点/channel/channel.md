channel类型  引用类型  ` var 变量 chan 元素类型`

### chan_01.go

通道的基本使用

channl的状态
- nil，未初始化的状态，只进行了声明，或者手动赋值为nil
- active，正常的channel，可读或者可写
- closed，已关闭，千万不要误认为关闭channel后，channel的值是nil

channel可进行3种操作：

- 读
- 写
- 关闭


### chan_02.go

channel 怎么会阻塞阻塞？


### chan_03.go

什么情况下关闭 channel 会造成 panic ？

总结：

- 应该只在发送端关闭 channel。（防止关闭后继续发送）
-  存在多个发送者时不要关闭发送者 channel，而是使用专门的 stop channel。（因为多个发送者都在发送，且不可能同时关闭多个发送者，否则会造成重复关闭。发送者和接收者多对一时，接收者关闭 stop channel；多对多时，由任意一方关闭 stop channel，双方监听 stop channel 终止后及时停止发送和接收）


这两点规律被称为“channel 关闭守则”。


### chan_04.go

#### 有没有必要关闭 channel？不关闭又如何？


- 情况一:channel 的发送次数等于接收次数
  - channel 的发送次数等于接收次数时，发送者携程 和接收者携程分别都会在发送或接收结束时结束 。
没有被关闭通道没有代码使用被垃圾收集器回收。因此这种情况下，不关闭 channel，没有任何副作用。

- 情况二：channel 的发送次数大于/小于接收次数
  - channel 的发送次数小于接收次数时，接收者 go routine 由于等待发送者发送一直阻塞。因此接收者 go routine 一直未退出，ich 也由于一直被接收者使用无法被垃圾回收。未退出的 go routine 和未被回收的 channel 都造成了内存泄漏的问题。 


#### 如何判断 channel 是否关闭？ 
多重返回
for-range

### chan_05.go 场景应用