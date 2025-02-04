# Number Classification

This API classifies numbers by determining whether they are prime, perfect, Armstrong, calculating the digit sum, and fetching a fun fact.

## Endpoint 

**GET** `/api/classify-number?number=371`

### Example Successful Response (200 OK)

```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": [
    "armstrong",
    "odd"
    ],
    "digit_sum": 11,
    "fun_fact": "371 is a boring number."
}

{
  "number": "alphabet",
  "error": true
}
