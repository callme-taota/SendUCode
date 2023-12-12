import 'package:flutter/material.dart';
import 'package:sender/pages/home.dart';
import 'package:sender/pages/setting.dart';

class Layout extends StatefulWidget {
  @override
  _LayoutState createState() => _LayoutState();
}

class _LayoutState extends State<Layout> {
  int _currentIndex = 0;

  // 定义底部导航栏的项目
  final List<Widget> _pages = [HomePage(), SettingPage()];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_currentIndex],
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _currentIndex,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home, color: Colors.grey),
            activeIcon: Icon(Icons.home, color: Colors.amber),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings, color: Colors.grey),
            activeIcon: Icon(Icons.settings, color: Colors.amber),
            label: 'Setting',
          ),
        ],
        selectedItemColor: Colors.amber, // 选中项的文字和图标颜色
        unselectedItemColor: Colors.grey, // 未选中项的文字和图标颜色
      ),
    );
  }
}
