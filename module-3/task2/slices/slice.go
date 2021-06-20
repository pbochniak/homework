package slices

func SortStringsInPlace(strings []string) {
	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i] > strings[j] {
				strings[i], strings[j] = strings[j], strings[i]
			}
		}
	}
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

func (that User) Compare(another User) int {
	if that.FirstName > another.FirstName {
		return 1
	}
	if that.FirstName < another.FirstName {
		return -1
	}
	if that.LastName > another.LastName {
		return 1
	}
	if that.LastName < another.LastName {
		return -1
	}
	return 0
}

func SortUsersPure(users []User) []User {
	temp_users := make([]User, len(users))
	copy(temp_users, users)
	for i := 0; i < len(temp_users); i++ {
		for j := i + 1; j < len(temp_users); j++ {
			if temp_users[i].Compare(temp_users[j]) > 0 {
				temp_users[i], temp_users[j] = temp_users[j], temp_users[i]
			}
		}
	}
	return temp_users
}

func SortUsersPureDesc(users []User) []User {
	temp_users := make([]User, len(users))
	copy(temp_users, users)
	for i := 0; i < len(temp_users); i++ {
		for j := i + 1; j < len(temp_users); j++ {
			if temp_users[i].Compare(temp_users[j]) < 0 {
				temp_users[i], temp_users[j] = temp_users[j], temp_users[i]
			}
		}
	}
	return temp_users
}
