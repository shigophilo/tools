using System;
using System.Security.Principal;
using System.DirectoryServices;
namespace ConsoleApp9
{
    class Program
    {
        static void Main(string[] args)
        {
            DirectoryEntry ldap_conn = new DirectoryEntry("LDAP://dc=mrhonest,dc=com");
            DirectorySearcher search = new DirectorySearcher(ldap_conn);
            String query = "(&(objectClass=computer))";//查找计算机
            search.Filter = query;
            foreach (SearchResult r in search.FindAll())
            {
                String mS_DS_CreatorSID = "";
                String computername = "";
                try
                {
                    computername = r.Properties["dNSHostName"][0].ToString();

                    mS_DS_CreatorSID = (new SecurityIdentifier((byte[])r.Properties["mS-DS-CreatorSID"][0], 0)).ToString();
                    //Console.WriteLine("{0} {1}\n", computername, mS_DS_CreatorSID);
                }
                catch
                {
                    ;
                }
                //再通过sid找用户名
                String UserQuery = "(&(objectClass=user))";
                DirectorySearcher search2 = new DirectorySearcher(ldap_conn);
                search2.Filter = UserQuery;

                foreach (SearchResult u in search2.FindAll())
                {
                    String user_sid = (new SecurityIdentifier((byte[])u.Properties["objectSid"][0], 0)).ToString();


                    if (user_sid == mS_DS_CreatorSID)
                    {
                        //Console.WriteLine("debug");
                        String username = u.Properties["name"][0].ToString();
                        Console.WriteLine("[*] [{0}] -> creator  [{1}]", computername, username);
                    }
                }

            }
        }
    }
}