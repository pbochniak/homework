package slices

import "sort"

func SortStringsInPlace(strings []string) {
	sort.Strings(strings)
}

func SortStringsPure(strings []string) []string {
	temp_strings := make([]string, len(strings))
	copy(temp_strings, strings)
	SortStringsInPlace(temp_strings)
	return temp_strings
}

type User struct {
	FirstName string
	LastName  string
}

func compareUsers(u1, u2 User) bool {
	if u1.FirstName > u2.FirstName {
		return false
	}
	if u1.FirstName < u2.FirstName {
		return true
	}
	if u1.LastName > u2.LastName {
		return false
	}
	if u1.LastName < u2.LastName {
		return true
	}
	return false
}

func SortUsersPure(users []User) []User {
	tempUsers := make([]User, len(users))
	copy(tempUsers, users)
	sort.Slice(tempUsers, func(i, j int) bool { return compareUsers(tempUsers[i], tempUsers[j]) })
	return tempUsers
}

func SortUsersPureDesc(users []User) []User {
	tempUsers := make([]User, len(users))
	copy(tempUsers, users)
	sort.Slice(tempUsers, func(i, j int) bool { return !compareUsers(tempUsers[i], tempUsers[j]) })
	return tempUsers
}
