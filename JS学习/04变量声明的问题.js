var a = "123";
var b = "456";
var c;
console.log(a);
console.log(b);
console.log(c);
//? 声明没有赋值的情况下，会默认赋值为undefined
//? JavaScript再解释执行之前，有预编译的过程，会将变量声明提升到最前面，但是不会将赋值提升
c = 789
console.log(d);
var d = 7890