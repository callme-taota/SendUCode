import 'package:flutter/material.dart';

class SettingItem extends StatelessWidget {
  final String title;
  final Widget widgetOnRight;
  final int margin;

  const SettingItem({
    required this.title,
    required this.widgetOnRight,
    required this.margin
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        SizedBox(height: 1.0 * margin), 
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(title),
            widgetOnRight,
          ],
        ),
        SizedBox(height: 1.0 * margin), 
      ],
    );
  }
}
