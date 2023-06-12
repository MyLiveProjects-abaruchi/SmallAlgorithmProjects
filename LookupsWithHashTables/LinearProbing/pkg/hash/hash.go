package hash

import (
    "fmt"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/LinearProbing/pkg/employee"
)

type LinearProbingHashTable struct {
    capacity   int
    numEntries int
    data       []*employee.Employee
}

func hash(value string) int {
    hash := 5381
    for _, ch := range value {
        hash = ((hash << 5) + hash) + int(ch)
    }

    if hash < 0 {
        hash = -hash
    }
    return hash
}

func NewChainingHashTable(capacity int) *LinearProbingHashTable {
    dataArray := make([]*employee.Employee, capacity)

    return &LinearProbingHashTable{
        capacity:   capacity,
        numEntries: 0,
        data:       dataArray,
    }
}

func (hashTable *LinearProbingHashTable) Dump() {
    if hashTable.capacity == 0 {
        return
    }
    for idx, bucket := range hashTable.data {
        if bucket == nil {
            fmt.Printf("%d: ---\n", idx)
        } else {
            fmt.Printf("%d: %s \t %s\n", idx, bucket.Name, bucket.Phone)
        }
    }
}

func (hashTable *LinearProbingHashTable) find(name string) (int, int) {
    if hashTable.numEntries == 0 {
        return -1, -1
    }

    keyHashIdx := hash(name) % (hashTable.capacity - 1)

    currIdx := keyHashIdx
    for i := 0; currIdx < hashTable.capacity; i++ {
        if hashTable.data[currIdx] == nil {
            return -1, 1
        }
        if hashTable.data[currIdx].Name == name {
            return keyHashIdx, i
        }
        currIdx += i
    }

    return -1, -1
}

func (hashTable *LinearProbingHashTable) Set(name string, phone string) {

    keyIdx, deltaIdx := hashTable.find(name)
    employeeIdx := keyIdx + deltaIdx

    if keyIdx >= 0 {
        hashTable.data[employeeIdx].Name = name
        hashTable.data[employeeIdx].Phone = phone
        return
    }

    employeeToAdd := employee.NewEmployee(name, phone)

    currIdx := hash(name) % (hashTable.capacity - 1)
    for i := 0; i < hashTable.capacity; i++ {
        if hashTable.data[currIdx] == nil {
            hashTable.data[currIdx] = employeeToAdd
            hashTable.numEntries += 1
            return
        }
        currIdx += i
    }
}

func (hashTable *LinearProbingHashTable) Get(name string) string {
    keyIdx, deltaIdx := hashTable.find(name)
    employeeIdx := keyIdx + deltaIdx

    if employeeIdx < 0 {
        return ""
    }

    currIdx := employeeIdx
    for i := 0; currIdx < hashTable.capacity; i++ {
        if hashTable.data[currIdx] == nil {
            return ""
        }
        if hashTable.data[currIdx].Name == name {
            return hashTable.data[currIdx].Phone
        }
        currIdx += i
    }

    return ""
}

func (hashTable *LinearProbingHashTable) Contains(name string) bool {
    keyIdx, _ := hashTable.find(name)
    return keyIdx >= 0
}

func (hashTable *LinearProbingHashTable) Delete(name string) {
    return
    //employeeBucket, employeeIdx := hashTable.find(name)
    //
    //if employeeIdx < 0 {
    //    return
    //}
    //
    //hashTable.buckets[employeeBucket] = append(hashTable.buckets[employeeBucket][:employeeIdx],
    //    hashTable.buckets[employeeBucket][employeeIdx+1:]...)
    //hashTable.numEntries--
}

func (hashTable *LinearProbingHashTable) DumpConcise() {
    // Loop through the array.
    for i, employeeStored := range hashTable.data {
        if employeeStored == nil {
            // This spot is empty.
            fmt.Printf(".")
        } else {
            // Display this entry.
            fmt.Printf("O")
        }
        if i%50 == 49 {
            fmt.Println()
        }
    }
    fmt.Println()
}

func (hashTable *LinearProbingHashTable) AveProbeSequenceLength() float32 {
    totalLength := 0
    numValues := 0
    for _, storedEmployee := range hashTable.data {
        if storedEmployee != nil {
            _, probeLength := hashTable.find(storedEmployee.Name)
            totalLength += probeLength
            numValues++
        }
    }
    return float32(totalLength) / float32(numValues)
}
