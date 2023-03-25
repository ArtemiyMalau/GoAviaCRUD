CREATE TABLE IF NOT EXISTS `account` (
  `id` int NOT NULL,
  `scheme_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `airline_provider` (
  `airline_id` int NOT NULL,
  `provider_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `airline` (
  `id` int NOT NULL,
  `code` varchar(8) NOT NULL,
  `name` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `provider` (
  `id` int NOT NULL,
  `code` varchar(8) NOT NULL,
  `name` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `scheme_provider` (
  `scheme_id` int NOT NULL,
  `provider_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `scheme` (
  `id` int NOT NULL,
  `name` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Индексы таблицы `account`
--
ALTER TABLE `account`
  ADD PRIMARY KEY (`id`),
  ADD KEY `account_ibfk_1` (`scheme_id`);

--
-- Индексы таблицы `airline`
--
ALTER TABLE `airline`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Индексы таблицы `airline_provider`
--
ALTER TABLE `airline_provider`
  ADD PRIMARY KEY (`airline_id`,`provider_id`),
  ADD KEY `provider_id` (`provider_id`);

--
-- Индексы таблицы `provider`
--
ALTER TABLE `provider`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Индексы таблицы `scheme`
--
ALTER TABLE `scheme`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Индексы таблицы `scheme_provider`
--
ALTER TABLE `scheme_provider`
  ADD PRIMARY KEY (`scheme_id`,`provider_id`),
  ADD KEY `provider_id` (`provider_id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `account`
--
ALTER TABLE `account`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `airline`
--
ALTER TABLE `airline`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `provider`
--
ALTER TABLE `provider`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `scheme`
--
ALTER TABLE `scheme`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `account`
--
ALTER TABLE `account`
  ADD CONSTRAINT `account_ibfk_1` FOREIGN KEY (`scheme_id`) REFERENCES `scheme` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `airline_provider`
--
ALTER TABLE `airline_provider`
  ADD CONSTRAINT `airline_provider_ibfk_1` FOREIGN KEY (`airline_id`) REFERENCES `airline` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `airline_provider_ibfk_2` FOREIGN KEY (`provider_id`) REFERENCES `provider` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `scheme_provider`
--
ALTER TABLE `scheme_provider`
  ADD CONSTRAINT `scheme_provider_ibfk_1` FOREIGN KEY (`scheme_id`) REFERENCES `scheme` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `scheme_provider_ibfk_2` FOREIGN KEY (`provider_id`) REFERENCES `provider` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;