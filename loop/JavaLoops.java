/*
 * run command
 * First compile java code -> javac JavaLoops.java
 * Run compiled java class -> java JavaLoops
 * If you change anything in the code, repeat first and second step to test your logic.
 */

class JavaLoops{
    
    public static void main(String[] args){

        String coolArray [] = {"i" , "am", "reading", "all" , "elements", "of", "this", "array"};
        String stars = "****************************************";

        // while loop
        System.out.println(stars);
        System.out.println("1) Let's start with 'while' loop");
        System.out.println(stars);

        int ind = 0;

        while(ind < coolArray.length){
            System.out.println(coolArray[ind]);
            ind += 1;
        }
        
        //-------------------------------------------------------------------------------------------//

        // do-while loop: Make sure array/list is not empty
        System.out.println(stars);
        System.out.println("2) Variation of 'while' loop - 'do-while'");
        System.out.println(stars);

        ind = 0;
        do{
            System.out.println(coolArray[ind]);
            ind += 1;
        }while(ind < coolArray.length);

        //-------------------------------------------------------------------------------------------//

        // for loop
        System.out.println(stars);
        System.out.println("3) Move on to 'for' loop");
        System.out.println(stars);

        for(int i = 0 ; i < coolArray.length ; i++){
            System.out.println(coolArray[i]);
        }

        //-------------------------------------------------------------------------------------------//

        // for loop access element by element
        System.out.println(stars);
        System.out.println("4) Variation of 'for' loop");
        System.out.println(stars);

        for(String ele : coolArray){
            System.out.println(ele);
        }

    }
}