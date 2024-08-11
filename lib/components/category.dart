import 'package:flutter/material.dart';

class Category extends StatelessWidget {
  const Category(
      {super.key,
      required this.color,
      required this.text,
      required this.ontap});
  final Color color;
  final String text;
  final Function() ontap;
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: ontap,
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
        padding: const EdgeInsets.only(left: 16),
        alignment: Alignment.centerLeft,
        width: double.infinity,
        height: 65,
        decoration: BoxDecoration(
          color: color,
          borderRadius: BorderRadius.circular(15),
        ),
        child: Text(
          text,
          style: const TextStyle(color: Colors.black, fontSize: 18),
        ),
      ),
    );
  }
}
