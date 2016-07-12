#include <iostream>
using namespace std;
extern "C" {

#define true 1
#define false 0

struct Stack
{
	int * base;
	int * top;
	int size;
};

/*
* 展示栈内容
*/
extern void Display(struct Stack stk);

/*
* 判断栈顶元素距离栈底高度，Go不能做指针运算，需要提供该函数
*/
int before(int * a,int* b){
	return a - b;
}

/*
* Go 不能直接索引，需要提供该函数
*/
int indexing(int * base,int length){
	return base[length];
}

// 初始化
void Init(struct Stack * stk,int size){
	stk->base = new int[size];
	stk->top = stk->base;
	stk->size = size;
}

// 栈容量
int Capacity(Stack stk){
	return stk.size;
}

// 栈大小
int Length(struct Stack stk){
	return before(stk.top,stk.base);
}

// 是否为空
bool IsEmpty(struct Stack stk){
	return Length(stk)<=0;
}

// 是否满栈
bool IsFull(struct Stack stk){
	if (Length(stk) >= Capacity(stk))
	{
		return true;
	}
	return false;
}

// 压栈
bool Push(struct Stack * stk,int v){
	if (IsFull(*stk))
	{
		return false;
	}
	*(stk->top) = v;
	(stk->top)++;
	return true;
}

// 弹栈
int Pop(struct Stack * stk,bool * ok){
	if (IsEmpty(*stk))
	{
		*ok = false;
		return -1;
	}
	*ok = true;
	(stk->top)--;
	return *(stk->top);
}

// 查看栈顶元素
int Seek(struct Stack stk,bool * ok){
if (IsEmpty(stk))
	{
		*ok = false;
		return -1;
	}
	*ok = true;
	return *(stk.top - 1);
}

}
