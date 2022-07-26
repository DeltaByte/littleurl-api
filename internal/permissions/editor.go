package permissions

import "github.com/deltabyte/littleurl-api/internal/entities"

const Editor = "editor"

type EditorRole struct {
	DomainID entities.DomainID
}

// domains
func (role *EditorRole) DomainRead(id string) bool  { return id == role.DomainID }
func (role *EditorRole) DomainWrite(id string) bool { return false }
