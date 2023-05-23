using System.Collections.ObjectModel; 
using System.Management.Automation; 
using System.Management.Automation.Runspaces; 
using System.IO;
using System;
using System.Text;
namespace PSLess
{
 class PSLess
 {
   static void Main(string[] args)
   {
     if(args.Length ==0)
         Environment.Exit(1);
     string script=LoadScript(args[0]);
     string s=RunScript(script);
     Console.WriteLine(s);
     Console.ReadKey();
   }
 private static string LoadScript(string filename) 
 { 
   string buffer ="";
   try {
    buffer = File.ReadAllText(filename);
    }
   catch (Exception e) 
   { 
     Console.WriteLine(e.Message);
     Environment.Exit(2);
    }
  return buffer;
 }
 private static string RunScript(string script) 
 { 
    Runspace MyRunspace = RunspaceFactory.CreateRunspace();
    MyRunspace.Open();
    Pipeline MyPipeline = MyRunspace.CreatePipeline(); 
    MyPipeline.Commands.AddScript(script);
    MyPipeline.Commands.Add("Out-String");
    Collection<PSObject> outputs = MyPipeline.Invoke();
    MyRunspace.Close();
   StringBuilder sb = new StringBuilder(); 
   foreach (PSObject pobject in outputs) 
   { 
       sb.AppendLine(pobject.ToString()); 
   }
    return sb.ToString(); 
  }
 }
}