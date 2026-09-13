package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, fmt.Errorf("Input must be positive number")
    }
    count := 0
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }
        count++
    }
    return count, nil
}
