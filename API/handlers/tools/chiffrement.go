package tools

func Chiffrement(email string, password string) string {
	// We get a password from the user and we will encrypt it
	entire_lenght := len(email) + len(email) + len(password)
	cesar := Cesar(email+email+password, entire_lenght)

	// We reverse the password and the pseudo

	reversed_password := Reverse(password)
	reversed_pseudo := Reverse(email)

	// And now we will use the vignere encryption method

	vignere := Vignere(reversed_pseudo+email+reversed_password+cesar, cesar)

	return vignere
}

func Cesar(entry string, key int) string {
	// We will use the Cesar encryption method
	output := ""
	// We transform the entry into a array of bytes

	array := []byte(entry)

	for i := 0; i < len(array); i++ {
		output = output + string(array[i]+byte(key))
	}

	return output
}

func Reverse(entry string) string {
	// We will reverse the entry
	output := ""

	for i := len(entry) - 1; i >= 0; i-- {
		output = output + string(entry[i])
	}

	return output
}

func Vignere(entry string, key string) string {
	// We will use the Vignere encryption method
	output := ""

	// We transform the entry into a array of bytes

	array := []byte(entry)
	key_array := []byte(key)

	for i := 0; i < len(array); i++ {
		output = output + string(array[i]+key_array[i%len(key_array)])
	}

	return output
}
