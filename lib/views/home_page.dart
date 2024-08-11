import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:go_project/components/category.dart';
import 'package:go_project/model/var.dart';
import 'package:go_project/views/Customers.dart';
import 'package:go_project/views/OrderItems.dart';
import 'package:go_project/views/Orders.dart';
import 'package:go_project/views/medicines.dart';
import 'package:http/http.dart' as http;

class MyHomePage extends StatelessWidget {
  const MyHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Center(child: Text("Table Name")),
          backgroundColor: const Color(0xff46322b),
        ),
        body: Column(
          children: [
            Category(
                color: const Color(0xfff4eedd),
                text: "Customers",
                ontap: () async {
                  final url = Uri.parse('http://$ip:$port/global');
                  final data = {'data': 'Customers'};

                  // Encode the data as JSON
                  String jsonData = jsonEncode(data);
                  try {
                    // var response = await http.post(url,
                    //     body: {'data': "Customers"}); // تعديل البيانات
                    var response = await http.post(url,
                        headers: <String, String>{
                          'Content-Type': 'application/json; charset=UTF-8',
                        },
                        body: jsonData);
                    print('Response status: ${response.statusCode}');
                    print('Response body: ${response.body}');
                  } catch (e) {
                    print('حدث خطأ أثناء إرسال البيانات: $e');
                  } // تعديل عنوان الخادم

                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) {
                        return Customers();
                      },
                    ),
                  );
                }),
            Category(
                color: const Color(0xfff4eedd),
                text: "Medicines",
                ontap: () async {
                  final url = Uri.parse('http://$ip:$port/global');
                  final data = {'data': 'medicines'};

                  // Encode the data as JSON
                  String jsonData = jsonEncode(data);
                  try {
                    // var response = await http.post(url,
                    //     body: {'data': "Customers"}); // تعديل البيانات
                    var response = await http.post(url,
                        headers: <String, String>{
                          'Content-Type': 'application/json; charset=UTF-8',
                        },
                        body: jsonData);
                    print('Response status: ${response.statusCode}');
                    print('Response body: ${response.body}');
                  } catch (e) {
                    print('حدث خطأ أثناء إرسال البيانات: $e');
                  } // تعديل عنوان الخادم

                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) {
                        return Medicines();
                      },
                    ),
                  );
                }),
            Category(
                color: const Color(0xfff4eedd),
                text: "Orders",
                ontap: () async {
                  final url = Uri.parse('http://$ip:$port/global');
                  final data = {'data': 'orders'};

                  // Encode the data as JSON
                  String jsonData = jsonEncode(data);
                  try {
                    // var response = await http.post(url,
                    //     body: {'data': "Customers"}); // تعديل البيانات
                    var response = await http.post(url,
                        headers: <String, String>{
                          'Content-Type': 'application/json; charset=UTF-8',
                        },
                        body: jsonData);
                    print('Response status: ${response.statusCode}');
                    print('Response body: ${response.body}');
                  } catch (e) {
                    print('حدث خطأ أثناء إرسال البيانات: $e');
                  } // تعديل عنوان الخادم

                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) {
                        return Orders();
                      },
                    ),
                  );
                }),
            Category(
                color: const Color(0xfff4eedd),
                text: "Order Items",
                ontap: () async {
                  final url = Uri.parse('http://$ip:$port/global');
                  final data = {'data': 'orderitems'};

                  // Encode the data as JSON
                  String jsonData = jsonEncode(data);
                  try {
                    // var response = await http.post(url,
                    //     body: {'data': "Customers"}); // تعديل البيانات
                    var response = await http.post(url,
                        headers: <String, String>{
                          'Content-Type': 'application/json; charset=UTF-8',
                        },
                        body: jsonData);
                    print('Response status: ${response.statusCode}');
                    print('Response body: ${response.body}');
                  } catch (e) {
                    print('حدث خطأ أثناء إرسال البيانات: $e');
                  } // تعديل عنوان الخادم

                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) {
                        return OrderItems();
                      },
                    ),
                  );
                }),
          ],
        ));
  }
}
