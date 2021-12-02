package e

func ErrMsg(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
