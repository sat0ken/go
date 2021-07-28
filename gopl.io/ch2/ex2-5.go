package ex2-5

func PopCount(x uint64) int {
    var pc int
    for ; x > 0; pc++ {
        x = x & (x - 1)
    }
    return pc
}
