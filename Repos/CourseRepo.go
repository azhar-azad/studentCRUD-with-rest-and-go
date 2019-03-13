package Repos

import (
	"fmt"
	"math"
	"strconv"
	"studentCgpa/Models"
)

type CourseRepo struct {
	CourseData []Models.Course
}

func NewCourseReop() *CourseRepo {
	return &CourseRepo{CourseData: make([]Models.Course, 0, 0)}
}

func (rep *CourseRepo) ShowCourseData() []Models.Course {
	return rep.CourseData
}

func (rep *CourseRepo) SaveCourse(course Models.Course) {
	rep.CourseData = append(rep.CourseData, course)
}

func (rep *CourseRepo) CalculateCgpa(stdId string) float64 {

	productValue := 0.0
	totalCredit := 0.0

	studentId, err := strconv.Atoi(stdId)
	if err != nil {
		fmt.Sprintf("Can't convert (string)StudentId to (int)StudentId")
		return 0.0
	}

	var tempData []Models.Course
	for _, val := range rep.CourseData {
		if val.StudentId == studentId {
			tempData = append(tempData, val)
		}
	}

	for _, val := range tempData {

		intCredit, err := strconv.ParseFloat(val.Credit, 64)
		if err != nil {
			fmt.Sprintf("Can't convert (string)Credit to (float64)Credit")
			return 0.0
		}

		intMark, err := strconv.ParseFloat(val.Mark, 64)
		if err != nil {
			fmt.Sprintf("Can't convert (string)Mark to (float64)Mark")
			return 0.0
		}

		productValue += intCredit * intMark
		totalCredit += intCredit
	}

	cg := productValue / totalCredit

	return math.Round(cg*100) / 100
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
