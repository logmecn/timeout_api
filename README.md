Time Out Trigger API
====
timeout_api is a Trigger which can send a defined URL, when time is out.

Installation
-----
```bash
go get github.com/logmecn/timeout_api
```

Usage
---
Run the execute file, it will work on 8001 port (default).
then post a json data for example:
```bash
curl -d '{"data":"aaaa=bbb","url":"http://127.0.0.1:8001","tout":5}' http://127.0.0.1
```
The program will visit URL: http://127.0.0.1:8001?aaaa=bbb after {tout} second. (in this case is 5s)

---
