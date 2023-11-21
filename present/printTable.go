package present

import (
	"fmt"

	"github.com/lastvoidtemplar/github-api-client/utils"
)

const intLen int = 6

func addSpaces(v any, width int) string {
	var str string
	switch t := v.(type) {
	case int:
		str = fmt.Sprintf("%d", t)
	case string:
		str = t
	case bool:
		str = fmt.Sprintf("%t", t)
	}
	width = width - len(str)
	half := width / 2
	otherHalf := width - half
	var frontSpaces string
	for i := 0; i < half; i++ {
		frontSpaces = frontSpaces + " "
	}
	var backSpaces string
	for i := 0; i < otherHalf; i++ {
		backSpaces = backSpaces + " "
	}
	return fmt.Sprintf("%s%s%s", frontSpaces, str, backSpaces)
}

func PrintTable(info []TableInfo) {
	widths := []int{
		utils.MaxStringLen(info, func(el *TableInfo) string { return el.Login }),
		utils.MaxStringLen(info, func(el *TableInfo) string { return el.Name }),
		len("Admin"),
		intLen,
		intLen,
		utils.MaxStringLen(info, func(el *TableInfo) string { return el.LangsDistro }),
		len("Followers"),
		len("Following"),
		utils.MaxStringLen(info, func(el *TableInfo) string { return el.YearActivityDistro }),
		len("CreatedAt"),
	}
	fmt.Printf("| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
		addSpaces("Login", widths[0]),
		addSpaces("Name", widths[1]),
		addSpaces("Admin", widths[2]),
		addSpaces("Repos", widths[3]),
		addSpaces("Forks", widths[4]),
		addSpaces("LangsDistro", widths[5]),
		addSpaces("Followers", widths[6]),
		addSpaces("Following", widths[7]),
		addSpaces("YearActivity", widths[8]),
		addSpaces("CreatedAt", widths[9]),
	)
	for ind := 0; ind < len(info); ind++ {
		fmt.Printf("| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			addSpaces(info[ind].Login, widths[0]),
			addSpaces(info[ind].Name, widths[1]),
			addSpaces(info[ind].SiteAdmin, widths[2]),
			addSpaces(info[ind].PublicRepos, widths[3]),
			addSpaces(info[ind].Forks, widths[4]),
			addSpaces(info[ind].LangsDistro, widths[5]),
			addSpaces(info[ind].Followers, widths[6]),
			addSpaces(info[ind].Following, widths[7]),
			addSpaces(info[ind].YearActivityDistro, widths[8]),
			addSpaces(info[ind].CreatedAt.Year(), widths[9]),
		)

	}
}
