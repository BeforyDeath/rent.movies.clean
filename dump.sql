DROP SCHEMA IF EXISTS movies CASCADE;
CREATE SCHEMA movies;

CREATE AGGREGATE movies.array_accum ( SFUNC = array_append, BASETYPE = ANYELEMENT, STYPE = ANYARRAY, INITCOND = '{}');

CREATE SEQUENCE movies.customer_id_seq;
CREATE TABLE movies.customer (
  id       INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('movies.customer_id_seq'),
  login    VARCHAR(64)         NOT NULL,
  pass     VARCHAR(64)         NOT NULL,
  name     VARCHAR(64),
  age      INTEGER,
  phone    VARCHAR(64),
  createAt TIMESTAMP           NOT NULL
);

CREATE SEQUENCE movies.genre_id_seq;
CREATE TABLE movies.genre (
  id   INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('movies.genre_id_seq'),
  name VARCHAR(64)         NOT NULL
);

CREATE SEQUENCE movies.movie_id_seq;
CREATE TABLE movies.movie (
  id          INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('movies.movie_id_seq'),
  name        VARCHAR(128)        NOT NULL,
  description TEXT,
  year        INTEGER             NOT NULL
);

CREATE SEQUENCE movies.movie_genre_id_seq;
CREATE TABLE movies.movie_genre (
  id      INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('movies.movie_genre_id_seq'),
  movieId INTEGER,
  genreId INTEGER,
  CONSTRAINT movie_genre_movie_id_fk FOREIGN KEY (movieId) REFERENCES movies.movie (id),
  CONSTRAINT movie_genre_genre_id_fk FOREIGN KEY (genreId) REFERENCES movies.genre (id)
);

CREATE SEQUENCE movies.rent_id_seq;
CREATE TABLE movies.rent (
  id         INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('movies.rent_id_seq'),
  customerId INTEGER,
  movieId    INTEGER,
  active     BOOLEAN             NOT NULL,
  createAt   TIMESTAMP           NOT NULL,
  closeAt    TIMESTAMP,
  CONSTRAINT rent_customer_id_fk FOREIGN KEY (customerId) REFERENCES movies.customer (id),
  CONSTRAINT rent_movie_id_fk FOREIGN KEY (movieId) REFERENCES movies.movie (id)
);

INSERT INTO movies.genre (id, name) VALUES (1, 'фантастика');
INSERT INTO movies.genre (id, name) VALUES (2, 'комедия');
INSERT INTO movies.genre (id, name) VALUES (3, 'приключение');
INSERT INTO movies.genre (id, name) VALUES (4, 'драма');
INSERT INTO movies.genre (id, name) VALUES (5, 'сериал');
INSERT INTO movies.genre (id, name) VALUES (6, 'боевик');

INSERT INTO movies.movie (id, name, description, year) VALUES (1, 'Планета Ка-Пэкс / K-PAX', 'Фантастическая драма «Планета Ка-Пэкс» снята по мотивам одноименного романа американского писателя Джин Брюэр. В центре сюжета картины загадочный мужчина, называющий себя Прот. Он появился из неоткуда.', 2001);
INSERT INTO movies.movie (id, name, description, year) VALUES (2, 'Смерч / Twister', 'Главная героиня фильма-катастрофы "Смерч" ("Twister") Джо работает метеорологом. Столь необычную профессию она выбрала неслучайно: когда-то торнадо убило ее отца. Вместе с мужем Биллом она разрабатывает новые методики прогнозирования урагана.', 1998);
INSERT INTO movies.movie (id, name, description, year) VALUES (3, 'Остров / The Island', 'Остросюжетный фантастический боевик «Остров», с Юэн Макгрегором и Скарлетт Йоханссон в главных ролях, снят известным голливудским кинорежиссером Майклом Бэйем. Действия фильма разворачиваются в 2019 году, после того как на Земле произошла глобальная катастрофа.', 2005);
INSERT INTO movies.movie (id, name, description, year) VALUES (4, 'Автостопом по галактике / The Hitchhiker''s Guide to the Galaxy', 'Художественный фильм Автостопом по галактике выпущен в прокат в 2005 г. Картина является экранизацией первой книги английского писателя серии научно-фантастических произведений Дугласа Адамса, вложившего в них тонкий сарказм и сатиру. Фильм, снятый как приключенческая комедия, показывает, если мы во Вселенной не одни, то какой может быть наша жизнь.', 2005);
INSERT INTO movies.movie (id, name, description, year) VALUES (5, 'Идиократия / Idiocracy', 'Можете быть спокойны, человечеству и через пятьсот лет ничего не угрожает — только глупость (в прямом смысле). Именно такой мир видит герой фильма "Идиократия" (Idiocracy) Джо Бауэрс.', 2006);
INSERT INTO movies.movie (id, name, description, year) VALUES (6, 'Клиника / Scrubs', 'Комедийно-драматический телесериал «Клиника» рассказывает о трудовых буднях молодого врача-интерна Джона Дорина и его приятеля Кристофера Терка.', 2001);
INSERT INTO movies.movie (id, name, description, year) VALUES (7, 'Мир Дикого запада / Westworld', 'Являясь адаптацией одноименного полнометражного фильма, фантастический сериал перенесёт нас в недалёкое будущее, где одним из самых популярных развлечений станут тематические парки, реалистично воссоздающие достаточно популярные периоды в истории с помощью специальных роботов, полностью имитирующих внешний вид и поведение людей того времени.', 2016);

INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (1, 1, 1);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (2, 1, 4);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (3, 2, 3);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (4, 2, 4);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (5, 3, 1);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (6, 3, 3);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (7, 3, 6);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (8, 4, 1);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (9, 4, 3);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (10, 4, 2);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (11, 5, 1);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (12, 5, 3);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (13, 5, 2);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (14, 6, 2);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (15, 6, 4);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (16, 6, 5);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (17, 7, 1);
INSERT INTO movies.movie_genre (id, movieId, genreId) VALUES (18, 7, 5);
