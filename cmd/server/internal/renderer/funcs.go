package renderer

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func gt(a, b float64) bool {
	return a > b
}

func formatNumber(value float64) string {
	i := int64(value)
	s := strconv.FormatInt(i, 10)
	var parts []string
	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}
	if s != "" {
		parts = append([]string{s}, parts...)
	}
	return strings.Join(parts, " ")
}

func formatDuration(value float64) string {
	i := int(value)
	hours := i / 3600
	minutes := (i % 3600) / 60
	seconds := i % 60
	return fmt.Sprintf("%0.2dh %0.2dm %0.2ds", hours, minutes, seconds)
}

func formatTime(value string) string {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return value
	}
	return t.Format("2006-01-02 15:04")
}

func getRankName(rank float64) string {
	switch int(rank) {
	case 0:
		return "RECRUIT"
	case 1:
		return "PRIVATE I"
	case 2:
		return "PRIVATE II"
	case 3:
		return "PRIVATE III"
	case 4:
		return "SPECIALIST I"
	case 5:
		return "SPECIALIST II"
	case 6:
		return "SPECIALIST III"
	case 7:
		return "CORPORAL I"
	case 8:
		return "CORPORAL II"
	case 9:
		return "CORPORAL III"
	case 10:
		return "SERGEANT I"
	case 11:
		return "SERGEANT II"
	case 12:
		return "SERGEANT III"
	case 13:
		return "STAFF SERGEANT I"
	case 14:
		return "STAFF SERGEANT II"
	case 15:
		return "STAFF SERGEANT III"
	case 16:
		return "MASTER SERGEANT I"
	case 17:
		return "MASTER SERGEANT II"
	case 18:
		return "MASTER SERGEANT III"
	case 19:
		return "FIRST SERGEANT I"
	case 20:
		return "FIRST SERGEANT II"
	case 21:
		return "FIRST SERGEANT III"
	case 22:
		return "WARRANT OFFICER I"
	case 23:
		return "WARRANT OFFICER II"
	case 24:
		return "WARRANT OFFICER III"
	case 25:
		return "CHIEF WARRANT OFFICER I"
	case 26:
		return "CHIEF WARRANT OFFICER II"
	case 27:
		return "CHIEF WARRANT OFFICER III"
	case 28:
		return "SECOND LIEUTENANT I"
	case 29:
		return "SECOND LIEUTENANT II"
	case 30:
		return "SECOND LIEUTENANT III"
	case 31:
		return "FIRST LIEUTENANT I"
	case 32:
		return "FIRST LIEUTENANT II"
	case 33:
		return "FIRST LIEUTENANT III"
	case 34:
		return "CAPTAIN I"
	case 35:
		return "CAPTAIN II"
	case 36:
		return "CAPTAIN III"
	case 37:
		return "MAJOR I"
	case 38:
		return "MAJOR II"
	case 39:
		return "MAJOR III"
	case 40:
		return "LIEUTENANT COLONEL I"
	case 41:
		return "LIEUTENANT COLONEL II"
	case 42:
		return "LIEUTENANT COLONEL III"
	case 43:
		return "COLONEL I"
	case 44:
		return "COLONEL II"
	case 45:
		return "COLONEL III"
	case 46:
		return "BRIGADIER GENERAL I"
	case 47:
		return "BRIGADIER GENERAL II"
	case 48:
		return "BRIGADIER GENERAL III"
	case 49:
		return "GENERAL"
	case 50:
		return "GENERAL OF THE ARMY"
	default:
		return fmt.Sprintf("%.0f", rank)
	}
}
