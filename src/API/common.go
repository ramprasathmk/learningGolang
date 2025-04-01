package API

// declare response status and employee errors
const (
	// status error
	lStatusE = "E"

	// status success
	lStatusS = "S"

	lDefaultErr = "something went wrong"

	lEmpNotFoundErr = "employee not found"

	lEmpExistErr = "employee already exist"
)

var empArr = EmployeesList{
	// employee 1
	Employee{
		Id:      101,
		Name:    "Rohan",
		Mobile:  "9988776601",
		Email:   "rohan@mycompany.com",
		Address: "Paris",
		Dob:     "20/11/2001",
	},

	// employee 2
	Employee{
		Id:      102,
		Name:    "Jane",
		Mobile:  "9988776602",
		Email:   "jane@mycompany.com",
		Address: "New York",
		Dob:     "02/01/2003",
	},

	// employee 3
	Employee{
		Id:      103,
		Name:    "Dave",
		Mobile:  "9988776603",
		Email:   "dave@mycompany.com",
		Address: "Bangkok",
		Dob:     "17/06/2000",
	},

	// employee 4
	Employee{
		Id:      104,
		Name:    "Ming",
		Mobile:  "9988776604",
		Email:   "ming@mycompany.com",
		Address: "Tokyo",
		Dob:     "04/11/2002",
	},
}
