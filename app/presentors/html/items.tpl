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
        {{ .Price}}
    </div>
{{end}}

</body>
</html>