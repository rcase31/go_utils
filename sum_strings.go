package utils

import "strconv"

func SumStrings(n string, m string) (*string, error) {

	iN, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}

	iM, err := strconv.Atoi(m)
	if err != nil {
		return nil, err
	}

	result := iM + iN
	strResult := strconv.Itoa(result)
	return &strResult, nil

}
