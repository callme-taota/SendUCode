import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ThemeProvider with ChangeNotifier {
  bool? _isDarkMode; // 使用 bool? 表示可以为 null
  bool? _isSystemMode; // 使用 bool? 表示可以为 null
  final String _darkModeKey = 'darkMode';
  final String _systemModeKey = 'systemMode';

  ThemeProvider() {
    _loadPreferences();
  }

  Future<void> loadPreferences() async {
    await _loadPreferences();
  }

  bool get isDarkMode => _isDarkMode ?? false; // 使用空值合并运算符处理 null
  bool get isSystemMode => _isSystemMode ?? true; // 使用空值合并运算符处理 null

  String getDefaultMode() {
    if (isSystemMode) return 'System';
    if (isDarkMode) return 'Dark Mode';
    return 'Light Mode';
  }

  void toggleTheme() {
    _isDarkMode = !_isDarkMode!;
    _savePreferences();
    notifyListeners();
  }

  void setDark() {
    _isDarkMode = true;
    _savePreferences();
    notifyListeners();
  }

  void setLight() {
    _isDarkMode = false;
    _savePreferences();
    notifyListeners();
  }

  void toggleSystemMode(bool value) {
    _isSystemMode = value;
    _savePreferences();
    notifyListeners();
  }

  Future<void> _loadPreferences() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    _isDarkMode = prefs.getBool(_darkModeKey) ?? false;
    _isSystemMode = prefs.getBool(_systemModeKey) ?? true;
    notifyListeners();
  }

  Future<void> _savePreferences() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setBool(_darkModeKey, _isDarkMode ?? false);
    prefs.setBool(_systemModeKey, _isSystemMode ?? true);
  }
}
