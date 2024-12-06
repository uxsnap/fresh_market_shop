-- +goose Up
-- +goose StatementBegin

-- Рыба / fish

insert into product_photos (id, product_uid, img_path) VALUES
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb101', 'assets/imgs/fish/salmon_1.png'),
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb102', 'assets/imgs/fish/tuna_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb103', 'assets/imgs/fish/trout_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb104', 'assets/imgs/fish/mackerel_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb105', 'assets/imgs/fish/hake_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb106', 'assets/imgs/fish/cod_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb107', 'assets/imgs/fish/chum_salmon_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb108', 'assets/imgs/fish/pollock_1.png'), 
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb109', 'assets/imgs/fish/halibut_1.png'),
(gen_random_uuid(), '38e27c8f-570c-4d50-bf33-a0bd03afb110', 'assets/imgs/fish/herring_1.png');

INSERT INTO recipes (uid, name, ccal, created_at, updated_at, cooking_time) VALUES
('479b421c-5784-4eaf-97fb-dee47c756f09', 'Филе индейки с кедровыми орехами', 170, NOW()::TIMESTAMP, NOW()::TIMESTAMP, '2 hours'::interval);

INSERT INTO recipes_products (recipe_uid, product_uid) VALUES 
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb402'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb511'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb312'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb309'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb310'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '9502434e-339f-4765-898d-48271e5a64ec'),
('479b421c-5784-4eaf-97fb-dee47c756f09', '38e27c8f-570c-4d50-bf33-a0bd03afb203');

INSERT INTO recipes_steps (recipe_uid, step, description) VALUES 
('479b421c-5784-4eaf-97fb-dee47c756f09', 1, 'Возьмите филе индейки и разрежьте на две части, в каждой сделайте надрез, чтобы получился кармашек. В отдельной посуде смешайте оливковое масло и паприку. Филе смажьте маринадом и оставьте мариноваться на некоторое время.'),
('479b421c-5784-4eaf-97fb-dee47c756f09', 2, 'Займитесь овощами. На плиту поставьте глубокую сковородку и включите огонь. Добавьте сливочное масло и растопите его. Затем выложите порезанные перцы, лук-порей и стебель сельдерея. Хорошо обжарьте, периодически помешивая. Когда овощи станут мягкими, добавьте кедровые орехи и немного потушите на среднем огне.'),
('479b421c-5784-4eaf-97fb-dee47c756f09', 3, 'Запеките филе. Включите духовку на 180 градусов. В кармашек, который получился на каждом кусочке мяса, положите начинку, посолите по вкусу. Закройте кармашек и перетяните ниткой, заверните в фольгу и положите в форму для запекания. Поставьте форму в духовку на 1 час. Затем разверните фольгу, полейте образовавшимся соком и поставьте в духовку еще на 20 минут.'),
('479b421c-5784-4eaf-97fb-dee47c756f09', 4, 'Подайте филе индейки на сервировочном блюде, полив апельсиновым соком и украсив ломтиками апельсина.');


-- +goose StatementEnd