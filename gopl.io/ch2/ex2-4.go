package ex2-4

func PopCount(x uint64) int {
    var pc int
    for i := 0; i < 64; i++ {
        pc += int(x & 1)
        x = x >> 1
    }
    return pc
}
