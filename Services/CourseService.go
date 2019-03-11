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

func (serv *CourseService) PassCourseData() []Models.Course {
	return serv.rep.GetCourseData()
}

func (serv *CourseService) ShowCourse() []Models.Course {
	return serv.rep.ShowCourseData()
}

func (serv *CourseService) SaveCourseData(c Models.Course)  {
	serv.rep.SaveCourse(c)
}

func (serv *CourseService) SaveData(course Models.Course) {
	serv.rep.SaveCourse(course)
}

func (serv *CourseService) GetCgpa() float64 {
	return serv.rep.CalculateCgpa()
}

//func (serv *StudentService) ShowCgpa() float64 {
//	return serv.rep.GetCgpa()
//}
