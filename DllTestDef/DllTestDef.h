// DllTestDef.h : DllTestDef DLL ����ͷ�ļ�
//

#pragma once

//#ifndef __AFXWIN_H__
//	#error "�ڰ������ļ�֮ǰ������stdafx.h�������� PCH �ļ�"
//#endif
//
//#include "resource.h"		// ������
////#include <string>
//
//
//// CDllTestDefApp
//// �йش���ʵ�ֵ���Ϣ������� DllTestDef.cpp
////
//
//class CDllTestDefApp : public CWinApp
//{
//public:
//	CDllTestDefApp();
//
//// ��д
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