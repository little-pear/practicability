## spring框架的两大核心机制

- IOC（控制反转）和DI（依赖注入）

- AOP（面向切面编程）

  Spring是一个企业级开发框架，是软件设计层面的框架，优势在于可以将应用程序进行分层，开发者可以自主选择组件。

  MVC：struts2、springmvc

  ORMapping：hibernate，mybatis，spring data

## 如何使用IOC

创建一个maven项目，在pom.xml添加依赖

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.bai</groupId>
    <artifactId>aispringioc</artifactId>
    <version>1.0-SNAPSHOT</version>

    <dependencies>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-context</artifactId>
            <version>5.2.7.RELEASE</version>
        </dependency>

        <!-- 简化实体类代码的开发 -->
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <version>1.18.12</version>
            <scope>provided</scope>
        </dependency>
    </dependencies>
</project>
```

### 配置文件

> 通过配置`bean`标签来完成对象的管理
>
> `id`：对象名
>
> `class`：对象的模板类。（所有交给IOC容器来管理的必须有无参构造函数，因为Spring底层是通过反射机制来创建对象的，调用的是无参构造。）
>
> 对象的成员变量通过`property`标签完成赋值。
>
> `name`：成员变量名。
>
> `value`：成员变量值（基本数据类型，String可以直接进行赋值，如果是其他的引用数据类型，不饿能通过value进行赋值）。
>
> `ref`：将IOC中的另外一个bean赋给当前的成员变量（DI）

### IOC底层原理

读取配置文件，解析XML。

通过反射机制实例化配置文件中所配置的所有的bean。

### scope作用域

Spring 管理的 bean 是根据 scope 来生成的，表示 bean 的作用域，共4种。 默认是singleton。

- singleton：单例，表示通过 IoC容器获取的 bean 是唯一的。
- prototype：原型，表示通过 IoC 容器获取的 bean 是不同的。
- request：请求，表示在一次 HTTP 请求内有效。
- session：会话，表示在一个用户会话内有效。

request 和 session 只适用于 Web 项目，大多数情况下，我们使用单例和原型较多。

prototype 模式当业务代码 IoC 容器中的 bean 时，Spring 才去调用无参构造创建对应的 bean。

singleton 模式无论在业务代码是否获取到 IoC 容器中的 bean ， Spring 在加载 spring.xml 时就会创建 bean 。

### Spring 的继承

与Java的继承不同，Java是类层面的继承，子类可以继承父类的内部结构信息；Spring  是对象层面的继承，子对象可以继承父对象的属性值。

```xml

<bean id="student2" class="com.bai.entity.Student" scope="prototype">
    <property name="id" value="1"></property>
    <property name="name" value="张三"></property>
    <property name="age" value="22"></property>
    <property name="addresses">
        <list>
            <ref bean="address"></ref>
            <ref bean="address2"></ref>
        </list>
    </property>
</bean>

<bean id="stu" class="com.bai.entity.Student" parent="student2">
    <property name="name" value="李四"></property>
</bean>

<bean id="address" class="com.bai.entity.Address">
    <property name="id" value="1"></property>
    <property name="name" value="科技路"></property>
</bean>

<bean id="address2" class="com.bai.entity.Address">
    <property name="id" value="2"></property>
    <property name="name" value="高新区"></property>
</bean>
```

Spring 继承的关注点在于具体的对象，而不在于类，即不同的两个类的实例化对象可以完成继承，前提是子对象必须包含父对象的所有属性，同时可以再次基础上添加更多的属性。



### Spring 的依赖

与继承类似，依赖也是描述 bean 和 bean 之间的一种关系，配置依赖之后，被依赖的 bean 一定先创建，再创建依赖的 bean ，A依赖于B，先创建B，在创建A。

```xml
<?xml version="1.0" encoding="utf-8" ?>
<beans xmlns="http://www.springframework.org/schema/beans"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns:context="http://www.springframework.org/schema/context"
            xmlns:p="http://www.springframework.org/schema/p"
        xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans-3.2.xsd
        http://www.springframework.org/schema/context
        http://www/springframework.org/schema/context/spring-context-4.3.xsd">

    <bean id="user" class="com.bai.entity.User" depends-on="student"></bean>

    <bean id="student" class="com.bai.entity.Student"></bean>

</beans>
```

### Spring 的 P 命名空间

P 命名空间是对 IoC / DI 的简化操作，使用 P 命名空间可以更加方便地完成 bean 的配置以及 bean 之间的依赖注入。

```xml
<?xml version="1.0" encoding="utf-8" ?>
<beans xmlns="http://www.springframework.org/schema/beans"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns:context="http://www.springframework.org/schema/context"
            xmlns:p="http://www.springframework.org/schema/p"
        xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans-3.2.xsd
        http://www.springframework.org/schema/context
        http://www/springframework.org/schema/context/spring-context-4.3.xsd">

    <bean id="student" class="com.bai.entity.Student" p:id="1" p:name="张三" p:age="18" p:address-ref="address"></bean>

    <bean id="address" class="com.bai.entity.Address" p:id="2" p:name="科技路"></bean>
</beans>
```

### Spring的工厂方法

IoC通过工厂模式创建 bean 的方式有两种。

- 静态工厂模式 （配置一个 bean，本身的静态工厂不需要进行实例化）
- 实例工厂模式 （配置两个 bean，实例工厂也需要进行一次实例化）

> 静态工厂模式

> 实体类

``` java
package com.bai.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author bai
 * @create 2020-06-18-11:07
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class Car {
    private long id;
    private String name;
}
```

> 静态构造方法

``` java
package com.bai.factory;

import com.bai.entity.Car;

import java.util.HashMap;
import java.util.Map;

/**
 * @author bai
 * @create 2020-06-18-11:08
 */
public class StaticCarFactory {
    private static Map<Long, Car> carMap;
    static {
        carMap = new HashMap<Long, Car>();
        carMap.put(1L,new Car(1L,"宝马"));
        carMap.put(2L,new Car(2L,"奔驰"));
    }

    public static Car getCar(long id) {
        return carMap.get(id);
    }
}
```

> 配置文件

```xml
<!-- 配置静态工厂创建 car -->
<bean id="car" class="com.bai.factory.StaticCarFactory" factory-method="getCar">
    <constructor-arg value="2"></constructor-arg>
</bean>
```

> 实例工厂方法

```java
package com.bai.factory;

import com.bai.entity.Car;

import java.util.HashMap;
import java.util.Map;

/**
 * @author bai
 * @create 2020-06-18-12:36
 */
public class InstanceCarFactory {
    private Map<Long, Car> carMap;
    public InstanceCarFactory() {
        carMap = new HashMap<Long, Car>();
        carMap.put(1L,new Car(1L,"宝马"));
        carMap.put(2L,new Car(2L,"奔驰"));
    }

    public Car getCar(long id) {
        return carMap.get(id);
    }
}
```

```xml
<!-- 配置实例工厂 bean -->
<bean id="carFactory" class="com.bai.factory.InstanceCarFactory"></bean>

<!-- 配置实例工厂创建的 Car -->
<bean id="car2" factory-bean="carFactory" factory-method="getCar">
    <constructor-arg value="1"></constructor-arg>
</bean>
```

### IoC 自动装载 （Autowire）

IoC 负责创建对象，DI负责完成对象的依赖注入，通过配置 property 标签的 ref 属性来完成，同时Spring 提供了更加简便的依赖注入的方式，：自动装载，不需要手动配置 property ，IoC 容器会自动选择 bean 完成注入。

自动装载的两种方式：

- byName： 通过属性名进行自动装载。
- byType： 通过属性的数据类型进行自动装载。

> byName

```xml
<bean id="car" class="com.bai.entity.Car">
    <property name="id" value="1"></property>
    <property name="name" value="宝马"></property>
</bean>

<bean id="person" class="com.bai.entity.Person" autowire="byName">
    <property name="id" value="11"></property>
    <property name="name" value="张三"></property>
</bean>
```

> byType

```xml
<bean id="cars" class="com.bai.entity.Car">
    <property name="id" value="1"></property>
    <property name="name" value="宝马"></property>
</bean>

<bean id="person" class="com.bai.entity.Person" autowire="byType">
    <property name="id" value="11"></property>
    <property name="name" value="张三"></property>
</bean>
```

byType 需要注意：如果同时存在两个及以上的符合条件的 bean 时，自动装载会抛出异常。

## AOP

AOP：Aspect Oriented Programming 面向切面编程。

AOP的优点：

- 降低模块之间的耦合度。
- 使系统更容易拓展。
- 更好的代码复用。
- 非业务代码更加集中，不分散，便于统一管理。
- 业务代码更加简洁纯粹，不参杂其他代码的影响。

AOP 使对面向对象编程的一个补充，在运行时，动态地将代码切入到类的指定方法、指定位置上的编程思想就是面向切面编程。将不同方法的同一个位置抽象成一个切面对象，对该切面对象进行编程就是AOP。

### 如何使用？？

- 创建 Maven 工程，pom.xml 添加

```xml
<dependencies>
    
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-context</artifactId>
        <version>5.2.7.RELEASE</version>
    </dependency>
    
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-aop</artifactId>
        <version>5.2.7.RELEASE</version>
    </dependency>

    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-aspects</artifactId>
        <version>5.2.7.RELEASE</version>
    </dependency>
</dependencies>
```

- 创建一个计算器的接口 Cal，定义4个方法。

```java
package com.bai.utils;

public interface Cal {
    public int add(int num1, int num2);
    public int sub(int num1, int num2);
    public int mul(int num1, int num2);
    public int div(int num1, int num2);
}
```

- 创建接口的实现类 CalImpl。

```java
package com.bai.utils.impl;

import com.bai.utils.Cal;

/**
 * @author bai
 * @create 2020-06-18-13:56
 */
public class CalImpl implements Cal {
    public int add(int num1, int num2) {
        System.out.println("add方法的参数是[" + num1 + "," + num2 + "]");
        int result = num1 + num2;
        System.out.println("add方法的结果是" + result);
        return result;
    }

    public int sub(int num1, int num2) {
        System.out.println("sub方法的参数是[" + num1 + "," + num2 + "]");
        int result = num1 - num2;
        System.out.println("sub方法的结果是" + result);
        return result;
    }

    public int mul(int num1, int num2) {
        System.out.println("mul方法的参数是[" + num1 + "," + num2 + "]");
        int result = num1 * num2;
        System.out.println("mul方法的结果是" + result);
        return result;
    }

    public int div(int num1, int num2) {
        System.out.println("div方法的参数是[" + num1 + "," + num2 + "]");
        int result = num1 / num2;
        System.out.println("div方法的结果是" + result);
        return result;
    }
}
```

上述代码中，日志信息和业务逻辑的耦合性很高，，不利于系统的维护，使用AOP可以进行优化，如何来实现AOP？使用动态代理的方式来实现。

给业务代码找一个代理，打印日志信息的工作交给代理来做，这样的话业务代码就只需要关注自身的业务即可。

```java
package com.bai.utils.impl;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.Arrays;

/**
 * @author bai
 * @create 2020-06-18-16:33
 */
public class MyInvocationHandler implements InvocationHandler {
    //接收委托对象
    private Object object = null;

    //返回代理对象
    public Object bind(Object object) {
        this.object = object;
        return Proxy.newProxyInstance(object.getClass().getClassLoader(),object.getClass().getInterfaces(),this);
    }

    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println(method.getName() + "方法的参数是" + Arrays.toString(args));
        Object result = method.invoke(this.object, args);
        System.out.println(method.getName() + "的结果是" + result);
        return result;
    }
}
```

以上是通过动态代理实现 AOP 的过程，比较复杂，不好理解，Spring 框架对 AOP 进行了封装，使用 SPring 框架可以用面向对象的思想来实现 AOP。

Spring 框架中不需要创建 InvocationHandler，只需要创建一个切面对象，将所有的非业务代码在切面对象中完成即可，Spring 框架底层会自动根据切面类以及目标类生成一个代理对象。

LoggerAspect

```java
package com.bai.aop;

import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.springframework.stereotype.Component;

import java.util.Arrays;

/**
 * @author bai
 * @create 2020-06-18-17:45
 */

@Aspect
@Component
public class LoggerAspect {

    @Before("execution(public int com.bai.utils.impl.CalImpl.*(..))")
    public void before(JoinPoint joinPoint) {
        //获取方法名
        String name = joinPoint.getSignature().getName();
        //获取参数
        String args = Arrays.toString(joinPoint.getArgs());
        System.out.println(name + "方法的参数是：" + args);
    }
}
```

LoggerAspect 类定义处添加的两个注解：

- `@Aspect` ：表示该类是切面类。
- `@Component` ：将该类的对象注入到 IoC 容器。

具体方法处添加的注解：

`@Before` ：表示方法执行的具体位置和时机。

CalImpl：也需要添加 `@component` ，交给 IoC 容器来管理。

```java
package com.bai.utils.impl;

import com.bai.utils.Cal;
import org.springframework.stereotype.Component;

/**
 * @author bai
 * @create 2020-06-18-13:56
 */
@Component
public class CalImpl implements Cal {
    public int add(int num1, int num2) {
        int result = num1 + num2;
        return result;
    }

    public int sub(int num1, int num2) {
        int result = num1 - num2;
        return result;
    }

    public int mul(int num1, int num2) {
        int result = num1 * num2;
        return result;
    }

    public int div(int num1, int num2) {
        int result = num1 / num2;
        return result;
    }
}
```

spring.xml中配置 AOP。

``` xml
<?xml version="1.0" encoding="utf-8" ?>
<beans xmlns="http://www.springframework.org/schema/beans"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns:context="http://www.springframework.org/schema/context"
            xmlns:aop="http://www.springframework.org/schema/aop"
            xmlns:p="http://www.springframework.org/schema/p"
        xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans-4.3.xsd
        http://www.springframework.org/schema/aop
        http://www.springframewok.org/schema/aop/spring-aop-4.3.xsd
        http://www.springframework.org/schema/context
        http://www/springframework.org/schema/context/spring-context-4.3.xsd">

<!-- 自动扫描 -->
<context:component-scan base-package="com.bai"></context:component-scan>

<!-- 使Aspect注解生效，为目标类自动生成代理对象 -->
<aop:aspectj-autoproxy></aop:aspectj-autoproxy>
</beans>
```

`context:component-scan`  将 `com.bai` 包中的所有类进行扫描，如果该类同时添加了 `@Component` 注解，则将该类扫描到 IoC 容器中，即 IoC 管理他的对象。

`aop:aspectj-autoproxy` 让 Spring 框架结合切面类和目标类自动生成动态代理对象。

- 切面：横切关注点被模块化的抽象对象。
- 通知：切面对象完成的工作。
- 目标：被通知的对象，即被横切的对象。
- 代理：切面、通知、目标混合之后的对象。
- 连接点：通知要插入业务代码的具体位置。
- 切点：AOP 通过切点定位到连接点。

















































