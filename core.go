package main
import (
    "strings"
    "strconv"
    "path/filepath"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

sizes:=map[string]int64{};counts:=map[string]int{}
for _,line:=range strings.Split(input,"\n"){parts:=strings.SplitN(line,"|",2);if len(parts)!=2{continue}
size,err:=strconv.ParseInt(strings.TrimSpace(parts[1]),10,64);if err!=nil||size<0{continue}
ext:=strings.ToLower(filepath.Ext(parts[0]));if ext==""{ext="(none)"};sizes[ext]+=size;counts[ext]++}
return map[string]interface{}{"bytes":sizes,"counts":counts}
}
