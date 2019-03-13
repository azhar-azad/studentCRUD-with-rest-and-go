package Repos

import "studentCgpa/Models"

type StudentRepo struct {
	StudentData map[string]Models.Student
}

func NewStudentRepo() *StudentRepo {
	return &StudentRepo{StudentData: make(map[string]Models.Student)}
}

func (rep *StudentRepo) ShowStudentData() map[string]Models.Student {
	return rep.StudentData
}

func (rep *StudentRepo) Save(s Models.Student) {
	rep.StudentData[s.Id] = s
}

func (rep *StudentRepo) FindOne(id string) Models.Student {
	return rep.StudentData[id]
}

func (rep *StudentRepo) Delete(id string) {
	delete(rep.StudentData, id)
}

func (rep *StudentRepo) GetNameById(id string) string {
	return rep.StudentData[id].Name
}

//func (rep *StudentRepo) GetCgpa() float64 {
//
//	productValue := 0.0
//	totalCredit := 0.0
//	for _, val := range rep.CourseData {
//		productValue += val.Credit * val.CreditHour
//		totalCredit += val.Credit
//	}
//
//	return productValue / totalCredit
//}
//
//func (rep *StudentRepo) GetCgpa2(selectedStudents []string) float64 {
//
//	productValue := 0.0
//	totalCredit := 0.0
//	for _, courseData := range rep.CourseData {
//		for _, studentId := range selectedStudents {
//			if courseData.StudentId == studentId {
//				productValue += courseData.Credit * courseData.CreditHour
//				totalCredit += courseData.Credit
//			}
//		}
//	}
//
//	return productValue / totalCredit
//}
