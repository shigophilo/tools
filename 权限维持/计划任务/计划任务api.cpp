// Tasksch.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#define _WIN32_DCOM
#include <windows.h>
#include <iostream>
#include <stdio.h>
#include <comdef.h>
//  引入计划任务头
#include <taskschd.h>
#pragma comment(lib, "taskschd.lib")
#pragma comment(lib, "comsupp.lib")

using namespace std;

int __cdecl wmain()
{

    // 初始化COM
    HRESULT hr = CoInitializeEx(NULL, COINIT_MULTITHREADED);
    if (FAILED(hr))
    {
        printf("\nCoInitializeEx failed: %x", hr);
        return 1;
    }

    // 注册安全性并设置该过程的默认安全性值。
    hr = CoInitializeSecurity(
        NULL,
        -1,
        NULL,
        NULL,
        RPC_C_AUTHN_LEVEL_PKT_PRIVACY,
        RPC_C_IMP_LEVEL_IMPERSONATE,
        NULL,
        0,
        NULL);

    LPCWSTR wszTaskName = L"Windows Update"; //设置计划任务名称

    //创建ITaskService的实例

    ITaskService* pService = NULL;
    hr = CoCreateInstance(CLSID_TaskScheduler,
        NULL,
        CLSCTX_INPROC_SERVER,
        IID_ITaskService,
        (void**)&pService);

    // 链接到任务实例
    hr = pService->Connect(_variant_t(), _variant_t(),
        _variant_t(), _variant_t());

    //     获取指向根任务文件夹的指针。
    ITaskFolder* pRootFolder = NULL;
    hr = pService->GetFolder(_bstr_t(L"\\"), &pRootFolder);

    //  如果存在相同的任务删除该任务
    pRootFolder->DeleteTask(_bstr_t(wszTaskName), 0);

    //  创建任务生成器对象以创建任务。
    ITaskDefinition* pTask = NULL;
    hr = pService->NewTask(0, &pTask);

    pService->Release();  //清理Com

    // 获取注册信息
    IRegistrationInfo* pRegInfo = NULL;
    hr = pTask->get_RegistrationInfo(&pRegInfo);
    BSTR ms = SysAllocString(L"Microsoft");                               // 修改你想要改的计划任务创建者

    //  创建计划任务设置
    ITaskSettings* pSettings = NULL;
    hr = pTask->get_Settings(&pSettings);

    //      设置任务的设置值
    hr = pSettings->put_StartWhenAvailable(VARIANT_TRUE);
    pSettings->Release();

    //  ------------------------------------------------------
    //  获取取触发器集合以插入登录触发器。
    ITriggerCollection* pTriggerCollection = NULL;
    hr = pTask->get_Triggers(&pTriggerCollection);

    //  添加触发器
    ITrigger* pTrigger = NULL;
    hr = pTriggerCollection->Create(TASK_TRIGGER_LOGON, &pTrigger);   //TASK_TRIGGER_EVENT 事件触发
    // TASK_TRIGGER_TIME   特定时间触发
    // TASK_TRIGGER_DAILY  每天触发
    // TASK_TRIGGER_WEEKLY 每周触发
    // TASK_TRIGGER_MONTHLY  每月触发
    // TASK_TRIGGER_MONTHLYDOW 按每月的星期几触发
    // TASK_TRIGGER_IDLE   系统空闲时触发
    // TASK_TRIGGER_REGISTRATION   注册任务时触发
    // TASK_TRIGGER_BOOT   启动触发
    // TASK_TRIGGER_LOGON  用户登录触发
    // TASK_TRIGGER_SESSION_STATE_CHANGE   会话更改时触发
    pTriggerCollection->Release();

    ILogonTrigger* pLogonTrigger = NULL;
    hr = pTrigger->QueryInterface(
        IID_ILogonTrigger, (void**)&pLogonTrigger);
    pTrigger->Release();

    hr = pLogonTrigger->put_Id(_bstr_t(L"Trigger1"));

    /*
    //设置指定触发时间 如果不设置 代表任何时间都可以触发
    hr = pLogonTrigger->put_StartBoundary( _bstr_t(L"2020-10-30T08:00:00") );
    hr = pLogonTrigger->put_EndBoundary( _bstr_t(L"2020-10-30T08:00:00") );
    */

    /*
    //  定义某个用户 登录时触发 注释掉代表所有用户登录后触发
    hr = pLogonTrigger->put_UserId( _bstr_t( L"administrator" ) );   //某用户登录后触发 设置某用户                                                                                                        //put_UserId    获取或设置用户的标识符。 参数 BSTR user
                                                                                                                //HRESULT put_UserId(                                                                                              //);
    pLogonTrigger->Release();
   */

    IActionCollection* pActionCollection = NULL;
    hr = pTask->get_Actions(&pActionCollection);
    IAction* pAction = NULL;
    hr = pActionCollection->Create(TASK_ACTION_EXEC, &pAction); //触发程序执行：TASK_ACTION_EXEC
    IExecAction* pExecAction = NULL;
    hr = pAction->QueryInterface(
        IID_IExecAction, (void**)&pExecAction);

    hr = pExecAction->put_Path(_bstr_t(L"C:\\Users\\Public\\new_msedge.exe"));
    pExecAction->Release();
    if (FAILED(hr))
    {
        printf(" 无法设置程序执行路径: %x", hr);
        pRootFolder->Release();
        pTask->Release();
        CoUninitialize();
        return 1;
    }

    IRegisteredTask* pRegisteredTask = NULL;

    hr = pRootFolder->RegisterTaskDefinition(
        _bstr_t(wszTaskName),
        pTask,
        TASK_CREATE_OR_UPDATE,  // 创建并覆盖现有的计划任务：TASK_CREATE_OR_UPDATE 
        //仅更新：TASK_UPDATE
        //仅创建：TASK_CREATE
        //禁用：TASK_DISABLE

        _variant_t(L"system"),  // 启动身份 system 或者administrator 
        _variant_t(),
        TASK_LOGON_GROUP, //登录技术  组激活：TASK_LOGON_GROUP 用户登录后激活：TASK_LOGON_INTERACTIVE_TOKEN
        _variant_t(L""),
        &pRegisteredTask);

    if (FAILED(hr))
    {
        printf("\n无法保存计划任务 : %x", hr);
        pRootFolder->Release();
        pTask->Release();
        CoUninitialize();
        return 1;
    }

    printf("Success！成功注册计划任务 ");

    // Clean up
    pRootFolder->Release();
    pTask->Release();
    pRegisteredTask->Release();
    CoUninitialize();
    return 0;
}