package ch5

import (
	"fmt"
)

func min(nums ...int) (int, error) {
	if nums == nil {
		err := fmt.Errorf("ch5.min: %S", "can't find min of nil array")
		return 0, err
	}
	if len(nums) == 0 {
		err := fmt.Errorf("ch5.min: %s", "can't find min of empty array")
		return 0, err
	}

	r := nums[0]
	for _, num := range nums {
		if r > num {
			r = num
		}
	}

	return r, nil
}

func max(nums ...int) (int, error) {
	if nums == nil {
		err := fmt.Errorf("ch5.max: %S", "can't find max of nil array")
		return 0, err
	}
	if len(nums) == 0 {
		err := fmt.Errorf("ch5.min: %s", "can't find max of empty array")
		return 0, err
	}

	r := nums[0]
	for _, num := range nums {
		if r < num {
			r = num
		}
	}

	return r, nil
}
