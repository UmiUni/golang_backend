user_registration_api.md API:

**User Signup**:
```$xslt
POST /signup HTTP/1.1
Host: 178.128.0.108:3001
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: e6223021-522c-45fe-9336-e5a9aeaf559e

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
Postman-Token: c21b47e0-1cea-4270-b05a-255a308f6652

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
Postman-Token: f9cd548e-c112-4f6f-ad6a-02ea0e8c25c0

{
	"Username": "mliu",
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
Postman-Token: d7c1297f-61c6-4d96-b889-a899248edc29

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
Postman-Token: 9f4d4eb5-5a91-4d7f-b648-107b9073aab3

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
