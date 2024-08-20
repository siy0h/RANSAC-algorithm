package A1;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Random;

public class PointCloud implements Iterable<Point3D> {
    //list to store all the points 
    private List<Point3D> cloud;
    

    public PointCloud(String filename) throws IOException {
        File file = new File(filename);
        BufferedReader br = new BufferedReader(new FileReader(file));
        cloud = new ArrayList<Point3D>();

        String temp[] = new String[3];
        //br.readLine(); // i tried to use this to skip the header but it wouldn't print
        // at all so I had to use the static isNumeric method instead.
        String st = br.readLine();

        while (st != null) {
            temp = st.split("\\s+"); // splits the line by whitespace and adds elements into an array

            if (isNumeric(temp[0]) && isNumeric(temp[1]) && isNumeric(temp[2])) { // makes sure the first line is of
                                                                                  // letters is not parsed
                Point3D point = new Point3D(Double.parseDouble(temp[0]), Double.parseDouble(temp[1]),
                        Double.parseDouble(temp[2])); // turn each element into a double
                cloud.add(point); // adds point into a cloud
            }
            st = br.readLine();

        }

        br.close();

    }

    public PointCloud() {
        cloud = new ArrayList<Point3D>(); // creates an empty list of points in case there is no file to read from
    }

    public boolean addPoint(Point3D pt) {
        if (cloud.add(pt)) {
            return true;
        }
        return false;

    }

    public void save(String filename) {
        try {
            File write = new File(filename + ".xyz"); // names the file
            FileWriter writer = new FileWriter(write);
            writer.write("x" + "	" + "y" + "	" + "z" + "\n"); // writes the points in the same format as the input
                                                                 // document
            for (int i = 0; i < cloud.size(); ++i) {
                writer.write(cloud.get(i).toString() + "\n");
            }
            writer.close();

            // exception handling
        } catch (IOException e) {
            System.out.println("There's a problem!");
        }
    }

    public Iterator<Point3D> iterator() {
        Iterator<Point3D> it = new Iterator<Point3D>() {
            private int index = 0;

            @Override
            public boolean hasNext() {
                return index > cloud.size();
            }

            @Override
            public Point3D next() {
                return cloud.get(index++);
            }

            @Override
            public void remove() {
                cloud.remove(cloud.get(index));

            }

        };
        return it;

    }

    public List<Point3D> getCloud() {
        return cloud;
    }

    public Point3D getPoint() {
        Random rand = new Random();
        return cloud.get(rand.nextInt(cloud.size()));
    }

    public static boolean isNumeric(String strNum) {
        // code from https://www.baeldung.com/java-check-string-number
        if (strNum == null) {
            return false;
        }
        try {
            double d = Double.parseDouble(strNum);
        } catch (NumberFormatException nfe) {
            return false;
        }
        return true;
    }

}
