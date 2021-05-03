<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Товар</title>
</head>
<body>
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
    {{ .Category}}
</div>
<div>
    {{ .Price.Base }}
</div>

<div>
    {{ .Price.Discounted }}
</div>
</body>
</html>