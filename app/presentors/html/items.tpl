<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Товары</title>
</head>
<body>
{{range .}}
    <div>
        {{ .Id}}
    </div>
    <div>
        {{ .Name}}
    </div>
    <div>
        {{ .Description}}
    </div>
    <div>
        {{ .Article}}
    </div>
    <div>
        {{ .Description}}
    </div>
    <div>
        {{ .Price.Base}}
    </div>
    <div>
        {{ .Price.Discounted}}
    </div>
{{end}}

</body>
</html>