var a = "100"
var b = 100
console.log(a - b)

//? isNaN 是不是一个非数字
//? isNaN()：任何不能呗转换为数值的值，都会让这个函数返回true，会对内容进行类型转换
console.log(isNaN(100))
console.log(isNaN("100"))
console.log(isNaN("blue"))
if ("100" == 100) {
    console.log(true)
} else {
    console.log(false)
}


//? 浮点数的运算是不精确的
var num1 = .1
var num2 = .2
console.log(num1 + num2)

//? 五种基本数据类型（string、number、boolean、null、undefined） 一种引用数据类型（object）