# flap

Working proof of the local Go server running within Flutter

## Prerequisites

- [Flutter](https://flutter.dev) < 2.0
- [Go](https://golang.org) < 1.16

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

- Hot reload doesn't work (workaround: run Go server independently on the development phase)

## Contribution

I welcome stuff implemented for other platforms.