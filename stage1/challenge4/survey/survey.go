package survey

func CountBigShoppers(N int, s []int) int {
	t_unbought := 0
	for _, t := range s {
		t_unbought += N - t
	}
	if N-t_unbought < 0 {
		return 0
	} else {
		return N - t_unbought
	}
}
