# Lyrid Golang 1.x Gin Template

This tools does a blind L7 reverse proxies a hostname to any other endpoint (that is internet accesible) 

You can either set the proxy using a file: map.json
or just 

**Disclaimer:** Keep in mind that some websites does not like to be "proxies", we are not responsible for any misuse of this tool.

## Run locally with:


```
go get
go run ./main.go
```

Open http://localhost:8000

## Edit the names (optional):
Open .lyrid-definition and change the App and Module name, because this will override another applications with the same name in the platform.

## Then submit to Lyrid Platform:

```
lc code submit
```
Wait until the cloud platform to finish with the build and the default deployment.

## Start hacking:

Edit the route url, settings, and views at /entry folder with your custom APIs.

Add more middlewares or your business logic in there.

## Use Google Sheet

https://benborgers.com/posts/google-sheets-json
