
func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	prevValue := <-inputStream
	outputStream <- prevValue

	for value := range inputStream {
		if prevValue == value {
			continue
		}

		outputStream <- value
		prevValue = value
	}
}





