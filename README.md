# Toko Ijah


# How to Install

- Install Dependencies

        go get ./...

- Start the app

        go run main.go

- Or simply run
        
        ./bin/toko-ijah

The server will be started at localhost port 9000

# ERD

![toko-ijah-erd](https://user-images.githubusercontent.com/16671326/60030545-5761cc80-96cd-11e9-9f7d-2cf426d9ae39.png)

# Endpoints

- /products (POST)

        Will add new product to the product table, and also a new row in `product_stock` table (initiated with 0 stock)

        payload example : 
        {
            "sku": "test-sku001",
            "name": "baju biru",
            "size": "XL",
            "color": "biru"
        }

        response example: 
        {
            "id": 1,
            "sku": "test-sku001",
            "name": "baju biru",
            "size": "XL",
            "color": "biru",
            "created_at": "2019-06-24T20:03:22.366488+07:00",
            "updated_at": "2019-06-24T20:03:22.366488+07:00",
            "deleted_at": null
        }

- /purchase (POST)

        The purchased items will add the product stock based on the product id,
        and also add a new row in `purchase_history` table

        payload example :
        {
            "number_ordered": 10,
            "number_received": 2,
            "buy_price": 2000,
            "product_id": 2
        }

        response example :
        {
            "id": 3,
            "number_ordered": 10,
            "number_received": 2,
            "buy_price": 2000,
            "total_price": 20000,
            "receipt_number": "",
            "product_id": 2,
            "created_at": "2019-06-24T20:03:36.297013+07:00",
            "updated_at": "2019-06-24T20:03:36.297013+07:00",
            "deleted_at": null
        }

- /purchase/:purchase_id (PUT)

        Will update the purchase data, but it will only add the number received, because directly
        the number of items that ordered when purchase is made is not the same as the number that received,
        so it will update the purchase data (increase the number received), and at the same time adding new
        row in `purchase_history`

        payload example : 
        {
            "number_ordered": 10,
            "number_received": 8,
            "buy_price": 2000,
            "product_id": 2
        }

        response example :
        {
            "id": 3,
            "number_ordered": 10,
            "number_received": 10,
            "buy_price": 2000,
            "total_price": 20000,
            "receipt_number": "",
            "product_id": 2,
            "created_at": "2019-06-24T20:03:36.297013+07:00",
            "updated_at": "2019-06-24T20:04:34.054655+07:00",
            "deleted_at": null
        }

- /orders (POST)

        Will create order document, and also decreasing the product stock if stock to be delivered is enough

        payload example:
        {
            "number_sold": 6,
            "sell_price": 6000,
            "product_id": 2
        }

        response example:
        {
            "id": 4,
            "number_sold": 6,
            "sell_price": 6000,
            "total_price": 36000,
            "receipt_number": "",
            "product_id": 2,
            "product_stock_id": 2,
            "created_at": "2019-06-24T20:05:02.794566+07:00",
            "updated_at": "2019-06-24T20:05:02.794566+07:00",
            "deleted_at": null
        }

# Export CSV

- /reports/product-values/export! (GET)

        Will return the exported csv of product values report download link as the response

        response example :
        {
            "download_link": "http://localhost:9000/report/downloads/product_values1561381636157388000.csv"
        }

- /reports/sales-detail/export! (GET)

        Will return the exported csv of product values report download link as the response

        response example :
        {
            "download_link": "http://localhost:9000/report/downloads/sales_detail1561381636157388000.csv"
        }

        
