package goapmongo

import (
	"github.com/go-ap/activitypub"
	"github.com/openshift/osin"
)

// FullStorage collects the interfaces required for a go-ap app to
// use a MongoDB database.
type FullStorage interface {

	/******************************************
	 * ClientSaver Interface
	 ******************************************/

	// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
	UpdateClient(c osin.Client) error

	// CreateClient stores the client in the database and returns an error, if something went wrong.
	CreateClient(c osin.Client) error

	// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
	RemoveClient(id string) error

	/******************************************
	 * ClientLister Interface
	 ******************************************/

	// ListClients lists existing clients
	ListClients() ([]osin.Client, error)

	// GetClient loads a single client by id
	GetClient(id string) (osin.Client, error)

	/******************************************
	 * osin.Storage Interface
	 ******************************************/

	// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
	// to avoid concurrent access problems.
	// This is to avoid cloning the connection at each method access.
	// Can return itself if not a problem.
	Clone() osin.Storage

	// Close the resources the Storage potentially holds (using Clone for example)
	Close()

	// SaveAuthorize saves authorize data.
	SaveAuthorize(*osin.AuthorizeData) error

	// LoadAuthorize looks up AuthorizeData by a code.
	// Client information MUST be loaded together.
	// Optionally can return error if expired.
	LoadAuthorize(code string) (*osin.AuthorizeData, error)

	// RemoveAuthorize revokes or deletes the authorization code.
	RemoveAuthorize(code string) error

	// SaveAccess writes AccessData.
	// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
	SaveAccess(*osin.AccessData) error

	// LoadAccess retrieves access data by token. Client information MUST be loaded together.
	// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
	// Optionally can return error if expired.
	LoadAccess(token string) (*osin.AccessData, error)

	// RemoveAccess revokes or deletes an AccessData.
	RemoveAccess(token string) error

	// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
	// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
	// Optionally can return error if expired.
	LoadRefresh(token string) (*osin.AccessData, error)

	// RemoveRefresh revokes or deletes refresh AccessData.
	RemoveRefresh(token string) error

	/******************************************
	 * processing.Store Interface
	 ******************************************/

	// Load returns an Item or an ItemCollection from an IRI
	Load(activitypub.IRI) (activitypub.Item, error)

	// Save saves the incoming ActivityStreams Object, and returns it together with any properties
	// populated by the method's side effects. (eg, Published property can point to the current time, etc.).
	Save(activitypub.Item) (activitypub.Item, error)

	// Delete deletes completely from storage the ActivityStreams Object
	Delete(activitypub.Item) error

	/******************************************
	 * st.PasswordChanger Interface
	 ******************************************/
	PasswordSet(activitypub.Item, []byte) error
	PasswordCheck(activitypub.Item, []byte) error

	/******************************************
	 * other FedBox interfaces we might need

	CreateService(activitypub.Service) error
	LoadMetadata(activitypub.IRI) (*Metadata, error)
	SaveMetadata(Metadata, activitypub.IRI) error
	SaveNaturalLanguageValues(activitypub.NaturalLanguageValues) error
	SaveMimeTypeContent(activitypub.MimeType, activitypub.NaturalLanguageValues) error
	Reset()
	IsLocalIRI(i activitypub.IRI) bool

	 ******************************************/
}
