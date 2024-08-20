random3points(_,[]).
random3points(Points, [H|T]):-
    member(H, Points),
    random3points(Points, T).


plane(Point3, [A,B,C,D]):-
    Point3 = [[X1, YI, Z1], [X2, Y2, Z2], [X3, Y3, Z3]],
    A is Y1*(Z2-Z3) + Y2*(Z3-Z1) + Y3*(Z1-Z2),
    B is Z1*(X2-X3) + Z2*(X3-X1) + Z3*(X1-X2),
    C is X1*(Y2-Y3) + X2*(Y3-Y1) + X3*(Y1-Y2),
    D is -(X1*(Y2*Z3-Y3*Z2) + X2*(Y3*Z1-Y1*Z3) + X3*(Y1*Z2-Y2*Z1)).

distance([X,Y,Z],[A,B,C,D],Distance) :-
    Distance is abs(A*X + B*Y + C*Z + D) / sqrt(A^2 + B^2 + C^2).

ransac_number_of_iterations(Confidence, Percentage, N) :-
    Power is 1 - Percentage,
    Base is 1 - Confidence,
    N is ceil(log(Base) / log(Power)).


support(_,[],_,0).
support(Plane, Points, Eps, N) :-
    Points = [H|T],
    distance(H, Plane, Distance),
    (Eps >= Distance ->
        support(Plane, T, Eps, Count),
        N is Count + 1;
        support(Plane, T, Eps, N)).

% Test case for plane
test_plane(Plane) :-
    plane([[1, 4, 6], [2, 9, 8], [0, 5, 2]], Plane).

% Test case for distance
test_distance(Distance) :-
    distance([5, 6, 1], [7, 2, 1, 1], Distance).

% Test case for ransac_number_of_iterations
test_ransac_number_of_iterations(N) :-
    ransac_number_of_iterations(0.97, 0.45, N).