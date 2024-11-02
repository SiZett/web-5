func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan) // Закрываем выходной канал при выходе из goroutine

		for {
			select {
			case firstValue, ok := <-firstChan:
				if !ok {
					return
				}
				resultChan <- firstValue * firstValue
				return // Выходим из goroutine после первой обработки

			case secondValue, ok := <-secondChan:
				if !ok {
					return
				}
				resultChan <- secondValue * 3
				return // Выходим из goroutine после первой обработки

			case <-stopChan:
				return
			}
		}
	}()

	return resultChan
}
   
   