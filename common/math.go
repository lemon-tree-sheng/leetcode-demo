package common

func Max(a, b, c int) int {
    if a > b {
        if a > c {
            return a
        } else {
            return c
        }
    } else {
        if b > c {
            return b
        } else {
            return c
        }
    }
}

func Min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        } else {
            return c
        }
    } else {
        if b < c {
            return b
        } else {
            return c
        }
    }
}