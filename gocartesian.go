package gocartesian

import (
    "fmt"
    "strings"
)

// Returns a function that, when called, will append the cartesian product from N number of arrays in s to 
// the array results.
func cartesianMaker(s [][]string, results *[][]string) (permute func(arr []string, n int)) {
    sets := s
    numSets := len(sets) - 1
    appendPerm := func(perm []string){
        *results = append(*results, perm)
    }
    
    return func(arr []string, n int) {

        // If for some reason we didn't get any sets at all, just return empty to avoid errors
        if len(sets) < 1 {
            appendPerm(arr)
            return
        }

        // Unexpected empty sets may break recursion, or worse. Handle them
        if len(sets[n]) == 0 {
            // Set is empty, just skip it if it's not the last one
            if n+1 <= numSets {
                permute(arr, n+1)
            } else {
                // Last set is empty, just append what we're left with from the previous one and exit
                appendPerm(arr)
                return
            }
        }

        // We're not out of range and we have some data, find the permutations
        for i, ln := 0, len(sets[n]); i < ln; i++ {
            tmp := append(arr, sets[n][i])
            if n < numSets {
                permute(tmp, n+1)
            } else {
                appendPerm(tmp)
            }
        }
    }
}

func getPermutations(sets [][]string) [][]string {
    permutations := [][]string{}
    permute := cartesianMaker(sets, &permutations)
    permute([]string{}, 0)
    return permutations
}

// We'll just expose this for convenience. We can use it both internally and as a way to
// access the raw data without printing anything
func GetCartesianProduct(sets [][]string) [][]string {
    return getPermutations(sets)
}

func PrintCartesianProduct(sets [][]string) {
    permutations := GetCartesianProduct(sets)

    // Print it pretty
    for _, ln := range permutations {
        fmt.Println(strings.Join(ln, ","))
    }
}