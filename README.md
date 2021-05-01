# flap

<p align="center">
<img width="150" height="150" src="https://raw.githubusercontent.com/capitalpidx/flap/master/art/logo.png"><br/>
Working proof of the Go server running inside Flutter
</p>

[Video in action](https://www.youtube.com/watch?v=4DPL5F2DVdc)

## Prerequisites

- [Flutter](https://flutter.dev) >2.0
- [Go](https://golang.org) >1.16

## Build Go server
```
cd go
```
macOS:
```
make macos
```
iOS:
```
make ios
```
iOS Simulator:
```
make ios-simulator
```

## Run app
```
cd ..
flutter devices
flutter run -d {target}
```

## Build app
```
flutter build {macos, ios}
```


## Known issues

- Hot reload doesn't work (workaround: run Go server independently in the development phase)

## Q&A

### Why .dylib instead of .a as an iOS library?
Because I can despite official docs prefer .a. In fact I haven't even tried .a.

## Contribution

I'd like to see stuff working on other platforms.
