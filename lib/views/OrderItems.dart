import 'package:flutter/material.dart';
import 'package:go_project/components/field.dart';
import 'package:go_project/model/var.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class OrderItems extends StatefulWidget {
  @override
  _OrderItemsState createState() => _OrderItemsState();
}

class _OrderItemsState extends State<OrderItems> {
  TextEditingController field1Controller = TextEditingController();
  TextEditingController field2Controller = TextEditingController();
  TextEditingController field3Controller = TextEditingController();
  TextEditingController field4Controller = TextEditingController();
  TextEditingController field5Controller = TextEditingController();
  TextEditingController field6Controller = TextEditingController();
  // TextEditingController field7Controller = TextEditingController();

  Future<void> OrderItemsToServer() async {
    var url = Uri.parse(
        'http://$ip:$port/customers'); // استبدل بعنوان الخادم ومنفذه
    dynamic field1Value = field1Controller.text;
    dynamic field2Value = field2Controller.text;
    dynamic field3Value = field3Controller.text;
    dynamic field4Value = field4Controller.text;
    dynamic field5Value = field5Controller.text;
    dynamic field6Value = field6Controller.text;
    // dynamic field7Value = field7Controller.text;
    //dynamic typeOperation = field4Controller.text;
    Map<String, dynamic> data = {
      'field1': field1Value,
      'field2': field2Value,
      'field3': field3Value,
      'field4': field4Value,
      'field5': field5Value,
      'field6': field6Value,
      // 'field7': field7Value,
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
        title: const Text('Medicines Table'),
      ),
      body: Center(
        child: ListView(
          children: <Widget>[
            FeildText(
                field1Controller: field1Controller, text: "Order Item ID"),
            FeildText(field1Controller: field2Controller, text: "Order ID"),
            FeildText(field1Controller: field3Controller, text: "Medicine ID"),
            FeildText(field1Controller: field4Controller, text: "Quantity"),
            FeildText(field1Controller: field5Controller, text: "Unit Price"),
            FeildText(
                field1Controller: field6Controller, text: "Type Operation"),
            ElevatedButton(
              onPressed: () {
                OrderItemsToServer();
              },
              child: const Text('Send Data To Server'),
            ),
          ],
        ),
      ),
    );
  }
}
