package repository

import (
	"context"
	"database/sql"
	"product_service/internal/domain"
	"product_service/internal/infrastructure/db/sqlc"
	"product_service/internal/usecase"
)

type productRepo struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewProductRepo(db *sql.DB, q *sqlc.Queries) usecase.ProductRepo {
	return &productRepo{q: q, db: db}
}

func (r *productRepo) GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error) {
	products, err := r.q.GetFilteredProducts(ctx, r.db, sqlc.GetFilteredProductsParams{Min: min, Max: max})
	if err != nil {
		return nil, err
	}
	var respProducts []domain.Product
	for _, p := range products {
		respProducts = append(respProducts, domain.Product{ID: int(p.ID), Name: p.Name, Price: p.Price, Stock: int(p.Stock)})
	}
	return respProducts, nil
}

func (r *productRepo) GetProductByID(ctx context.Context, id int) (domain.Product, error) {
	p, err := r.q.GetProductByID(ctx, r.db, int64(id))
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{ID: int(p.ID), Name: p.Name, Price: p.Price, Stock: int(p.Stock)}, nil
}

func (r *productRepo) GetProductByIDForAdmin(ctx context.Context, id int) (domain.Product, error) {
	p, err := r.q.GetProductByIDForAdmin(ctx, r.db, int64(id))
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{ID: int(p.ID), Name: p.Name, Price: p.Price, Stock: int(p.Stock), IsDeleted: p.Isdeleted}, nil
}

func (r *productRepo) GetAllProductsForAdmin(ctx context.Context, min, max float64) ([]domain.Product, error) {
	products, err := r.q.GetProductsForAdmin(ctx, r.db, sqlc.GetProductsForAdminParams{Min: min, Max: max})
	if err != nil {
		return nil, err
	}
	var respProducts []domain.Product
	for _, p := range products {
		respProducts = append(respProducts, domain.Product{ID: int(p.ID), Name: p.Name, Price: p.Price, Stock: int(p.Stock), IsDeleted: p.Isdeleted})
	}
	return respProducts, nil
}

func (r *productRepo) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	product, err := r.q.CreateProduct(ctx, r.db, sqlc.CreateProductParams{Name: p.Name, Price: p.Price, Stock: int64(p.Stock)})
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{ID: int(product.ID), Name: product.Name, Price: p.Price, Stock: int(product.Stock)}, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	product, err := r.q.UpdateProductByID(ctx, r.db, sqlc.UpdateProductByIDParams{ID: int64(p.ID), Name: p.Name, Price: p.Price, Stock: int64(p.Stock)})
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{ID: int(product.ID), Name: product.Name, Price: p.Price, Stock: int(product.Stock)}, nil
}

func (r *productRepo) DeleteProductByID(ctx context.Context, id int) error {
	return r.q.DeleteProductByID(ctx, r.db, int64(id))
}
