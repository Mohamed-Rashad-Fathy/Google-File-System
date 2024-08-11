import 'package:flutter/material.dart';

class FeildText extends StatelessWidget {
  const FeildText(
      {super.key, required this.field1Controller, required this.text});
  final TextEditingController field1Controller;
  final String text;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(10),
      child: TextFormField(
        controller: field1Controller,
        decoration: InputDecoration(
            label: Text(
              text,
              style: const TextStyle(fontSize: 15),
            ),
            border:
                OutlineInputBorder(borderRadius: BorderRadius.circular(15))),
      ),
    );
  }
}
