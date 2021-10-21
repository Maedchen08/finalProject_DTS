# final project DTS

## Antar Jemput BRI Link

_kelompok C_

1. Rahma Asdarina Badegeil
2. Mochamad Sofi Maulana
3. Muhammad Zikri
4. Eka Nur Rahmawati

================================================================================================

## setup

1. clone this repository
2. type `make run` to run the apps
3. Hit the Server `localhost:80/transaction`
4. Unit Test `make coverage`
5. see coverage all test in html `make coverage-out`

### Tasks

We define routes for handling operations:

| Method | Route                  | Action                                |
| ------ | ---------------------- | ------------------------------------- |
| POST   | /register              | registration account                  |
| POST   | /login                 | create token JWT                      |
| POST   | /agents                | get add agent                         |
| GET    | /agents                | get all agents                        |
| GET    | /agents/1              | get agents by id                      |
| POST   | /rating/transaksi      | give rating after transaction         |
| GET    | /rating/agent/2        | view rating agent where id = 2        |
| GET    | /cariagent             | search agent whent status 202 and 200 |
| POST   | /customer              | add customer                          |
| GET    | /customer              | get all customer                      |
| GET    | /customer/1            | get customer when id = 1              |
| POST   | /st                    | add service transaction               |
| GET    | /st                    | get all service transaction           |
| GET    | /st/1                  | get service by id                     |
| POST   | /ts                    | add type service                      |
| GET    | /ts                    | get all type service transaction      |
| GET    | /ts/1                  | get type service by id                |
| POST   | /transaksi/create      | add new transaction                   |
| POST   | /transaksi/dikofirmasi | change status to confirm              |
| POST   | /transaksi/dibatalkan  | change status to reject               |
| POST   | /transaksi/selesai     | change status to done                 |
| GET    | /transaksi             | get all transaction                   |
| GET    | /transaksi/1           | get transaction by id                 |
| GET    | /transaksi/customer/2  | get all transaction customer id = 2   |
| GET    | /transaksi/agent/2     | get all transaksi where id = 2        |
| DELETE | /transaksi/1           | delete transaction where id = 1       |
| GET    | /users                 | get all users                         |

Access API via `http://localhost:80/{route}`

1. POST `/register`

Request:

```
{
    "role": 2,
    "customer_id": 2,
    "agent_id": 2,
    "username": "maulana",
    "password": "maulana"
}
```

Response:
status code: 200

```
{
       "data": {
        "ID": 0,
        "CreatedAt": "2021-10-21T12:40:57.98+07:00",
        "UpdatedAt": "2021-10-21T12:40:57.98+07:00",
        "DeletedAt": null,
        "id": 37,
        "role": 2,
        "customer_id": 2,
        "agent_id": 2,
        "username": "maulana",
        "password": "$2a$14$bdzA66IKLom/Nk4fhvfFp.diYjCGWP1LjMFZ3JnHxYFlZlZ3qSFK6"
    },
    "message": "Success"
}
```

2. POST `/login`

Request:

```
{
    "username": "maulana",
    "password": "maulana"
}
```

Response:
status code: 200

```
{
    "status": 200,
    "message": "Login Succeeded",
    "Id": 37,
    "role": 2,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzUwNTQzMzUsImxvZ2luX2FzIjoyLCJ1c2VybmFtZSI6Im1hdWxhbmEifQ.IlQoBAZAOUyIRDULdXZtRS4iIS4lrzbJb_j9ELv-2oU",
    "username": "maulana"
}
```

3. POST `/agents`

Request:

```
{

    "agent_name": "Arifin",
    "alamat_agen": "jl Margitosan",
    "agent_province_id": 35,
    "agent_regency_id": 3516,
    "agent_district_id": 3516140,
    "no_wa": "087483920123",
    "longitude_agent": 7,
    "latitude_agent": 9.9

}
```

Response
Status Code : 200

```
{
    "message": "Success",
    "status": 201
    "data": {
        "ID": 0,
        "CreatedAt": "2021-10-21T11:26:12.284+07:00",
        "UpdatedAt": "2021-10-21T11:26:12.284+07:00",
        "DeletedAt": null,
        "id": 12,
        "agent_name": "Arifin",
        "alamat_agen": "jl Margitosan",
        "agent_province_id": 35,
        "agent_regency_id": 3516,
        "agent_district_id": 3516140,
        "no_wa": "087483920123",
        "longitude_agent": 7,
        "latitude_agent": 9.9
    },
}
```

4. GET `/agents`

Response
Status code: 200

```
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 1,
            "agent_name": "bambang",
            "alamat_agen": "Jl. Raya Soreang No. 300",
            "agent_province_id": 32,
            "agent_regency_id": 3204,
            "agent_district_id": 3204190,
            "no_wa": "089525315860",
            "longitude_agent": 7.8,
            "latitude_agent": 9.8
        },

        {
            "ID": 0,
            "CreatedAt": "2021-10-21T04:26:12.284Z",
            "UpdatedAt": "2021-10-21T04:26:12.284Z",
            "DeletedAt": null,
            "id": 12,
            "agent_name": "Arifin",
            "alamat_agen": "jl Margitosan",
            "agent_province_id": 35,
            "agent_regency_id": 3516,
            "agent_district_id": 3516140,
            "no_wa": "087483920123",
            "longitude_agent": 7,
            "latitude_agent": 9.9
        }
    ],
    "message": "success",
    "status": 200
}
```

5. GET `/agent/1`

Response
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 1,
        "agent_name": "bambang",
        "alamat_agen": "Jl. Raya Soreang No. 300",
        "agent_province_id": 32,
        "agent_regency_id": 3204,
        "agent_district_id": 3204190,
        "no_wa": "089525315860",
        "longitude_agent": 7.8,
        "latitude_agent": 9.8
    },
}
```

6. POST `/rating/transaksi`

Request:

```
{
    "id":1,
    "rating":5,
    "rating_comment":"final done"
}
```

Response:
Status code: 200

```
{
    "status": 200
    "message": "data has been update",
    "result": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-21T04:39:36.01Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_id": 1,
        "customers_id": 1,
        "agents_id": 3,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 100000,
        "status_transaction": 3,
        "rating": 5,
        "rating_comment": "final done",
        "longitude_cust": 7.9,
        "latitude_cust": 5.7
    },
}
```

7. GET `/rating/agent/2`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "agents_id": 2,
        "total": 2,
        "avg_rating": 4.5
    },
}
```

8. POST `/cariagent`

Request:

```
{
    "type_transaction_id": 1,
    "customers_id": 1,
    "address": "Jl. Soreang No.181",
    "transaction_prov_id": 32,
    "transaction_regency_id": 3204,
    "transaction_district_id": 3204191,
    "amount": 100000,
    "status_transaction": 2,
    "longitude_agent": 7,
    "latitude_agent": 9.9
}
```

Response:
Status code: 202

```
{
    "msg": "Data agen tidak ditemukan",
    "status": 202
}
```

9. POST `/customer`

Request:

```
{
    "customer_name":"Nirmala",
    "no_wa":"087692883910"
}
```

Response:
Status code: 200

```
{
    "message": "Success",
    "status": 200
    "data": {
        "name": "Nirmala",
        "no wa": "087692883910"
    },
}
```

10. GET `/customer`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 1,
            "customer_name": "rahma",
            "no_wa": "084736273849"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 4,
            "customer_name": "ekanur",
            "no_wa": "0893726361812"
        },
    ],
}
```

11. GET `/customer/1`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "name": "rahma",
        "no wa": "084736273849"
    },
}
```

12. POST `/st`

Request:

```
{
    "id":3,
    "service_name":"Mini Atm Bri"
}
```

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 3,
        "service_name": "Mini ATM BRI"
    },
}
```

13. GET `/st`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 1,
            "service_name": "Laku Pandai"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 2,
            "service_name": "Tunai"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 3,
            "service_name": "Mini ATM BRI"
        }
    ],
}
```

14. GET `/st/1`

Response:
Status code: 200

```
{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 1,
        "service_name": "Laku Pandai"
    },
    "message": "success",
    "status": 200
}
```

15. POST `/ts`

Request:

```
{
    "id":15,
    "type_transaction_name":"Setor-Pasti",
    "service_transaction_id":3
}
```

Response:
Status code: 200

```
{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 15,
        "type_transaction_name": "Setor-Pasti",
        "service_transaction_id": 3
    },
    "message": "success",
    "status": 200
}
```

16. GET `/ts`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data":    [
    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_name": "Cash-in & Out",
        "service_transaction_id": 1
    },

    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 4,
        "type_transaction_name": "Tarik Tunai",
        "service_transaction_id": 1
    },

]
}
```

17. GET `/ts/1`

Response:
Status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_name": "Cash-in & Out",
        "service_transaction_id": 1
    },
}
```

18. POST `/transaksi/create`

Request:

```
{
    "type_transaction_id": 1,
    "customers_id": 1,
    "agents_id": 3,
    "address": "Jl. Soreang No.181",
    "transaction_prov_id": 32,
    "transaction_regency_id": 3204,
    "transaction_district_id": 3204190,
    "amount": 100000,
    "status_transaction": 3,
    "longitude_cust": 7.9,
    "latitude_cust": 5.7
}
```

Response:
status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-13T08:59:15.967Z",
        "DeletedAt": null,
        "id": 2,
        "type_transaction_id": 1,
        "customers_id": 2,
        "agents_id": 2,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 200000,
        "status_transaction": 3,
        "longitude_cust": 7.9,
        "latitude_cust": 5.7
      }
}
```

19. POST `/transaksi/dikonfirmasi`

Request:

```
{
    "id":1
}
```

Response:
Status code: 200

```
{
    "status": 200
    "message": "data has been update",
    "result": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-21T06:50:03.506Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_id": 1,
        "customers_id": 1,
        "agents_id": 3,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 100000,
        "status_transaction": 1,
        "longitude_cust": 7.9,
        "latitude_cust": 5.7
    },
}
```

20. POST `/transaksi/dibatalkan`

Request:

```
{
    "id":13
}
```

Response:
status code: 200

```
{
    "status": 200
    "message": "data has been update",
    "result": {
        "ID": 0,
        "CreatedAt": "2021-10-09T14:41:55.851Z",
        "UpdatedAt": "2021-10-21T06:54:16.591Z",
        "DeletedAt": null,
        "id": 13,
        "type_transaction_id": 11,
        "customers_id": 1,
        "agents_id": 1,
        "address": "Jl. Mulyorejo utara gang 3 no 7",
        "transaction_prov_id": 35,
        "transaction_regency_id": 3578,
        "transaction_district_id": 3578090,
        "amount": 9000,
        "status_transaction": 2,
        "rating": 4,
        "longitude_cust": 0,
        "latitude_cust": 0
    },
}
```

21. POST `/transaksi/selesai`

Request:

```
{
    "id":12
}
```

Response:
status code: 200

```
{
    "status": 200,
    "message": "data has been update",
    "time": "Thu Oct 21 13:57:02 2021"
    "result": {
        "ID": 0,
        "CreatedAt": "2021-10-09T14:37:08.965Z",
        "UpdatedAt": "2021-10-21T06:57:02.42Z",
        "DeletedAt": null,
        "id": 12,
        "type_transaction_id": 10,
        "customers_id": 3,
        "agents_id": 3,
        "address": "Jl. Mulyorejo utara gang 3 no 6",
        "transaction_prov_id": 35,
        "transaction_regency_id": 3578,
        "transaction_district_id": 3578090,
        "amount": 9000,
        "status_transaction": 3,
        "rating": 4,
        "rating_comment": "",
        "longitude_cust": 0,
        "latitude_cust": 0
    },
}
```

22. GET `/transaksi`

Response:
status code : 200

```
{
    "status": 200,
    "message": "data has been update",
    "result": [
        [
    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-21T06:50:03.506Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_id": 1,
        "customers_id": 1,
        "agents_id": 3,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 100000,
        "status_transaction": 3,
        "rating": 5,
        "rating_comment": "final done",
        "longitude_cust": 7.9,
        "latitude_cust": 5.7
    },
    {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-13T08:59:15.967Z",
        "DeletedAt": null,
        "id": 2,
        "type_transaction_id": 1,
        "customers_id": 2,
        "agents_id": 2,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 200000,
        "status_transaction": 3,
        "rating": 4,
        "rating_comment": "final done",
        "longitude_cust": 0,
        "latitude_cust": 0
    },

]

}
```

23. GET `/transaksi/1`

Response:
status code: 200

```
{
    "message": "success",
    "status": 200
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-10-21T06:50:03.506Z",
        "DeletedAt": null,
        "id": 1,
        "type_transaction_id": 1,
        "customers_id": 1,
        "agents_id": 3,
        "address": "Jl. Soreang No.181",
        "transaction_prov_id": 32,
        "transaction_regency_id": 3204,
        "transaction_district_id": 3204190,
        "amount": 100000,
        "status_transaction": 1,
        "rating": 5,
        "rating_comment": "final done",
        "longitude_cust": 7.9,
        "latitude_cust": 5.7
    },
}
```

24. GET `/transaksi/customer/2`

Response :
status code : 200

```
{
    "message": "success",
    "status": 200
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "2021-10-13T08:59:15.967Z",
            "DeletedAt": null,
            "id": 2,
            "type_transaction_id": 1,
            "customers_id": 2,
            "agents_id": 2,
            "address": "Jl. Soreang No.181",
            "transaction_prov_id": 32,
            "transaction_regency_id": 3204,
            "transaction_district_id": 3204190,
            "amount": 200000,
            "status_transaction": 3,
            "rating": 4,
            "rating_comment": "final done",
            "longitude_cust": 0,
            "latitude_cust": 0
        },
        {
            "ID": 0,
            "CreatedAt": "2021-10-09T11:31:58.115Z",
            "UpdatedAt": "2021-10-09T14:26:17.344Z",
            "DeletedAt": null,
            "id": 3,
            "type_transaction_id": 8,
            "customers_id": 2,
            "agents_id": 3,
            "address": "Jl. H. Dahlan no 75",
            "transaction_prov_id": 31,
            "transaction_regency_id": 3171,
            "transaction_district_id": 3171020,
            "amount": 8900000,
            "status_transaction": 2,
            "rating": 4,
            "rating_comment": "",
            "longitude_cust": 0,
            "latitude_cust": 0
        }
    ],
}
```

25. GET `/transaksi/agent/2`

Response :
status code: 200

```
{
    "message": "success",
    "status": 200
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "2021-10-13T08:59:15.967Z",
            "DeletedAt": null,
            "id": 2,
            "type_transaction_id": 1,
            "customers_id": 2,
            "agents_id": 2,
            "address": "Jl. Soreang No.181",
            "transaction_prov_id": 32,
            "transaction_regency_id": 3204,
            "transaction_district_id": 3204190,
            "amount": 200000,
            "status_transaction": 3,
            "rating": 4,
            "rating_comment": "final done",
            "longitude_cust": 0,
            "latitude_cust": 0
        },
        {
            "ID": 0,
            "CreatedAt": "2021-10-20T01:46:57.261Z",
            "UpdatedAt": "2021-10-21T04:35:21.862Z",
            "DeletedAt": null,
            "id": 16,
            "type_transaction_id": 10,
            "customers_id": 5,
            "agents_id": 2,
            "address": "Jl. Soreang No.181",
            "transaction_prov_id": 32,
            "transaction_regency_id": 3204,
            "transaction_district_id": 3204190,
            "amount": 200000,
            "status_transaction": 3,
            "rating": 5,
            "rating_comment": "final done",
            "longitude_cust": 0,
            "latitude_cust": 0
        }
    ],
}
```

26. DELETE `/transaksi/1`

Response:
status code : 200

```
{
    "message": "success delete transaction",
    "status": 200
    "data": {
        "id_transaction": 1
    },
}
```

27. GET `/users`
    Response:
    status code : 200

```
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2021-10-20T02:36:40.308Z",
            "UpdatedAt": "2021-10-20T02:36:40.308Z",
            "DeletedAt": null,
            "id": 31,
            "role": 2,
            "customer_id": 6,
            "agent_id": 0,
            "username": "wulan",
            "password": "$2a$14$WQ5I8xDhOe7DxRh1fy0L1.qfGFi17eCgaqNUUpu4MIWdkUSw7z1PG"
        },

        {
            "ID": 0,
            "CreatedAt": "2021-10-21T05:40:57.98Z",
            "UpdatedAt": "2021-10-21T05:40:57.98Z",
            "DeletedAt": null,
            "id": 37,
            "role": 2,
            "customer_id": 2,
            "agent_id": 2,
            "username": "maulana",
            "password": "$2a$14$bdzA66IKLom/Nk4fhvfFp.diYjCGWP1LjMFZ3JnHxYFlZlZ3qSFK6"
        }
    ],
    "message": "success",
    "status": 200
}
```

[golang]: https://golang.org/
[fiber]: https://github.com/gofiber/fiber/
[gomock]: https://github.com/golang/mock/
[jsonwebtoken]: https://jwt.io/
