package gocartesian

import (
    "fmt"
    "strings"
)

type Cartesian struct {
    sets [][]string
    results [][]string
}

func (c *Cartesian) appendPermutation(perm []string) {
    c.results = append(c.results, perm)
}

func (c *Cartesian) permute(arr []string, n int) {

    // If for some reason we didn't get any sets at all, just return empty to avoid errors
    if len(c.sets) < 1 {
        c.appendPermutation(arr)
        return
    }

    // Unexpected empty sets may break recursion, or worse. Just skip them, unless it's the
    // last set in which case we want to append whetever we're left with and exit
    if len(c.sets[n]) == 0 {
        if n+1 <= len(c.sets)-1 {
            c.permute(arr, n+1)
        } else {
            c.appendPermutation(arr)
            return
        }
    }

    // We're not out of range and we have some data, find the permutations
    for i, ln := 0, len(c.sets[n]); i < ln; i++ {
        tmp := append(arr, c.sets[n][i])
        if n < len(c.sets)-1 {
            c.permute(tmp, n+1)
        } else {
            c.appendPermutation(tmp)
        }
    }
}

func (c *Cartesian) getCartesianProduct() [][]string {
    c.permute([]string{}, 0)
    return c.results
}

/*-------- Public --------*/

// We'll just expose this for convenience. We can use it both internally and as a way to
// access the raw data without printing anything
func GetCartesianProduct(sets [][]string) [][]string {
    cart := Cartesian{sets, [][]string{}}
    return cart.getCartesianProduct()
}

// Will print the cartesian product of N number of arrays of strings in sets
func PrintCartesianProduct(sets [][]string) {
    permutations := GetCartesianProduct(sets)

    // Print it pretty
    for _, ln := range permutations {
        fmt.Println(strings.Join(ln, ","))
    }
}