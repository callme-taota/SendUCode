import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:sender/stroe/ThemeProvider.dart';

class SettingPage extends StatefulWidget {
  @override
  _SettingPageState createState() => _SettingPageState();
}

mixin _SettingPageMixin <T extends StatefulWidget> on State<T> {
  bool isDarkMode = ThemeProvider().isDarkMode;
  final List<String> dropdownItems = ['Local Network', 'Composite Connection', 'Server Only'];
  void DarkModeChange(value){
    setState(() {
      isDarkMode = value;
      Provider.of<ThemeProvider>(context, listen: false).toggleTheme();
    });
  }
}


class _SettingPageState extends State<SettingPage> with _SettingPageMixin {
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Settings'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SettingItem(
              title: 'Dark Mode',
              widgetOnRight: Switch(
                value: isDarkMode,
                onChanged: (value) {
                  DarkModeChange(value);
                },
              ),
            ),
            SettingItem(
              title: 'Scan QR Code',
              widgetOnRight: ElevatedButton(
                onPressed: () {
                  // 处理扫码操作
                },
                child: Text('Scan'),
              ),
            ),
            SettingItem(
              title: 'Delete User Records',
              widgetOnRight: ElevatedButton(
                onPressed: () {
                  // 处理删除用户记录操作
                },
                child: Text('Delete'),
              ),
            ),
            SettingItem(
              title: 'Connection Mode',
              widgetOnRight: DropdownButton<String>(
                items: dropdownItems
                    .map((String value) {
                  return DropdownMenuItem<String>(
                    value: value,
                    child: Text(value),
                  );
                }).toList(),
                onChanged: (value) {
                  
                },
                value: 'Local Network', // 初始值，你需要根据实际情况设置
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class SettingItem extends StatelessWidget {
  final String title;
  final Widget widgetOnRight;

  const SettingItem({
    required this.title,
    required this.widgetOnRight,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(title),
        widgetOnRight,
      ],
    );
  }
}
