package main

import "fmt"

func main(){

jb := []string{"james","bond","is","awesome"}
xp := []string{"xbox","is","cool"}
// slice of slice of string
wp := [][]string{jb,xp}

fmt.Println(wp)



}