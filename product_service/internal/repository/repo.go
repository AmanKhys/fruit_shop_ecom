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

func NewProductRepo(q *sqlc.Queries, db *sql.DB) usecase.ProductRepo {
	return &productRepo{q: q, db: db}
}

func (r *productRepo) GetFilteredProducts(ctx context.Context, min, max float64) ([]domain.Product, error) {
	products, err := r.q.GetFilteredProducts(ctx, r.db, sqlc.GetFilteredProductsParams{Min: min, Max: max})
	if err != nil {
		return nil, err
	}
	var respProducts = make([]domain.Product, len(products))
	for _, p := range products {
		price := p.Price.(float64)
		respProducts = append(respProducts, domain.Product{ID: int(p.ID), Name: p.Name, Price: price, Stock: int(p.Stock)})
	}
	return respProducts, nil
}

func (r *productRepo) GetProductByID(ctx context.Context, id int) (domain.Product, error) {
	p, err := r.q.GetProductByID(ctx, r.db, int64(id))
	if err != nil {
		return domain.Product{}, err
	}
	price := p.Price.(float64)
	return domain.Product{ID: int(p.ID), Name: p.Name, Price: price, Stock: int(p.Stock)}, nil
}

func (r *productRepo) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := r.q.GetProducts(ctx, r.db)
	if err != nil {
		return nil, err
	}
	var respProducts = make([]domain.Product, len(products))
	for _, p := range products {
		price := p.Price.(float64)
		respProducts = append(respProducts, domain.Product{ID: int(p.ID), Name: p.Name, Price: price, Stock: int(p.Stock)})
	}
	return respProducts, nil
}

func (r *productRepo) GetAllProductsForAdmin(ctx context.Context, userID int) ([]domain.Product, error) {
	products, err := r.q.GetProductsForAdmin(ctx, r.db)
	if err != nil {
		return nil, err
	}
	var respProducts = make([]domain.Product, len(products))
	for _, p := range products {
		price := p.Price.(float64)
		respProducts = append(respProducts, domain.Product{ID: int(p.ID), Name: p.Name, Price: price, Stock: int(p.Stock), IsDeleted: p.Isdeleted})
	}
	return respProducts, nil
}

func (r *productRepo) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	product, err := r.q.CreateProduct(ctx, r.db, sqlc.CreateProductParams{Name: p.Name, Price: p.Price, Stock: int64(p.Stock)})
	if err != nil {
		return domain.Product{}, err
	}
	price := product.Price.(float64)
	return domain.Product{ID: int(product.ID), Name: product.Name, Price: price, Stock: int(product.Stock)}, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	product, err := r.q.UpdateProductByID(ctx, r.db, sqlc.UpdateProductByIDParams{ID: int64(p.ID), Name: p.Name, Price: p.Price, Stock: int64(p.Stock)})
	if err != nil {
		return domain.Product{}, err
	}
	price := product.Price.(float64)
	return domain.Product{ID: int(product.ID), Name: product.Name, Price: price, Stock: int(product.Stock)}, nil
}
