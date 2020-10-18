<?php
//自定义函数
//在一个字符串中,匹配一个字符出现的次数.
function getCharNum($char,$str){
	$num = 0;
	//对字符串进行遍历
	for($i=0;$i<strlen($str);$i++){
		//遍历到的字符
		$curChar = $str[$i];
		if($curChar == $char){
			$num++;
		}
	}
	return $num;
}
$re = getCharNum('a','afsafawf');
var_dump($re);