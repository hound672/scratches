package main

import (
	"context"
	"log"
	"runtime"
)

type SomeModel struct {
	Field1 string
	Field2 int
}

type SomeRepo interface {
	ReadAll(ctx context.Context, limit int, offset int) ([]*SomeModel, error)
	ReadAll2(ctx context.Context, limit int, offset int)
	ReadCount(ctx context.Context) (int, error)
}

type someRepo struct {
}

func (r *someRepo) ReadAll(ctx context.Context, limit int, offset int) ([]*SomeModel, error) {
	return nil, nil
}

func (r *someRepo) ReadCount(ctx context.Context) (int, error) {
	return 0, nil
}

func (r *someRepo) ReadAll2(ctx context.Context, limit int, offset int) {

}

//

func genFunc(ctx context.Context, limit int, offset int, args ...any) {
	log.Printf("limit: %d; offset: %d; args: %v", limit, offset, args)
}

func paginator[T any, Ptr interface{ *T }](
	ctx context.Context,
	repoReader func(ctx context.Context, limit int, offset int, args ...any) ([]Ptr, error),
	args ...any,
) ([]Ptr, error){
	log.Printf("ctx: %v", ctx)
	rows, err := repoReader(ctx, 1, 2, args...)
	log.Printf("rows: %v; err: %v", rows, err)

	return rows, nil
}

func paginator2(
	ctx context.Context,
	repoReader func(ctx context.Context, limit int, offset int, args []any),
	args ...any,
) {
	//log.Printf("ctx: %v", ctx)
	//rows, err := repoReader(ctx, 1, 2, args...)
	//log.Printf("rows: %v; err: %v", rows, err)
	//
	//return rows, nil
}

//

func main() {
	log.Printf("Hello world: %v", runtime.Version())

	var repo SomeRepo
	repo = &someRepo{}

	log.Printf("repo: %v", repo)

	ctx := context.Background()
	//genFunc(ctx, 1, 2)

	//res, err := paginator(ctx, repo.ReadAll)
	//log.Printf("res: %v; err:= %v", res, err)

	paginator2(ctx, repo.ReadAll2)
}
