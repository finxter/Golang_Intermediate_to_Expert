package animals

// Lion represents information about lions.
type Lion struct {
    Name         string
    RoarStrength int
    age          int
}

// Embedding two structs
// Animal represents information about all animals.
type animal struct {
    Name string
    Age  int
}

// Cat represents information about cats.
type Cat struct {
    animal
    MeowStrength int
}