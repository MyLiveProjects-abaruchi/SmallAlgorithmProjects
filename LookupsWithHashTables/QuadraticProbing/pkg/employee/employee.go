package employee

type Employee struct {
    Phone   string
    Name    string
    Deleted bool
}

func NewEmployee(name, phone string) *Employee {
    return &Employee{
        Name:    name,
        Phone:   phone,
        Deleted: false,
    }
}
