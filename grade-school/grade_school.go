package school

import (
	"sort"
)

type Grade struct {
	ID       int
	Students []string
}

type School struct {
	Grades map[int]*Grade
}

type ByID []Grade

func (a ByID) Len() int {
	return len(a)
}

func (a ByID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func New() *School {
	return &School{make(map[int]*Grade)}
}

func (s *School) Enrollment() []Grade {
	res := make([]Grade, 0, len(s.Grades))
	for _, grade := range s.Grades {
		res = append(res, *grade)
	}

	sort.Sort(ByID(res))
	return res
}

func (s *School) Add(student string, grade int) {
	gradePointer, ok := s.Grades[grade]

	if !ok {
		s.Grades[grade] = &Grade{ID: grade}
	}

	gradePointer = s.Grades[grade]
	gradePointer.Students = append(gradePointer.Students, student)
	sort.Strings(gradePointer.Students)
}

func (s *School) Grade(grade int) []string {
	gradePointer, ok := s.Grades[grade]
	if !ok {
		return []string{}
	}

	return gradePointer.Students
}
