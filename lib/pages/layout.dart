import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:sender/pages/home.dart';
import 'package:sender/pages/setting.dart';
import 'package:sender/stroe/ThemeProvider.dart';

class Layout extends StatefulWidget {
  @override
  _LayoutState createState() => _LayoutState();
}

class _LayoutState extends State<Layout> {
  int _currentIndex = 0;
  void changeCurrentIndex(number) {
    setState(() {
      _currentIndex = number;
    });
    return;
  }

  // 定义底部导航栏的项目
  final List<Widget> _pages = [HomePage(), SettingPage()];

  @override
  Widget build(BuildContext context) {
    bool isDarkMode =
        MediaQuery.of(context).platformBrightness == Brightness.dark;
    return Scaffold(
      body: Stack(
        children: [
          _pages[_currentIndex],
          Positioned(
            left: 80,
            right: 80,
            bottom: 60,
            child: Container(
              height: 56,
              decoration: BoxDecoration(
                color: Provider.of<ThemeProvider>(context).isSystemMode
                    ? (isDarkMode
                        ? Colors.black.withOpacity(0.9)
                        : Colors.white.withOpacity(0.9))
                    : (Provider.of<ThemeProvider>(context).isDarkMode
                        ? Colors.black.withOpacity(0.9)
                        : Colors.white.withOpacity(0.9)),
                borderRadius: BorderRadius.all(Radius.circular(25)),
                boxShadow: [
                  BoxShadow(
                    color: Colors.amber.withOpacity(0.3),
                    blurRadius: 16.0,
                  ),
                ],
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: <Widget>[
                  IconButton(
                      icon: Icon(Icons.home, color: Colors.amber),
                      onPressed: () => changeCurrentIndex(0)),
                  IconButton(
                      icon: Icon(Icons.settings, color: Colors.amber),
                      onPressed: () => changeCurrentIndex(1)),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
