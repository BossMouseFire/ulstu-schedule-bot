package repository

type StudentRepository interface {
	AddStudent(firstName, lastName string, userId int, groupName string) error
	CheckStudent(userId int) (bool, error)
	GetStudentGroup(userId int) (string, error)
	UpdateStudentGroup(userId int, newGroupName string) error
}
