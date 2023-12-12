import 'package:flutter/material.dart';

class ThemeProvider with ChangeNotifier {
  bool _isDarkMode = false;
  bool _isSystemMode = true;

  bool get isDarkMode => _isDarkMode;
  bool get isSystemMode => _isSystemMode;

  void toggleTheme() {
    _isDarkMode = !_isDarkMode;
    notifyListeners();
  }

  void setDark() {
    _isDarkMode = true;
    notifyListeners();
  }

  void setLight() {
    _isDarkMode = false;
    notifyListeners();
  }

  void toggleSystemMode(bool value) {
    _isSystemMode = value;
    notifyListeners();
  }
}

