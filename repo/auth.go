package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type AuthRepoItf interface {
	IsEmailExist(ctx context.Context, email string) (bool, error)
	IsIdExist(ctx context.Context, credId int64) (bool, error)
	GetCredential(ctx context.Context, email string) (*entity.Credential, error)
	GetCredentialById(ctx context.Context, credId int64) (*entity.Credential, error)
	InsertCredential(ctx context.Context, cred *entity.Credential) (int64, error)
	DeleteCredential(ctx context.Context, c *entity.Credential) error
	UpdateCredential(ctx context.Context, cred *entity.Credential) error
	UpdateVerified(ctx context.Context, credId int64) error
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepoItf {
	return &AuthRepoImpl{db: db}
}

func (r *AuthRepoImpl) IsEmailExist(ctx context.Context, email string) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("credentials").
		Where(
			sb.Equal("email", email),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *AuthRepoImpl) IsIdExist(ctx context.Context, credId int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("credentials").
		Where(
			sb.Equal("id", credId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *AuthRepoImpl) GetCredential(ctx context.Context, email string) (*entity.Credential, error) {
	query := `SELECT c.id, c.password, r.id FROM credentials c 
    LEFT JOIN roles r ON r.id = c.role_id 
	WHERE c.email = $1 AND c.deleted_at IS NULL AND r.deleted_at IS NULL`

	cred := entity.Credential{Email: email}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&cred.Id, &cred.Password, &cred.RoleId)
	if err != nil {
		return nil, err
	}

	return &cred, nil
}

func (r *AuthRepoImpl) GetCredentialById(ctx context.Context, credId int64) (*entity.Credential, error) {
	query := `SELECT c.email, c.password, r.id FROM credentials c 
    LEFT JOIN roles r ON r.id = c.role_id 
	WHERE c.id = $1 AND c.deleted_at IS NULL AND r.deleted_at IS NULL`

	cred := entity.Credential{Id: credId}
	err := r.db.QueryRowContext(ctx, query, credId).Scan(&cred.Email, &cred.Password, &cred.RoleId)
	if err != nil {
		return nil, err
	}

	return &cred, nil
}

func (r *AuthRepoImpl) InsertCredential(ctx context.Context, cred *entity.Credential) (int64, error) {
	query := `INSERT INTO credentials (email, password, role_id) VALUES ($1, $2, $3) RETURNING id`

	var credId int64
	tx := extractTx(ctx)
	err := tx.QueryRowContext(ctx, query, cred.Email, cred.Password, cred.RoleId).Scan(&credId)
	if err != nil {
		return 0, err
	}

	return credId, nil
}

func (r *AuthRepoImpl) DeleteCredential(ctx context.Context, c *entity.Credential) error {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	ub.Update("credentials").
		Set(ub.Assign("deleted_at", "NOW()")).
		Where(
			ub.Equal("id", c.Id),
			ub.IsNull("deleted_at"),
		)

	tx := extractTx(ctx)
	query, args := ub.Build()
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepoImpl) UpdateCredential(ctx context.Context, cred *entity.Credential) error {
	query := `UPDATE credentials SET password = $1, updated_at = NOW() WHERE email = $2 AND deleted_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, cred.Password, cred.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepoImpl) UpdateVerified(ctx context.Context, credId int64) error {
	query := `UPDATE users u SET is_verified = TRUE, updated_at = NOW() FROM credentials c 
  	WHERE c.id = u.credential_id AND c.id = $1 
   	AND u.deleted_at IS NULL AND c.deleted_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, credId)
	if err != nil {
		return err
	}

	return nil
}
