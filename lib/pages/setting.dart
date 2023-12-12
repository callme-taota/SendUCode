import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:sender/widgets/SettingItem.dart';
import 'package:sender/stroe/ThemeProvider.dart';

class SettingPage extends StatefulWidget {
  @override
  _SettingPageState createState() => _SettingPageState();
}

mixin _SettingPageMixin<T extends StatefulWidget> on State<T> {
  final List<String> modeItems = ['System', 'Light Mode', 'Dark Mode'];
  String sysMode = "System";
  final List<String> dropdownItems = [
    'Local Network',
    'Composite Connection',
    'Server Only'
  ];
  String userName = "";

  void darkModeChange(value) {
    setState(() {
      sysMode = value;
      switch (value) {
        case "System":
          Provider.of<ThemeProvider>(context, listen: false)
              .toggleSystemMode(true);
          return;
        case "Light Mode":
          Provider.of<ThemeProvider>(context, listen: false)
              .toggleSystemMode(false);
          Provider.of<ThemeProvider>(context, listen: false).setLight();
          return;
        case "Dark Mode":
          Provider.of<ThemeProvider>(context, listen: false)
              .toggleSystemMode(false);
          Provider.of<ThemeProvider>(context, listen: false).setDark();
          return;
      }
    });
  }
}

class _SettingPageState extends State<SettingPage> with _SettingPageMixin {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Center(child: Text('Settings')),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SettingItem(
                title: 'Dark Mode',
                margin: 6,
                widgetOnRight: DropdownButton<String>(
                  items: modeItems.map((String value) {
                    return DropdownMenuItem<String>(
                      value: value,
                      child: Text(value),
                    );
                  }).toList(),
                  onChanged: (value) {
                    darkModeChange(value);
                  },
                  value: sysMode,
                )),
            SettingItem(
              title: 'Scan QR Code',
              margin: 6,
              widgetOnRight: ElevatedButton(
                onPressed: () {
                  // 处理扫码操作
                },
                child: const Text('Scan'),
              ),
            ),
            SettingItem(
              title: 'Delete User Records',
              margin: 6,
              widgetOnRight: ElevatedButton(
                onPressed: () {
                  // 处理删除用户记录操作
                },
                child: const Text('Delete'),
              ),
            ),
            SettingItem(
              title: 'Connection Mode',
              margin: 6,
              widgetOnRight: DropdownButton<String>(
                items: dropdownItems.map((String value) {
                  return DropdownMenuItem<String>(
                    value: value,
                    child: Text(value),
                  );
                }).toList(),
                onChanged: (value) {},
                value: 'Local Network', // 初始值，你需要根据实际情况设置
              ),
            ),
            SettingItem(title: "Current username",margin: 10, widgetOnRight: Text(userName)),
          ],
        ),
      ),
    );
  }
}
