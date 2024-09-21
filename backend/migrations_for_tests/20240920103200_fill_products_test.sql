-- +goose Up

INSERT INTO categories (uid, name, description) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб', 'Хлебное описание');

INSERT INTO products (uid, category_uid, name, description, ccal, price, stock_quantity, created_at, updated_at) VALUES 
('38e27c8f-570c-4d50-bf33-a0bd03afb99e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Ржаной хлеб', 'Полезный хлеб из ржаной муки.', 40, 89, 1, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb88e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Цельнозерновой хлеб', 'Хлеб из цельнозерновой муки с высоким содержанием клетчатки.', 45, 99, 10, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb77e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Овсяные хлебцы', 'Лёгкие и полезные хлебцы из овсяной муки.', 38, 45, 7, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb66e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб из полбы', 'Хлеб из древней пшеничной культуры с высоким содержанием белка.', 42, 70, 3, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb55e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Гречневый хлеб', 'Хлеб из муки гречихи, без глютена.', 37, 105, 4, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb44e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Рисовые хлебцы', 'Полезный продукт из цельного риса.', 35, 40, 2, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb33e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Амарантовый хлеб', 'Хлеб из амарантовой муки, богат витаминами.', 43, 75, 14, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb22e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Льняные хлебцы', 'Полезные хлебцы с высоким содержанием клетчатки.', 40, 50, 8, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb11e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Кукурузные хлебцы', 'Лёгкий продукт без глютена.', 37, 39, 2, NOW()::TIMESTAMP, NOW()::TIMESTAMP),
('38e27c8f-570c-4d50-bf33-a0bd03afb00e', '38e27c8f-570c-4d50-bf33-a0bd03afb13e', 'Хлеб с семенами льна', 'Полезный хлеб с высоким содержанием клетчатки и омега-3.', 44, 85, 4, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- +goose Down

DELETE FROM products WHERE uid IN ('38e27c8f-570c-4d50-bf33-a0bd03afb99e', '38e27c8f-570c-4d50-bf33-a0bd03afb88e', '38e27c8f-570c-4d50-bf33-a0bd03afb77e', '38e27c8f-570c-4d50-bf33-a0bd03afb66e', '38e27c8f-570c-4d50-bf33-a0bd03afb55e', '38e27c8f-570c-4d50-bf33-a0bd03afb44e', '38e27c8f-570c-4d50-bf33-a0bd03afb33e', '38e27c8f-570c-4d50-bf33-a0bd03afb22e', '38e27c8f-570c-4d50-bf33-a0bd03afb11e', '38e27c8f-570c-4d50-bf33-a0bd03afb00e');

DELETE FROM categories WHERE uid='38e27c8f-570c-4d50-bf33-a0bd03afb13e';

