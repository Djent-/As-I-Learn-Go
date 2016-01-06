<html>
<head>
<title>Upload File</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://localhost:9090/upload" method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{{.}}" />
	<input type="submit" name="upload" />
</body>
</html>