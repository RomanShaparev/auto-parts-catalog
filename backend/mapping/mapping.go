package mapping

func Map[T, U any](ts []T, f func(T) U) []U {
	if ts == nil {
		return nil
	}

	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
