VERSION 5.00
Begin VB.Form FrmMain 
   Caption         =   "Email群发器"
   ClientHeight    =   8160
   ClientLeft      =   120
   ClientTop       =   510
   ClientWidth     =   15510
   Icon            =   "FrmMain.frx":0000
   LinkTopic       =   "Form1"
   MaxButton       =   0   'False
   ScaleHeight     =   8160
   ScaleWidth      =   15510
   StartUpPosition =   2  '屏幕中心
   Begin VB.CommandButton Command9 
      Caption         =   "添加邮箱"
      Height          =   375
      Left            =   9840
      TabIndex        =   40
      Top             =   6840
      Width           =   2535
   End
   Begin VB.TextBox Text1 
      Height          =   375
      Left            =   9840
      TabIndex        =   39
      Top             =   6240
      Width           =   2535
   End
   Begin VB.CommandButton Command8 
      Caption         =   "删除邮箱"
      Height          =   315
      Left            =   11400
      TabIndex        =   38
      Top             =   7440
      Width           =   975
   End
   Begin VB.Frame Frame3 
      Caption         =   "主题列表"
      Height          =   4095
      Left            =   9720
      TabIndex        =   28
      Top             =   120
      Width           =   5415
      Begin VB.CommandButton Command4 
         Caption         =   "清空"
         Height          =   495
         Left            =   3720
         TabIndex        =   33
         Top             =   2640
         Width           =   1215
      End
      Begin VB.CommandButton Command3 
         Caption         =   "删除"
         Height          =   495
         Left            =   3720
         TabIndex        =   32
         Top             =   1920
         Width           =   1215
      End
      Begin VB.CommandButton Command1 
         Caption         =   "添加"
         Height          =   495
         Left            =   3720
         TabIndex        =   31
         Top             =   1080
         Width           =   1215
      End
      Begin VB.TextBox txtMaintitle 
         Height          =   420
         Left            =   240
         TabIndex        =   30
         Top             =   360
         Width           =   4680
      End
      Begin VB.ListBox List1 
         Height          =   2760
         Left            =   240
         TabIndex        =   29
         Top             =   1080
         Width           =   3255
      End
   End
   Begin VB.CommandButton Command2 
      Caption         =   "清空邮箱"
      Height          =   315
      Left            =   9840
      TabIndex        =   27
      Top             =   7440
      Width           =   1065
   End
   Begin VB.OptionButton Option1 
      Caption         =   "随机字符"
      Height          =   375
      Left            =   14160
      TabIndex        =   26
      Top             =   5040
      Width           =   1095
   End
   Begin VB.TextBox Textyanshi 
      Height          =   270
      Left            =   13200
      TabIndex        =   23
      Text            =   "5"
      Top             =   5040
      Width           =   495
   End
   Begin VB.TextBox txtXlsfile 
      Height          =   300
      Left            =   10800
      TabIndex        =   19
      Top             =   4560
      Width           =   2745
   End
   Begin VB.CommandButton CmdDr 
      Caption         =   "导入邮箱"
      Height          =   495
      Left            =   9840
      TabIndex        =   18
      Top             =   5640
      Width           =   2505
   End
   Begin VB.ComboBox cmbSheet 
      Height          =   300
      Left            =   10800
      Style           =   2  'Dropdown List
      TabIndex        =   16
      Top             =   5160
      Width           =   1185
   End
   Begin VB.CommandButton cmdSelectExclsFile 
      Caption         =   "选择EXCEL文件"
      Height          =   315
      Left            =   13680
      TabIndex        =   15
      Top             =   4560
      Width           =   1425
   End
   Begin VB.CommandButton cmdSend 
      Cancel          =   -1  'True
      Caption         =   "发送"
      Height          =   705
      Left            =   6720
      TabIndex        =   14
      Top             =   7080
      Width           =   2520
   End
   Begin VB.ListBox ListMailBox 
      Height          =   2220
      Left            =   12480
      TabIndex        =   13
      Top             =   5520
      Width           =   2655
   End
   Begin VB.Frame Frame2 
      Caption         =   "邮件内容"
      Height          =   6975
      Left            =   90
      TabIndex        =   7
      Top             =   915
      Width           =   9360
      Begin VB.CommandButton Command10 
         Caption         =   "清空编辑"
         Height          =   615
         Left            =   6120
         TabIndex        =   41
         Top             =   3840
         Width           =   2055
      End
      Begin VB.CommandButton Command7 
         Caption         =   "清空列表信息"
         Height          =   495
         Left            =   6120
         TabIndex        =   37
         Top             =   5400
         Width           =   2055
      End
      Begin VB.CommandButton Command6 
         Caption         =   "删除"
         Height          =   495
         Left            =   6120
         TabIndex        =   36
         Top             =   4680
         Width           =   2055
      End
      Begin VB.CommandButton Command5 
         Caption         =   "添加"
         Height          =   495
         Left            =   6120
         TabIndex        =   35
         Top             =   3120
         Width           =   2055
      End
      Begin VB.ListBox List2 
         Height          =   2940
         Left            =   840
         TabIndex        =   34
         Top             =   3120
         Width           =   5055
      End
      Begin VB.CommandButton CmdSelectFile 
         Caption         =   "选择附件"
         Height          =   315
         Left            =   4680
         TabIndex        =   12
         Top             =   6480
         Width           =   1140
      End
      Begin VB.TextBox txtFileName 
         Height          =   300
         Left            =   840
         TabIndex        =   11
         Top             =   6480
         Width           =   3645
      End
      Begin VB.TextBox txtBody 
         Height          =   2640
         Left            =   840
         MultiLine       =   -1  'True
         ScrollBars      =   2  'Vertical
         TabIndex        =   9
         Top             =   240
         Width           =   8400
      End
      Begin VB.Label Label6 
         AutoSize        =   -1  'True
         Caption         =   "附件："
         Height          =   180
         Left            =   240
         TabIndex        =   10
         Top             =   6480
         Width           =   540
      End
      Begin VB.Label Label5 
         AutoSize        =   -1  'True
         Caption         =   "编辑："
         Height          =   180
         Left            =   180
         TabIndex        =   8
         Top             =   765
         Width           =   540
      End
   End
   Begin VB.Frame Frame1 
      Caption         =   "你的邮箱资料"
      Height          =   735
      Left            =   60
      TabIndex        =   0
      Top             =   75
      Width           =   9435
      Begin VB.ComboBox cmbEmail 
         Height          =   300
         Left            =   2610
         Style           =   2  'Dropdown List
         TabIndex        =   21
         Top             =   255
         Width           =   1155
      End
      Begin VB.TextBox txtSmtpServer 
         Height          =   300
         Left            =   7500
         TabIndex        =   5
         Text            =   "smtp.163.com"
         Top             =   255
         Width           =   1740
      End
      Begin VB.TextBox txtUserPwd 
         Height          =   270
         IMEMode         =   3  'DISABLE
         Left            =   4650
         PasswordChar    =   "*"
         TabIndex        =   4
         Top             =   270
         Width           =   1665
      End
      Begin VB.TextBox txtUserName 
         Height          =   270
         Left            =   945
         TabIndex        =   1
         Top             =   270
         Width           =   1605
      End
      Begin VB.Label Label3 
         AutoSize        =   -1  'True
         Caption         =   "Smtp服务器："
         Height          =   180
         Left            =   6375
         TabIndex        =   6
         Top             =   315
         Width           =   1080
      End
      Begin VB.Label Label2 
         AutoSize        =   -1  'True
         Caption         =   "密  码："
         Height          =   180
         Left            =   3840
         TabIndex        =   3
         Top             =   315
         Width           =   720
      End
      Begin VB.Label Label1 
         AutoSize        =   -1  'True
         Caption         =   "用户名："
         Height          =   180
         Left            =   195
         TabIndex        =   2
         Top             =   315
         Width           =   720
      End
   End
   Begin VB.Label Label4 
      Height          =   255
      Left            =   120
      TabIndex        =   42
      Top             =   7920
      Width           =   8655
   End
   Begin VB.Label Label11 
      Caption         =   "秒"
      Height          =   255
      Left            =   13800
      TabIndex        =   25
      Top             =   5040
      Width           =   375
   End
   Begin VB.Label Label10 
      Caption         =   "延时:"
      BeginProperty Font 
         Name            =   "宋体"
         Size            =   12
         Charset         =   134
         Weight          =   400
         Underline       =   0   'False
         Italic          =   0   'False
         Strikethrough   =   0   'False
      EndProperty
      Height          =   255
      Left            =   12360
      TabIndex        =   24
      Top             =   5040
      Width           =   735
   End
   Begin VB.Label Label9 
      Caption         =   "延时：1秒=1000毫秒"
      BeginProperty Font 
         Name            =   "宋体"
         Size            =   9.75
         Charset         =   134
         Weight          =   400
         Underline       =   0   'False
         Italic          =   0   'False
         Strikethrough   =   0   'False
      EndProperty
      Height          =   375
      Left            =   360
      TabIndex        =   22
      Top             =   4320
      Width           =   1455
   End
   Begin VB.Label Label8 
      AutoSize        =   -1  'True
      Caption         =   "Excel文件："
      Height          =   180
      Left            =   9720
      TabIndex        =   20
      Top             =   4560
      Width           =   990
   End
   Begin VB.Label Label7 
      AutoSize        =   -1  'True
      Caption         =   "选择工作表："
      Height          =   180
      Left            =   9720
      TabIndex        =   17
      Top             =   5160
      Width           =   1080
   End
End
Attribute VB_Name = "FrmMain"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = False
Private Declare Sub Sleep Lib "kernel32" (ByVal dwMilliseconds As Long)

Private Declare Function GetOpenFileName Lib "comdlg32.dll" Alias "GetOpenFileNameA" (pOpenfilename As OPENFILENAME) As Long
Private Type OPENFILENAME
    lStructSize As Long
    hwndOwner As Long
    hInstance As Long
    lpstrFilter As String
    lpstrCustomFilter As String
    nMaxCustFilter As Long
    nFilterIndex As Long
    lpstrFile As String
    nMaxFile As Long
    lpstrFileTitle As String
    nMaxFileTitle As Long
    lpstrInitialDir As String
    lpstrTitle As String
    flags As Long
    nFileOffset As Integer
    nFileExtension As Integer
    lpstrDefExt As String
    lCustData As Long
    lpfnHook As Long
    lpTemplateName As String
End Type

Dim MailPassWord, MailUserName As String
Dim SendError As Boolean

Private Sub LoadSheet()
    On Error GoTo step_error
    Dim XLS As New Excel.Application
    Dim WRK As Excel.Workbook
    Dim SHT As Excel.Worksheet
    Dim RNG As Excel.Range

    Dim ArrayCells() As Variant

    If cmbSheet.ListIndex <> -1 Then
        '建立Excel新实例
        Set XLS = CreateObject("Excel.Application")
        '打开 XLS 文件
        Set WRK = XLS.Workbooks.Open(txtXlsfile.Text, False, True)
        '把当前选择的工作表赋值给SHT
        Set SHT = WRK.Worksheets(cmbSheet.List(cmbSheet.ListIndex))

        '得到当前工作表的使用范围
        Set RNG = SHT.UsedRange

        '重新分配数组
        ReDim ArrayCells(1 To RNG.Rows.Count, 1 To RNG.Columns.Count)

        '在使用范围内使用新的数组传值
        ArrayCells = RNG.Value

        '关闭工作表
        WRK.Close False
        '退出 Excel
        XLS.Quit

        '变量释放
        Set XLS = Nothing
        Set WRK = Nothing
        Set SHT = Nothing
        Set RNG = Nothing

        '网格数据显示设置

        For r = 0 To UBound(ArrayCells, 1) - 1
            For C = 0 To UBound(ArrayCells, 2) - 1
                ListMailBox.AddItem CStr(ArrayCells(r + 1, C + 1))
            Next
        Next
        
        
    Else
        MsgBox "请选择一个工作表!"
        cmbSheet.SetFocus
    End If
    Exit Sub
step_error:
    MsgBox Err.Number & " - " & Err.Description & "(" & "可能你打开的是空表" & ")", , "提示"
    Resume Next
End Sub

Private Sub CmdDr_Click()
LoadSheet
End Sub

Private Sub cmdSelectExclsFile_Click()
    Dim OFName As OPENFILENAME
    Dim XLS As Object
    Dim WRK As Object
    Dim SHT As Object

    OFName.lStructSize = Len(OFName)
    '父窗体
    OFName.hwndOwner = Me.hWnd
    '程序实例
    OFName.hInstance = App.hInstance
    '文件过滤
    OFName.lpstrFilter = "Excel 文件 (*.xls)" + Chr$(0) + "*.xls" + Chr$(0) + "所有文件 (*.*)" + Chr$(0) + "*.*" + Chr$(0)
    '建立文件缓冲区
    OFName.lpstrFile = Space$(254)
    '返回文件的最大长度
    OFName.nMaxFile = 255
    '建立文件标题缓存区
    OFName.lpstrFileTitle = Space$(254)
    '返回文件标题最大长度
    OFName.nMaxFileTitle = 255
    '默认目录
    OFName.lpstrInitialDir = App.Path
    '对话框标题
    OFName.lpstrTitle = "打开 XLS 文件"
    '无标志
    OFName.flags = 0

    '显示对话框
    If GetOpenFileName(OFName) Then
        txtXlsfile.Text = Trim$(OFName.lpstrFile)

        cmbSheet.Clear
        '建立Excel新实例
        Set XLS = CreateObject("Excel.Application")

        '打开XLS文件. UpdateLink = False 和 ReadOnly = True.
        Set WRK = XLS.Workbooks.Open(txtXlsfile.Text, False, True)
        '读取xls文件中的工作表
        For Each SHT In WRK.Worksheets
            '加载到列表框
            cmbSheet.AddItem SHT.Name
        Next

        cmbSheet.ListIndex = 0
        

        '关闭并不保存
        WRK.Close False
        '退出MS Excel
        XLS.Quit

        '释放变量
        Set XLS = Nothing
        Set WRK = Nothing
        Set SHT = Nothing
    End If

End Sub

Private Sub CmdSelectFile_Click()
    
    Dim OFName As OPENFILENAME
    OFName.lStructSize = Len(OFName)
    '父窗体
    OFName.hwndOwner = Me.hWnd
    '程序实例
    OFName.hInstance = App.hInstance
    '文件过滤
    OFName.lpstrFilter = "所有文件 (*.*)" + Chr$(0) + "*.*" + Chr$(0)
    '建立文件缓冲区
    OFName.lpstrFile = Space$(254)
    '返回文件的最大长度
    OFName.nMaxFile = 255
    '建立文件标题缓存区
    OFName.lpstrFileTitle = Space$(254)
    '返回文件标题最大长度
    OFName.nMaxFileTitle = 255
    '默认目录
    OFName.lpstrInitialDir = App.Path
    '对话框标题
    OFName.lpstrTitle = "选择附件"
    '无标志
    OFName.flags = 0

    '显示对话框
    If GetOpenFileName(OFName) Then
        txtFileName.Text = Trim$(OFName.lpstrFile)
    End If
End Sub


Private Sub cmdSend_Click()
 Dim i As Integer
 Dim xin As Integer
 Dim yanshi
 Dim youxiang
 Dim neirong
 Dim suiji
 yanshi = Textyanshi.Text * 1000
 If txtUserName.Text = "" Then
    MsgBox "请输入你的Email地址！", vbCritical
    txtUserName.SetFocus
    Exit Sub
 End If
 If txtUserPwd.Text = "" Then
    MsgBox "请输入你的密码！", vbCritical
    txtUserPwd.SetFocus
    Exit Sub
 End If
 
 If txtSmtpServer.Text = "" Then
    MsgBox "请输入smtp服务器！如：smtp.163.com", vbCritical
    txtSmtpServer.SetFocus
    Exit Sub
 End If
 
 If List1.ListCount = 0 Then
    MsgBox "请输入邮件主题！", vbCritical
    txtMaintitle.SetFocus
    Exit Sub
 End If
 
 If List2.ListCount = 0 Then
    MsgBox "请输入正文！", vbCritical
    txtBody.SetFocus
    Exit Sub
 End If
 
 If ListMailBox.ListCount = 0 Then
    MsgBox "没有可发送的邮箱地址！", vbCritical
    Exit Sub
 End If
 
 MailUserName = txtUserName.Text
 MailPassWord = txtUserPwd.Text
 
 'Me.MousePointer = 11 VbHourglass 11 沙漏（表示等待状态）
 xin = 0
 With ListMailBox
      For i = 0 To .ListCount - 1
       Randomize
 n = Int((List1.ListCount - 1) * Rnd) '随机取列表中的一行
 m = Int((List2.ListCount - 1) * Rnd)
Randomize Timer
HzAsc1 = Int(Rnd() * 10) + 127
HzAsc2 = Int(Rnd() * 10) + 127
HzAsc3 = Int(Rnd() * 10) + 127
HzAsc4 = Int(Rnd() * 10) + 127
HzAsc5 = Int(Rnd() * 10) + 127
HzAsc6 = Int(Rnd() * 10) + 127
HzAsc7 = Int(Rnd() * 10) + 127
suiji = ChrB(HzAsc1) + ChrB(HzAsc2) + ChrB(HzAsc3) + ChrB(HzAsc4) + ChrB(HzAsc5) + ChrB(HzAsc7) + ChrB(HzAsc2) + ChrB(HzAsc4) + ChrB(HzAsc4) + ChrB(HzAsc3) + ChrB(HzAsc6) + ChrB(HzAsc1) + ChrB(HzAsc5) + ChrB(HzAsc2) + ChrB(HzAsc7) + ChrB(HzAsc1) + ChrB(HzAsc7) + ChrB(HzAsc6) + ChrB(HzAsc3) + ChrB(HzAsc5) + ChrB(HzAsc4) + ChrB(HzAsc7) + ChrB(HzAsc2) + ChrB(HzAsc4) + ChrB(HzAsc5)
  suijiji = ChrB(HzAsc1) + ChrB(HzAsc2) + ChrB(HzAsc3) + ChrB(HzAsc4) + ChrB(HzAsc5)
  If Option1.Value = True Then
 neirong = List2.List(m) + suiji
 youxiang = List1.List(n) + suijiji
 Else
 neirong = List2.List(m)
 youxiang = List1.List(n)
 End If
          SendMail youxiang, neirong, txtFileName.Text, .List(i)
          xin = xin + 1
          Sleep yanshi
          FrmMain.Caption = xin & ":" & ListMailBox.List(i) & "发送完成"
      Next
      MsgBox "发送完毕,希望没被遮蔽", vbCritical
 End With
' Me.MousePointer = 0
End Sub

Private Sub Command1_Click()
If txtMaintitle.Text = "" Then
MsgBox "请输入内容", vbCritical
Else
List1.AddItem txtMaintitle.Text
End If
End Sub

Private Sub Command10_Click()
txtBody.Text = ""
End Sub

Private Sub Command2_Click()
ListMailBox.Clear
End Sub

Private Sub Command3_Click()
If List1.SelCount <> 0 Then List1.RemoveItem List1.ListIndex
End Sub

Private Sub Command4_Click()
List1.Clear
End Sub

Private Sub Command5_Click()
If txtBody.Text = "" Then
MsgBox "请输入内容", vbCritical
Else
List2.AddItem txtBody.Text
End If
End Sub

Private Sub Command6_Click()
If List2.SelCount <> 0 Then List2.RemoveItem List2.ListIndex
End Sub

Private Sub Command7_Click()
List2.Clear
End Sub


Private Sub Command8_Click()
If ListMailBox.SelCount <> 0 Then ListMailBox.RemoveItem ListMailBox.ListIndex
End Sub

Private Sub Command9_Click()
ListMailBox.AddItem Text1.Text
End Sub

Private Sub Form_Load()
    Shell "regsvr32 jmail.dll /s", vbNormalFocus  '注释：注册控件,无弹出对话框
    SendError = False
    MailUserName = ""
    MailPassWord = ""
    cmbEmail.AddItem "@163.com"
    cmbEmail.AddItem "@126.com"
    cmbEmail.AddItem "@sohu.com"
    cmbEmail.AddItem "@QQ.com"
    cmbEmail.ListIndex = 0
End Sub
Public Sub SendMail(Optional ByVal sSubject As String, _
                    Optional ByVal sBody As String, _
                    Optional ByVal sFileName As String, Optional ByVal MailTo As String)

    On Error GoTo ToExit '打开错误陷阱
    '------------------------------------------------

    Dim Jmail
    Dim ErrorTimes As Integer
    ErrorTimes = 0
    Set Jmail = CreateObject("jmail.Message")
    If sFileName <> "" Then Jmail.AddAttachment sFileName             '附件

    Jmail.Charset = "gb2312"
    Jmail.Silent = False
    Jmail.Priority = 3  '邮件状态,1-5 1为最高
    Jmail.MailServerUserName = MailUserName         '发件人Email帐号,自己改
    Jmail.MailServerPassWord = MailPassWord        '发件人Email密码,自己改

    Jmail.FromName = txtUserName            '发信人姓名,自己改
    Jmail.From = MailUserName + cmbEmail.Text   '发邮件地址,自己改

    Jmail.Subject = sSubject                  '主题
    Jmail.AddRecipient MailTo        '收信人地址
    Jmail.Body = sBody                      '信件正文

    Jmail.Send ("" & txtSmtpServer.Text & "")    'SMTP服务器，如smtp.sohu.com
    DoEvents

    Set Jmail = Nothing
    '------------------------------------------------
    Exit Sub
    '----------------
ToExit:
    ErrorTimes = ErrorTimes + 1
    If ErrorTimes < 5 Then Resume
    Select Case Jmail.ErrorCode
    Case 550
        MsgBox MailTo + "该邮件地址不存在，请更改后再发", , "提示"
    Case 535
        MsgBox "发件人的用户名或密码错误，请改正后再发", , "提示"
    Case Else
    Label4.Caption = Jmail.ErrorMessage
    End Select
    SendError = True
End Sub


