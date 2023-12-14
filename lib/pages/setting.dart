import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:sender/widgets/SettingItem.dart';
import 'package:sender/stroe/ThemeProvider.dart';
import 'package:barcode_scan2/barcode_scan2.dart';

class SettingPage extends StatefulWidget {
  @override
  _SettingPageState createState() => _SettingPageState();
}

mixin _SettingPageMixin<T extends StatefulWidget> on State<T> {
  final List<String> modeItems = ['System', 'Light Mode', 'Dark Mode'];
  String sysMode = "System";
  final List<String> netWorkItems = [
    'Local Network',
    'Composite Connection',
    'Server Only'
  ];
  String SelectedNetwork = "Local Network";
  String ScanResult = "";
  String userName = "";

  void _setDefaultMode() async {
    String defaultMode =
        await Provider.of<ThemeProvider>(context, listen: false)
            .getDefaultMode();
    setState(() {
      sysMode = defaultMode;
    });
  }

  @override
  void initState() {
    super.initState();
    _setDefaultMode();
  }

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

  void _scanQRCode() async {
    var options = const ScanOptions(
        strings: {
          'cancel': "取消扫描",
          'flash_on': "打开闪光灯",
          'flash_off': "关闭闪光灯",
        },
        android: AndroidOptions(
          aspectTolerance : 0.5,
          useAutoFocus :true,
        ),
        restrictFormat:[BarcodeFormat.qr]
        );
    var result = await BarcodeScanner.scan(options: options);
    setState(() {
      ScanResult = result.rawContent;
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
                  _scanQRCode();
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
                items: netWorkItems.map((String value) {
                  return DropdownMenuItem<String>(
                    value: value,
                    child: Text(value),
                  );
                }).toList(),
                onChanged: (value) {},
                value: SelectedNetwork, // 初始值，你需要根据实际情况设置
              ),
            ),
            SettingItem(
                title: "Current username",
                margin: 10,
                widgetOnRight: Text(userName)),
          ],
        ),
      ),
    );
  }
}
