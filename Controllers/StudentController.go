package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"studentCgpa/Models"
	"studentCgpa/Services"
)

type StudentController struct {
	serv *Services.StudentService
	rendr *render.Render
}

func NewStudentController(s *Services.StudentService, r *render.Render) *StudentController {
	return &StudentController{serv: s, rendr: r}
}

func (ctrl *StudentController) ShowHomePage(w http.ResponseWriter, req *http.Request) {
	data := struct {
		Student map[string]Models.Student
	}{
		Student: ctrl.serv.ShowData(),
	}
	ctrl.rendr.HTML(w, http.StatusOK, "layout", data)
}

func (ctrl *StudentController) ShowFormPage(w http.ResponseWriter, req *http.Request) {

	ctrl.rendr.HTML(w, http.StatusOK, "home", nil)
}

type Student struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
}

func (ctrl *StudentController) Insert(w http.ResponseWriter, req *http.Request) {

	byteCode, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var student Models.Student

	err = json.Unmarshal(byteCode, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("student: ", student, "\n")

	ctrl.serv.SaveData(student)

	//http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func (ctrl *StudentController) Edit(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Edit starts...")
	params := mux.Vars(req)
	receivedId := params["id"]

	data := struct {
		Student Models.Student
	}{
		Student: ctrl.serv.ReturnOneEntry(receivedId),
	}

	ctrl.rendr.HTML(w, http.StatusOK, "edit", data)
	fmt.Println("Edit ends...\n")
}

func (ctrl *StudentController) Update(w http.ResponseWriter, req *http.Request) {

	byteCode, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var student Models.Student

	err = json.Unmarshal(byteCode, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.serv.SaveData(student)

	//json.NewEncoder(w).Encode(&student)
}

func (ctrl *StudentController) Delete(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	receivedId := params["id"]
	ctrl.serv.DeleteOneEntry(receivedId)
}
//
//func (ctrl *StudentController) CalculateCgpa(w http.ResponseWriter, req *http.Request) {
//
//	ctrl.rendr.HTML(w, http.StatusOK, "layout", ctrl.serv.ShowCgpa())
//}



