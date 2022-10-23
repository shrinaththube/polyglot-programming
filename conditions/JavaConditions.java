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

    }
    
}