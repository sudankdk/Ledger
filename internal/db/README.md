Double-entry convention

This project uses a simple numeric sign convention for entries:

- Debit entries store a positive `amount` value.
- Credit entries store a negative `amount` value.

`DoTransaction(ctx, amount, desc, debitAccountID, creditAccountID)` will:

1. Create a `transactions` record with `description = desc`.
2. Create a debit `entries` row for `debitAccountID` with `amount` (positive).
3. Create a credit `entries` row for `creditAccountID` with `-amount` (negative).

Both entries are created inside a single DB transaction to ensure atomicity.
