INSERT INTO media (id, url, mime_type, account_id)
VALUES ('7ceefef0-2a3e-4e47-a07d-0709c70d1cfa', 'https://rt-online.ru/wp-content/uploads/2021/02/viktorina.jpg', 'image/jpeg', 'd0fbc24f-5061-4d10-b92c-d386c8eba600'),
       ('44c78b42-e931-4d8c-8451-cc59383e5af6', 'https://www.mgpu.ru/wp-content/uploads/2019/02/Viktorina.jpg', 'image/jpeg', 'd0fbc24f-5061-4d10-b92c-d386c8eba600'),
       ('922d0089-ec67-47a7-ad25-dc8e36d7c93d', 'https://gameorg.ru/media/games/viktorina_22.jpg', 'image/jpeg', 'd0fbc24f-5061-4d10-b92c-d386c8eba600');

INSERT INTO package (name, account_id, is_published, language_id, cover)
VALUES ('Русский пакет 1', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', FALSE, 1, '7ceefef0-2a3e-4e47-a07d-0709c70d1cfa'),
       ('Русский пакет 2', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', FALSE, 1, '44c78b42-e931-4d8c-8451-cc59383e5af6'),
       ('Русский пакет 3', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', TRUE, 1, '922d0089-ec67-47a7-ad25-dc8e36d7c93d'),
       ('English package 1', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', FALSE, 2, '7ceefef0-2a3e-4e47-a07d-0709c70d1cfa'),
       ('English package 2', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', FALSE, 2, '44c78b42-e931-4d8c-8451-cc59383e5af6'),
       ('English package 3', 'd0fbc24f-5061-4d10-b92c-d386c8eba600', TRUE, 2, '922d0089-ec67-47a7-ad25-dc8e36d7c93d');

INSERT INTO package_tag (package_id, tag_id)
VALUES (1, 1),
       (1, 2),
       (1, 6),
       (1, 12),
       (2, 3),
       (2, 11),
       (2, 10),
       (3, 4),
       (3, 12),
       (3, 8),
       (3, 7),
       (4, 13),
       (4, 14),
       (4, 15),
       (5, 16),
       (5, 17),
       (5, 18),
       (5, 19),
       (6, 20),
       (6, 21),
       (6, 22),
       (6, 23),
       (6, 24);

INSERT INTO stage (name, is_final, "order", package_id)
VALUES ('Этап 1', FALSE, 0, 1),
       ('Этап 2', FALSE, 1, 1),
       ('Этап 3', FALSE, 2, 1),
       ('Этап 4', FALSE, 3, 1),
       ('Финальный этап', TRUE, 5, 1),

       ('Этап 1', FALSE, 0, 2),
       ('Финал', TRUE, 5, 2),

       ('Этап 1', FALSE, 0, 3),
       ('Этап 2', FALSE, 1, 3),
       ('Этап 3', FALSE, 2, 3),
       ('Этап 4 финал', TRUE, 5, 3),

       ('Stage 1', FALSE, 0, 4),
       ('Stage 2', FALSE, 1, 4),
       ('Stage 3', FALSE, 2, 4),
       ('Stage 4', FALSE, 3, 4),
       ('Final stage!', TRUE, 5, 4),

       ('Stage 1', FALSE, 0, 5),
       ('Stage 2 final', TRUE, 5, 5),

       ('Stage 1', FALSE, 0, 6),
       ('Stage 2', FALSE, 1, 6),
       ('Final', TRUE, 5, 6);

INSERT INTO stage_content (stage_id, topic_id, question_id, type, cost, interval, comment, secret_topic, secret_cost, is_keepable, is_visible)
VALUES (1, 1, 1, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL),
       (1, 2, 2, 'DEFAULT', 300, 15, NULL, NULL, NULL, NULL, NULL),
       (1, 3, 3, 'DEFAULT', 400, 15, 'this question is sus', NULL, NULL, NULL, NULL),
       (2, 1, 4, 'DEFAULT', 500, 15, NULL, NULL, NULL, NULL, NULL),
       (2, 2, 5, 'DEFAULT', 1000, 30, 'host comment', NULL, NULL, NULL, NULL),
       (2, 3, 6, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (3, 1, 7, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (3, 2, 8, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (3, 3, 9, 'SAFE', 50, 15, NULL, NULL, NULL, NULL, NULL),
       (4, 1, 10, 'SAFE', 150, 15, NULL, NULL, NULL, NULL, NULL),
       (4, 2, 11, 'SAFE', 350, 20, NULL, NULL, NULL, NULL, NULL),
       (4, 3, 12, 'SAFE', 500, 25, NULL, NULL, NULL, NULL, NULL),
       (5, 1, 1, 'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (5, 2, 2, 'SUPERSECRET', 1000, 15, NULL, 'super secret topic 3', 10000, FALSE, TRUE),
       (5, 3, 3, 'SUPERSECRET', 2000, 20, NULL, 'super secret topic 3', 500, TRUE, TRUE),
       (5, 4, 4, 'SUPERSECRET', 2000, 20, NULL, 'super secret topic 3', 500, TRUE, TRUE),
       (5, 5, 5, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),

       (6, 1, 1, 'DEFAULT', 500, 15, NULL, NULL, NULL, NULL, NULL),
       (6, 2, 2, 'SAFE', 50, 15, NULL, NULL, NULL, NULL, NULL),
       (6, 3, 3, 'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (6, 4, 4, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (6, 5, 5, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (7, 1, 6, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (7, 2, 7, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (7, 3, 8, 'BET', 1000, 30, 'bet comment', NULL, NULL, NULL, NULL),
       (7, 4, 9, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (7, 5, 10, 'BET', 150, 15, NULL, NULL, NULL, NULL, NULL),

       (8, 1, 1, 'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (8, 2, 2, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (8, 3, 3, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (9, 1, 4, 'SAFE', 50, 15, NULL, NULL, NULL, NULL, NULL),
       (9, 2, 5, 'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (9, 3, 6, 'BET', 1000, 30, 'bet comment', NULL, NULL, NULL, NULL),
       (10, 1, 7, 'BET', 150, 15, NULL, NULL, NULL, NULL, NULL),
       (10, 2, 8, 'BET', 150, 15, NULL, NULL, NULL, NULL, NULL),
       (10, 3, 9, 'DEFAULT', 1000, 30, 'host comment', NULL, NULL, NULL, NULL),
       (11, 1, 10, 'DEFAULT', 500, 15, NULL, NULL, NULL, NULL, NULL),
       (11, 2, 11, 'DEFAULT', 400, 15, 'this question is sus', NULL, NULL, NULL, NULL),
       (11, 3, 12, 'DEFAULT', 300, 15, NULL, NULL, NULL, NULL, NULL),
       (11, 4, 1, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL),
       (11, 5, 2, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),

       (12, 6, 13, 'DEFAULT', 500, 15, NULL, NULL, NULL, NULL, NULL),
       (12, 7, 14,  'DEFAULT', 500, 15, NULL, NULL, NULL, NULL, NULL),
       (12, 8, 15, 'DEFAULT', 400, 15, 'this question is sus', NULL, NULL, NULL, NULL),
       (13, 6, 16, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL),
       (13, 7, 17, 'BET', 1000, 30, 'bet comment', NULL, NULL, NULL, NULL),
       (13, 8, 18,  'SAFE', 50, 15, NULL, NULL, NULL, NULL, NULL),
       (14, 6, 19,  'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (14, 7, 20, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (14, 8, 21, 'SAFE', 800, 40, NULL, NULL, NULL, NULL, NULL),
       (15, 6, 22, 'BET', 150, 15, NULL, NULL, NULL, NULL, NULL),
       (15, 7, 23, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL),
       (15, 8, 24, 'SECRET', 1000, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),
       (16, 6, 13, 'SAFE', 350, 20, NULL, NULL, NULL, NULL, NULL),
       (16, 7, 14, 'SAFE', 777, 20, NULL, NULL, NULL, NULL, NULL),
       (16, 8, 15, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (16, 9, 16, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (16, 10, 17, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),

       (17, 6, 13, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL),
       (17, 7, 14, 'DEFAULT', 1000, 15, NULL, NULL, NULL, NULL, NULL),
       (17, 8, 15, 'DEFAULT', 2000, 15, NULL, NULL, NULL, NULL, NULL),
       (17, 9, 16, 'DEFAULT', 3000, 15, NULL, NULL, NULL, NULL, NULL),
       (17, 10, 17, 'SUPERSECRET', 5000, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (18, 6, 18, 'SUPERSECRET', 600, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (18, 7, 19, 'SUPERSECRET', 500, 45, NULL, 'super secret topic 3', 500, TRUE, FALSE),
       (18, 8, 20, 'SAFE', 200, 20, NULL, NULL, NULL, NULL, NULL),
       (18, 9, 21, 'SAFE', 350, 20, NULL, NULL, NULL, NULL, NULL),
       (18, 10, 22, 'SECRET', 1500, 20, NULL, 'secret topic 2', 50, FALSE, FALSE),

       (19, 6, 13, 'BET', 450, 15, NULL, NULL, NULL, NULL, NULL),
       (19, 7, 14, 'BET', 1250, 30, NULL, NULL, NULL, NULL, NULL),
       (19, 8, 15, 'BET', 1850, 30, NULL, NULL, NULL, NULL, NULL),
       (19, 9, 16, 'SAFE', 50, 15, NULL, NULL, NULL, NULL, NULL),
       (20, 6, 17, 'SAFE', 500, 30, NULL, NULL, NULL, NULL, NULL),
       (20, 7, 18, 'SAFE', 1000, 30, NULL, NULL, NULL, NULL, NULL),
       (20, 8, 19, 'SECRET', 500, 15, NULL, 'secret topic 1', 5000, FALSE, FALSE),
       (20, 9, 20, 'SUPERSECRET', 3000, 30, 'super secret host comment', 'super secret topic 3', 50, TRUE, TRUE),
       (21, 6, 21, 'SUPERSECRET', 5000, 33, 'super secret host comment', 'super secret topic 3', 50, TRUE, TRUE),
       (21, 7, 22, 'SECRET', 1333, 15, NULL, 'secret topic 1', 5000, FALSE, FALSE),
       (21, 8, 23, 'DEFAULT', 100, 15, NULL, NULL, NULL, NULL, NULL);







