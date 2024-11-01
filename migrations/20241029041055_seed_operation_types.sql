-- Add the initial operation types to the database.
INSERT INTO
    operation_types (
        id,
        description,
        is_debit,
        create_time,
        update_time
    )
VALUES
    (1, 'Normal Purchase', true, now(), now()),
    (
        2,
        'Purchase with installments',
        true,
        now(),
        now()
    ),
    (3, 'Withdrawal', false, now(), now()),
    (4, 'Credit Voucher', false, now(), now());