package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Data struct for JSON decoding
type Data struct {
	Field1 string `json:"data"`
	// Field2 string `json:"field2"`
	// Field3 string `json:"field3"
}

type CustomerData struct {
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
	// Field5 string json:"field5"
}
type MemedicineData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
}
type OrderItemData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
}

var table Data
var tablename string

func main() {
	http.HandleFunc("/table", recetablename)
	fmt.Println("العبد يعمل على المنفذ 9090...")
	http.ListenAndServe("192.168.137.253:9090", nil)

}

func receiveDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helo")
	// قراءة البيانات من الجسم ك JSON

	// طباعة البيانات

	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/pharmacy_db") // استبدل "user" و "password" و "database" بمعلومات الاتصال الفعلية
	if err != nil {
		http.Error(w, "Data Not Receive successfuly", http.StatusInternalServerError)
		return
	}
	fmt.Println(tablename)

	switch tablename {
	case "Customers":

		var cus CustomerData
		err := json.NewDecoder(r.Body).Decode(&cus)
		if err != nil {
			http.Error(w, "فشل في قراءة البيانات", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Println(cus)
		idInt, _ := strconv.Atoi(cus.Field1)
		name := cus.Field2
		email := cus.Field3
		add := cus.Field4
		phone := cus.Field5
		TypeOperation := strings.ToUpper(cus.Field6)
		switch TypeOperation {

		case "INSERT":
			_, err = db.Exec("INSERT INTO customers(customer_id,name,email,address,phone) VALUES(?,?,?,?,?)", idInt, name, email, add, phone)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
				return

			}

		case "DELETE":
			_, err = db.Exec("DELETE FROM customers WHERE customer_id=?", idInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Delete data From the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}

		case "UPDATE":
			_, err = db.Exec("UPDATE customers SET name=?, email=?, address=?, phone=? WHERE customer_id=?", name, email, add, phone, idInt)
			if err != nil {
				fmt.Println("Error inserting data into database:", err)
				http.Error(w, "Failed to Update data into the database"+"Verify Your ID", http.StatusInternalServerError)
				return
			}
			defer db.Close()

		}
	case "orders":
		var ord OrderData
		err := json.NewDecoder(r.Body).Decode(&ord)
		if err != nil {
			http.Error(w, "فشل في قراءة البيانات", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Println(ord)
		idOrderInt, _ := strconv.Atoi(ord.Field1)
		idCustomerInt, _ := strconv.Atoi(ord.Field2)
		// name := orderdata.Field2
		total_amount_string := ord.Field3
		total_amount, _ := strconv.ParseFloat(total_amount_string, 64)
		// pricestring := orderdata.Field4
		// price, _ := strconv.ParseFloat(pricestring, 64)
		// quantity_available, _ := strconv.Atoi(orderdata.Field5)
		TypeOperation := strings.ToUpper(ord.Field4)
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
		}
	case "medicines":
		var medicinesData MemedicineData
		err := json.NewDecoder(r.Body).Decode(&medicinesData)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Println(medicinesData)
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
		}
	case "orderitems":
		var orderitemdata OrderItemData
		err := json.NewDecoder(r.Body).Decode(&orderitemdata)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Println(orderitemdata)
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
		}

		defer db.Close()

	}
	// إرسال استجابة
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "تم استلام البيانات بنجاح")
}
func recetablename(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&table)
	if err != nil {
		http.Error(w, "فشل في قراءة البيانات", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	tablename = table.Field1
	fmt.Println(tablename)
	if strings.Contains(tablename, "Customers") {
		http.HandleFunc("/receiveData", receiveDataHandler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "Customer")
	} else if strings.Contains(tablename, "medicines") {
		http.HandleFunc("/receiveData", receiveDataHandler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "medicines")
	} else if strings.Contains(tablename, "orders") {
		http.HandleFunc("/receiveData", receiveDataHandler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "Orders")
	} else if strings.Contains(tablename, "orderitems") {
		http.HandleFunc("/receiveData", receiveDataHandler)
		fmt.Fprintf(w, "Data received contains the keyword: %s", "orderitems")
	}

}
