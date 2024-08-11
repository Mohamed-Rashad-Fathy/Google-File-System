package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

// Data struct for JSON decoding
type CustomerData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
	// وهكذا لبقية الحقول
}
type MemedicineData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
}
type OrderData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	// Field5 string `json:"field5"`
}
type OrderItemData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
	// Field7 string `json:"field7"`
	// Field5 string `json:"field5"`
}
type tabelName struct {
	Name string `json:"data"`
}

// master
var my_ip string = "192.168.36.149"
var my_port string = "8080"

// var slave1_ip="192.168.45.158"
var slave2_ip = "192.168.36.139"
var slaves_ports = "9090"

func main() {

	http.HandleFunc("/global", gloabelhande)
	// ابدأ الخادم HTTP
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(my_ip+":"+my_port, nil)
}

// type tabelName struct {
// 	Name string `json:"data"`
// }

var requestData tabelName
var receivedData string
var table string

func gloabelhande(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = SendToSlaveTable(requestData)
	if err != nil {
		http.Error(w, "Field to Send Data to slave ", http.StatusInternalServerError)
		return
	}

	// Extract the "data" field from the received data
	receivedData = requestData.Name

	// Do whatever processing you need with the received data
	fmt.Println("Received data:", receivedData)
	if strings.Contains(receivedData, "Customers") {
		// Register the "/customers" handler
		http.HandleFunc("/customers", handler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "Customers")
	} else if strings.Contains(receivedData, "medicines") {
		http.HandleFunc("/customers", handler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "Medicines")

	} else if strings.Contains(receivedData, "order") {
		http.HandleFunc("/customers", handler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "Order")

	} else if strings.Contains(receivedData, "orderitems") {
		http.HandleFunc("/customers", handler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "orderitems")

	} else {
		// If the data doesn't contain the keyword "Customers", handle it differently
		// For example, you can register another handler or directly process the request here
		// ...
		fmt.Fprintf(w, "Data received does not contain the keyword: %s", "Costomer OR medicines OR Order OR OrderItem")
	}

	// Send a response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Received data successfully"}`)
	// ارسل استجابة إلى الطلب
}

func handler(w http.ResponseWriter, r *http.Request) {
	// قراءة البيانات من الجسم ك JSON

	table = requestData.Name
	fmt.Println(table)

	// الاتصال بقاعدة البيانات وإدخال البيانات

	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/pharmacy_db")
	if err != nil {
		http.Error(w, "Data Not Receive successfuly", http.StatusInternalServerError)
		return
	}

	switch table {
	case "Customers":
		var data CustomerData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		err = SendToSlaveData(data)
		if err != nil {
			http.Error(w, "Field to Send Data to slave ", http.StatusInternalServerError)
			return
		}
		// طباعة البيانات
		fmt.Println("البيانات المستقبلة:", data)
		// ageInt, _ := strconv.Atoi(data.Field2)

		idInt, _ := strconv.Atoi(data.Field1)
		name := data.Field2
		email := data.Field3
		add := data.Field4
		phone := data.Field5
		TypeOperation := strings.ToUpper(data.Field6)
		switch TypeOperation {
		case "INSERT":
			_, err = db.Exec("INSERT INTO customers(customer_id,name,email,address,phone) VALUES(?,?,?,?,?)", idInt, name, email, add, phone)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
				return
			}
			defer db.Close()

		case "DELETE":
			_, err = db.Exec("DELETE FROM customers WHERE customer_id=?", idInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Delete data From the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		case "UPDATE":
			_, err = db.Exec("UPDATE customers SET name=?, email=?, address=?, phone=? WHERE customer_id=?", name, email, add, phone, idInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Update data into the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		default:
			fmt.Println("There is NO Query Called ", TypeOperation, "Plese Enter DELETE OR INSERT OR UPDATE")
			http.Error(w, "There is NO Query Called "+TypeOperation+"Plese Enter DELETE OR INSERT OR UPDATE", http.StatusInternalServerError)

		}
	case "medicines":
		var medicinesData MemedicineData
		err := json.NewDecoder(r.Body).Decode(&medicinesData)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		err = SendToSlaveData(medicinesData)
		if err != nil {
			http.Error(w, "Field to Send Data to slave ", http.StatusInternalServerError)
			return
		}
		// طباعة البيانات
		fmt.Println("Received Data for Medicines Data :", medicinesData)

		idInt, _ := strconv.Atoi(medicinesData.Field1)
		name := medicinesData.Field2
		description := medicinesData.Field3
		pricestring := medicinesData.Field4
		price, _ := strconv.ParseFloat(pricestring, 64)
		quantity_available, _ := strconv.Atoi(medicinesData.Field5)
		TypeOperation := strings.ToUpper(medicinesData.Field6)
		switch TypeOperation {
		case "INSERT":
			_, err = db.Exec("INSERT INTO medicines(medicine_id,name,description,price,quantity_available) VALUES(?,?,?,?,?)", idInt, name, description, price, quantity_available)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
				return
			}
			defer db.Close()

		case "DELETE":
			_, err = db.Exec("DELETE FROM medicines WHERE medicine_id=?", idInt)
			if err != nil {
				fmt.Println("Error Delete data From database:", err)
				http.Error(w, "Failed to Delete data From the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		case "UPDATE":
			_, err = db.Exec("UPDATE customers SET name=?, description=?, price=?, quantity_available=? WHERE medicine_id=?", name, description, price, quantity_available, idInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Update data into the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		default:
			fmt.Println("There is NO Query Called ", TypeOperation, "Plese Enter DELETE OR INSERT OR UPDATE")
			http.Error(w, "There is NO Query Called "+TypeOperation+"Plese Enter DELETE OR INSERT OR UPDATE", http.StatusInternalServerError)
		}
	case "orders":
		var orderdata OrderData
		err := json.NewDecoder(r.Body).Decode(&orderdata)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		err = SendToSlaveData(orderdata)
		if err != nil {
			http.Error(w, "Field to Send Data to slave ", http.StatusInternalServerError)
			return
		}
		// طباعة البيانات
		fmt.Println("Data Received : ", orderdata)

		idOrderInt, _ := strconv.Atoi(orderdata.Field1)
		idCustomerInt, _ := strconv.Atoi(orderdata.Field2)
		// name := orderdata.Field2
		total_amount_string := orderdata.Field3
		total_amount, _ := strconv.ParseFloat(total_amount_string, 64)
		// pricestring := orderdata.Field4
		// price, _ := strconv.ParseFloat(pricestring, 64)
		// quantity_available, _ := strconv.Atoi(orderdata.Field5)
		TypeOperation := strings.ToUpper(orderdata.Field4)
		switch TypeOperation {
		case "INSERT":
			_, err = db.Exec("INSERT INTO orders(order_id,customer_id,total_amount) VALUES(?,?,?)", idOrderInt, idCustomerInt, total_amount)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
				return
			}
			defer db.Close()

		case "DELETE":
			_, err = db.Exec("DELETE FROM orders WHERE order_id=? AND customer_id=?", idOrderInt, idCustomerInt)
			if err != nil {
				fmt.Println("Error Delete data From database:", err)
				http.Error(w, "Failed to Delete data From the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		case "UPDATE":
			_, err = db.Exec("UPDATE orders SET total_amount=? WHERE order_id=?", total_amount, idOrderInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Update data into the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		default:
			fmt.Println("There is NO Query Called ", TypeOperation, "Plese Enter DELETE OR INSERT OR UPDATE")
			http.Error(w, "There is NO Query Called "+TypeOperation+"Plese Enter DELETE OR INSERT OR UPDATE", http.StatusInternalServerError)

		}
	case "orderitems":
		var orderitemdata OrderItemData
		err := json.NewDecoder(r.Body).Decode(&orderitemdata)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		err = SendToSlaveData(orderitemdata)
		if err != nil {
			http.Error(w, "Field to Send Data to Slaves", http.StatusInternalServerError)
			return
		}
		// طباعة البيانات
		fmt.Println("Data Received for Order Item Data :", orderitemdata)

		order_item_id, _ := strconv.Atoi(orderitemdata.Field1)
		order_id, _ := strconv.Atoi(orderitemdata.Field2)
		medicine_id, _ := strconv.Atoi(orderitemdata.Field3)
		quantity, _ := strconv.Atoi(orderitemdata.Field4)
		// name := orderdata.Field2
		unit_price_string := orderitemdata.Field5
		unit_price, _ := strconv.ParseFloat(unit_price_string, 64)
		total_price := unit_price * float64(quantity)
		TypeOperation := strings.ToUpper(orderitemdata.Field6)
		switch TypeOperation {
		case "INSERT":
			_, err = db.Exec("INSERT INTO orderitems(order_item_id,order_id,medicine_id,quantity,unit_price,total_price) VALUES(?,?,?,?,?,?)", order_item_id, order_id, medicine_id, quantity, unit_price, total_price)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
				return
			}
			defer db.Close()

		case "DELETE":
			_, err = db.Exec("DELETE FROM orderitems WHERE order_item_id=?", order_item_id)
			if err != nil {
				fmt.Println("Error Delete data From database:", err)
				http.Error(w, "Failed to Delete data From the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		case "UPDATE":
			_, err = db.Exec("UPDATE orderitems SET quantity,unit_price,total_price WHERE order_item_id=?", quantity, unit_price, total_price, order_item_id)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Update data into the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()
		default:
			fmt.Println("There is NO Query Called ", TypeOperation, "Plese Enter DELETE OR INSERT OR UPDATE")
			http.Error(w, "There is NO Query Called "+TypeOperation+"Plese Enter DELETE OR INSERT OR UPDATE", http.StatusInternalServerError)

		}
		// query := fmt.Sprintf("INSERT INTO test(name,age,id) VALUES(?,?,?)", test)

	}
	// إرسال استجابة
	fmt.Fprintf(w, "Data Receive SuccessFuly")
}
func SendToSlaveTable(requestData interface{}) error {
	// تحويل البيانات إلى JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	// إرسال البيانات إلى العبد
	// resp, err := http.Post(fmt.Sprintf("http://%s:%s/receiveData", slave1_ip,slaves_ports), "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	resp1, err1 := http.Post(fmt.Sprintf("http://%s:%s/table", slave2_ip, slaves_ports), "application/json", bytes.NewBuffer(jsonData))
	if err1 != nil {
		return err1
	}
	defer resp1.Body.Close()

	return nil
}
func SendToSlaveData(requestData interface{}) error {
	// تحويل البيانات إلى JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// إرسال البيانات إلى العبد
	// resp, err := http.Post(fmt.Sprintf("http://%s:%s/receiveData", slave1_ip,slaves_ports), "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	resp1, err1 := http.Post(fmt.Sprintf("http://%s:%s/receiveData", slave2_ip, slaves_ports), "application/json", bytes.NewBuffer(jsonData))
	if err1 != nil {
		return err1
	}
	defer resp1.Body.Close()

	return nil
}
