package lenconv

func MToF(m Meter) Foot {
    return Foot(m / OneFootInMeters)
}

func FToM(f Foot) Meter {
    return Meter(f * OneFootInMeters)
}
