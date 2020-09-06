# Spring Boot 项目

- 特点

1、不需要 web.xml

2、不需要 springmvc.xml

3、不需要 tomcat、Spring Boot 内嵌了 tomcat

4、不需要配置 JSON 解析，支持 REST 架构

5、个性化配置非常简单

- 如何使用

1、创建 Maven 工程，导入相关依赖。

```xml
<!-- 继承父包 -->
<parent>
    <artifactId>spring-boot-starter-parent</artifactId>
    <groupId>org.springframework.boot</groupId>
    <version>2.3.1.RELEASE</version>
</parent>

<dependencies>
    <!-- web启动jar -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <version>1.18.12</version>
        <scope>provided</scope>
    </dependency>
</dependencies>
```

2、创建 Student 实体类

```java
package com.bai.entity;

import lombok.Data;

@Data
public class Student {
    private long id;
    private String name;
    private int age;
}
```

3、StudentRepository

```java
package com.bai.Repository;

import com.bai.entity.Student;

import java.util.Collection;

public interface StudentRepository {
    public Collection<Student> findAll();
    public Student findById(long id);
    public void saveOrUpdate(Student student);
    public void deleteById(long id);
}
```

4、StudnetRepositoryImpl

```java
package com.bai.Repository.impl;

import com.bai.Repository.StudentRepository;
import com.bai.entity.Student;
import org.springframework.stereotype.Repository;

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;

@Repository
public class StudentRepositoryImpl implements StudentRepository {

    private static Map<Long, Student> studentMap;

    static {
        studentMap = new HashMap<>();
        studentMap.put(1L,new Student(1L,"张三",22));
        studentMap.put(2L,new Student(2L,"李四",20));
        studentMap.put(3L,new Student(3L,"王五",18));
    }

    @Override
    public Collection<Student> findAll() {
        return studentMap.values();
    }

    @Override
    public Student findById(long id) {
        return studentMap.get(id);
    }

    @Override
    public void saveOrUpdate(Student student) {
        studentMap.put(student.getId(), student);
    }

    @Override
    public void deleteById(long id) {
        studentMap.remove(id);
    }
}
```

5、StudentHandler

```java
package com.bai.controller;

import com.bai.Repository.StudentRepository;
import com.bai.entity.Student;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.Collection;

@RestController
@RequestMapping("/student")
public class StydentHandler {

    @Autowired
    private StudentRepository studentRepository;

    @GetMapping("/findAll")
    public Collection<Student> findAll() {
        return studentRepository.findAll();
    }

    @GetMapping("/findById/{id}")
    public Student findById(@PathVariable("id") long id) {
        return studentRepository.findById(id);
    }

    @PostMapping("/save")
    public void save(@RequestBody Student student) {
        studentRepository.saveOrUpdate(student);
    }

    @PutMapping("/update")
    public void update(@RequestBody Student student) {
        studentRepository.saveOrUpdate(student);
    }

    @DeleteMapping("/deleteById/{id}")
    public void deleteById(@PathVariable("id") long id) {
        studentRepository.deleteById(id);
    }
}
```

6、application.yml

```yaml
server:
  port: 9090
```

7、启动类

```java
package com.bai;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class,args);
    }
}
```

`@SpringBootApplication` 表示当前类是 Spring Boot 的入口，Application 类的存放位置必须其他业务相关类的存放位置的父级。

### Spring Boot 整合 JSP

- pom.xml

```xml
<parent>
	<artifactId>spring-boot-starter-parent</artifactId>
	<groupId>org.springframework.boot</groupId>
	<version>2.3.0.RELEASE</version>
</parent>

<dependencies>
	<dependency>
  		<!-- Web -->
 		<groupId>org.springframework.boot</groupId>
  		<artifactId>spring-boot-starter-web</artifactId>
	</dependency>

	<!-- 整合 JSP -->
    <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-tomcat</artifactId>
    </dependency>
    <dependency>
      <groupId>org.apache.tomcat.embed</groupId>
      <artifactId>tomcat-embed-jasper</artifactId>
    </dependency>

    <!-- JSTL -->
    <dependency>
      <groupId>jstl</groupId>
      <artifactId>jstl</artifactId>
      <version>1.2</version>
    </dependency>
    
    <dependency>
      <groupId>org.projectlombok</groupId>
      <artifactId>lombok</artifactId>
      <version>1.18.12</version>
      <scope>provided</scope>
    </dependency>
</dependencies>
```

- 创建配置文件 application.yml

```yaml
server:
  port: 8181
spring:
  mvc:
    view:
      prefix: /
      suffix: .jsp
```

- 创建 Handler

```java
package com.bai.controller;

import com.bai.Repository.StudentRepository;
import com.bai.entity.Student;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

/**
 * @author bai
 * @create 2020-06-20-15:36
 */
@Controller
@RequestMapping("/hello")
public class HelloHandler {

    @Autowired
    private StudentRepository studentRepository;

    @GetMapping("/index")
    public ModelAndView index() {
        ModelAndView modelAndView = new ModelAndView();
        modelAndView.setViewName("index");
        modelAndView.addObject("list",studentRepository.findAll());
        return modelAndView;
    }

    @GetMapping("/deleteById/{id}")
    public String deleteById(@PathVariable("id") long id) {
        studentRepository.deleteById(id);
        return "redirect:/hello/index";
    }

    @PostMapping("/save")
    public String save(Student student) {
        studentRepository.saveOrUpdate(student);
        return "redirect:/hello/index";
    }

    @PostMapping("/update")
    public String update(Student student) {
        studentRepository.saveOrUpdate(student);
        return "redirect:/hello/index";
    }

    @GetMapping("findById/{id}")
    public ModelAndView findById(@PathVariable("id") long id) {
        ModelAndView modelAndView = new ModelAndView();
        modelAndView.setViewName("update");
        modelAndView.addObject("student",studentRepository.findById(id));
        return modelAndView;
    }
}
```

- JSP

```jsp
<%--
  Created by IntelliJ IDEA.
  User: 24378
  Date: 2020/6/20
  Time: 15:49
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ page isELIgnored="false" %>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
    <h1>学生信息</h1>
    <table>
        <tr>
            <th>学生编号</th>
            <th>学生姓名</th>
            <th>学生年龄</th>
            <th>操作</th>
        </tr>
        <c:forEach items="${list}" var="student">
            <tr>
                <td>${ student.id }</td>
                <td>${ student.name }</td>
                <td>${ student.age }</td>
                <td>
                    <a href="/hello/findById/${ student.id }">修改</a>
                    <a href="/hello/deleteById/${ student.id }">删除</a>
                </td>
            </tr>
        </c:forEach>
    </table>
    <a href="/save.jsp">添加学生</a>
</body>
</html>
```

```jsp
<%--
  Created by IntelliJ IDEA.
  User: 24378
  Date: 2020/6/20
  Time: 16:08
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
    <form action="/hello/save" method="post">
        ID:<input type="text" name="id" /><br>
        name:<input type="text" name="name" /><br>
        age:<input type="text" name="age" /><br>
        <input type="submit" value="提交" />
</body>
</html>
```

```jsp
<%--
  Created by IntelliJ IDEA.
  User: 24378
  Date: 2020/6/20
  Time: 16:08
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
    <form action="/hello/update" method="post">
        ID:<input type="text" name="id" value="${ student.id }" readonly /><br>
        name:<input type="text" name="name" value="${ student.name }" /><br>
        age:<input type="text" name="age" value="${ student.age }" /><br>
        <input type="submit" value="提交" />
</body>
</html>
```

### Spring Boot HTML

Spring Boot 可以结合 Thymeleaf 模板来整合 HTML，使用原生的 HTML 来作为视图。

Thymeleaf 模板是面向 Web 和独立环境的 Java 模板引擎，能够处理 HTML、XML、JavaScript、CSS 等。

```html
<P th:text="${ message }"></P>
```

- pom.xml 添加相关配置

```xml
<!-- 继承父包 -->
<parent>
    <artifactId>spring-boot-starter-parent</artifactId>
    <groupId>org.springframework.boot</groupId>
    <version>2.3.1.RELEASE</version>
</parent>

<dependencies>
    <!-- web启动jar -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>

    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <version>1.18.12</version>
        <scope>provided</scope>
    </dependency>

    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-thymeleaf</artifactId>
    </dependency>
</dependencies>
```

- application.yml

```yaml
server:
  port: 9090
spring:
  thymeleaf:
    prefix: calsspath:/templates/
    suffix: .html
    mode: HTML5
    encoding: utf-8
```

- Handler

```java
package com.bai.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/index")
public class IndexHandler {

    @GetMapping("/index")
    public String index() {
        System.out.println("index..");
        return  "index";
    }
}
```

- HTML

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <h1>Hello World</h1>
</body>
</html>
```

如果希望客户端可以直接访问 HTML 资源，将这些资源放置在 static 路径下即可，否则必须通过 Handler 的后台映射下才可以访问静态资源。