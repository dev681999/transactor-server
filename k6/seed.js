import http from "k6/http";
import faker from "https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js";
import { expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";

export const options = {
  iterations: 100,
};

const operationTypes = [
  {
    id: 1,
    is_debit: true,
  },
  {
    id: 2,
    is_debit: true,
  },
  {
    id: 3,
    is_debit: false,
  },
  {
    id: 4,
    is_debit: false,
  },
];

export default function () {
  const body = JSON.stringify({
    name: faker.name.findName(),
    document_number: faker.finance.account(),
  });
  const response = http.post("http://localhost:8080/api/v1/accounts", body, {
    headers: {
      Authorization: __ENV.API_KEY,
      "Content-Type": "application/json",
    },
  });

  expect(response.status, "response status").to.equal(201);

  const accountId = response.json().id;

  console.log("account created id: " + accountId);

  for (let i = 1; i <= 10; i++) {
    const operationType = operationTypes[faker.random.number(3)];
    const amount =
      (faker.finance.amount(1000, 9999) / 100) *
      (operationType.is_debit ? -1 : 1);
    const operationTypeId = operationType.id;

    const body = JSON.stringify({
      account_id: accountId,
      amount: amount,
      operation_type_id: operationTypeId,
    });
    const response = http.post(
      "http://localhost:8080/api/v1/transactions",
      body,
      {
        headers: {
          Authorization: "strongapikey",
          "Content-Type": "application/json",
        },
      }
    );

    expect(response.status, "response status").to.equal(201);

    const transactionId = response.json().id;
    console.log("transaction created id: " + transactionId);
  }
}
