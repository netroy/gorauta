package gorauta

import (
  "math/rand"
  "time"
)

// Return a random entry in an array
func random(array []string) string {
  size := len(array)
  // if there is only one element, return that
  if size == 1 {
    return array[0]
  }
  // else return a random element
  rand.Seed(time.Now().Unix())
  return array[rand.Intn(size)]
}
