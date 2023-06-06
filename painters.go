package main

type(
    Painter struct {
        FirstName string
        LastName string
    }
)

func GetPainters()(painters []Painter) {
    aivazovski := Painter{ FirstName: "Ivan", LastName: "Aivazovski"}

    painters = append(painters, aivazovski,
        Painter{ FirstName: "Ivan", LastName: "Shishkin"},
    )

    return painters
}
