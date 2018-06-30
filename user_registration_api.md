user_registration_api.md API:

Sign up:
```
curl -X POST \
  http://178.128.0.108:3001/signup \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: a8b8f207-21b3-44a4-962a-77560b15993e' \
  -d '{
    "Username": "superchaoran",
    "Password": "8515111q",
    "Email": "superchaoran@gmail.com"
}' | jq .
```
Sign up possibilities:
```
POST http://178.128.0.108:3001/signup
raw JSON request:
{
    "Username": "superchaoran",
    "Password": "8515111q",
    "Email": "superchaoran@gmail.com"
}
raw JSON response (success):
{
    "Status": "OK",
    "AccountType": "user",
    "Email": "superchaoran@gmail.com",
    "AuthToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA1NjY5NzksImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoic3VwZXJjaGFvcmFuIn0.ZGX1xnj_ZMyA5qS_apgsJqKboijUSJrqzT_XyNKbvU8",
    "IsLoggedIn": true
}
raw JSON response (error0):
{
    "error": "email already registered"
}
```

* 我们的sign up/sign in api, (email, username, password) 都是required；
email不能重复, username可以重复; sign in的时候必须用email sign in。

Sign in:
```
curl -X POST \
  http://178.128.0.108:3001/login \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 49f0c56a-c810-4d75-8b5e-4ffb811ae0df' \
  -d '{
    "Email": "superchaoran@gmail.com",
    "Password": "8515111"
}' | jq .
```
Sign in Possibilities:
```

raw JSON request:
{
    "Email": "superchaoran@gmail.com",
    "Password": "8515111q"
}
raw JSON response (success):
{
    "Status": "OK",
    "AccountType": "user",
    "Email": "superchaoran@gmail.com",
    "AuthToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA1Njc1MjAsImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoic3VwZXJjaGFvcmFuIn0.86i5NNArj12EACL0uPji2uKE26omLDStfb3iOx8wLgE",
    "IsLoggedIn": true
}
raw JSON response (error0):
{
    "error": "please verify your email"
}
raw JSON response (error1):
{
    "error": "invalid password"
}
```

Email activation example:
```
http://178.128.0.108:3001/activate?email=superchaoran@gmail.com&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA1NjY5NzksImlzcyI6ImpvZ2NoYXQuY29tIiwic3ViIjoic3VwZXJjaGFvcmFuIn0.ZGX1xnj_ZMyA5qS_apgsJqKboijUSJrqzT_XyNKbvU8
raw JSON response (success):
{"Congratulations":"You've activated your email."}
raw JSON response (error0):
{"error":"email already activated"}
```
