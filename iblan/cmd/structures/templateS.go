package structures

import "os"

var TempPath = os.Getenv("TemplatesPath")

var Elements = []string{
	"/home/voodie/iblan/ui/index.html",
	TempPath + "/navbar.tmpl",
	TempPath + "/footer.tmpl",
	TempPath + "/header.tmpl",
}
