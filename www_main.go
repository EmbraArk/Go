package main

import (
	"fmt"
	"net/http"
)

//	http://localhost:8080/contacts
//
// 3 урок
type User struct {
	name                 string
	year                 uint16
	money                int16   // балланс
	avr_grades, hapinnes float64 // happinness: [0-1]
}

// start: метод структуры
func (user User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s; "+
		"\nHe is: %d. \nAnd he has money: %d ", user.name, user.year, user.money)
}

func (user *User) setNewName(newName string) {
	user.name = newName

}

// end: метод структуры

func home_page(w http.ResponseWriter, r *http.Request) {
	// w: параметр - адрес запрашиваемой страницы
	// r: параметр -

	bob := User{"Bob", 25, 0, 4.4, 0.7}
	bob.setNewName("Kali")
	fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, "Go")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page) // отслеживание главной страницы
	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":8080", nil) // порт, настройки
}

// var Bob User // создание объекта на основе структуры
func main() {
	handleRequest()

}
