package hash

import (
    "fmt"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/QuadraticProbing/pkg/employee"
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
    // The hash table is empty
    if hashTable.numEntries == 0 {
        return -1, -1
    }

    keyHashIdx := hash(name) % (hashTable.capacity - 1)

    deletedIdx := -1
    currIdx := 0
    for i := 1; currIdx < hashTable.capacity; i++ {
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
            return keyHashIdx, i
        }
        currIdx = keyHashIdx + (i * i)
    }
    // This return indicates the hash is full
    return -1, -1
}

func (hashTable *LinearProbingHashTable) Set(name string, phone string) {

    if hashTable.numEntries == 0 {
        employeeIdx := hash(name) % hashTable.capacity
        hashTable.data[employeeIdx] = employee.NewEmployee(name, phone)
        hashTable.numEntries += 1
    }

    keyIdx, deltaIdx := hashTable.find(name)

    if hashTableIsFull(keyIdx, deltaIdx) {
        return
    }

    newEmployee := employee.NewEmployee(name, phone)
    hashTable.data[keyIdx] = newEmployee

    if deltaIdx == -1 {
        hashTable.numEntries += 1
    }
    return
}

func (hashTable *LinearProbingHashTable) Get(name string) string {
    var itemIdx int
    keyIdx, deltaIdx := hashTable.find(name)

    if hashTableIsFull(keyIdx, deltaIdx) {
        return ""
    }

    if keyIdx != -1 && deltaIdx != -1 {
        return hashTable.data[itemIdx].Phone
    }

    return ""
}

func (hashTable *LinearProbingHashTable) Contains(name string) bool {
    keyIdx, delta := hashTable.find(name)

    return keyIdx != -1 && delta != -1
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
    if keyIdx != -1 && delta != -1 {
        hashTable.data[keyIdx+delta].Deleted = true
    }
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

func hashTableIsFull(keyIdx, delta int) bool {
    return keyIdx == -1 && delta == -1
}
