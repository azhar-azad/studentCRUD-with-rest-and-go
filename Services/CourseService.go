package Services

import (
	"studentCgpa/Models"
	"studentCgpa/Repos"
)

type CourseService struct {
	rep *Repos.CourseRepo
}

func NewCourseService(r *Repos.CourseRepo) *CourseService {
	return &CourseService{rep: r}
}

func (serv *CourseService) ShowCourse() []Models.Course {
	return serv.rep.ShowCourseData()
}

func (serv *CourseService) SaveCourseData(c Models.Course) {
	serv.rep.SaveCourse(c)
}

func (serv *CourseService) SaveData(course Models.Course) {
	serv.rep.SaveCourse(course)
}

func (serv *CourseService) GetCgpa(stdId string) float64 {
	return serv.rep.CalculateCgpa(stdId)
}

//func (serv *StudentService) ShowCgpa() float64 {
//	return serv.rep.GetCgpa()
//}
