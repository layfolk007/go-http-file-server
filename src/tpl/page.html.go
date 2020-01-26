package tpl

import (
	"../serverErrHandler"
	"./util"
	"html/template"
	"path"
)

const pageTplStr = `
{{$subItemPrefix := .SubItemPrefix}}
<!DOCTYPE html>
<html lang="">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="initial-scale=1"/>
<meta name="format-detection" content="telephone=no"/>
<meta name="renderer" content="webkit"/>
<meta name="wap-font-scale" content="no"/>
<title>{{.Path}}</title>
<link rel="stylesheet" type="text/css" href="{{.RootRelPath}}?assert=main.css"/>
</head>
<body class="{{if .IsRoot}}root-dir{{else}}sub-dir{{end}}">
<ol class="path-list">
{{range .Paths}}
<li><a href="{{.Path}}">{{fmtFilename .Name}}</a></li>
{{end}}
</ol>
{{if .CanUpload}}
<div class="upload">
<form method="POST" enctype="multipart/form-data">
<input type="file" name="files" class="files" multiple="multiple" accept="*/*"/>
<input type="submit" value="Upload"/>
</form>
</div>
{{end}}
{{if .CanArchive}}
<div class="archive">
<a href="{{$subItemPrefix}}?tar" download="{{.ItemName}}.tar">.tar</a>
<a href="{{$subItemPrefix}}?tgz" download="{{.ItemName}}.tar.gz">.tar.gz</a>
<a href="{{$subItemPrefix}}?zip" download="{{.ItemName}}.zip">.zip</a>
</div>
{{end}}
<ul class="item-list">
<li class="dir parent">
<a href="{{if .IsRoot}}./{{else}}../{{end}}">
<span class="name">../</span>
<span class="size"></span>
<span class="time"></span>
</a>
</li>
{{range .SubItems}}{{with .Html}}
<li class="{{if .IsDir}}dir{{else}}file{{end}}">
<a href="{{$subItemPrefix}}{{.Link}}{{if .IsDir}}/{{end}}">
<span class="name">{{.Name}}{{if .IsDir}}/{{end}}</span>
<span class="size">{{if not .IsDir}}{{.Size}}{{end}}</span>
<span class="time">{{.ModTime}}</span>
</a>
</li>
{{end}}{{end}}
</ul>
<script type="text/javascript" src="{{.RootRelPath}}?assert=main.js"></script>
</body>
</html>
`

var defaultPageTpl *template.Template

func init() {
	pageTpl := template.New("page")
	pageTpl = addFuncMap(pageTpl)

	var err error
	defaultPageTpl, err = pageTpl.Parse(pageTplStr)
	if serverErrHandler.CheckError(err) {
		defaultPageTpl = template.Must(pageTpl.Parse("Builtin Template Error"))
	}
}

func LoadPage(tplPath string) (*template.Template, error) {
	if len(tplPath) == 0 {
		return defaultPageTpl, nil
	}

	var err error
	pageTpl := template.New(path.Base(tplPath))
	pageTpl = addFuncMap(pageTpl)
	pageTpl, err = pageTpl.ParseFiles(tplPath)
	if err != nil {
		pageTpl = defaultPageTpl
	}

	return pageTpl, err
}

func addFuncMap(tpl *template.Template) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"fmtFilename": util.FormatFilename,
		"fmtSize":     util.FormatSize,
		"fmtTime":     util.FormatTime,
	})
}
