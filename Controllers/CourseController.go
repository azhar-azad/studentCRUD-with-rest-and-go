package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"strconv"
	"studentCgpa/Models"
	"studentCgpa/Services"
)

type CourseController struct {
	serv  *Services.CourseService
	rendr *render.Render
}

func NewCourseController(s *Services.CourseService, r *render.Render) *CourseController {
	return &CourseController{serv: s, rendr: r}
}

func (ctrl *CourseController) ShowCoursePage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	c := params["id"]

	intc, _ := strconv.Atoi(c)

	coursedata := ctrl.serv.ShowCourse()

	var courseList []Models.Course

	for _, item := range coursedata {
		if item.StudentId == intc {
			courseList = append(courseList, item)
		}
	}

	data := struct {
		CourseList []Models.Course
	}{
		CourseList: courseList,
	}

	ctrl.rendr.HTML(w, http.StatusOK, "courselayout", data)
}

func (ctrl *CourseController) ShowCourseFormPage(w http.ResponseWriter, req *http.Request) {

	ctrl.rendr.HTML(w, http.StatusOK, "courseform", nil)
}

type Course struct {
	CourseId   string `json:"courseid"`
	CourseName string `json:"coursename"`
	Credit     string `json:"credit"`
	Mark       string `json:"mark"`
	StudentId  int    `json:"studentid"`
}

func (ctrl *CourseController) AddCourse(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	c := params["id"]

	data := struct {
		StudentId string
	}{
		StudentId: c,
	}

	ctrl.rendr.HTML(w, http.StatusOK, "courseform", data)
	//http.Redirect(w, req, "/courselayout", http.StatusMovedPermanently)
}

func (ctrl *CourseController) SaveCourse(w http.ResponseWriter, req *http.Request) {

	byteCode, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var c Course

	err = json.Unmarshal(byteCode, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	course := Models.Course{
		Id:        c.CourseId,
		Name:      c.CourseName,
		Credit:    c.Credit,
		Mark:      c.Mark,
		StudentId: c.StudentId,
	}

	ctrl.serv.SaveData(course)
}

type Data struct {
	//Name string
	CGPA float64
}

func (ctrl *CourseController) ShowCgpa(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	c := params["id"]

	cg := ctrl.serv.GetCgpa(c)
	fmt.Println("Cgpa: ", cg)

	data := Data{
		//Name: name,
		CGPA: cg,
	}

	ctrl.rendr.HTML(w, http.StatusOK, "gpa", data)
}
