import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Scanner;

/**
 * run command
 * First compile java code -> javac JavaLoops.java
 * Run compiled java class -> java JavaLoops
 * If you change anything in the code, repeat first and second step to test your logic.
 */

 // https://www.geeksforgeeks.org/different-ways-reading-text-file-java/
public class JavaRead {

    public static void main(String[] args) {
        String filename = "sample-read.txt";
        String stars = "****************************************";

        System.out.println(stars);
        System.out.println("1) Read line by line");
        System.out.println(stars);

        File fileoperator = new File(filename);
        
        try (Scanner sccanner = new Scanner(fileoperator)) {
            
            while (sccanner.hasNextLine())
                System.out.println(sccanner.nextLine());

        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

        System.out.println(stars);
        System.out.println("2) Read everything use delimiter");
        System.out.println(stars);

        try (Scanner sccanner = new Scanner(fileoperator)) {            
            
            sccanner.useDelimiter("\\Z");
            System.out.println(sccanner.next());

        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }        
        System.out.println(stars);
        System.out.println("3) Read everything");
        System.out.println(stars);

        try {
            String content  = new String(Files.readAllBytes(Paths.get(filename)));
            System.out.println(content);
        } catch (IOException e) {
            e.printStackTrace();
        }
        
    }
    
}

