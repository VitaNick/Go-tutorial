package main

func main() {
	// p := organization.NewPerson("Nick", "Vanden Eynde", organization.NewEuropeanUnionIdentifier("123-45-6789", "Belgium"))
	// err := p.SetTwitterHandler(organization.TwitterHandler("@ninjawulf98"))
	// fmt.Printf("%T\n", organization.TwitterHandler("test"))

	// if err != nil {
	// 	fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	// }

	// p.First = "Collin"
	// println(p.First)
	// println(p.Name.First)
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectUrl())
	// println(p.ID())
	// println(p.Country())
	// println(p.FullName())

	name1 := Name{First: "James", Last: "Wilson"}
	name2 := OtherName{First: "James", Last: "Wilson"}

	if name1 == name2 {
		println("We match")
	}
}

type Name struct {
	First string
	Last  string
}

type OtherName struct {
	First string
	Last  string
}
