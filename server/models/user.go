package models

type User struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Username string         `json:"username" gorm:"unique"`
	Password string         `json:"password"`
	Courses  string         `json:"courses"`
	Hubs     map[string]int `json:"hubs" gorm:"-"`
}

type CourseRequest struct {
	UserID      uint   `json:"userID"`
	CourseCodes string `json:"courseCodes"`
}

func GetHubs() map[string]int {
	hubs := map[string]int{
		"Philosophical Inquiry and Life's Meanings":     1,
		"Aesthetic Exploration":                         1,
		"Historical Consciousness":                      1,
		"Scientific Inquiry I":                          1,
		"Social Inquiry I":                              1,
		"Scientific Inquiry II":                         1,
		"Social Inquiry II":                             1,
		"Quantitative Reasoning I":                      1,
		"Quantitative Reasoning II":                     1,
		"The Individual in Community":                   1,
		"Global Citizenship and Intercultural Literacy": 2,
		"Ethical Reasoning":                             1,
		"First-Year Writing Seminar":                    1,
		"Writing, Research, and Inquiry":                1,
		"Writing-Intensive Course":                      2,
		"Oral and/or Signed Communication":              1,
		"Digital/Multimedia Expression":                 1,
		"Critical Thinking":                             2,
		"Research and Information Literacy":             2,
		"Teamwork/Collaboration":                        2,
		"Creativity/Innovation":                         2,
		"Life Skills":                                   0,
	}
	return hubs
}

/*
Hub names and required credits:
Philosophical Inquiry and Life’s Meanings	1
Aesthetic Exploration	1
Historical Consciousness	1
Scientific Inquiry I	1
Social Inquiry I	1
Scientific Inquiry II or Social Inquiry II	1
Quantitative Reasoning I	1
Quantitative Reasoning II	1
The Individual in Community	1
Global Citizenship and Intercultural Literacy	2
Ethical Reasoning	1
First-Year Writing Seminar	1
Writing, Research, and Inquiry	1
Writing-Intensive Course	2
Oral and/or Signed Communication	1
Digital/Multimedia Expression	1
Critical Thinking	2
Research and Information Literacy	2
Teamwork/Collaboration	2
Creativity/Innovation	2
Life Skills	0
*/
