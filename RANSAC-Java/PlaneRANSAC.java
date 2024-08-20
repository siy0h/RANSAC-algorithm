package A1;

import java.io.IOException;

public class PlaneRANSAC {

    private PointCloud ref;
    private double epsilon;

    public PlaneRANSAC(PointCloud pc) {
        ref = pc;
        epsilon = 0;
    }

    public void setEps(double eps) {
        epsilon = eps;
    }

    public double getEps() {
        return epsilon;
    }

    public int getNumberOfIterations(double confidence, double percentageOfPointsOnPlane) {
        // makes sure that the input is correct
        if (percentageOfPointsOnPlane < 0 || percentageOfPointsOnPlane > 100) {
            throw new NumberFormatException("please enter a percentage between 0-100");
        }
        double k = 0;

        k = Math.log((1 - confidence / Math.log(1 - (Math.pow(percentageOfPointsOnPlane / 100, 3)))));

        return (int) k;
    }

    public void run(int NumberOfIterations, String filename) throws IOException {
        if (NumberOfIterations > 0) {

            PointCloud support = new PointCloud(); // creates an empty point cloud. this corresponds to a support
                                                   // value of zero

            for (int i = 0; i < NumberOfIterations; ++i) {
                PointCloud temp = new PointCloud(); // creates a tempoary cloud that will be compared with the
                                                    // support
                Plane3D plane = new Plane3D(ref.getPoint(), ref.getPoint(), ref.getPoint()); // draws a random plane

                for (int j = 0; j < ref.getCloud().size(); ++j) { // loop goes through all the points in the
                                                                  // original point cloud
                    Point3D x = ref.getCloud().get(j);
                    //System.out.println(x.toString());


                    if (plane.getDistance(x) < epsilon) { // compares the distance of plane from the point to the
                                                          // epsilon value
                        //System.out.println(x.toString());
                        temp.addPoint(x); // vales lower than the epsilon value get added to the temporary
                                          // PointCloud
                    }

                }
                //System.out.println(temp.getCloud().size());

                if (temp.getCloud().size() > support.getCloud().size()) { // comparing the size of the temporary
                                                                          // cloud to the support tells us what the
                                                                          // dominant plane should be
                    //System.out.println(support.getCloud().size());
                    support = temp; // with each iteration, we reassign the support value to ensure we're always
                                    // working with the plane that has the most points
                }

            }

            if (support != null) {
                // saves the points of the biggest point cloud to the solution file
                support.save(filename);
                for (int k = 0; k < support.getCloud().size(); k++) {
                    Point3D removal = support.getCloud().get(k);
                    ref.getCloud().remove(removal);// removes the points of the biggest point cloud from the
                                                   // original point cloud
                }

            }
        }
        return;

    }

    public static void main(String args[]) {
        //three most dominant planes for PointCloud3.xyz
        try {
            PointCloud pc = new PointCloud("PointCloud3.xyz");
            PlaneRANSAC ransac = new PlaneRANSAC(pc);
            ransac.setEps(0.1);

            ransac.run(168, "PointCloud3_p1");

            ransac.run(168, "PointCloud3_p2");

            ransac.run(168, "PointCloud3_p3");

        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

    }

}
