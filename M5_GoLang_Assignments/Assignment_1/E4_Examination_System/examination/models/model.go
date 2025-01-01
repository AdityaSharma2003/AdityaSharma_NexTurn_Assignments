package models

type Question struct {
	ID        int
	Question  string
	Options   map[int]string
	Answer    int
	Weightage int
}
