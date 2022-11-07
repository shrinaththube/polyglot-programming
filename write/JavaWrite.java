import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;

/**
 * JavaWrite
 */

// https://www.fillingthejars.com/november-quotes/
// https://everydaypower.com/november-quotes/
// https://www.w3schools.com/java/java_files_create.asp

// "My lovely November. Have you seen my heart, somewhere in your castle of yellow leaves?" – A Waltz for Zizi

// "November with uncanny witchery in its changed trees. With murky red sunsets flaming in smoky crimson behind the westering hills. With dear days when the austere woods were beautiful and gracious in a dignified serenity of folded hands and closed eyes–days full of a fine, pale sunshine that sifted through the late, leafless gold of the juniper-trees and glimmered among the grey beeches, lighting up evergreen banks of moss and washing the colonnades of the pines." — Lucy Maud Montgomery

// 12. "October extinguished itself in a rush of howling winds and driving rain and November arrived, cold as frozen iron, with hard frosts every morning and icy drafts that bit at exposed hands and faces." – J.K. Rowling

// 13. "November comes and November goes. With the last red berries and the first white snows. With night coming early and dawn coming late, and ice in the bucket and frost by the gate. The fires burn and the kettles sing, and earth sinks to rest until next spring." – Clyde Watson

public class JavaWrite {

    public static void main(String[] args) {
        
        String filename = "sample-write.txt";

        FileWriter fileWriter;
        try {
            fileWriter = new FileWriter(filename, true);
            fileWriter.write("\"In November you begin to know how long the winter will be.\" —Martha Gellhorn\n\n");
            fileWriter.close();
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        
        try {
            fileWriter = new FileWriter(filename,true);
            PrintWriter printWriter = new PrintWriter(fileWriter);
            // printWriter.print("Some String");
            printWriter.printf("%s\n\n%s\n\n%s\n\n%s\n\n", "\"The month of November makes me feel that life is passing more quickly. In an effort to slow it down, I try to fill the hours more meaningfully.\"  —Henry Rollins",
            "\"November’s sky is chill and drear, November’s leaf is red and sear.\" —Sir Walter Scott",
            "\"How sad would be November if we had no knowledge of the spring!\" —Edwin Way Teale, Circle of the Seasons",
            "\"Welcome sweet November, the season of senses and my favorite month of all.\" —Gregory F. Lenz");
            printWriter.close();    
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

    }
    
}