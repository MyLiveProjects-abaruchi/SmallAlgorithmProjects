package employee

type Employee struct {
    Phone string
    Name  string
}

func NewEmployee(name, phone string) *Employee {
    return &Employee{
        Name:  name,
        Phone: phone,
    }
}
