// DllTestDef.cpp : ���� DLL �ĳ�ʼ�����̡�
//

#include "stdafx.h"
#include "DllTestDef.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <corecrt_malloc.h>
//#include <string.h>

//#ifdef _DEBUG
//#define new DEBUG_NEW
//#endif
//
////
////TODO:  ����� DLL ����� MFC DLL �Ƕ�̬���ӵģ�
////		��Ӵ� DLL �������κε���
////		MFC �ĺ������뽫 AFX_MANAGE_STATE ����ӵ�
////		�ú�������ǰ�档
////
////		����: 
////
////		extern "C" BOOL PASCAL EXPORT ExportedFunction()
////		{
////			AFX_MANAGE_STATE(AfxGetStaticModuleState());
////			// �˴�Ϊ��ͨ������
////		}
////
////		�˺������κ� MFC ����
////		������ÿ��������ʮ����Ҫ��  ����ζ��
////		��������Ϊ�����еĵ�һ�����
////		���֣������������ж������������
////		������Ϊ���ǵĹ��캯���������� MFC
////		DLL ���á�
////
////		�й�������ϸ��Ϣ��
////		����� MFC ����˵�� 33 �� 58��
////
//
//// CDllTestDefApp
//
//BEGIN_MESSAGE_MAP(CDllTestDefApp, CWinApp)
//END_MESSAGE_MAP()
//
//
//// CDllTestDefApp ����
//
//CDllTestDefApp::CDllTestDefApp()
//{
//	// TODO:  �ڴ˴���ӹ�����룬
//	// ��������Ҫ�ĳ�ʼ�������� InitInstance ��
//}
//
//
//// Ψһ��һ�� CDllTestDefApp ����
//
//CDllTestDefApp theApp;
//
//
//// CDllTestDefApp ��ʼ��
//
//BOOL CDllTestDefApp::InitInstance()
//{
//	CWinApp::InitInstance();
//
//	return TRUE;
//}

int  add(int x, int y)
{
	return x + y;
}
char*  addstr(char* x, char* y)
{
	int len = strlen(x) + strlen(y);
	char* m = (char*)malloc(len * sizeof(char));
	sprintf(m, "%s_%s", x,y);
	return m;
}

int  addstr2(char* x, char* y)
{
	int nx = atoi(x);
	int ny = atoi(y);
	return nx + ny;
	//return strlen(x) + strlen(y);
}

int my_strcpy(char *dest, const char *src)    //constʹ�ں����в����޸�*src��ԭ�ȵ�ֵ
{                                //����ԭʼ��strDest
	if (dest == NULL || src == NULL)
	{
		return 0;
	}
	while ((*dest++ = *src++) != '\0');
	return 1;
}