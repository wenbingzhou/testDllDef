// DllTestDef.h : DllTestDef DLL 的主头文件
//

#pragma once

//#ifndef __AFXWIN_H__
//	#error "在包含此文件之前包含“stdafx.h”以生成 PCH 文件"
//#endif
//
//#include "resource.h"		// 主符号
////#include <string>
//
//
//// CDllTestDefApp
//// 有关此类实现的信息，请参阅 DllTestDef.cpp
////
//
//class CDllTestDefApp : public CWinApp
//{
//public:
//	CDllTestDefApp();
//
//// 重写
//public:
//	virtual BOOL InitInstance();
//
//	DECLARE_MESSAGE_MAP()
//};
#ifndef TEST_API
#define TEST_API __declspec(dllexport)
#endif
extern "C"  TEST_API int  add(int x, int y);
extern "C"  TEST_API char*  addstr(char* x, char* y);
extern "C"  TEST_API int  addstr2(char* x, char* y);
extern "C"  TEST_API int my_strcpy(char *dest, const char *src);