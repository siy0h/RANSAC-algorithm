# RANSAC-algorithm
This project implements an algorithm to detect the three most dominant planes in a cloud of 3D points using the RANSAC algorithm. The dominant plane is identified as the one containing the largest number of points, with a distance threshold (ε) defining point inclusion.

Application in LiDAR:
The algorithm is particularly relevant for autonomous vehicles equipped with LiDAR (Light Detection And Ranging) sensors. LiDAR systems scan the environment by emitting laser beams and capturing the reflected light to generate a point cloud representing the surrounding scene. In this context, the project aims to identify the main planar structures—such as roads and building facades—within the captured point cloud, aiding in scene understanding and vehicle navigation.
