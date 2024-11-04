package repositories

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/movierepo"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/sessionrepo"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/threaterrepo"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories/userrepo"
	"github.com/gocraft/dbr/v2"
)

type RepoProvider interface {
	Movie() ports.MovieRepository
	Session() ports.SessionRepository
	Threater() ports.ThreaterRepository
	User() ports.UserRepository
}

type Provider struct {
	MovieRepository    ports.MovieRepository
	SessionRepository  ports.SessionRepository
	ThreaterRepository ports.ThreaterRepository
	UserRepository     ports.UserRepository
}

func (p Provider) Movie() ports.MovieRepository       { return p.MovieRepository }
func (p Provider) Session() ports.SessionRepository   { return p.SessionRepository }
func (p Provider) Threater() ports.ThreaterRepository { return p.ThreaterRepository }
func (p Provider) User() ports.UserRepository         { return p.UserRepository }

func New(dbrConn *dbr.Session) *Provider {
	return &Provider{
		MovieRepository:    movierepo.New(dbrConn),
		SessionRepository:  sessionrepo.New(dbrConn),
		ThreaterRepository: threaterrepo.New(dbrConn),
		UserRepository:     userrepo.New(dbrConn),
	}
}
