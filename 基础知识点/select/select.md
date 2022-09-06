### select

当case上读一个通道时，如果这个通道是nil，则该case永远阻塞。这个功能有1个妙用，select通常处理的是多个通道，当某个读通道关闭了，但不想select再继续关注此case，而是关注其他case，把该通道设置为nil即可。


#### 如何跳出for-select
在select内的break并不能跳出for-select循环.
- 在满足条件的case内，使用return介绍协程，如果有结尾工作，尝试交给defer。
- 在select外for内使用break挑出循环，如combine函数。
- 使用goto，goto没有那么可怕，适当使用。

####  select{}永远阻塞 

select{}的效果等价于直接从刚创建的通道读数据：
```go
ch := make(chan int)
<-ch
```
