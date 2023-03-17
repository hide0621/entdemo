package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/user"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//DBと繋ぐ（SQLite）
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	//スキーマを作る
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {

	//Whereで条件付けしたselectクエリを実行
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
