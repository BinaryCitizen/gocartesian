gocartesian
===========

gocartesian.PrintCartesianProduct will take an array of N arrays of strings and print their cartesian product
gocartesian.GetCartesianProduct does the same as above, but returns the result as an array of strings instead of printing it

    sets := [][]string{[]string{"1", "2", "3"}, []string{"A", "B"}, []string{"X", "Y", "Z"}}
  	gocartesian.PrintCartesianProduct(sets)
  	/*
  	  1,A,X
      1,A,Y
      1,A,Z
      1,B,X
      1,B,Y
      1,B,Z
      2,A,X
      2,A,Y
      2,A,Z
      2,B,X
      2,B,Y
      2,B,Z
      3,A,X
      3,A,Y
      3,A,Z
      3,B,X
      3,B,Y
      3,B,Z
  	*/
