#include <stdio.h>
#include <Windows.h>
#include <iostream>

#define SERVICE_NAME "Myservice"
#pragma comment(lib, "advapi32.lib")

void ServiceInstall()
{
    SC_HANDLE schSCManager;
    SC_HANDLE schService;
    char binpath[] = "c:\\programdata\\config.exe";
    char szPath[MAX_PATH] = { 0 };
    if (!GetModuleFileNameA(NULL, szPath, MAX_PATH))
        return;

    schSCManager = OpenSCManager(NULL, NULL, SC_MANAGER_ALL_ACCESS);
    if (!schSCManager)
        return;

    schService = CreateServiceA(schSCManager,
        SERVICE_NAME,
        SERVICE_NAME,
        SERVICE_ALL_ACCESS,
        SERVICE_WIN32_OWN_PROCESS,
        SERVICE_AUTO_START,
        SERVICE_ERROR_NORMAL,
        binpath,
        NULL,
        NULL,
        NULL,
        NULL,
        NULL);

    if (!schService)
    {
        CloseServiceHandle(schSCManager);
        return;
    }

    CloseServiceHandle(schSCManager);
    CloseServiceHandle(schService);
}
int main(int argc, char** argv)
{
    ServiceInstall();
    //   std::cout << "Hello World!\n";
    return 0;

}