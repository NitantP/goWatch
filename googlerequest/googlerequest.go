package googlerequest
import "strconv"


func BuildURL(stock string, interval int) string {

	return "https://www.google.com/finance/getprices?q=" + stock + "&i=" + strconv.Itoa(interval) + "&f=d,c,o,h,l"
}
