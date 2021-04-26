# flap

Working proof of the local Go server running within Flutter

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

### App crashes when I hit fetch button repeatedly and quickly on iOS
Yeah please fix them.

## Contribution

I'd like to see stuff working on other platforms.
