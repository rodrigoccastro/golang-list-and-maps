package main

import (
	"fmt"
	"sort"
)

type Client struct {
	Id   int
	Name string
}

type ByName []Client

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByNameDesc []Client

func (a ByNameDesc) Len() int           { return len(a) }
func (a ByNameDesc) Less(i, j int) bool { return a[i].Name > a[j].Name }
func (a ByNameDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	list();
	maps();
}

func list() {
	list := make([]Client, 0)
	list = append(list, Client{Id: 1, Name: "teste 5"})
	list = append(list, Client{Id: 2, Name: "teste 4"})
	list = append(list, Client{Id: 3, Name: "teste 3"})
	list = append(list, Client{Id: 4, Name: "teste 2"})
	list = append(list, Client{Id: 5, Name: "teste 1"})
	fmt.Println("Original List:", list)

	// Sort the list by Name
	sort.Sort(ByName(list))
	fmt.Println("Sorted by Name:", list)

	// Sort the list by Name in descending order
	sort.Sort(ByNameDesc(list))
	fmt.Println("Sorted by Name desc:", list)

	// Get an item by ID
	obj := getItemById(list, 1)
	fmt.Println("Item with ID 1:", obj)

	// Remove an item by ID
	list = removeItemById(list, 1)
	fmt.Println("List after removing item with ID 1:", list)

	// Remove an item by ID
	list = removeItemById(list, 5)
	fmt.Println("List after removing item with ID 5:", list)

	// Remove an item by ID
	list = removeItemById(list, 3)
	fmt.Println("List after removing item with ID 3:", list)
}

func getItemById(list []Client, id int) *Client {
	for i := range list {
		if list[i].Id == id {
			return &list[i]
		}
	}
	return nil
}

func removeItemById(list []Client, id int) []Client {
	for i, item := range list {
		if item.Id == id {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func maps() {
    // Create an empty map with string keys and string values
    myMap := make(map[string]string)

    // Add key-value pairs to the map
    myMap["abc"] = "city 90"
    myMap["def"] = "city 80"
    myMap["ghi"] = "city 70"
    fmt.Println("all map", myMap)

	sorted := getListKeysSorted(myMap,false)
    fmt.Println("all keys sorted asc", sorted)

	sorted = getListKeysSorted(myMap,true)
    fmt.Println("all keys sorted desc", sorted)

	sortedMapAsc := sortMapByValue(myMap);
	fmt.Println("Sorted map asc:", sortedMapAsc)

	sortedMapDesc := sortMapByValueReverse(myMap);
	fmt.Println("Sorted map desc:", sortedMapDesc)


	// Access the value associated with the "abc" key
    fmt.Println("Value of abc:", myMap["abc"])

    // Delete a key-value pair from the map
    delete(myMap, "abc")
    fmt.Println("after delete abc", myMap)

    // Loop through the map using a for loop
    fmt.Println("item by item")
    for key, value := range myMap {
        fmt.Println("Key:", key, "Value:", value)
    }

}

func getListKeysSorted(myMap map[string]string, reverse bool) []string {
	// Convert the map keys to a slice
	keys := make([]string, 0, len(myMap))
	for key := range myMap {
		keys = append(keys, key)
	}

	// Sort the slice of keys
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}

	return keys
}

type KeyValuePair struct {
	Key   string
	Value string
}

func sortMapByValueReverse(myMap map[string]string) map[string]string {
	// Convert the map into a slice of key-value pairs
	var keyValuePairs []KeyValuePair
	for key, value := range myMap {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// Define a custom sorting function based on values in reverse order
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// Create a new sorted map from the sorted slice
	sortedMap := make(map[string]string)
	for _, kvp := range keyValuePairs {
		sortedMap[kvp.Key] = kvp.Value
	}

	return sortedMap
}

func sortMapByValue(myMap map[string]string) []KeyValuePair {
	var keyValuePairs []KeyValuePair
	for key, value := range myMap {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// Define a custom sorting function based on values
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value < keyValuePairs[j].Value
	})

	// Print the sorted slice
	return keyValuePairs;
}



