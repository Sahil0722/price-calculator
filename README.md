# Price Calculator

A tool that calculates "tax included" prices for a given list of prices & tax rates

## Project Description

The application processes input data consisting of prices and corresponding tax rates, and produces an output that includes the calculated prices with tax included. It demonstrates how to work with structs, functions, methods, modules, multiple packages, and JSON file handling in Go.

## Input Data

The input data consists of two lists:
- **Prices**: A list of base prices.
- **Tax Rates**: A list of corresponding tax rates (as decimals).

**Example Input:**

*Tax Rates File (`tax_rates.txt`)*
*Prices File (`prices.txt`)*

## Output Data

The output is a JSON file for each tax rate, including the prices and the mapping of input prices to their corresponding tax-included prices.
