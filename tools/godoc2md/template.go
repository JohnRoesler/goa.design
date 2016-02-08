package main

var pkgTemplate = `{{with .PDoc}}{{if $.IsMain}}
# {{ base .ImportPath }}

{{comment_md .Doc}}
{{else}}
# {{ .Name }}
    import "{{.ImportPath}}"

{{comment_md .Doc}}
{{example_html $ ""}}

{{with .Consts}}## Constants
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}
{{with .Vars}}## Variables
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}
{{range .Funcs}}{{$name_html := html .Name}}## func {{$name_html}}{{anchor $.PDoc.Name .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_html $ .Name}}{{end}}
{{range .Types}}{{$tname := .Name}}{{$tname_html := html .Name}}## type {{$tname_html}}{{anchor $.PDoc.Name .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}

{{range .Consts}}{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}

{{range .Vars}}{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}

{{example_html $ $tname}}

{{range .Funcs}}{{$name_html := html .Name}}### func {{$name_html}}{{anchor $.PDoc.Name $tname .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_html $ .Name}}{{end}}

{{range .Methods}}{{$name_html := html .Name}}### func ({{md .Recv}}) {{$name_html}}{{anchor $.PDoc.Name $tname .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}{{example_html $ $name}}
{{end}}{{end}}{{end}}

{{with $.Notes}}
{{range $marker, $content := .}}
## {{noteTitle $marker | html}}s
<ul style="list-style: none; padding: 0;">
{{range .}}
<li><a href="{{posLink_url $ .}}">&#x261e;</a> {{html .Body}}</li>
{{end}}
</ul>
{{end}}
{{end}}
{{end}}

{{with .PAst}}{{node_html $ . false}}{{end}}`
