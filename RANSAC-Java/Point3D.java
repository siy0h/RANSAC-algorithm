package A1;

public class Point3D {
    private double x;
    private double y;
    private double z;

    public Point3D() {
        x = 0.0;
        y = 0.0;
        z = 0.0;

    }

    public Point3D(double xcord, double ycord, double zcord) {
        x = xcord;
        y = ycord;
        z = zcord;

    }

    public void setX(double num) {
        x = num;
    }

    public void setY(double num) {
        y = num;
    }

    public void setZ(double num) {
        z = num;
    }

    public double getX() {
        return x;
    }

    public double getY() {
        return y;
    }

    public double getZ() {
        return z;
    }

    // prints it as the same format of the .xyz file
    public String toString() {
        return String.valueOf(this.getX()) + "		" + String.valueOf(this.getY()) + "		"
                + String.valueOf(this.getZ());
    }

}
