package graduation

import (
	"strings"
	"time"
)

func formatDate(date int64) string {
	dateWithDot := strings.SplitN(time.UnixMilli(date).Local().String(), "+", 2)[0]
	return strings.Split(dateWithDot, ".")[0]

}

func verifyPassingGrade(grade int) bool {
	return grade >= 70
}
