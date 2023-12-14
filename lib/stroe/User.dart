import 'package:shared_preferences/shared_preferences.dart';

class User {
  String userID;
  String session;

  User({required this.userID, required this.session});

  factory User.fromMap(Map<String, dynamic> map) {
    return User(
      userID: map['userID'],
      session: map['session'],
    );
  }

  Map<String, dynamic> toMap() {
    return {
      'userID': userID,
      'session': session,
    };
  }

  // 保存用户信息到本地存储
  Future<void> saveToPrefs() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setString('userId', userID);
    prefs.setString('session', session);
  }

  // 从本地存储加载用户信息
  static Future<User?> loadFromPrefs() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String? userId = prefs.getString('userId');
    String? session = prefs.getString('session');

    if (userId != null && session != null) {
      return User(userID: userId, session: session);
    } else {
      return null;
    }
  }
}
