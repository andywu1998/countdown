package main

import (
	"flag"
	tm "github.com/buger/goterm"
	"strings"
	"time"
)

var (
	one = []string{
		".----------------.",
		"| .--------------. |",
		"| |     __       | |",
		"| |    /  |      | |",
		"| |    `| |      | |",
		"| |     | |      | |",
		"| |    _| |_     | |",
		"| |   |_____|    | |",
		"| |              | |",
		"| '--------------' |",
		" '----------------' ",
	}

	numsGraph = [][]string{
		{
			" .----------------. ",
			"| .--------------. |",
			"| |     ____     | |",
			"| |   .'    '.   | |",
			"| |  |  .--.  |  | |",
			"| |  | |    | |  | |",
			"| |  |  `--'  |  | |",
			"| |   '.____.'   | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |     __       | |",
			"| |    /  |      | |",
			"| |    `| |      | |",
			"| |     | |      | |",
			"| |    _| |_     | |",
			"| |   |_____|    | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |    _____     | |",
			"| |   / ___ `.   | |",
			"| |  |_/___) |   | |",
			"| |   .'____.'   | |",
			"| |  / /____     | |",
			"| |  |_______|   | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |    ______    | |",
			"| |   / ____ `.  | |",
			"| |   `'  __) |  | |",
			"| |   _  |__ '.  | |",
			"| |  | \\____) |  | |",
			"| |   \\______.'  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |   _    _     | |",
			"| |  | |  | |    | |",
			"| |  | |__| |_   | |",
			"| |  |____   _|  | |",
			"| |      _| |_   | |",
			"| |     |_____|  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |   _______    | |",
			"| |  |  _____|   | |",
			"| |  | |____     | |",
			"| |  '_.____''.  | |",
			"| |  | \\____) |  | |",
			"| |   \\______.'  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |    ______    | |",
			"| |  .' ____ \\   | |",
			"| |  | |____\\_|  | |",
			"| |  | '____`'.  | |",
			"| |  | (____) |  | |",
			"| |  '.______.'  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |   _______    | |",
			"| |  |  ___  |   | |",
			"| |  |_/  / /    | |",
			"| |      / /     | |",
			"| |     / /      | |",
			"| |    /_/       | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |     ____     | |",
			"| |   .' __ '.   | |",
			"| |   | (__) |   | |",
			"| |   .`____'.   | |",
			"| |  | (____) |  | |",
			"| |  `.______.'  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
		{
			" .----------------. ",
			"| .--------------. |",
			"| |    ______    | |",
			"| |  .' ____ '.  | |",
			"| |  | (____) |  | |",
			"| |  '_.____. |  | |",
			"| |  | \\____| |  | |",
			"| |   \\______,'  | |",
			"| |              | |",
			"| '--------------' |",
			" '----------------' ",
		},
	}
)

func graphToString(in []string) string {
	return strings.Join(in, "\n")
}

func PrintGraph(num int) string {
	digitals := []int{}
	if num == 0 {
		digitals = append(digitals, 0)
	}
	for num > 0 {
		digitals = append(digitals, num%10)
		num /= 10
	}
	graphs := [][]string{}
	maxRows := 0
	for i := len(digitals) - 1; i >= 0; i-- {
		dig := digitals[i]
		graph := numsGraph[dig]
		graphs = append(graphs, graph)
		if len(graph) > maxRows {
			maxRows = len(graph)
		}
	}

	return merge(graphs, maxRows)
}

func merge(graphs [][]string, maxRows int) string {
	lines := []string{}
	for i := 0; i < maxRows; i++ {
		lines = append(lines, "")
	}
	for _, graph := range graphs {
		for i := 0; i < len(graph); i++ {
			lines[i] += graph[i]
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	tm.Clear() // Clear current screen

	second := flag.Int("n", 300, "0")
	flag.Parse()
	total := time.Duration((*second)) * time.Second
	process := ""
	cnt := *second
	tm.Flush()

	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)
		tm.Println("Current Time:", time.Now().Format(time.RFC1123))
		tm.Println(tm.Color("duration", cnt%8), total)
		total = total - time.Second

		tm.Println(process)
		process = process + "->"
		tm.Println(PrintGraph(cnt))

		if (*second-cnt+1)%10 == 0 {
			process = ""
		}
		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second)
		tm.Clear()
		if total < 0 {
			break
		}
		cnt--
		tm.Flush()
	}
}
