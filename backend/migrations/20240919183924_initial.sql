-- +goose Up
-- SQL в этой секции будет выполнен для обновления БД

-- Создание таблицы "users"
CREATE TABLE users (
    uid uuid PRIMARY KEY,
    email VARCHAR(30) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    birthday TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы "categories" для классификации товаров
CREATE TABLE categories (
    uid uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы "products" для хранения данных о товарах
CREATE TABLE products (
    uid uuid PRIMARY KEY,
    category_uid uuid NOT NULL REFERENCES categories(uid) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    ccal INT,
		weight INT NOT null,
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT false
);

create index ix_products_category_uid on products (category_uid);

-- Создание таблицы "products_count" для хранения количества продуктов в наличии
CREATE TABLE products_count (
    product_uid uuid NOT NULL REFERENCES products(uid) ON DELETE CASCADE,
    stock_quantity INT NOT NULL
);

CREATE TABLE product_photos (
    id uuid NOT NULL,
    product_uid uuid NOT NULL REFERENCES products(uid) ON DELETE SET NULL,
    img_path TEXT NOT NULl
);

INSERT INTO categories (uid, name, description, created_at,updated_at) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб', 'Хлебное описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Рыба', 'Рыбное описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('4fdf5c09-7c04-4853-bab3-240ae3671538', 'Фрукты', 'Фруктовое описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Овощи', 'Овощи описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Мясной отдел', 'Мясной отдел описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('4cf8fc1e-1745-4939-b434-5067684f65fb', 'Бакалея', 'Бакалея описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Заправка', 'Заправка описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Молочные продукты', 'Молочные продукты описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Топпинги', 'Топпинги описание', NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Хлеб
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb99e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Ржаной хлеб', 'Полезный хлеб из ржаной муки.', 40, 89, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb88e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Цельнозерновой хлеб', 'Хлеб из цельнозерновой муки с высоким содержанием клетчатки.', 45, 99, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb77e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Овсяные хлебцы', 'Лёгкие и полезные хлебцы из овсяной муки.', 38, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb66e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Хлеб из полбы', 'Хлеб из древней пшеничной культуры с высоким содержанием белка.', 42, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb55e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Гречневый хлеб', 'Хлеб из муки гречихи, без глютена.', 37, 105, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb44e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Рисовые хлебцы', 'Полезный продукт из цельного риса.', 35, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb33e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Амарантовый хлеб', 'Хлеб из амарантовой муки, богат витаминами.', 43, 75, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb22e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Льняные хлебцы', 'Полезные хлебцы с высоким содержанием клетчатки.', 40, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb11e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Кукурузные хлебцы', 'Лёгкий продукт без глютена.', 37, 39, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb00e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 20, 'Хлеб с семенами льна', 'Полезный хлеб с высоким содержанием клетчатки и омега-3.', 44, 85, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

insert into product_photos (id, product_uid, img_path) VALUES
('e1c22b00-e924-4c37-a639-8bccc1aab8a2', '38e27c8f-570c-4d50-bf33-a0bd03afb99e', 'assets/imgs/bread/rye_bread_1.png'),
('4527a2db-ed25-415e-9b9f-19d316eb1a0f', '38e27c8f-570c-4d50-bf33-a0bd03afb99e', 'assets/imgs/bread/rye_bread_2.png'),
('277308e7-0b27-4f2f-a910-b294d7708a92', '38e27c8f-570c-4d50-bf33-a0bd03afb88e', 'assets/imgs/bread/wheat_bread_1.png'),
('a1e90029-9c8b-4042-aaaa-fe2cd1ba081d', '38e27c8f-570c-4d50-bf33-a0bd03afb88e', 'assets/imgs/bread/wheat_bread_2.png'),
('44fe01a9-fbc6-4c52-905a-ce089f284a53', '38e27c8f-570c-4d50-bf33-a0bd03afb77e', 'assets/imgs/bread/oat_bread_1.png'),
('db417c1d-1537-4676-8143-3f4705a3ecf1', '38e27c8f-570c-4d50-bf33-a0bd03afb77e', 'assets/imgs/bread/oat_bread_2.png'),
('6014ec53-81e3-4fbc-9437-c591a8f0f7bb', '38e27c8f-570c-4d50-bf33-a0bd03afb66e', 'assets/imgs/bread/spelled_bread_1.png'),
('40ff9c43-9aa1-40e1-9f3c-b5c6a5dd5cb8', '38e27c8f-570c-4d50-bf33-a0bd03afb66e', 'assets/imgs/bread/spelled_bread_2.png'),
('7a73a017-8f4a-4432-a69c-4ad6816f2db2', '38e27c8f-570c-4d50-bf33-a0bd03afb55e', 'assets/imgs/bread/buckwheat_bread_1.png'),
('5ffb3ffb-6ae2-4322-b90d-d98be012a773', '38e27c8f-570c-4d50-bf33-a0bd03afb44e', 'assets/imgs/bread/rice_cake_1.png'),
('cf1107a8-872d-4518-b2cc-3c706e39eb72', '38e27c8f-570c-4d50-bf33-a0bd03afb44e', 'assets/imgs/bread/rice_cake_2.png'),
('6135439c-2492-4d3f-92e2-2792bae3a142', '38e27c8f-570c-4d50-bf33-a0bd03afb33e', 'assets/imgs/bread/amaranth_bread_1.png'),
('2b22b276-5abe-46f3-b8e2-742f575a81dc', '38e27c8f-570c-4d50-bf33-a0bd03afb33e', 'assets/imgs/bread/amaranth_bread_2.png'),
('728538c9-702d-44b0-8787-fefdd83994c4', '38e27c8f-570c-4d50-bf33-a0bd03afb22e', 'assets/imgs/bread/linen_cakes_1.png'),
('1588e609-0a2f-4ada-85ad-b4462e65fd6d', '38e27c8f-570c-4d50-bf33-a0bd03afb22e', 'assets/imgs/bread/linen_cakes_2.png'),
('755f8d83-0077-4ffb-ba0b-6de23052a3c9', '38e27c8f-570c-4d50-bf33-a0bd03afb11e', 'assets/imgs/bread/corn_cakes_1.png'),
('bb565e29-30f5-45a1-8a28-d7a21383e4f5', '38e27c8f-570c-4d50-bf33-a0bd03afb00e', 'assets/imgs/bread/flax_seeds_bread_1.png');

-- Рыба
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb101', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Лосось', 'Жирная рыба, богатая омега-3 жирными кислотами.', 208, 220, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb102', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Тунец', 'Полезная и диетическая рыба.', 144, 180, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb103', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Форель', 'Рыба, богатая белком и омега-3 кислотами.', 187, 200, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb104', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Скумбрия', 'Полезная жирная рыба с высоким содержанием витаминов.', 205, 150, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb105', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Хек', 'Нежирная рыба с низким содержанием калорий.', 86, 100, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb106', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Треска', 'Нежирная рыба, богатая витаминами.', 82, 120, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb107', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Кета', 'Рыба с умеренным содержанием жира и высоким содержанием белка.', 127, 130, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb108', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Минтай', 'Низкокалорийная рыба, идеальная для диеты.', 72, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb109', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Палтус', 'Белая рыба с высоким содержанием белка.', 186, 210, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb110', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 100, 'Сельдь', 'Жирная рыба, богатая кальцием и витамином D.', 187, 110, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Фрукты
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb201', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Яблоко', 'Источник клетчатки и витамина C.', 52, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb202', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Банан', 'Богат калием и быстрыми углеводами.', 96, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb203', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Апельсин', 'Цитрусовый фрукт, богатый витамином C.', 47, 42, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb204', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Груша', 'Сладкий фрукт, богатый клетчаткой.', 57, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb205', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Киви', 'Богатый витамином C и антиоксидантами.', 61, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb206', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Гранат', 'Фрукт с высоким содержанием антиоксидантов.', 83, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb207', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Малина', 'Ягода с низким содержанием сахара и высоким содержанием клетчатки.', 52, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb208', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Черника', 'Антиоксидантная ягода для поддержания здоровья глаз.', 57, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb209', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Персик', 'Сочный фрукт, богатый витаминами.', 39, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb210', '4fdf5c09-7c04-4853-bab3-240ae3671538', 20, 'Грейпфрут', 'Противовирусное средство с ярким вкусом.', 42, 65, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Овощи
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb301', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Помидор', 'Богатый антиоксидантами овощ.', 1, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb302', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Огурец', 'Низкокалорийный овощ, идеальный для салатов.', 0.4, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb303', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Морковь', 'Богатая бета-каротином и клетчаткой.', 1.2, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb304', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Брокколи', 'Крестоцветный овощ, богатый витаминами и минералами.', 1, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb305', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Кабачок', 'Низкокалорийный овощ с высоким содержанием клетчатки.', 0.8, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb306', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Цветная капуста', 'Полезный овощ, богатый витаминами и антиоксидантами.', 0.8, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb307', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Шпинат', 'Листовой овощ, богатый железом и витаминами.', 0.6, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb308', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Редис', 'Овощ с низким содержанием калорий и высоким содержанием витамина C.', 0.8, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb309', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Зелёный лук', 'Ароматный овощ, богатый витаминами A и C.', 0.4, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb310', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Сельдерей', 'Низкокалорийный овощ с высоким содержанием клетчатки.', 0.4, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb311', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Салат', 'Низкокалорийный продукт, богатый витаминами A, C и K, антиоксидантами и клетчаткой.', 0.15, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb312', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 20, 'Болгарский перец', 'Сладкий овощ с высоким содержанием витамина C и антиоксидантов.', 8, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Мясной отдел
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb401', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Куриная грудка', 'Нежирное мясо, богатое белком.', 220, 300, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb402', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Индейка', 'Диетическое мясо с низким содержанием жира.', 200, 350, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb403', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Телятина', 'Нежирное и питательное мясо.', 240, 600, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb404', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Говяжья вырезка', 'Постное мясо с высоким содержанием белка.', 280, 700, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb405', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Кролик', 'Легкоусвояемое мясо с низким содержанием жира.', 250, 450, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb406', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Баранина', 'Нежное мясо с насыщенным вкусом.', 300, 550, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb407', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Утиная грудка', 'Богатое белком мясо с умеренным содержанием жира.', 350, 800, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb408', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Филе оленя', 'Диетическое мясо с насыщенным вкусом.', 260, 1200, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb409', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Говяжий язык', 'Мясо с высоким содержанием белка и железа.', 300, 400, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb410', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Печень говяжья', 'Полезный субпродукт, богатый железом и витаминами группы B.', 200, 250, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb411', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 100, 'Печень куриная', 'Полезный субпродукт, богатый железом и витаминами группы B.', 200, 250, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Бакалея
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb501', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Киноа', 'Псевдозерновая культура, богатая белком и клетчаткой.', 3, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb502', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Коричневый рис', 'Полезный цельнозерновой рис с низким гликемическим индексом.', 2.8, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb503', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Булгур', 'Пшеница, обработанная на пару, легкоусвояемая и питательная.', 2.8, 25, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb504', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Овсяные хлопья', 'Цельнозерновые овсяные хлопья для быстрого приготовления.', 3.2, 20, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb505', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Гречневая крупа', 'Полезная крупа с высоким содержанием белка и клетчатки.', 2.6, 18, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb506', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Чечевица', 'Богатый белком и железом бобовый продукт.', 3.6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb507', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Перловка', 'Цельнозерновая перловая крупа, полезная для пищеварения.', 2.4, 16, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb508', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Кускус', 'Быстро приготовляемая крупа из твёрдых сортов пшеницы.', 2.6, 26, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb509', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Амарант', 'Безглютеновая культура с высоким содержанием белка.', 3.2, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb510', '4cf8fc1e-1745-4939-b434-5067684f65fb', 20, 'Ячневая крупа', 'Дроблёная ячменная крупа для полезного питания.', 2.2, 18, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Заправка
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb511', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Оливковое масло', 'Холодного отжима, богато антиоксидантами и омега-9.', 36, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb512', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Льняное масло', 'Масло с высоким содержанием омега-3 жирных кислот.', 36, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb513', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Масло грецкого ореха', 'Богатое полезными жирами и антиоксидантами.', 36, 120, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb514', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Масло авокадо', 'Высокое содержание полезных мононенасыщенных жиров.', 36, 140, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb515', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Кунжутное масло', 'Ароматное масло, полезное для суставов и костей.', 36, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb516', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Кокосовое масло', 'Полезное масло для приготовления и увлажнения кожи.', 36, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb517', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Миндальное масло', 'Богатое витаминами Е и омега-6 жирными кислотами.', 36, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb518', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Масло тыквенных семян', 'Полезное масло с ярким ароматом и высоким содержанием цинка.', 36, 85, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb519', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Масло из расторопши', 'Обладает антиоксидантными и противовоспалительными свойствами.', 36, 55, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb520', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Масло виноградных косточек', 'Легкое масло с высоким содержанием витамина E.', 36, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb521', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 20, 'Бальзамический уксус', 'Ароматная приправа, богатая антиоксидантами и минералами.', 14, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Молочные продукты
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec3', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Нежирное молоко', 'Молоко с низким содержанием жира, богатое кальцием и витаминами.', 2, 20, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec4', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Греческий йогурт', 'Йогурт с высоким содержанием белка и низким содержанием сахара.', 4.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec5', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Творог 5% жирности', 'Творог с низким содержанием жира, богатый белком и кальцием.', 2, 25, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec6', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Кефир 1% жирности', 'Пробиотический напиток, полезный для пищеварения.', 2, 22, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec7', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Молочный напиток из овса', 'Веганский молочный напиток с добавлением витаминов и минералов.', 4, 24, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec9', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Брынза', 'Соленый сыр, богатый белком и кальцием.', 5, 35, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9eca', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Рикотта', 'Мягкий сыр с низким содержанием жира и высоким содержанием белка.', 6, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecb', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Сыр моцарелла', 'Сыр с низким содержанием жира и мягкой текстурой.', 5.6, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecc', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 20, 'Сметана 10% жирности', 'Кисломолочный продукт с умеренным содержанием жира.', 4, 23, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Топинги
INSERT INTO products (uid, category_uid, weight, name, description, ccal, price, created_at, updated_at) VALUES
('da44e640-c3ee-4bf2-9647-b312117b7435', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Миндаль', 'Орех с высоким содержанием витаминов и полезных жиров.', 2.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('7de410b4-5caf-4bfa-9aa8-632af1d047c2', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Льняные семена', 'Семена с высоким содержанием омега-3 жирных кислот и клетчатки.', 3.2, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('1b00cc80-896a-455a-aadc-142c5879e6c6', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Чиа', 'Семена с высоким содержанием клетчатки и омега-3.', 3.2, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('9021cbed-ed00-490f-b93a-3ea0fa600c58', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Грецкие орехи', 'Орехи с высоким содержанием полезных жиров и антиоксидантов.', 4, 35, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('b557bd28-0870-4533-968a-3a3788c58912', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Кокосовая стружка', 'Сухая стружка кокоса, богата клетчаткой и полезными жирами.', 6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('9502434e-339f-4765-898d-48271e5a64ec', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Кедровые орехи', 'Орехи с высоким содержанием витаминов и минералов.', 4.2, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('96e091ed-981a-4c4d-afdc-e5bae7151df8', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Пекан', 'Орехи с насыщенным вкусом и высоким содержанием полезных жиров.', 5, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('06784003-3c44-4913-adb6-7e433f82d19e', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Семена подсолнечника', 'Семена с высоким содержанием витамина E и полезных жиров.', 3, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('fd99b689-646d-47e0-ac98-287b60be6f45', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Семена тыквы', 'Семена с высоким содержанием минералов и омега-3 жирных кислот.', 3.6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('a323efa1-4788-459d-a1c4-c1710d289edc', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Кунжут', 'Семена с высоким содержанием кальция и полезных жиров.', 3.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('a323efa1-4788-459d-a1c4-c1710d289edd', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 20, 'Мед', 'Натуральный подсластитель, богатый антиоксидантами и полезными ферментами.', 64, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

INSERT INTO public.products_count (product_uid,stock_quantity) VALUES
	('38e27c8f-570c-4d50-bf33-a0bd03afb99e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb88e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb77e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb66e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb55e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb44e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb33e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb22e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb11e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb00e',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb101',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb102',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb103',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb104',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb105',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb106',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb107',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb108',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb109',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb110',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb201',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb202',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb203',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb204',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb205',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb206',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb207',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb208',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb209',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb210',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb301',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb302',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb303',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb304',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb305',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb306',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb307',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb308',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb309',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb310',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb401',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb402',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb403',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb404',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb405',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb406',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb407',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb408',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb409',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb410',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb501',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb502',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb503',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb504',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb505',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb506',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb507',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb508',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb509',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb510',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb511',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb512',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb513',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb514',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb515',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb516',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb517',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb518',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb519',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb520',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec3',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec4',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec5',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec6',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec7',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec9',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9eca',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecb',100),
	('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecc',100),
	('da44e640-c3ee-4bf2-9647-b312117b7435',100),
	('7de410b4-5caf-4bfa-9aa8-632af1d047c2',100),
	('1b00cc80-896a-455a-aadc-142c5879e6c6',100),
	('9021cbed-ed00-490f-b93a-3ea0fa600c58',100),
	('b557bd28-0870-4533-968a-3a3788c58912',100),
	('9502434e-339f-4765-898d-48271e5a64ec',100),
	('96e091ed-981a-4c4d-afdc-e5bae7151df8',100),
	('06784003-3c44-4913-adb6-7e433f82d19e',100),
	('fd99b689-646d-47e0-ac98-287b60be6f45',100),
	('38e27c8f-570c-4d50-bf33-a0bd03afb312',100),
	('a323efa1-4788-459d-a1c4-c1710d289edc',10);

-- +goose Down
-- SQL в этой секции будет выполнен для отката изменений

DROP INDEX IF EXISTS ix_products_category_uid;

-- Удаление всех созданных таблиц
DROP TABLE IF EXISTS product_photos;
DROP TABLE IF EXISTS products_count;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS cities_handbook;
DROP TABLE IF EXISTS users;