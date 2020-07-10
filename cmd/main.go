package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"srt2fcpxml/core"
	"strings"
	"github.com/asticode/go-astisub"
)

func main() {
	srtFile := flag.String("srt", "", "srt 字幕文件")
	frameDuration := flag.Int("fd", 25, "帧率目前只支持整数24、25、30、50、60 (default 25)")
	width := flag.Int("w", 1920, "分辨率宽 (default 1920)")
	height := flag.Int("h", 1080, "分辨率高 (default 1080)")
	lineBreak := flag.String("line", " - ", "如果单条字幕有多行，行和行之间以line进行拼接 当传\\\\n字符时会以换行符进行拼接 (default - )")
	flag.Parse()
	*lineBreak = strings.Replace(string(*lineBreak), "\\n", "\n", -1 )
	f, _ := astisub.OpenFile(*srtFile)
	out := `<?xml version="1.0" encoding="UTF-8" ?>
	<!DOCTYPE fcpxml>
	
	`
	if len(*srtFile) == 0 {
		flag.PrintDefaults()
		os.Exit(20)
	}

	project, path := getPath(*srtFile)
	result, _ := core.Srt2FcpxmlExport(project, *frameDuration, *width, *height, *lineBreak, f)
	out += strings.Replace(string(result), "&#xA;", "\n", -1 )
	targetFile := fmt.Sprintf("%s/%s.fcpxml", path, project)
	fd, err := os.Create(targetFile)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = fd.Write([]byte(out))
	if err != nil {
		fmt.Println(err)
	}
}

func getPath(filePath string) (projectName, targetPath string) {
	path, _ := filepath.Abs(filePath)
	parts := strings.Split(path, "/")
	projectName = func(file string) string {
		parts := strings.Split(file, ".")
		return strings.Join(parts[0:len(parts)-1], ".")
	}(parts[len(parts)-1])
	targetPath = func(parts []string) string {
		return strings.Join(parts, "/")
	}(parts[0 : len(parts)-1])
	return
}
