```
POST /insert_news HTTP/1.1
Host: localhost:3001
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: fefda49c-6bfa-4359-acec-de23d0b44fd0

{
    "Domain": "soccer",
    "Timestamp": "20180701",
    "Author": "Mengxiong Liu",
    "Summary": "My summary",
    "Title": "My title",
    "URL": "My URL"
}
```

```
curl -X POST \
  http://localhost:3001/insert_news \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 1ac9cf6f-cbbf-4bd4-b6db-c7c6841e798f' \
  -d '{
    "Domain": "soccer",
    "Timestamp": "20180701",
    "Author": "Mengxiong Liu",
    "Summary": "My summary",
    "Title": "My title",
    "URL": "My URL"
}'
```
