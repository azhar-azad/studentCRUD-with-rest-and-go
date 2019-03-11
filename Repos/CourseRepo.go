package Repos

import (
	"fmt"
	"strconv"
	"studentCgpa/Models"
)

type CourseRepo struct {
	CourseData []Models.Course
}



func NewCourseReop() *CourseRepo {
	return &CourseRepo{CourseData: make([]Models.Course, 0, 0)}
}

func (rep *CourseRepo)GetCourseData() []Models.Course {
	return rep.CourseData
}

func (rep *CourseRepo) ShowCourseData() []Models.Course {
	return rep.CourseData
}

func (rep *CourseRepo) SaveCourse(course Models.Course) {
	rep.CourseData = append(rep.CourseData, course)
}

func (rep *CourseRepo) CalculateCgpa() float64 {
	productValue := 0.0
	totalCredit := 0.0

	for _, val := range rep.CourseData {
		intCredit, err := strconv.ParseFloat(val.Credit, 64)
		if err != nil {
			fmt.Println("Can't convert (string)Credit to (float64)Credit")
			return 0.0
		}

		intMark, err := strconv.ParseFloat(val.Mark, 64)
		if err != nil {
			fmt.Println("Can't convert (string)Mark to (float64)Mark")
			return 0.0
		}

		productValue += intCredit * intMark
		totalCredit += intCredit
	}

	return productValue / totalCredit
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
