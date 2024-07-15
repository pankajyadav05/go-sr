<div align="center">
	<a target="_blank" href="https://gitforcetalent.com">
        <picture>
            <source media="(prefers-color-scheme: dark)" srcset="https://gitforcetalent.com/_next/image?url=%2Fimages%2Flogo-light.png&w=1920&q=75">
            <source media="(prefers-color-scheme: light)" srcset="https://gitforcetalent.com/_next/image?url=%2Fimages%2Flogo.png&w=1920&q=75">
            <img alt="https://gitforcetalent.com" src="https://gitforcetalent.com/_next/image?url=%2Fimages%2Flogo.png">
        </picture>
	</a>
    <br />
    <br />
</div>

---

---

# Assignment: Golang Order Intake & Processing API

The API should be able to handle trade orders. A trade order should contain the following information (Define the structs to represent the data structures before starting):

1. `client_id`: The ID of the client who is placing the order.
2. `asset_id`: The ID of the stock being traded.
3. `volume`: The number of units shares to be traded.

When a trade order is received, the API should:

1. Verify that the client exists and has a valid account balance.
2. Verify that the asset exists in the asset catalog.
3. Check if the client has enough funds to complete the purchase.
4. Deduct the total cost from the client's account balance.
5. Update the owned assets table to reflect the new asset ownership.
6. Update the total available count of the asset in the asset catalog.

## Dummy JSON Payload:

```json
{
  "client_id": 1,
  "asset_id": 2,
  "volume": 25
}
```

In this example payload:

- The client with ID 1 (John Doe) is placing an order to buy 25 units (shares) of the asset with ID 2 (Tesla Inc.).
- The API should verify that John Doe exists, has enough funds in his account balance, and then proceed with the purchase.
- After the purchase, John Doe's account balance should be updated, and the owned assets table should reflect that he now owns 75 shares of Tesla Inc. (assuming he previously owned 50 shares).
- The total available count of Tesla Inc. in the asset catalog should be reduced by 25.
