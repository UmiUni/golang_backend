user_registration_api.md API:

**User Signup**:
```$xslt
POST /signup HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
{
	"Email": "liumengxiong1218@gmail.com"
}
```

**Activate Email**:
```$xslt
POST /activate HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
{
	"Username": "mliu",
	"Password": "mengxiong",
	"Email": "liumengxiong1218@gmail.com",
	"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEwMTc0MTMsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoibGl1bWVuZ3hpb25nMTIxOEBnbWFpbC5jb20ifQ.ft10FCf6ONPg64f7oPqWY6Y1Dgo5Sx_tpLObbT084do"
}
```

**User Login**:
```$xslt
POST /login HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
{
	"Password": "mengxiong1218",
	"Email": "liumengxiong1218@gmail.com"
}
```

**User Request Password Reset**:
```$xslt
POST /reset_request HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
{
	"Email": "liumengxiong1218@gmail.com"
}
```
* Backend will generate a new password reset token, and send user an email with password reset link and the generated verification token.

**User Reset Password**:
```$xslt
POST /reset_password HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
{
	"Email": "liumengxiong1218@gmail.com",
	"Password": "mengxiong1218",
	"Token": "eyJbbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzEwMTc2NzIsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoibGl1bWVuZ3hpb25nMTIxOEBnbWFpbC5jb20ifQ.Zel4ZEG2ALmYxw8kLVbsbobj_foUB1TuTqshvgNybkI"
}
```

* 我们的sign up/sign in api, (email, username, password) 都是required；
* email, username都不能重复; sign in的时候必须用email sign in。
* 前端调用api时候最好软性要求用户使用公司、edu注册。虽然现在后端还没有做validation必须要求公司使用公司, edu邮箱注册。 
* company or edu email (e.g. @airbnb.com @stanford.edu)
