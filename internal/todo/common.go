package todo

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)


var debug_flag = false

type DataItem struct {
	ID          int
	Description *string
	CreatedAt   *string
	IsComplete  *bool
}

func CheckError(err error, prefix string) bool {
	if err != nil {
		if debug_flag {
			fmt.Println(prefix + ": " + err.Error())
		}

		return true
	}

	return false
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)

	return !errors.Is(error, os.ErrNotExist)
}

func getTimeDiff(timeStr string) string {
	s1, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return "invalid time"
	}
	s2 := time.Now()

	s3 := s2.Sub(s1)
	minutes := int(s3.Minutes())

	timeDiff := "a few seconds ago"
	if minutes >= 525600 {
		years := minutes / 525600
		if years == 1 {
			timeDiff = "a year ago"
		} else {
			timeDiff = fmt.Sprintf("%v years ago", years)
		}
	} else if minutes >= 43800 {
		months := minutes / 43800
		if months == 1 {
			timeDiff = "a month ago"
		} else {
			timeDiff = fmt.Sprintf("%v months ago", months)
		}
	} else if minutes >= 10080 {
		weeks := minutes / 10080
		if weeks == 1 {
			timeDiff = "a week ago"
		} else {
			timeDiff = fmt.Sprintf("%v weeks ago", weeks)
		}
	} else if minutes >= 1440 {
		days := minutes / 1440
		if days == 1 {
			timeDiff = "a day ago"
		} else {
			timeDiff = fmt.Sprintf("%v days ago", days)
		}
	} else if minutes >= 60 {
		hours := minutes / 60
		if hours == 1 {
			timeDiff = "an hour ago"
		} else {
			timeDiff = fmt.Sprintf("%v hours ago", hours)
		}
	} else if minutes > 0 {
		if minutes == 1 {
			timeDiff = "an minute ago"
		} else {
			timeDiff = fmt.Sprintf("%v minutes ago", minutes)
		}
	}

	return timeDiff
}


func DisplayItems(items []DataItem, hideDone bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	if !hideDone {
		fmt.Fprintf(w, "%v \t%v\t%v\t%v \n ", "ID", "Task", "Created", "Done")

		for _, item := range items {
			if item.ID > 0 {
				fmt.Fprintf(w, "%v \t%v\t%v\t%v \n ", item.ID, *item.Description, getTimeDiff(*item.CreatedAt), *item.IsComplete)
			}
		}
	}else{
		fmt.Fprintf(w, "%v \t%v\t%v \n ", "ID", "Task", "Created")

		for _, item := range items {
			if item.ID > 0 {
				fmt.Fprintf(w, "%v \t%v\t%v \n ", item.ID, *item.Description, getTimeDiff(*item.CreatedAt))
			}
		}
	}

	w.Flush()
}
