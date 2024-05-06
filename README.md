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

```diff
You can use a JSON object(s) for DB without needing any real DB
```


1. **Order Intake & Processing**:

   - [dummy data] Store clients profiles in DB. Profiles should contain information like the client's name, contact details, account balance, owned asstes etc.
   - [dummy data] Store catalog of assets (e.g., stocks, bonds, currencies) in DB. This catalog should contain details about each asset, such as name, description, current price, total available count etc.
   - API that can receive trade orders and processes them (one trade order should include clienId, Asset being traded, Volume etc.). This involves checking the client's profile, verifying the assets being traded, updating the client's account balance, upating asset's available count and other edge cases.

2. **Distributed Architecture**: Design and implement a distributed architecture for processing trade orders. Split the order processing workload across multiple nodes (servers) based on a suitable strategy (e.g., asset classes, client groups, or geographical regions).

3. **Documentation**: A README file explaining the design and implementation of your distributed real-time trading system, including any assumptions and limitations.

## Submission:

Once you've completed the assignment, please submit your work by providing:

1. Provide the GitHub repository URL as a PRIVATE Git repository (also provide access) to pankaj@gitforcetalent.com and cc prithvi@gitforcetalent.com, harika@gitforcetalent.com.
2. Source code for all components, including order intake, client profiles, asset catalog, order processing, and distributed architecture.
3. Instructions on how to set up and run your solution, including any required dependencies or configurations.
4. Provide any additional notes or considerations regarding the implementation.

## Evaluation Criteria:

Check full details about Evaluation criteria here [marks.md](MARKS.md)
