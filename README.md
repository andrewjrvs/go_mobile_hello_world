# go_mobile_hello_world

## setup

```
# go get golang.org/x/mobile/cmd/gomobile
# gomobile init
```

https://pkg.go.dev/golang.org/x/mobile/example/basic

Testing steps

run the following code - then copy the basic.apk onto a mobile device
```
go get -d golang.org/x/mobile/example/basic
gomobile build golang.org/x/mobile/example/basic # will build an APK
```