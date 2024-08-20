package A1;
//import java.math.*;

public class Plane3D {
    private double a;
    private double b;
    private double c;
    private double d;

    public Plane3D(Point3D p1,Point3D p2, Point3D p3){

        double[] equation = equationFinder(p1, p2, p3); //this gives the coefficients for the equations

        a = equation[0];
        b = equation[1];
        c = equation[2];
        d = equation[3];

    }

    public Plane3D(double p1, double p2, double p3, double p4){
        a = p1;
        b = p2;
        c = p3;
        d = p4;
    }

    public double getDistance(Point3D pt){
        //https://www.geeksforgeeks.org/distance-between-a-point-and-a-plane-in-3-d/
        double distance = Math.abs(a*pt.getX() +b*pt.getY() + c*pt.getZ() + d) / 
                Math.sqrt(a*a+ b*b+ c*c);
        return distance;
    }

    private double[] equationFinder(Point3D p1, Point3D p2, Point3D p3){
        //finds the equations from the points
        //https://www.geeksforgeeks.org/program-to-find-equation-of-a-plane-passing-through-3-points/
        
        double a1 = p2.getX() - p1.getX();
        double b1 = p2.getY() - p1.getY();
        double c1 = p2.getZ() - p1.getZ();
        double a2 = p3.getX() - p1.getZ();
        double b2 = p3.getY() - p1.getZ();
        double c2 = p3.getZ() - p1.getZ();
        double a = b1 * c2 - b2 * c1;
        double b = a2 * c1 - a1 * c2;
        double c = a1 * b2 - b1 * a2;
        double d = (- a * p1.getX() - b * p1.getY() - c * p1.getZ());
        double[] result = {a,b,c,d};
        // System.out.println("equation of plane is " + a +
        //                 " x + " + b + " y + " + c +
        //                 " z + " + d + " = 0.");
        return result;
                
    }

    // public static void main(String[] args){
    //     String hello = "hey";
    //     System.out.println("Hello World" +  hello);
    //     // Point3D a,b,c;
    //     // a = new Point3D(1,2,3);
    //     // b = new Point3D(4,5,6);
    //     // c = new Point3D(7,8,9);
    //     // Plane3D obj = new Plane3D(a,b,c);

    //     }
    
}
