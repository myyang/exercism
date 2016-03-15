// Package school provides school
package school

import (
	"log"
	"sort"
)

// Grade contains grade indication and students' name
type Grade struct {
	grade int
	names []string
}

// School contains several grades
type School map[int]Grade

// New a school
func New() *School {
	return &School{}
}

// Enrollment show school status
func (sch *School) Enrollment() []Grade {
	var grades []Grade
	keys := []int{}
	for k, _ := range *sch {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		grade, ok := (*sch)[v]
		if !ok {
			log.Fatal("Error grade: ", v, ok, grade)
			return nil
		}
		sort.Strings(grade.names)
		grades = append(grades, grade)
	}
	return grades
}

// Add new student
func (sch *School) Add(name string, grade int) {
	_, ok := (*sch)[grade]
	if !ok {
		(*sch)[grade] = Grade{grade, []string{name}}
	} else {
		(*sch)[grade] = Grade{grade, append((*sch)[grade].names, name)}
	}
}

// Grade return students in that grade
func (sch *School) Grade(grade int) []string {
	return (*sch)[grade].names
}
