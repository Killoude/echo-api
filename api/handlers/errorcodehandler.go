package handlers

const (
	//NoDataFound データが無いとき
	NoDataFound = 40001
)

var statusText = map[int]string{
	NoDataFound: "No Data Found",
}

//StatusText set text to errorcode
func StatusText(code int) string {
	return statusText[code]
}
