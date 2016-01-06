<head>
<title>Duplicate Submissions</title>
</head>
<body>
<form action="http://localhost:9090/login" method="post">
<input type="checkbox" name="interest" value="Football">Football
<input type="checkbox" name="interest" value="Basketball">Basketball
<input type="checkbox" name="interest" value="Tennis">Tennis
Username:<input type="text" name="username">
Password:<input type="password" name="password">
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="Login">
</form>
</body>