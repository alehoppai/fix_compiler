package main

func PrependArray(arr []string, to []string) []string {
    for _, item := range arr {
        to = append(to, "")
        copy(to[1:], to)
        to[0] = item
    }

    return arr
}
