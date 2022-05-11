package repository

import (
	"context"
	"fmt"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	single_transaction "github.com/paul-ss/pgram-backend/internal/pkg/database/single-transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	err := single_transaction.InitDriver("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	code := m.Run()

	if err = single_transaction.TeardownDriver(); err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(code)
}

func TestStore(t *testing.T) {
	conn, err := single_transaction.Open()
	require.Nil(t, err)
	defer func() {
		assert.Nil(t, conn.Close(context.Background()))
	}()

	repo := &Repository{conn}
	req := &domain.PostStoreR{
		UserId:  1,
		GroupId: 2,
		Content: "content",
		Created: time.Now().Truncate(time.Microsecond),
		Image:   "image",
	}

	var postId int64
	err = repo.db.QueryRow(context.Background(), "select count(*) + 1 from posts").Scan(&postId)
	require.Nil(t, err)

	expected := &domain.Post{
		Id:      postId,
		UserId:  req.UserId,
		GroupId: req.GroupId,
		Content: req.Content,
		Created: req.Created,
		Image:   req.Image,
	}

	res, err := repo.Store(context.Background(), req)
	require.Nil(t, err)

	assert.Equal(t, *expected, *res)

	got := &domain.Post{}
	err = repo.db.QueryRow(context.Background(),
		`select id, user_id, group_id, content, created, image
			from posts
			where created = $1`, expected.Created).Scan(&got.Id, &got.UserId, &got.GroupId, &got.Content, &got.Created, &got.Image)
	require.Nil(t, err)

	assert.Equal(t, *expected, *got)
}
