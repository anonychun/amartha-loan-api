# Amartha Loan API

This project was generated using [E-Corp](https://github.com/anonychun/ecorp).

Clone this repository, then run `cp .env.sample .env` to create your own `.env` file. Fill in the environment variables as needed.

To run the application with hot reload, use the command below:

```bash
./bin/dev
```

Before doing so, run the database setup command below to create the database, apply migrations, and seed it with default data:

```bash
./bin/db setup
```

## Solution Overview

### Borrower

- Create a borrower account to propose a loan: `POST /api/v1/borrower/auth/signup`
- Propose a loan: `POST /api/v1/borrower/loan` (requires borrower authentication; you are automatically authenticated after signup)

### Admin

- Sign in as an admin: `POST /api/v1/admin/auth/signin` (use the default admin credentials from the seed data)
- Approve a proposed loan: `POST /api/v1/admin/loan/:id/approve`
- After the loan is fully invested, disburse it: `POST /api/v1/admin/loan/:id/disburse`

### Investor

- Create an investor account to invest in a loan: `POST /api/v1/investor/auth/signup`
- Invest in an approved loan: `POST /api/v1/investor/loan/:id/invest` (requires investor authentication; you are automatically authenticated after signup)

## Worker

There is one scheduled job that runs every 5 minutes to check for loans that have been fully funded but have not yet triggered an investor notification. When such loans are found, the worker sends a notification email to the investor.

To run the worker, use the command below:

```bash
./bin/worker
```

## Running Tests

To run the tests, use the command below:

```bash
./bin/test
```

You can also run specific test files using:

```bash
./bin/test path/to/test/file.test.ts
```
