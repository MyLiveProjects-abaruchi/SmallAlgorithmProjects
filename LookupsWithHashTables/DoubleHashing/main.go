package main

import (
    "fmt"
    "math/rand"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/DoubleHashing/pkg/employee"
    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/DoubleHashing/pkg/hash"
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
    }

    hash_table := hash.NewChainingHashTable(10)
    for _, storedEmployee := range employees {
        hash_table.Set(storedEmployee.Name, storedEmployee.Phone)
    }
    hash_table.Dump()

    fmt.Printf("Table contains Sally Owens: %t\n", hash_table.Contains("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.Contains("Dan Deever"))
    fmt.Println("Deleting Dan Deever")
    hash_table.Delete("Dan Deever")
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.Contains("Dan Deever"))
    fmt.Printf("Sally Owens: %s\n", hash_table.Get("Sally Owens"))
    fmt.Printf("Fred Franklin: %s\n", hash_table.Get("Fred Franklin"))
    fmt.Println("Changing Fred Franklin")
    hash_table.Set("Fred Franklin", "202-555-0100")
    fmt.Printf("Fred Franklin: %s\n", hash_table.Get("Fred Franklin"))

    //Look at clustering.
    rand.Seed(12345)
    big_capacity := 1000
    big_hash_table := hash.NewChainingHashTable(big_capacity)
    num_items := int(float32(big_capacity) * 0.9)
    for i := 0; i < num_items; i++ {
        str := fmt.Sprintf("%d-%d", i, rand.Intn(1000000))
        big_hash_table.Set(str, str)
    }
    big_hash_table.DumpConcise()
    fmt.Printf("Average probe sequence length: %f\n",
        big_hash_table.AveProbeSequenceLength())
}
