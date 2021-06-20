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

// to implement sort.Interface
type UserSlice []User

func (users UserSlice) Len() int {
	return len(users)
}

func (users UserSlice) Less(i, j int) bool {
	if users[i].FirstName > users[j].FirstName {
		return false
	}
	if users[i].FirstName < users[j].FirstName {
		return true
	}
	if users[i].LastName > users[j].LastName {
		return false
	}
	if users[i].LastName < users[j].LastName {
		return true
	}
	return false
}

func (users UserSlice) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func SortUsersPure(users []User) []User {
	temp_users := make([]User, len(users))
	copy(temp_users, users)
	sort.Sort(UserSlice(temp_users))
	return temp_users
}

func SortUsersPureDesc(users []User) []User {
	temp_users := make([]User, len(users))
	copy(temp_users, users)
	sort.Sort(sort.Reverse(UserSlice(temp_users)))
	return temp_users
}
