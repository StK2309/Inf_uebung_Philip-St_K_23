package pointer

// Aufgabe 4: new / Heap-Allocation
// Erstelle und returne eine *Person mit den übergebenen Werten.
// (Nutze new(Person) oder &Person{...})
func NewPerson(name string, alter int) *Person {
	// TODO: implement
	newPerson := new(Person)
	newPerson.Name = name
	newPerson.Alter = alter
	return newPerson

	//p := &Person{
	//	Name:  name,
	//	Alter: alter,
	//}
	//return p
}
