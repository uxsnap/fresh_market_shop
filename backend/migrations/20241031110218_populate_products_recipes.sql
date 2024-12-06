-- +goose Up
-- +goose StatementBegin

-- Молочные продукты / dairy

insert into product_photos (id, product_uid, img_path) VALUES
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec3', 'assets/imgs/dairy/milk_1.png'),
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec4', 'assets/imgs/dairy/greek_yog_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec5', 'assets/imgs/dairy/cottage_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec6', 'assets/imgs/dairy/kefir_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec7', 'assets/imgs/dairy/oat_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec9', 'assets/imgs/dairy/feta_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9eca', 'assets/imgs/dairy/ricotta_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecb', 'assets/imgs/dairy/mozzarella_1.png'), 
(gen_random_uuid(), 'f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecc', 'assets/imgs/dairy/soar_creme_1.png');

INSERT INTO recipes (uid, name, ccal, created_at, updated_at, cooking_time) VALUES
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 'Теплый салат с куриной печенью и апельсинами', 185, NOW()::TIMESTAMP, NOW()::TIMESTAMP, '20 minutes'::interval);

INSERT INTO recipes_products (recipe_uid, product_uid) VALUES 
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '38e27c8f-570c-4d50-bf33-a0bd03afb411'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '38e27c8f-570c-4d50-bf33-a0bd03afb203'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '9502434e-339f-4765-898d-48271e5a64ec'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '38e27c8f-570c-4d50-bf33-a0bd03afb311'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '38e27c8f-570c-4d50-bf33-a0bd03afb511'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', '38e27c8f-570c-4d50-bf33-a0bd03afb521'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 'a323efa1-4788-459d-a1c4-c1710d289edd');

INSERT INTO recipes_steps (recipe_uid, step, description) VALUES 
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 1, 'Куриную печень очистите от прожилок и нарежьте небольшими кусочками. Вымойте, посолите по вкусу, посыпьте черным молотым перцем. Обжарьте на подсолнечном масле со всех сторон в течение 3-4 минут. Внутри печень должна быть сочной.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 2, 'Апельсин очистите от кожуры, разделите на дольки и удалите белую пленку.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 3, 'Кедровые орешки слегка поджарьте на сухой сковороде.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 4, 'Приготовьте салатную заправку. Смешайте оливковое масло, бальзамический уксус и мед.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 5, 'Салатные листья порвите руками и выложите на плоское блюдо, сверху положите дольки апельсина.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 6, 'Добавьте кусочки теплой куриной печени и посыпьте кедровыми орешками.'),
('7f14286d-58e0-4778-b7d5-826a784fdfe7', 7, 'Полейте салатной заправкой и готовый теплый салат с куриной печенью и апельсинами подавайте сразу же. Приятного аппетита!');


-- +goose StatementEnd