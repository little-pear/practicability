var str = "我是一个字符串"
var num = 789
var bool = true
var undefine
var kong = ""
console.log(typeof a)
console.log(typeof num)
console.log(typeof bool)
console.log(typeof undefine )
console.log(typeof kong)
console.log("this is \"abc\"") //? \作为转义字符，本身并没有任何的的意义，若要显示\则需要写入两个\\
console.log("this is \\")
/**
 *  引号不能进行嵌套：双引号不能再放单引号，单引号不能再放单引号，但是单引号里可以嵌套双引号
*   转义字符
    在字符串中我们可以使用\作为转义字符，当表示一些特殊符号时可以使用\进行转义
    \" 表示 "
    \' 表示 '
    \n 表示 换行
    \r 表示 回车
    \t 表示 制表符
    \b 表示 空格
    \\ 表示 \

    任何的数据类型+字符串都会转换成字符串

    将其他数值转化成字符串的方法有三种
    1、拼串
    2、toString()
    3、String()
*/

var a = '100';
var b = 100;
console.log(typeof(a+b));

console.log('typeof 的使用');
console.log(typeof undefined);
console.log(typeof null);
