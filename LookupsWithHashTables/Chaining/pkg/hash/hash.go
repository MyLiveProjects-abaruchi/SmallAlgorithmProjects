package hash

import (
    "fmt"

    "github.com/MyLiveProjects-abaruchi/SmallAlgorithmProjects/LookupsWithHashTables/Chaining/pkg/employee"
)

type ChainingHashTable struct {
    numBuckets int
    numEntries int
    buckets    [][]*employee.Employee
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

func NewChainingHashTable(numBuckets int) *ChainingHashTable {
    newBuckets := make([][]*employee.Employee, numBuckets)

    return &ChainingHashTable{
        numBuckets: numBuckets,
        numEntries: 0,
        buckets:    newBuckets,
    }

}

func (hashTable *ChainingHashTable) Dump() {
    if hashTable.numBuckets == 0 {
        return
    }
    for idx, bucket := range hashTable.buckets {
        fmt.Printf("Bucket %d:\n", idx)

        for _, contact := range bucket {
            fmt.Printf("   %s: %s\n", contact.Name, contact.Phone)
        }
    }
}

func (hashTable *ChainingHashTable) find(name string) (int, int) {
    hashBucketIDX := hash(name) % hashTable.numBuckets

    if hashTable.numEntries == 0 {
        return hashBucketIDX, -1
    }

    for idx, contact := range hashTable.buckets[hashBucketIDX] {
        if contact.Name == name {
            return hashBucketIDX, idx
        }
    }

    return hashBucketIDX, -1
}

func (hashTable *ChainingHashTable) Set(name string, phone string) {

    employeeBucket, employeeIdx := hashTable.find(name)

    if employeeIdx >= 0 {
        hashTable.buckets[employeeBucket][employeeIdx].Name = name
        hashTable.buckets[employeeBucket][employeeIdx].Phone = phone
        return
    }

    employeeToAdd := employee.NewEmployee(name, phone)
    hashTable.buckets[employeeBucket] = append(hashTable.buckets[employeeBucket], employeeToAdd)
    hashTable.numEntries++
}

func (hashTable *ChainingHashTable) Get(name string) string {
    employeeBucket, employeeIdx := hashTable.find(name)

    if employeeIdx < 0 {
        return ""
    }

    return hashTable.buckets[employeeBucket][employeeIdx].Phone
}

func (hashTable *ChainingHashTable) Contains(name string) bool {
    _, employeeIdx := hashTable.find(name)

    return employeeIdx >= 0
}

func (hashTable *ChainingHashTable) Delete(name string) {
    employeeBucket, employeeIdx := hashTable.find(name)

    if employeeIdx < 0 {
        return
    }

    hashTable.buckets[employeeBucket] = append(hashTable.buckets[employeeBucket][:employeeIdx],
        hashTable.buckets[employeeBucket][employeeIdx+1:]...)
    hashTable.numEntries--
}
