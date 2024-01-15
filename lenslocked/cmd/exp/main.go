package main

import (
	"html/template"
	"os"
)

type User struct {
	Name       string
	Bio        string
	Age        int
	VisitDates []int
	Meta       UserMeta
}
type UserMeta struct {
	Data map[int]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:       "Jeppe Korbo",
		Bio:        `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Age:        123,
		VisitDates: []int{1, 2, 3},
		Meta: UserMeta{
			Data: map[int]string{1: "hello", 2:"there"},
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
