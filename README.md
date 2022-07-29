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

## tests

### fyne

```
apt-get install xorg-dev
go get fyne.io/fyne/v2/cmd/fyne_demo/
```

### qt

<https://github.com/therecipe/qt/wiki/Installation>

```
sudo apt-get install libqt5quickcontrols2-5 libqt5multimedia5 libqt5webengine5 libqt5quick5 libqt5qml5
export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false
github.com/therecipe/examples/basic/widgets
```

for windows?
```
apt-get install gcc-multilib
apt-get install gcc-mingw-w64
#GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc 
```
https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1