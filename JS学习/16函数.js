function addN(n) {
  var sum = 0;
  for (var i = 0; i <= 100; i++) {
    sum += i;
  }
  console.log(sum);
}


console.log("你好")
sayHello()

//? 定义函数
function sayHello() {
    console.log("欢迎")
    console.log("Welcome")
}

(function() {
    console.log("Hello World")
})()