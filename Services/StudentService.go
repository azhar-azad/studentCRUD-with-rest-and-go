package Services

import (
	"studentCgpa/Models"
	"studentCgpa/Repos"
)

type StudentService struct {
	rep *Repos.StudentRepo
}

func NewStudentService(r *Repos.StudentRepo) *StudentService {
	return &StudentService{rep: r}
}

func (serv *StudentService) ShowData() map[string]Models.Student {
	return serv.rep.ShowStudentData()
}

func (serv *StudentService) SaveData(s Models.Student) {
	serv.rep.Save(s)
}

func (serv *StudentService) ReturnOneEntry(id string) Models.Student {
	return serv.rep.FindOne(id)
}

func (serv *StudentService) DeleteOneEntry(id string) {
	serv.rep.Delete(id)
}

//func (serv *StudentService) ShowCgpa() float64 {
//	return serv.rep.GetCgpa()
//}
