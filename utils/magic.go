// utils/magic.go
package utils

import "strings"

// MagicNumber struct with a single field Number.
type MagicNumber struct {
    Number int
}

// Multiply method for MagicNumber to multiply Number with n and update it.
func (m *MagicNumber) Multiply(n int) {
    m.Number *= n
}

// MagicSum returns the sum of n + n.
func MagicSum(n int) int {
    return n + n
}

// MagicPow returns n raised to the power of n.
func MagicPow(n int) int {
    result := 1
    for i := 0; i < n; i++ {
        result *= n
    }
    return result
}

// MagicOdd returns true if n is an odd number, otherwise false.
func MagicOdd(n int) bool {
    return n%2 != 0
}

// MagicGrade returns a string based on the grade of n.
func MagicGrade(n int) string {
    grades := map[int]string{
        0: "Zero",
        1: "Bad",
        2: "Ok",
        3: "Nice",
        4: "Awesome",
        5: "Exceptional",
    }
    if grade, exists := grades[n]; exists {
        return grade
    }
    return "Unknown"
}

// MagicName returns a slice of strings with the name repeated n times.
func MagicName(n int) []string {
    name := "radit"  // Replace "radit" with your actual name
    return strings.Split(strings.Repeat(name+",", n), ",")[:n]
}

// MagicTria returns the sum of numbers from 1 to n.
func MagicTria(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// MagicChange modifies n by doubling its value in place.
func MagicChange(n *int) {
    *n *= 2
}
