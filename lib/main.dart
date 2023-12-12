import 'package:flutter/material.dart';
import 'package:sender/pages/layout.dart';
import 'package:provider/provider.dart';
import 'package:sender/stroe/ThemeProvider.dart';

void main() {
  runApp(
    ChangeNotifierProvider(
      create: (context) => ThemeProvider(),
      child: MyApp(),
    ),
  );
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Sender',
      theme: ThemeData(colorScheme: const ColorScheme.light()), // 亮色主题
      darkTheme: ThemeData(colorScheme: const ColorScheme.dark()), // 主题模式
       themeMode: Provider.of<ThemeProvider>(context).isSystemMode
        ? ThemeMode.system
        : (Provider.of<ThemeProvider>(context).isDarkMode
            ? ThemeMode.dark
            : ThemeMode.light),
      home: Layout(),
    );
  }
}
