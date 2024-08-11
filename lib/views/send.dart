import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';

class SendData extends StatefulWidget {
  @override
  _SendDataState createState() => _SendDataState();
}

class _SendDataState extends State<SendData> {
  TextEditingController field1Controller = TextEditingController();
  TextEditingController field2Controller = TextEditingController();
  TextEditingController field3Controller = TextEditingController();

  Future<void> sendDataToServer() async {
    var url = Uri.parse('http://192.168.137.183:8080'); // استبدل بعنوان الخادم ومنفذه
    dynamic field1Value = field1Controller.text;
    dynamic field2Value = field2Controller.text;
    dynamic field3Value = field3Controller.text;
    // يمكنك تكرار هذه الخطوة لبقية الحقول

    Map<String, dynamic> data = {
      'field1': field1Value,
      'field2': field2Value,
      'field3': field3Value,
      // وهكذا لبقية الحقول
    };
    String jsonString = json.encode(data);

    try {
      var response = await http.post(
        url,
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonString,
      );

      if (response.statusCode == 200) {
        print('تم إرسال البيانات بنجاح');
      } else {
        print('فشل في إرسال البيانات. الرمز: ${response.statusCode}');
      }
    } catch (e) {
      print('حدث خطأ أثناء إرسال البيانات: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('إرسال البيانات إلى الخادم'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            TextField(
              controller: field1Controller,
              decoration: InputDecoration(
                hintText: 'ادخل القيمة للحقل الأول',
              ),
            ),
            TextField(
              controller: field2Controller,
              decoration: InputDecoration(
                hintText: 'ادخل القيمة للحقل الثاني',
              ),
            ),
            TextField(
              controller: field3Controller,
              decoration: InputDecoration(
                hintText: 'ادخل القيمة للحقل الثاني',
              ),
            ),
            // وهكذا لبقية الحقول
            ElevatedButton(
              onPressed: () {
                sendDataToServer();
              },
              child: Text('إرسال البيانات إلى الخادم'),
            ),
          ],
        ),
      ),
    );
  }
}