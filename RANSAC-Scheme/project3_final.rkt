#lang scheme
(require scheme/math)

(define (readXYZ fileIn)
  (let ((sL (map (lambda s (string-split (car s)))
                          (cdr (file->lines fileIn)))))
    (map (lambda (L)
           (map (lambda (s)
(if (eqv? (string->number s) #f)
    s
    (string->number s))) L)) sL)))


  ;input: list[list[int]]
  ;output: list[int]
(define (pick3 Ps)
  ((lambda (x y z) (list x y z)) (list-ref Ps (random (length Ps))) (list-ref Ps (random (length Ps))) (list-ref Ps (random (length Ps)))))

  ;input: list[list[int]]
  ;output: list[int]
(define (plane P1 P2 P3)
  (let ((a1 (- (list-ref P2 0) (list-ref P1 0)))
        (b1 (- (list-ref P2 1) (list-ref P1 1)))
        (c1 (- (list-ref P2 2) (list-ref P1 2)))
        (a2 (- (list-ref P3 0) (list-ref P1 0)))
        (b2 (- (list-ref P3 1) (list-ref P1 1)))
        (c2 (- (list-ref P3 2) (list-ref P1 2))))
    ((lambda (a b c d) (list a b c d))
     (- (* b1 c2) (* b2 c1))
     (- (* a2 c1) (* a1 c2))
     (- (* a1 b2) (* b1 a2))
     (- (- (- (* (- (* b1 c2) (* b2 c1)) (list-ref P1 0))) (* (- (* b1 c2) (* b2 c1)) (list-ref P1 1))) (* (- (* a1 b2) (* b1 a2)) (list-ref P1 2))))))



  ;input: int int
  ;output int
(define (ransacNumberOfIterations confidence percentage)
  (let ((x (log (/ (- 1 confidence) (log (- 1 (expt (/ percentage 100) 3))))))) x))



 ; a helper function
  ; input: int list
  ;output: int
(define (count number L)
  (if (null? L)
      0
      (if (<= (car L) number)
          (+ 1 (count number (cdr L)))
          (count number (cdr L)))))


(define (distance point plane)
  (lambda (x y z)
    (abs (/ (+ (+ (* (list-ref plane 0) x) (* (list-ref plane 1) y))
                 (+ (* (list-ref plane 2) z) (list-ref plane 3)))
            (sqrt (+ (+ (expt (list-ref plane 0) 2) (expt (list-ref plane 1) 2))
                     (expt (list-ref plane 2) 2))))))  (list-ref point 0) (list-ref point 1) (list-ref point 2))




;input: in,list,list[int], list[int]
 ;output: int 
(define (support eps points plane)
  (let ((values(map (lambda (point)
    (let ((x (list-ref point 0))
          (y (list-ref point 1))
          (z (list-ref point 2)))
      (abs (/ (+ (+ (* (list-ref plane 0) x) (* (list-ref plane 1) y))
                   (+ (* (list-ref plane 2) z) (list-ref plane 3)))
              (sqrt (+ (+ (expt (list-ref plane 0) 2) (expt (list-ref plane 1) 2))
                       (expt (list-ref plane 2) 2)))))) )points)))
    (cons (count eps values) plane)))


(define (planeRANSAC filename confidence percentage eps)
    (let ((Ps (readXYZ filename))
        (k (ransacNumberOfIteration confidence percentage)))
    (dominantPlane Ps k)))

(define (dominantPlane Ps k eps)
  (let (count 0)
    ;i'm so sorry but i didnt do it! my brain hurts
  ))



; running the program
(define x (readXYZ "Point_Cloud_1_No_Road_Reduced.xyz"))
(define y (pick3 x))
(define z (plane (list-ref y 0) (list-ref y 1) (list-ref y 2)))
(define a (support 10 x z))

  
