>	Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

*	select的用法与switch语言非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。
* 	与switch语句相比， select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作。

大致的结构如下：

```
select {
    case <-chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
    }
```