package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	man1 := Man{Name: "kyle", LastName: "Busch", Age: 36, Gender: "male", Crimes: 2}
	man2 := Man{Name: "Anna", LastName: "Clark", Age: 25, Gender: "female", Crimes: 0}
	man3 := Man{Name: "kurt", LastName: "Busch", Age: 40, Gender: "male", Crimes: 1}
	man4 := Man{Name: "Kevin", LastName: "Harvick", Age: 47, Gender: "male", Crimes: 3}
	man5 := Man{Name: "Maria", LastName: "Smit", Age: 36, Gender: "female", Crimes: 5}

	people := map[string]Man{
		man1.Name: man1,
		man2.Name: man2,
		man3.Name: man3,
		man4.Name: man4,
		man5.Name: man5,
	}

	suspects := []string{"kyle", "Anna", "Kevin", "kurt", "Bart"}

	var mostCriminal Man
	var mostCriminalFound bool
	for i, name := range suspects {
		person, ok := people[name]
		if !ok {
			fmt.Printf("Подозреваемого %s нет в базе данных \n", suspects[i])
			continue
		}

		if person.Crimes > mostCriminal.Crimes {
			mostCriminal = person
			mostCriminalFound = true
		}
	}

	if mostCriminalFound {
		fmt.Printf("Из списка подозреваемых больше преступлений у: %s %s : %d\n", mostCriminal.Name, mostCriminal.LastName, mostCriminal.Crimes)
	} else {
		fmt.Println("В списке подозреваемых нет никого из базы данных")
	}

	fmt.Println("Всего в базе данных:")

	for name, man := range people {
		fmt.Printf("Имя: %s, Фамилия: %s, Возраст: %d, Пол: %s, Преступлений: %d\n", name, man.LastName, man.Age, man.Gender, man.Crimes)
	}
}
