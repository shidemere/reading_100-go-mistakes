package main

import "errors"

/*
	Тут по сути своей говорится не столько про Go, сколько про программирование в целом. 
	Необходимо снижать когнитивную нагрузку на мозг, и поэтому избегать больших уровней вложенности. 
	То, как всё выстроено в join1 гораздо сложнее для понимания чем join2, хотя функционал тот же.
	Эксперт по разработке на Go писал о том, что необходимо следовать "Happy Path" - это когда по левому краю у меня успех
	а во вложенных условиях потенциальные ошибки
*/

func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", nil
}
