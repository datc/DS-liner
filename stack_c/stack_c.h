#ifndef _TEST_H
#define _TEST_H

#ifdef __cplusplus
extern "C" {
#endif

typedef int bool;
#define true 1
#define false 0
struct Stack
{
	int * base;
	int * top;
	int size;
};

void Display(struct Stack stk);
int before(int * a,int* b);
int indexing(int * base,int length);

void Init(struct Stack * stk,int size);
int Capacity(struct Stack stk);
int Length(struct Stack stk);
bool IsEmpty(struct Stack stk);
bool IsFull(struct Stack stk);
bool Push(struct Stack * stk,int v);
int Pop(struct Stack * stk,bool * ok);
int Seek(struct Stack stk,bool * ok);

#ifdef __cplusplus
}
#endif

#endif
