package auth

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/answersuck/vault/internal/config"
	"github.com/answersuck/vault/pkg/logging"

	"github.com/answersuck/vault/internal/domain/account"
	"github.com/answersuck/vault/internal/domain/session"
)

type (
	AccountService interface {
		GetById(ctx context.Context, accountId string) (*account.Account, error)
		GetByEmail(ctx context.Context, email string) (*account.Account, error)
		GetByNickname(ctx context.Context, nickname string) (*account.Account, error)
	}

	SessionService interface {
		Create(ctx context.Context, accountId string, d session.Device) (*session.Session, error)
		Terminate(ctx context.Context, sessionId string) error
	}

	TokenManager interface {
		New(subject, audience string, expiration time.Duration) (string, error)
		Parse(token, audience string) (string, error)
	}
)

type service struct {
	log     logging.Logger
	cfg     *config.AccessToken
	token   TokenManager
	account AccountService
	session SessionService
}

type Deps struct {
	Logger         logging.Logger
	Config         *config.Aggregate
	Token          TokenManager
	AccountService AccountService
	SessionService SessionService
}

func NewService(d *Deps) *service {
	return &service{
		log:     d.Logger,
		cfg:     &d.Config.AccessToken,
		token:   d.Token,
		account: d.AccountService,
		session: d.SessionService,
	}
}

func (s *service) Login(ctx context.Context, login, password string, d session.Device) (*session.Session, error) {
	var a *account.Account

	_, err := mail.ParseAddress(login)
	if err != nil {
		a, err = s.account.GetByNickname(ctx, login)
		if err != nil {
			return nil, fmt.Errorf("authService - Login - s.account.GetByNickname: %w", err)
		}
	} else {
		a, err = s.account.GetByEmail(ctx, login)
		if err != nil {
			return nil, fmt.Errorf("authService - Login - s.account.GetByEmail: %w", err)
		}
	}

	if err = a.ComparePasswords(password); err != nil {
		return nil, fmt.Errorf("authService - Login - a.CompareHashAndPassword: %w", err)
	}

	sess, err := s.session.Create(ctx, a.Id, d)
	if err != nil {
		return nil, fmt.Errorf("authService - Login - s.session.Create: %w", err)
	}

	return sess, nil
}

func (s *service) NewToken(ctx context.Context, accountId, password, audience string) (string, error) {
	a, err := s.account.GetById(ctx, accountId)
	if err != nil {
		return "", fmt.Errorf("authService - NewSecurityToken - s.account.GetByID: %w", err)
	}

	if err = a.ComparePasswords(password); err != nil {
		return "", fmt.Errorf("authService - NewSecurityToken - a.CompareHashAndPassword: %w", err)
	}

	t, err := s.token.New(accountId, audience, s.cfg.Expiration)
	if err != nil {
		return "", fmt.Errorf("authService - NewSecurityToken - s.token.New: %w", err)
	}

	return t, nil
}

func (s *service) ParseToken(ctx context.Context, token, audience string) (string, error) {
	accountId, err := s.token.Parse(token, audience)
	if err != nil {
		return "", fmt.Errorf("authService - ParseSecurityToken - s.token.Parse: %w", err)
	}

	return accountId, nil
}