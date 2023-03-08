(ns adt
  (:require [clojure.java.io :as io]))

(defprotocol Shape
  (area [_])
  (circ [_]) ;; NOTE: added in Part 3
  )

(deftype Circle [radius]
  Shape
  (area [_] (* Math/PI radius radius))
  (circ [_] (* 2 Math/PI radius)) ;; NOTE: added in Part 3
  )

(deftype Rect [length height]
  Shape
  (area [_] (* length height))
  (circ [_] (+ length height length height)) ;; NOTE: added in Part 3
  )

(deftype Square [side] ;; NOTE: added in Part 4
  Shape
  (area [_] (* side side))
  (circ [_] (* 4 side))
  )

(println "Part 2")
(-> (Circle. 3) area println)
(-> (Rect. 3 5) area println)

(println "Part 3")
(-> (Circle. 3) ((juxt area circ)) println)
(-> (Rect. 3 5) ((juxt area circ)) println)

(println "Part 4")
(-> (Circle. 3) ((juxt area circ)) println)
(-> (Rect. 3 5) ((juxt area circ)) println)
(-> (Square. 3) ((juxt area circ)) println)

;; Part 3
;; Had to update the interface plus both implementations.

;; Part 4
;; In Haskell, you'd have to add a new 'branch' to the sumtype itself, 
;; then add the implementation to area and circ. Interestingly, neither change
;; would be allowed by the compiler independently.
;; While in C/Java, you *don't* update the interface (equiv to sumtype) and instead
;; just add the definition of Square
