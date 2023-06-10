package main

import (
    "fmt"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/Chaining/pkg/employee"
    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/Chaining/pkg/hash"
)

func main() {
    // Make some names.
    employees := []employee.Employee{
        {Name: "Ann Archer", Phone: "202-555-0101"},
        {Name: "Bob Baker", Phone: "202-555-0102"},
        {Name: "Cindy Cant", Phone: "202-555-0103"},
        {Name: "Dan Deever", Phone: "202-555-0104"},
        {Name: "Edwina Eager", Phone: "202-555-0105"},
        {Name: "Fred Franklin", Phone: "202-555-0106"},
        {Name: "Gina Gable", Phone: "202-555-0107"},
        {Name: "Herb Henshaw", Phone: "202-555-0108"},
        {Name: "Ida Iverson", Phone: "202-555-0109"},
        {Name: "Jeb Jacobs", Phone: "202-555-0110"},
    }

    hash_table := hash.NewChainingHashTable(10)
    for _, emply := range employees {
        hash_table.Set(emply.Name, emply.Phone)
    }
    hash_table.Dump()

    fmt.Printf("Table contains Sally Owens: %t\n", hash_table.Contains("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.Contains("Dan Deever"))
    fmt.Println("Deleting Dan Deever")
    hash_table.Delete("Dan Deever")
    fmt.Printf("Sally Owens: %s\n", hash_table.Get("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.Contains("Dan Deever"))
    fmt.Printf("Fred Franklin: %s\n", hash_table.Get("Fred Franklin"))
    fmt.Println("Changing Fred Franklin")
    hash_table.Set("Fred Franklin", "202-555-0100")
    fmt.Printf("Fred Franklin: %s\n", hash_table.Get("Fred Franklin"))
}
