package hash

import (
    "fmt"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/DoubleHashing/pkg/employee"
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

func hash2(value string) int {
    hash := 0
    for _, ch := range value {
        hash += int(ch)
        hash += hash << 10
        hash ^= hash >> 6
    }

    // Make sure the result is non-negative.
    if hash < 0 {
        hash = -hash
    }

    // Make sure the result is not 0.
    if hash == 0 {
        hash = 1
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
            continue
        }
        if bucket.Deleted {
            fmt.Printf("%d: xxx\n", idx)
            continue
        }
        fmt.Printf("%d: %s \t %s\n", idx, bucket.Name, bucket.Phone)
    }
}

func (hashTable *LinearProbingHashTable) find(name string) (int, int) {

    hash1Val := hash(name) % 100000
    hash2Val := hash2(name) % 100000

    deletedIdx := -1
    currIdx := 0
    for i := 0; currIdx < hashTable.capacity; i++ {
        currIdx = (hash1Val + i*hash2Val) % (hashTable.capacity - 1)

        // Found a NIL position (the item isn't in hash table)
        if hashTable.data[currIdx] == nil {
            // Return the deleted item we found (if we found)
            if deletedIdx != -1 {
                return deletedIdx, -1
            }
            // Return current index to be used to store, but indicates we didn't find the item.
            return currIdx, -1
        }

        // Store the first deleted index we find.
        if hashTable.data[currIdx].Deleted && deletedIdx == -1 {
            deletedIdx = currIdx
            continue
        }

        // We found the element, return the index + delta
        if hashTable.data[currIdx].Name == name {
            return currIdx, i
        }
    }
    // This return indicates the hash is full
    return -1, -1
}

func (hashTable *LinearProbingHashTable) Set(name string, phone string) {
    keyIdx, delta := hashTable.find(name)
    newEmployee := employee.NewEmployee(name, phone)

    if keyIdx == -1 {
        return
    }

    if delta <= 0 {
        hashTable.numEntries += 1
    }

    hashTable.data[keyIdx] = newEmployee

    if hashTable.numEntries == 0 {
        employeeIdx := hash(name) % hashTable.capacity
        hashTable.data[employeeIdx] = employee.NewEmployee(name, phone)
        hashTable.numEntries += 1
    }

}

func (hashTable *LinearProbingHashTable) Get(name string) string {

    keyIdx, delta := hashTable.find(name)

    if delta >= 0 {
        return hashTable.data[keyIdx].Phone
    }
    return ""
}

func (hashTable *LinearProbingHashTable) Contains(name string) bool {
    _, delta := hashTable.find(name)

    return delta >= 0
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

func (hashTable *LinearProbingHashTable) Delete(name string) {
    keyIdx, delta := hashTable.find(name)

    if delta <= 0 {
        return
    }

    hashTable.data[keyIdx].Deleted = true
    hashTable.capacity -= 1
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
