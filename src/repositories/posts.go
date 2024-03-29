package repositories

import (
	"api/src/models"
	"database/sql"
)

type Post struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *Post {
	return &Post{db}
}

func (p Post) Create(post models.Post) (uint64, error) {
	statement, err := p.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (p Post) GetPost(postId uint64) (models.Post, error) {
	row, err := p.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.id = ?
	`, postId)

	if err != nil {
		return models.Post{}, err
	}

	defer row.Close()

	var post models.Post
	if row.Next() {
		if err = row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (p Post) GetPosts(userID uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		left join followers f on f.user_id = p.author_id
		where u.id = ? or f.follower_id = ?
		order by p.createdAt desc
	`, userID, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p Post) Update(postId uint64, post models.Post) error {
	statement, err := p.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

func (p Post) Delete(postId uint64) error {
	statement, err := p.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (p Post) GetPostsByUser(userID uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?
		order by p.createdAt desc
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p Post) LikePost(postId uint64) error {
	statement, err := p.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (p Post) DislikePost(postId uint64) error {
	statement, err := p.db.Prepare("update posts set likes = case when likes > 0 then likes - 1 else 0 end where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
