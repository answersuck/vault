package psql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	"github.com/answersuck/vault/internal/domain/account"

	"github.com/answersuck/vault/pkg/logging"
	"github.com/answersuck/vault/pkg/postgres"
)

const (
	accountTable                   = "account"
	accountAvatarTable             = "account_avatar"
	accountVerificationCodeTable   = "account_verification_code"
	accountPasswordResetTokenTable = "account_password_reset_token"
)

type accountRepo struct {
	l logging.Logger
	c *postgres.Client
}

func NewAccountRepo(l logging.Logger, c *postgres.Client) *accountRepo {
	return &accountRepo{
		l: l,
		c: c,
	}
}

func (r *accountRepo) Save(ctx context.Context, a *account.Account, code string) (string, error) {
	sql := fmt.Sprintf(`
		WITH a AS (
			INSERT INTO %s(nickname, email, password, is_verified, updated_at, created_at)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id AS account_id
		),
		av AS (
			INSERT INTO %s(code, account_id)
			VALUES($7, (SELECT account_id FROM a) )
		),
		aa AS (
			INSERT INTO %s(url, account_id)
			VALUES($8, (SELECT account_id FROM a) )
		)
		SELECT account_id FROM a
	`, accountTable, accountVerificationCodeTable, accountAvatarTable)

	r.l.Info("psql - account - Save: %s", sql)

	err := r.c.Pool.QueryRow(ctx, sql,
		a.Nickname,
		a.Email,
		a.PasswordHash,
		a.Verified,
		a.UpdatedAt,
		a.CreatedAt,
		a.VerificationCode,
		a.AvatarURL,
	).Scan(&a.Id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("psql - account - Save - r.c.Pool.QueryRow.Scan: %w", account.ErrAlreadyExist)
			}

		}

		return nil, fmt.Errorf("psql - account - Save - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return a, nil
}

func (r *accountRepo) FindById(ctx context.Context, accountId string) (*account.Account, error) {
	sql := fmt.Sprintf(`
		SELECT
			a.username,
			a.email,
			a.password,
			a.created_at,
			a.updated_at,
			a.is_verified,
			aa.url AS avatar_url
		FROM %s AS a
		LEFT JOIN %s AS aa
		ON aa.account_id = a.id
		WHERE
			a.id = $1
			AND a.is_archived = $2
	`, accountTable, accountAvatarTable)

	r.l.Info("psql - account - FindById: %s", sql)

	a := account.Account{Id: accountId}

	if err := r.c.Pool.QueryRow(ctx, sql, accountId, false).Scan(
		&a.Nickname,
		&a.Email,
		&a.PasswordHash,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.Verified,
		&a.AvatarURL,
	); err != nil {

		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("psql - account - FindById - r.c.Pool.QueryRow.Scan: %w", account.ErrNotFound)
		}

		return nil, fmt.Errorf("psql - account - FindById - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return &a, nil
}

func (r *accountRepo) FindByEmail(ctx context.Context, email string) (*account.Account, error) {
	sql := fmt.Sprintf(`
		SELECT
			a.id,
			a.username,
			a.password,
			a.created_at,
			a.updated_at,
			a.is_verified,
			aa.url AS avatar_url
		FROM %s AS a
		LEFT JOIN %s AS aa
		ON aa.account_id = a.id
		WHERE
			a.email = $1
			AND a.is_archived = $2
	`, accountTable, accountAvatarTable)

	r.l.Info("psql - account - FindByEmail: %s", sql)

	a := account.Account{Email: email}

	if err := r.c.Pool.QueryRow(ctx, sql, email, false).Scan(
		&a.Id,
		&a.Nickname,
		&a.PasswordHash,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.Verified,
		&a.AvatarURL,
	); err != nil {

		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("psql - account - FindByEmail - r.c.Pool.QueryRow.Scan: %w", account.ErrNotFound)
		}

		return nil, fmt.Errorf("psql - account - FindByEmail - r.client.Pool.QueryRow.Scan: %w", err)
	}

	return &a, nil
}

func (r *accountRepo) FindByNickname(ctx context.Context, username string) (*account.Account, error) {
	sql := fmt.Sprintf(`
		SELECT
			a.id,
			a.email,
			a.password,
			a.created_at,
			a.updated_at,
			a.is_verified,
			aa.url AS avatar_url
		FROM %s AS a
		LEFT JOIN %s AS aa
		ON aa.account_id = a.id
		WHERE
			a.username = $1
			AND a.is_archived = $2
	`, accountTable, accountAvatarTable)

	r.l.Info("psql - account - FindByUsername: %s", sql)

	a := account.Account{Nickname: username}

	err := r.c.Pool.QueryRow(ctx, sql, username, false).Scan(
		&a.Id,
		&a.Email,
		&a.PasswordHash,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.Verified,
		&a.AvatarURL,
	)
	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("psql - account - FindByUsername - r.c.Pool.QueryRow.Scan: %w", account.ErrNotFound)
		}

		return nil, fmt.Errorf("psql - account - FindByUsername - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return &a, nil
}

func (r *accountRepo) SetArchived(ctx context.Context, accountId string, archived bool, updatedAt time.Time) error {
	sql := fmt.Sprintf(`
		UPDATE %s
		SET
			is_archived = $1,
			updated_at = $2
		WHERE
			id = $3
			AND is_archived = $4
	`, accountTable)

	r.l.Info("psql - account - SetArchived: %s", sql)

	ct, err := r.c.Pool.Exec(ctx, sql, archived, updatedAt, accountId, !archived)
	if err != nil {
		return fmt.Errorf("psql - account - SetArchived - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - account - SetArchived - r.c.Pool.Exec: %w", account.ErrNotArchived)
	}

	return nil
}

func (r *accountRepo) Verify(ctx context.Context, code string, verified bool, updatedAt time.Time) error {
	sql := fmt.Sprintf(`
		UPDATE %s AS a
		SET
			is_verified = $1,
			updated_at = $2
		FROM (
			SELECT av.code, a.id AS account_id
			FROM %s AS av
			INNER JOIN %s AS a
			ON av.account_id = a.id
			WHERE av.code = $3
		) AS sq
		WHERE
			a.is_verified = $4
			AND a.id = sq.account_id;
	`, accountTable, accountVerificationCodeTable, accountTable)

	r.l.Info("psql - account - Verify: %s", sql)

	ct, err := r.c.Pool.Exec(ctx, sql, verified, updatedAt, code, !verified)
	if err != nil {
		return fmt.Errorf("psql - account - Verify - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - account - Verify - r.c.Pool.Exec: %w", account.ErrAlreadyVerified)
	}

	return nil
}

func (r *accountRepo) FindVerification(ctx context.Context, accountId string) (account.VerificationDTO, error) {
	sql := fmt.Sprintf(`
		SELECT a.email, a.is_verified, av.code AS verification_code
		FROM %s AS a
		LEFT JOIN %s AS av
		ON av.account_id = a.id
		WHERE a.id = $1
	`, accountTable, accountVerificationCodeTable)

	r.l.Info("psql - account - FindVerification: %s", sql)

	var v account.VerificationDTO

	if err := r.c.Pool.QueryRow(ctx, sql, accountId).Scan(
		&v.Email,
		&v.Verified,
		&v.Code,
	); err != nil {

		if err == pgx.ErrNoRows {
			return account.VerificationDTO{},
				fmt.Errorf("psql - account - FindVerification - r.c.Pool.QueryRow.Scan: %w", account.ErrVerificationNotFound)
		}

		return account.VerificationDTO{}, fmt.Errorf("psql - account - FindVerification - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return v, nil
}

func (r *accountRepo) SavePasswordResetToken(ctx context.Context, email, token string) error {
	sql := fmt.Sprintf(`
		INSERT INTO %s (token, account_id)
		VALUES($1, (SELECT id AS account_id FROM %s WHERE email = $2))
	`, accountPasswordResetTokenTable, accountTable)

	r.l.Info("psql - account - SavePasswordResetToken: %s", sql)

	if _, err := r.c.Pool.Exec(ctx, sql, token, email); err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return fmt.Errorf("psql - account - SavePasswordResetToken - r.c.Pool.Exec: %w", account.ErrPasswordResetTokenAlreadyExist)
			case pgerrcode.ForeignKeyViolation:
				return fmt.Errorf("psql - account - SavePasswordResetToken - r.c.Pool.Exec: %w", account.ErrNotFound)
			}
		}

		return fmt.Errorf("psql - account - SavePasswordResetToken - r.c.Pool.Exec: %w", err)
	}

	return nil
}

func (r *accountRepo) FindPasswordResetToken(ctx context.Context, token string) (account.PasswordResetToken, error) {
	sql := fmt.Sprintf(`
		SELECT t.token, t.created_at, a.id
		FROM %s AS t
		INNER JOIN %s AS a
		ON a.id = t.account_id
		WHERE t.token = $1
	`, accountPasswordResetTokenTable, accountTable)

	r.l.Info("psql - account - FindPasswordResetToken: %s", sql)

	var t account.PasswordResetToken

	if err := r.c.Pool.QueryRow(ctx, sql, token).Scan(&t.Token, &t.CreatedAt, &t.AccountId); err != nil {

		if err == pgx.ErrNoRows {
			return account.PasswordResetToken{},
				fmt.Errorf("psql - account - FindPasswordResetToken - r.c.Pool.QueryRow.Scan: %w", account.ErrPasswordResetTokenNotFound)
		}

		return account.PasswordResetToken{}, fmt.Errorf("psql - account - FindPasswordResetToken - r.client.Pool.QueryRow.Scan: %w", err)
	}

	return t, nil
}

func (r *accountRepo) SetPassword(ctx context.Context, dto account.SetPasswordDTO) error {
	sql := fmt.Sprintf(`
		WITH a AS (
			UPDATE %s
			SET password = $1, updated_at = $2
			WHERE id = $3
		)
		DELETE FROM %s
		WHERE
			account_id = $4
			AND token = $5
	`, accountTable, accountPasswordResetTokenTable)

	r.l.Info("psql - account - UpdatePasswordWithToken: %s", sql)

	ct, err := r.c.Pool.Exec(
		ctx,
		sql,
		dto.PasswordHash,
		dto.UpdatedAt,
		dto.AccountId,
		dto.AccountId,
		dto.Token,
	)
	if err != nil {
		return fmt.Errorf("psql - account - SetPassword - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - account - SetPassword - r.c.Pool.Exec: %w", account.ErrPasswordNotSet)
	}

	return nil
}
