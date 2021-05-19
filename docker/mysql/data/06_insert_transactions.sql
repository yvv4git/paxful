-- use paxful-test;

INSERT INTO transactions(idempotence_key, expired, attempt) VALUE(LEFT(UUID(), 50), NOW() + INTERVAL 1 DAY, 0);