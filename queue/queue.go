package queue

func Qadd[T interface{}](queue []T, element T, less func(T, T) bool) (result []T) {
	result = append(queue, element)
	for i := len(result) - 1; i > 0; i = i / 2 {
		if less(result[i], result[i/2]) {
			tmp := result[i/2]
			result[i/2] = result[i]
			result[i] = tmp
		}
	}
	return
}

func Qpop[T interface{}](queue []T, less func(T, T) bool) (result *T, newq []T, ok bool) {
	if len(queue) == 0 {
		return nil, queue, false
	}

	ok = true
	r1 := queue[0]
	result = &r1

	queue[0] = queue[len(queue)-1]

	for i := 0; 2*i+1 < len(queue); {
		var j int

		if 2*i+1 == len(queue)-1 {
			j = 2*i + 1
		} else if less(queue[2*i+1], queue[2*i+2]) {
			j = 2*i + 1
		} else {
			j = 2*i + 2
		}

		if !less(queue[i], queue[j]) {
			tmp := queue[i]
			queue[i] = queue[j]
			queue[j] = tmp
		}

		i = j
	}

	newq = queue[0 : len(queue)-1]

	return
}
