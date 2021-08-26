package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-hobby-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-hobby-api/internal/flusher Flusher
