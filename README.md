### Http microservice for retrieving mortgage calculations

Endpoint: https://localhost:8080
Methd: GET 
Params: 

| name               | type    | required |
| ------------------ | ------- | -------- |
| capital            | float64 | true     |
| terms              | int     | true     |
| interestType       | float64 | true     |
| amortizationAmount | float64 | false    |
| year               | int     | false    |
| month              | int     | false    |

Response: 
```JSON
{
    "interestSavingsForPrice": 12189.52,
    "monthlyPrice": 194.22,
    "pendingPayments": 302,
    "timeSavingsMonth": 10,
    "timeSavingsYear": 14,
    "totalTimeInterest": 22542.83
}
```