var a = 123;
console.log(typeof(a + ''))
console.log(typeof(toString(a)))
var b = true
console.log(typeof(b + ''))
console.log(Number(a))
//? 注意：null和undefined这俩个值没有toString()方法，所以它们不能用方法二，如果调用，会报错。
console.log(String(null))
console.log(null)
//? 使用String()函数做强制类型转换时：对于Number和Boolean而言。实际上就是在调用toString()方法
console.log('-------------------')
var d
console.log(d)
//? prompt()：用户的输入，用户不管输入什么都是字符串类型
//? 转换成Number类型
var e = '123'
console.log(Number(a))
var f = null
console.log(Number(f))
var g = undefined
console.log(Number(g))
//? parseInt() 的作用是向下取整 只保留字符串最开头的数字，后面的中文自动消失
console.log(parseInt('2017年的第一场雪'))
//? 数字转换成布尔，除了0和NaN，其余都是true，字符串转换成布尔，除了空串，其余都是true
var h = Boolean(0)
console.log(h)
var i = Boolean(NaN)
console.log(i)
var j = Boolean("")
console.log(j)