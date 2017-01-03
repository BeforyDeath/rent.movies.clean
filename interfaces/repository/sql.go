package repository

var sql = map[string]string{
	`genreOne`: `SELECT id, name FROM movies.genre WHERE id = $1`,
	`genreAll`: `SELECT id, name FROM movies.genre LIMIT $1 OFFSET $2`,

	`movieOne`:    `SELECT id, name FROM movies.movie WHERE id = $1`,
	`movieSelect`: `SELECT m.id, m.name, m.year, array_to_string(movies.array_accum(g.name), ', ') as genre, m.description `,
	`movieCount`:  `SELECT COUNT(DISTINCT (m.id)) `,
	`movieFrom`:   `FROM movies.movie m, movies.movie_genre mg, movies.genre g WHERE mg.movieId = m.id AND mg.genreId = g.id `,
	`movieYear`:   `AND m.year = $4 `,
	`movieGenre`:  `AND m.id IN (SELECT m.id FROM movies.movie m, movies.movie_genre mg, movies.genre g WHERE mg.movieId = m.id AND mg.genreId = g.id AND g.name = $3 GROUP BY m.id) `,
	`movieGroup`:  `GROUP BY m.id `,
	`movieLimit`:  `ORDER BY m.id DESC LIMIT $1 OFFSET $2`,

	`customerValid`:  `SELECT id FROM movies."user" WHERE login = $1`,
	`customerLogin`:  `SELECT id, login, name, age, phone, createat FROM movies."user" WHERE login=$1 AND pass=$2`,
	`customerInsert`: `INSERT INTO movies."user" (login, pass, name, age, phone, createat) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,

	`rentValid`:  `SELECT createAt FROM movies.rent WHERE userId = $1 AND movieId = $2 AND active = $3`,
	`rentInsert`: `INSERT INTO movies.rent (userId, movieId, active, createAt) VALUES ($1, $2, $3, $4) RETURNING id`,
	`rentUpdate`: `UPDATE movies.rent SET active=$3, closeAt=$4 WHERE active = true AND userId = $1 AND movieId = $2`,
	`rentSelect`: `SELECT r.active, r.createAt, r.closeAt, m.id, m.name, m.year, array_to_string(movies.array_accum(g.name), ', ') AS genre, m.description `,
	`rentCount`:  `SELECT COUNT(DISTINCT (r.id)) `,
	`rentFrom`:   `FROM movies.rent r, movies.movie m, movies.movie_genre mg, movies.genre g WHERE r.movieId = m.id AND mg.movieId = m.id AND mg.genreId = g.id AND r.userId = $3 AND r.active = $4 `,
	`rentGroup`:  `GROUP BY r.id, m.id `,
	`rentLimit`:  `LIMIT $1 OFFSET $2`,
}
