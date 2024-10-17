-- +goose Up
-- SQL в этой секции будет выполнен для обновления БД

-- Создание таблицы "users"
CREATE TABLE users (
    uid uuid PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы "addresses" для хранения адресов доставки пользователей с широтой и долготой
CREATE TABLE delivery_addresses (
    uid uuid PRIMARY KEY,
    user_uid uuid NOT NULL REFERENCES users(uid) ON DELETE CASCADE,
    latitude NUMERIC(10, 8) NOT NULL,  -- Широта с точностью до 8 знаков после запятой
    longitude NUMERIC(11, 8) NOT NULL, -- Долгота с точностью до 8 знаков после запятой
    city_name VARCHAR(30) NOT NULL,
    street_name VARCHAR(30) NOT NULL,
    house_number INT NOT NULL,
    building INT NOT NULL,
    floor INT NOT NULL,
    apartment INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы "payment_cards" для хранения карт оплаты
CREATE TABLE payment_cards (
    uid uuid PRIMARY KEY,
    user_uid uuid NOT NULL REFERENCES users(uid) ON DELETE CASCADE,
    card_number VARCHAR(16) NOT NULL,
    card_holder_name VARCHAR(100) NOT NULL,
    expiration_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL
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
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

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

-- Создание таблицы "recipes" для хранения рецептов, где используются продукты
-- products [{
--    "name": "Хлеб",
--    "quantity": 0.5,
--    "measure": "gramm"
--}]
CREATE TABLE recipes (
    uid uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cooking_time INT,
    products json,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO categories (uid, name, description) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб', 'Хлебное описание'),
('e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Рыба', 'Рыбное описание'),
('4fdf5c09-7c04-4853-bab3-240ae3671538', 'Фрукты', 'Фруктовое описание'),
('82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Овощи', 'Овощи описание'),
('23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Мясной отдел', 'Мясной отдел описание'),
('4cf8fc1e-1745-4939-b434-5067684f65fb', 'Бакалея', 'Бакалея описание'),
('36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Заправка', 'Заправка описание'),
('038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Молочные продукты', 'Молочные продукты описание'),
('1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Топпинги', 'Топпинги описание');

-- Связь продуктов с рецептами (многие ко многим)
-- CREATE TABLE recipe_products (
--     recipe_uid uuid NOT NULL REFERENCES recipes(uid) ON DELETE CASCADE,
--     product_uid uuid NOT NULL,
--     quantity NUMERIC(10, 2) NOT NULL,
--     name VARCHAR(255) NOT NULL,
--     PRIMARY KEY (recipe_uid, product_uid)
-- );

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
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb99e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Ржаной хлеб', 'Полезный хлеб из ржаной муки.', 40, 89, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb88e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Цельнозерновой хлеб', 'Хлеб из цельнозерновой муки с высоким содержанием клетчатки.', 45, 99, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb77e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Овсяные хлебцы', 'Лёгкие и полезные хлебцы из овсяной муки.', 38, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb66e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб из полбы', 'Хлеб из древней пшеничной культуры с высоким содержанием белка.', 42, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb55e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Гречневый хлеб', 'Хлеб из муки гречихи, без глютена.', 37, 105, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb44e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Рисовые хлебцы', 'Полезный продукт из цельного риса.', 35, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb33e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Амарантовый хлеб', 'Хлеб из амарантовой муки, богат витаминами.', 43, 75, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb22e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Льняные хлебцы', 'Полезные хлебцы с высоким содержанием клетчатки.', 40, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb11e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Кукурузные хлебцы', 'Лёгкий продукт без глютена.', 37, 39, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb00e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб с семенами льна', 'Полезный хлеб с высоким содержанием клетчатки и омега-3.', 44, 85, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Рыба
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb101', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Лосось', 'Жирная рыба, богатая омега-3 жирными кислотами.', 208, 220, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb102', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Тунец', 'Полезная и диетическая рыба.', 144, 180, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb103', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Форель', 'Рыба, богатая белком и омега-3 кислотами.', 187, 200, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb104', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Скумбрия', 'Полезная жирная рыба с высоким содержанием витаминов.', 205, 150, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb105', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Хек', 'Нежирная рыба с низким содержанием калорий.', 86, 100, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb106', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Треска', 'Нежирная рыба, богатая витаминами.', 82, 120, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb107', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Кета', 'Рыба с умеренным содержанием жира и высоким содержанием белка.', 127, 130, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb108', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Минтай', 'Низкокалорийная рыба, идеальная для диеты.', 72, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb109', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Палтус', 'Белая рыба с высоким содержанием белка.', 186, 210, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb110', 'e069195a-ab82-4a17-ad9d-111a1dee2afd', 'Сардины', 'Жирная рыба, богатая кальцием и витамином D.', 140, 110, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Фрукты
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb201', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Яблоко', 'Источник клетчатки и витамина C.', 52, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb202', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Банан', 'Богат калием и быстрыми углеводами.', 96, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb203', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Апельсин', 'Цитрусовый фрукт, богатый витамином C.', 47, 42, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb204', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Груша', 'Сладкий фрукт, богатый клетчаткой.', 57, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb205', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Киви', 'Богатый витамином C и антиоксидантами.', 61, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb206', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Гранат', 'Фрукт с высоким содержанием антиоксидантов.', 83, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb207', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Малина', 'Ягода с низким содержанием сахара и высоким содержанием клетчатки.', 52, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb208', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Черника', 'Антиоксидантная ягода для поддержания здоровья глаз.', 57, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb209', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Персик', 'Сочный фрукт, богатый витаминами.', 39, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb210', '4fdf5c09-7c04-4853-bab3-240ae3671538', 'Грейпфрут', 'Противовирусное средство с ярким вкусом.', 42, 65, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Овощи
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb301', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Помидор', 'Богатый антиоксидантами овощ.', 1, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb302', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Огурец', 'Низкокалорийный овощ, идеальный для салатов.', 0.4, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb303', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Морковь', 'Богатая бета-каротином и клетчаткой.', 1.2, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb304', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Брокколи', 'Крестоцветный овощ, богатый витаминами и минералами.', 1, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb305', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Кабачок', 'Низкокалорийный овощ с высоким содержанием клетчатки.', 0.8, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb306', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Цветная капуста', 'Полезный овощ, богатый витаминами и антиоксидантами.', 0.8, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb307', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Шпинат', 'Листовой овощ, богатый железом и витаминами.', 0.6, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb308', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Редис', 'Овощ с низким содержанием калорий и высоким содержанием витамина C.', 0.8, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb309', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Зелёный лук', 'Ароматный овощ, богатый витаминами A и C.', 0.4, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb310', '82ebbd1c-dcfe-4d12-8729-bd778de4365c', 'Сельдерей', 'Низкокалорийный овощ с высоким содержанием клетчатки.', 0.4, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Мясной отдел
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb401', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Куриная грудка', 'Нежирное мясо, богатое белком.', 220, 300, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb402', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Индейка', 'Диетическое мясо с низким содержанием жира.', 200, 350, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb403', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Телятина', 'Нежирное и питательное мясо.', 240, 600, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb404', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Говяжья вырезка', 'Постное мясо с высоким содержанием белка.', 280, 700, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb405', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Кролик', 'Легкоусвояемое мясо с низким содержанием жира.', 250, 450, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb406', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Баранина', 'Нежное мясо с насыщенным вкусом.', 300, 550, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb407', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Утиная грудка', 'Богатое белком мясо с умеренным содержанием жира.', 350, 800, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb408', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Филе оленя', 'Диетическое мясо с насыщенным вкусом.', 260, 1200, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb409', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Говяжий язык', 'Мясо с высоким содержанием белка и железа.', 300, 400, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb410', '23f1cf15-95ac-4abf-bda2-37c62fa24e5d', 'Печень говяжья', 'Полезный субпродукт, богатый железом и витаминами группы B.', 200, 250, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Бакалея
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb501', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Киноа', 'Псевдозерновая культура, богатая белком и клетчаткой.', 3, 80, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb502', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Коричневый рис', 'Полезный цельнозерновой рис с низким гликемическим индексом.', 2.8, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb503', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Булгур', 'Пшеница, обработанная на пару, легкоусвояемая и питательная.', 2.8, 25, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb504', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Овсяные хлопья', 'Цельнозерновые овсяные хлопья для быстрого приготовления.', 3.2, 20, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb505', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Гречневая крупа', 'Полезная крупа с высоким содержанием белка и клетчатки.', 2.6, 18, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb506', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Чечевица', 'Богатый белком и железом бобовый продукт.', 3.6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb507', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Перловка', 'Цельнозерновая перловая крупа, полезная для пищеварения.', 2.4, 16, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb508', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Кускус', 'Быстро приготовляемая крупа из твёрдых сортов пшеницы.', 2.6, 26, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb509', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Амарант', 'Безглютеновая культура с высоким содержанием белка.', 3.2, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb510', '4cf8fc1e-1745-4939-b434-5067684f65fb', 'Ячневая крупа', 'Дроблёная ячменная крупа для полезного питания.', 2.2, 18, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Заправка
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('38e27c8f-570c-4d50-bf33-a0bd03afb511', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Оливковое масло', 'Холодного отжима, богато антиоксидантами и омега-9.', 36, 60, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb512', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Льняное масло', 'Масло с высоким содержанием омега-3 жирных кислот.', 36, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb513', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Масло грецкого ореха', 'Богатое полезными жирами и антиоксидантами.', 36, 120, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb514', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Масло авокадо', 'Высокое содержание полезных мононенасыщенных жиров.', 36, 140, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb515', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Кунжутное масло', 'Ароматное масло, полезное для суставов и костей.', 36, 70, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb516', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Кокосовое масло', 'Полезное масло для приготовления и увлажнения кожи.', 36, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb517', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Миндальное масло', 'Богатое витаминами Е и омега-6 жирными кислотами.', 36, 90, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb518', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Масло тыквенных семян', 'Полезное масло с ярким ароматом и высоким содержанием цинка.', 36, 85, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb519', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Масло из расторопши', 'Обладает антиоксидантными и противовоспалительными свойствами.', 36, 55, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb520', '36adad5c-8c71-4beb-ade3-ae6b0a10d998', 'Масло виноградных косточек', 'Легкое масло с высоким содержанием витамина E.', 36, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Молочные продукты
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec3', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Нежирное молоко', 'Молоко с низким содержанием жира, богатое кальцием и витаминами.', 2, 20, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec4', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Греческий йогурт', 'Йогурт с высоким содержанием белка и низким содержанием сахара.', 4.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec5', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Творог 5% жирности', 'Творог с низким содержанием жира, богатый белком и кальцием.', 2, 25, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec6', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Кефир 1% жирности', 'Пробиотический напиток, полезный для пищеварения.', 2, 22, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec7', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Молочный напиток из овса', 'Веганский молочный напиток с добавлением витаминов и минералов.', 4, 24, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec8', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Кottage cheese', 'Творог с нежной текстурой, богатый белком.', 3, 28, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ec9', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Брынза', 'Соленый сыр, богатый белком и кальцием.', 5, 35, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9eca', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Рикотта', 'Мягкий сыр с низким содержанием жира и высоким содержанием белка.', 6, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecb', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Сыр моцарелла', 'Сыр с низким содержанием жира и мягкой текстурой.', 5.6, 45, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('f1c52c01-d7be-43aa-b9f6-f30e5d7f9ecc', '038f69de-cb6e-4d72-87f4-6dfdf0e4f290', 'Сметана 10% жирности', 'Кисломолочный продукт с умеренным содержанием жира.', 4, 23, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- Топинги
INSERT INTO products (uid, category_uid, name, description, ccal, price, created_at, updated_at) VALUES
('da44e640-c3ee-4bf2-9647-b312117b7435', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Миндаль', 'Орех с высоким содержанием витаминов и полезных жиров.', 2.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('7de410b4-5caf-4bfa-9aa8-632af1d047c2', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Льняные семена', 'Семена с высоким содержанием омега-3 жирных кислот и клетчатки.', 3.2, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('1b00cc80-896a-455a-aadc-142c5879e6c6', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Чиа', 'Семена с высоким содержанием клетчатки и омега-3.', 3.2, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('9021cbed-ed00-490f-b93a-3ea0fa600c58', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Грецкие орехи', 'Орехи с высоким содержанием полезных жиров и антиоксидантов.', 4, 35, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('b557bd28-0870-4533-968a-3a3788c58912', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Кокосовая стружка', 'Сухая стружка кокоса, богата клетчаткой и полезными жирами.', 6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('9502434e-339f-4765-898d-48271e5a64ec', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Кедровые орехи', 'Орехи с высоким содержанием витаминов и минералов.', 4.2, 50, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('96e091ed-981a-4c4d-afdc-e5bae7151df8', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Пекан', 'Орехи с насыщенным вкусом и высоким содержанием полезных жиров.', 5, 40, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('06784003-3c44-4913-adb6-7e433f82d19e', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Семена подсолнечника', 'Семена с высоким содержанием витамина E и полезных жиров.', 3, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('fd99b689-646d-47e0-ac98-287b60be6f45', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Семена тыквы', 'Семена с высоким содержанием минералов и омега-3 жирных кислот.', 3.6, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('a323efa1-4788-459d-a1c4-c1710d289edc', '1a1e32b7-5c0f-4409-a82c-3e5c973868e6', 'Кунжут', 'Семена с высоким содержанием кальция и полезных жиров.', 3.4, 30, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- +goose Down
-- SQL в этой секции будет выполнен для отката изменений

-- Удаление всех созданных таблиц
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS product_photos;
DROP TABLE IF EXISTS products_count;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS payment_cards;
DROP TABLE IF EXISTS delivery_addresses;
DROP TABLE IF EXISTS cities_handbook;
DROP TABLE IF EXISTS users;