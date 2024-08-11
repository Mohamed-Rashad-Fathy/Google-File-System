import 'package:flutter/material.dart';
import 'package:go_project/components/field.dart';
import 'package:go_project/model/var.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class Orders extends StatefulWidget {
  @override
  _OrdersState createState() => _OrdersState();
}

class _OrdersState extends State<Orders> {
  TextEditingController field1Controller = TextEditingController();
  TextEditingController field2Controller = TextEditingController();
  TextEditingController field3Controller = TextEditingController();
  TextEditingController field4Controller = TextEditingController();
  // TextEditingController field5Controller = TextEditingController();

  Future<void> orderstoserver() async {
    var url =
        Uri.parse('http://$ip:$port/customers'); // استبدل بعنوان الخادم ومنفذه
    dynamic field1Value = field1Controller.text;
    dynamic field2Value = field2Controller.text;
    dynamic field3Value = field3Controller.text;
    dynamic field4Value = field4Controller.text;
    // dynamic field5Value = field5Controller.text;
    //dynamic typeOperation = field4Controller.text;
    // يمكنك تكرار هذه الخطوة لبقية الحقول
    Map<String, dynamic> data = {
      'field1': field1Value,
      'field2': field2Value,
      'field3': field3Value,
      'field4': field4Value,
      // 'field5': field5Value,
      //'field4':typeOperation,
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
        title: const Text('Order Table'),
      ),
      body: Center(
        child: ListView(
          children: <Widget>[
            FeildText(field1Controller: field1Controller, text: "Order ID"),
            FeildText(field1Controller: field2Controller, text: "Customer ID"),
            // FeildText(field1Controller: field3Controller, text: "Order Date"),
            FeildText(field1Controller: field3Controller, text: "Total Amount"),
            FeildText(
                field1Controller: field4Controller, text: "Type Operation"),
            ElevatedButton(
              onPressed: () {
                orderstoserver();
              },
              child: const Text('إرسال البيانات إلى الخادم'),
            ),
          ],
        ),
      ),
    );
  }
}
