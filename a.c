#include <stdio.h>

void main()
{
   /* 我的第一个 C 程序 */
	int a,b,tmp;
	printf("请输入两个整数");
   scanf("%d%d",&a,&b);
	if(a>b)
	{tmp=a;a=b;b=tmp;}
	printf("a=%d\nb=%d",a,b);
}
