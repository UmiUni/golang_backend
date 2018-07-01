
# Insert_news endpoint:

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
All fields: Domain, Timestamp, Author, Summary, Title, URL are required fields.
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



# Get_news endpoint:

```
GET /get_news?domain=soccer HTTP/1.1
Host: localhost:3001
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: cf9bb5cb-7a3e-4466-90f2-622ea79848b2
```

## Get news by domain:
```
curl -X GET \
  'http://localhost:3001/get_news?domain=soccer' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: ceeaca02-5dd7-4764-9380-414263e77031' 
```
## Get news by id:
```
curl -X GET \
  'http://localhost:3001/get_news?id=22b35308-afcf-47dc-9b26-11e46700273a' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: ceeaca02-5dd7-4764-9380-414263e77031' 
```

