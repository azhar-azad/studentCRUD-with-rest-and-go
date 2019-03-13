package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"studentCgpa/Controllers"
	"studentCgpa/Repos"
	"studentCgpa/Services"
)

func main() {

	fmt.Println("Program starts...")

	rend := render.New(render.Options{
		Directory:  "Templates",
		Extensions: []string{".tmpl", ".html"},
	})

	studentRepo := Repos.NewStudentRepo()
	studentService := Services.NewStudentService(studentRepo)
	studentController := Controllers.NewStudentController(studentService, rend)

	courseRepo := Repos.NewCourseReop()
	courseService := Services.NewCourseService(courseRepo)
	courseController := Controllers.NewCourseController(courseService, rend)

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", studentController.ShowHomePage).Methods(http.MethodGet)
	myRouter.HandleFunc("/home", studentController.ShowFormPage).Methods(http.MethodGet)
	myRouter.HandleFunc("/insert", studentController.Insert).Methods(http.MethodPost)
	myRouter.HandleFunc("/edit/{id}", studentController.Edit).Methods(http.MethodGet)
	myRouter.HandleFunc("/update", studentController.Update).Methods(http.MethodPut)
	myRouter.HandleFunc("/delete/{id}", studentController.Delete).Methods(http.MethodDelete)

	myRouter.HandleFunc("/addcourse/{id}", courseController.AddCourse).Methods(http.MethodGet)
	myRouter.HandleFunc("/savecourse", courseController.SaveCourse).Methods(http.MethodPost)
	myRouter.HandleFunc("/courselayout/{id}", courseController.ShowCoursePage).Methods(http.MethodGet)
	myRouter.HandleFunc("/courseform", courseController.ShowCourseFormPage).Methods(http.MethodGet)
	myRouter.HandleFunc("/calculatecgpa/{id}", courseController.ShowCgpa).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
