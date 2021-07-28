package weightconv

func KToP(k Kilogram) Pound {
    return Pound(k / OnePoundInKilograms)
}

func PToK(p Pound) Kilogram {
    return Kilogram(p * OnePoundInKilograms)
}
