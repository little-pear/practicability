# JDK8 中新增的日期时间 API

## 新时间日期 API

- java.time - 包含值的对象的基础包
- java.time.chrono - 提供对不同的日历系统的访问
- java.time.format - 格式化和解析时间和日期
- java.time.temporal - 包括底层框架和扩展特性
- java.time.zone - 包含时区支持的类

> 说明：大多数开发者只会用到基础包和 format 包，也可能会用到 temporal 包。因此，尽管有68个新的公开类型，大多数开发者，大概将只会用到其中的三分之一。

- LocalDate、LocalTime、LocalDateTime 类是其中较重要的几个类，它们的实例时**不可变的对象**，分别表示使用 ISO-8601 日历系统的日期、时间、日期和时间。它们提供了简单的本地日期或时间，并不包含当前的时间信息，也不包含与时区相关的信息。

> LocalDate 代表 IOS 格式（yyyy-MM-dd）的日期，可以存储生日、纪念日等日期。
>
> LocalTime 表示一个时间，而不是日期。
>
> LocalDateTime 是用来表示日期和时间的，**这是一个最常见的类之一**。

***注：ISO-8601 日历系统是国际标准组织指定的现代公民的日期和时间的表示法，也就是公历***

LocalDateTime 相对于LocalDate 和 LocalTime 使用频率更高一些。