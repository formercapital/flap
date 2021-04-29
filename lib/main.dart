import 'package:flutter/material.dart';
import 'dart:io';
import 'dart:ffi' as ffi;
import 'package:path_provider/path_provider.dart';
import 'package:ffi/ffi.dart';
import 'dart:developer';
import 'dart:isolate';
import 'package:http/http.dart' as http;

void main() {
  runApp(MyApp());
}

class Config {
  static const bool localDev = false;
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'local Go server'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

typedef serve_func = ffi.Void Function(ffi.Int32, ffi.Pointer<Utf8>);
typedef Serve = void Function(int, ffi.Pointer<Utf8>);

class _MyHomePageState extends State<MyHomePage> {
  int _port = 9090;
  bool _running = false;
  String _body = "";

  void _startServer() async {
    if (_running || Config.localDev) return;
    Directory supportDir = await getApplicationSupportDirectory();
    String dir = supportDir.path;
    log('application support directory: $dir');
    int p = await getUnusedPort();
    setState(() {
      _port = p;
      _running = true;
    });
    await Isolate.spawn(_launchServer, [p, dir]);
  }

  static void _launchServer(List<Object> args) async {
    int p = args[0];
    String dir = args[1];

    final dylib = () {
      if (Platform.isIOS) {
        return ffi.DynamicLibrary.open('server-ios.dylib');
      }
      return ffi.DynamicLibrary.open('server-macos.dylib');
    }();
    final Serve serve =
        dylib.lookup<ffi.NativeFunction<serve_func>>('serve').asFunction();
    serve(p, dir.toNativeUtf8());
    log('finished go server');
  }

  Future<int> getUnusedPort() async {
    ServerSocket socket = await ServerSocket.bind(
        InternetAddress.loopbackIPv4 ?? InternetAddress.anyIPv4, 0);
    var port = socket.port;
    socket.close();
    log('go server listening on port: $port');
    return port;
  }

  void _fetch() async {
    final res = await http.get(Uri.http('localhost:$_port', '/'));
    if (res.statusCode == 200) {
      setState(() {
        _body = res.body;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            if (_running || Config.localDev) Text('server port: $_port'),
            if (_running || Config.localDev)
              TextButton(onPressed: _fetch, child: Text('fetch')),
            Text('$_body'),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _startServer,
        tooltip: 'Launch server',
        child: Icon(Icons.launch),
      ),
    );
  }
}
