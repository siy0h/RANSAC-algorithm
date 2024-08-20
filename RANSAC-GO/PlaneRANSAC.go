package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	//"time"
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}
type Plane3D struct {
	A float64
	B float64
	C float64
	D float64
}
type Plane3DwSupport struct {
	Plane3D
	SupportSize int
}

// reads an XYZ file and returns a slice of Point3D
func ReadXYZ(filename string) []Point3D {
	var pointsFromFile []Point3D

	file, ferr := os.Open(filename)

	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	scan.Scan()

	for scan.Scan() {
		line := scan.Text()
		temp_points := strings.Fields(line)

		X, err := strconv.ParseFloat(temp_points[0], 64)
		if err != nil {
			panic(err)
		}
		Y, _ := strconv.ParseFloat(temp_points[1], 64)
		Z, _ := strconv.ParseFloat(temp_points[2], 64)

		pointToAdd := Point3D{X, Y, Z}

		pointsFromFile = append(pointsFromFile, pointToAdd)

	}
	return pointsFromFile
}

// saves a slice of Point3D into an XYZ file
func SaveXYZ(filename string, points []Point3D) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	_, err2 := file.WriteString("X 	,	Y	,	Z	" + "\n")
	if err2 != nil {
		panic(err)
	}

	for i := 0; i < len(points); i++ {

		_, err2 := file.WriteString(fmt.Sprintf("%f", points[i].X) + "	" + fmt.Sprintf("%f", points[i].Y) + "	" + fmt.Sprintf("%f", points[i].Z) + "\n")

		if err2 != nil {
			panic(err)
		}
	}

	defer file.Close()
}

// // computes the distance between points p1 and p2
func (p1 *Point3D) GetDistance(p2 *Point3D) float64 {
	var result float64
	a := p2.X - p1.X
	b := p2.Y - p1.Y
	c := p2.Z - p1.Z

	result = math.Sqrt((math.Pow(a, 2) + math.Pow(b, 2) + math.Pow(c, 2)))

	return result
}

// computes the plane defined by a set of 3 points
func GetPlane(points []Point3D) Plane3D {
	if len(points) > 3 || points == nil || len(points) < 3 {
		panic("please give me an array of size 3... also learn how to make a try catch in Go")
	}
	p1 := points[0]
	p2 := points[1]
	p3 := points[2]
	a1 := p2.X - p1.X
	b1 := p2.Y - p1.Y
	c1 := p2.Z - p1.Z
	a2 := p3.X - p1.X
	b2 := p3.Y - p1.Y
	c2 := p3.Z - p1.Z
	a := b1*c2 - b2*c1
	b := a2*c1 - a1*c2
	c := a1*b2 - b1*a2
	d := (-a*p1.X - b*p1.Y - c*p1.Z)

	//fmt.Printf("equation of plane is %f + x  + %f +  y  + %f + z  + %f + = 0. \n", a, b, c, d)

	return Plane3D{a, b, c, d}

}

// // computes the number of required RANSAC iterations
func GetNumberOfIterations(confidence float64, percentageOfPointsOnPlane float64) int {
	if percentageOfPointsOnPlane < 0 || percentageOfPointsOnPlane > 100 {
		panic("please enter a percentage between 0-100")
	}

	k := math.Log((1 - confidence/math.Log(1-(math.Pow(percentageOfPointsOnPlane/100, 3)))))
	var j int = int(k)

	return j
}

// // computes the support of a plane in a set of points
func GetSupport(plane Plane3D, points []Point3D, eps float64) Plane3DwSupport {

	support := Plane3DwSupport{plane, 0}
	for i := 0; i < len(points); i++ {
		distance := math.Abs(plane.A*points[i].X+plane.B*points[i].Y+plane.C*points[i].Z+plane.D) /
			math.Sqrt(math.Pow(plane.A, 2)+math.Pow(plane.B, 2)+math.Pow(plane.C, 2))

		if distance < eps {
			support.SupportSize++
		}
	}

	return support
}

// extracts the points that supports the given plane and returns them as a slice of points
func GetSupportingPoints(plane Plane3D, points []Point3D, eps float64) []Point3D {
	support := Plane3DwSupport{plane, 0}
	var suporting_points []Point3D
	for i := 0; i < len(points); i++ {
		distance := math.Abs(plane.A*points[i].X+plane.B*points[i].Y+plane.C*points[i].Z+plane.D) /
			math.Sqrt(math.Pow(plane.A, 2)+math.Pow(plane.B, 2)+math.Pow(plane.C, 2))
		if distance < eps {
			support.SupportSize++
			suporting_points = append(suporting_points, points[i])
    
		}

	}   

	return suporting_points
}

// creates a new slice of points in which all points belonging to the plane have been removed
func RemovePlane(plane Plane3D, points []Point3D, eps float64) []Point3D {
	support := Plane3DwSupport{plane, 0}
	var final_points []Point3D
	for i := 0; i < len(points); i++ {
		distance := math.Abs(plane.A*points[i].X+plane.B*points[i].Y+plane.C*points[i].Z+plane.D) /
			math.Sqrt(math.Pow(plane.A, 2)+math.Pow(plane.B, 2)+math.Pow(plane.C, 2))

		if distance > eps {
			support.SupportSize++
			final_points = append(final_points, points[i])
		}
	}
	return final_points
}

func randomPointGenerator(array []Point3D) chan Point3D {
	rval := make(chan Point3D)
	i := rand.Intn(len(array))
	go func() {
		rval <- array[i]
		close(rval)
	}()

	return rval
}

func tripleRandomGenerator(inputStream chan Point3D) chan []Point3D {
	rval := make(chan []Point3D, 3)
	var points []Point3D

	for i := 0; i < 3; i++ {
		points = append(points, <-inputStream)
	}
	rval <- points

	return rval

}

func TakeN(array []Point3D, N int) chan []Point3D {
	rval := make(chan []Point3D, N)
	temp := make(chan Point3D, N*3)
	var test Point3D
	var test2 []Point3D

	for i := 0; i < N*3; i++ {
		test = <-randomPointGenerator(array)
		//fmt.Println(test)
		temp <- test
	}

	for i := 0; i < N; i++ {
		test2 = <-tripleRandomGenerator(temp)
		//fmt.Println(i)
		rval <- test2
	}

	return rval
}

func planeEstimator(array []Point3D) chan Plane3D {
	if len(array) > 3 {
		panic("please insert an array of size 3")
	}
	rval := make(chan Plane3D, 1)

	go func() {
		rval <- GetPlane(array)
		close(rval)
	}()

	return rval

}

func supportingPointFinder(eps float64, plane Plane3D, point_cloud []Point3D) chan Plane3DwSupport {
	rval := make(chan Plane3DwSupport, 1)
	rval <- GetSupport(plane, point_cloud, eps)

	return rval

}

func dominantPlaneFinder(instances chan Plane3DwSupport, support *Plane3DwSupport) {
	var instance Plane3DwSupport
	for i := 0; i < len(instances); i++ {
		instance = <-instances
		if instance.SupportSize > support.SupportSize {
			*support = instance
		}
	}

}

func main() {

	// read file
	filename := "PointCloud1.xyz"
	eps := 1.1
	point_cloud := ReadXYZ(filename)

	// find number of iterations
	iterations := GetNumberOfIterations(40, 95)
	//fmt.Println(iterations)

	planes := make(chan Plane3D, iterations)
	var temp_plane Plane3D

	temp := TakeN(point_cloud, iterations)

	//fmt.Println("jello")

	// create channel of random points
	for i := 0; i < iterations; i++ {
		temp_plane = <-planeEstimator(<-temp)
		planes <- temp_plane
	}
	//fmt.Println("jello2")

	// find the number of supports
	planeswSupport := make(chan Plane3DwSupport, iterations)

	var temp_planeSupport Plane3DwSupport

	for i := 0; i < iterations; i++ {
		temp_planeSupport = <-supportingPointFinder(eps, <-planes, point_cloud)
		planeswSupport <- temp_planeSupport
		//fmt.Println(temp_planeSupport.SupportSize)

	}

	//finds the dominant plane
	var planeTemp Plane3DwSupport
	planeTemp = <-planeswSupport
	//fmt.Println(planeTemp)

	dominantPlaneFinder(planeswSupport, &planeTemp)
	SaveXYZ(filename+"solution.xyz", RemovePlane(planeTemp.Plane3D, point_cloud, eps))

}


// (define (support plane points)
//   ;in: int list[int] list[list[int]]
//   (let (formula
//   ((lambda (x y z)
//     (abs (/ (+ (* (list-ref plane 0) x) (* (list-ref plane 1) y) (* (list-ref plane 2) z)  (list-ref plane 3))
//             (sqrt (+ (+ (expt (list-ref plane 0) 2) (expt (list-ref plane 1) 2)) (expt (list-ref plane 2) 2)))))))
//   (let k (map formula points)
     
  
//   )

