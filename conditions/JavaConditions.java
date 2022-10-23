import java.util.Random;

/*
 * run command
 * First compile java code -> javac JavaConditions.java
 * Run compiled java class -> java JavaConditions
 * If you change anything in the code, repeat first and second step to test your logic.
 */

public class JavaConditions {

    public static void main(String[] args) {
     
        boolean sun_is_bright = true;
        boolean moon_is_cool = true;

        // in km
        int earth_diameter = 12742;
        int moon_diameter = 3475;

        if(sun_is_bright){
            System.out.println("Shine Shine and Shine");
        }

        if(moon_is_cool){
            System.out.println("Chill vibes");
        }

        if (earth_diameter > moon_diameter){
            System.out.println("Mother Earth is definitely bigger than moon");
            sun_is_bright = false;
        } else{
            System.out.println("Hmmm. Something looks odd");
        }

        if(!sun_is_bright){
            System.out.println("I am white dwarf now");
        }

        Random randomChoice = new Random();
        int max = 8;
        int min = 1;
        int planetNo = randomChoice.nextInt((max - min) + 1) + min;

        //https://www.space.com/16080-solar-system-planets.html
        switch (planetNo) {
            case 1:
                System.out.println("I am Mercury. Closest to the sun");
                break;
            case 2:
                System.out.println("I am Venus. Second planet from the sun");
                break;
            case 3:
                System.out.println("Hey you know me. You live here. I am Earth");
                break;
            case 4:
                System.out.println("Most popular one for traveling talks. Many are trying to land on me. I am Mars");
                break;
            case 5:
                System.out.println("I am Jupiter. I am the largest plannet.");
                break;
            case 6:
                System.out.println("I am Saturn. I have largest ring.");
                break;
            case 7:
                System.out.println("I am Uranus. I have 27 moons and I have blue-green color.");
                break;
            case 8:
                System.out.println("I am Neptune. I am coldest one. You guys found me by math haha!");
                break;
            default:
                System.out.println("Don't call me planet. I am not discovered yet.");
                break;
        }

    }
    
}