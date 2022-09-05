## 函数

### func_01.go

值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。


### func_02.go
闭包

### func_03.go
defer的使用

### func_04.go
异常处理

painc

1、内置函数

2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行

3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行

4、直到goroutine整个退出，并报告错误


recover：

1、内置函数

2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为

3、一般的调用建议
        a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
        b). 可以获取通过panic传递的error

注意：

1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。

2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。

3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。