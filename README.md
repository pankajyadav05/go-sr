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

# Assignment: Golang Distributed Real-time Trading System

Create a system in Golang that can handle a lot of trade orders at once, coming from different places, and process them quickly.

## Requirements:

### You can use a JSON object(s) for DB without needing any real DB.

1. **Order Intake**: Implement a component that can receive trade orders from different sources (e.g., trading platforms, brokers, mobile apps). This component should be able to handle multiple orders coming in simultaneously.

2. **Client Profiles**: Implement a component that stores and manages profiles for the institution's clients. These profiles should contain information like the client's name, contact details, account balance, and trading preferences.

3. **Asset Catalog**: Implement a component that stores and manages the catalog of assets (e.g., stocks, bonds, currencies) that the institution trades. This catalog should contain details about each asset, such as name, description, current price, and trading volume.

4. **Order Processing**: Implement a component that takes incoming trade orders and processes them. This involves checking the client's profile, verifying the assets being traded, updating the client's account balance, and calculating any fees or charges.

5. **Distributed Architecture**: Design and implement a distributed architecture for processing trade orders. Split the order processing workload across multiple nodes (servers) based on a suitable strategy (e.g., asset classes, client groups, or geographical regions).

## Deliverables:

1. A README file explaining the design and implementation of your distributed real-time trading system, including any assumptions and limitations.
2. Source code for all components, including order intake, client profiles, asset catalog, order processing, and distributed architecture.
3. Instructions on how to set up and run your solution, including any required dependencies or configurations.
4. Example usage scenarios demonstrating the key features of your real-time trading system.

## Evaluation Criteria:

Check full details about Evaluation criteria here [marks.md](MARKS.md)
