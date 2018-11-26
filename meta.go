package flavor

/**
 *
 * @author purvades
 */

// Meta defines metadata info about the flavor
type Meta struct {
	Schema      *Schema      `json:"schema,omitempty"`
	ID          string       `json:"id,omitempty"`
	Author      *Author      `json:"author,omitempty"`
	Realm       string       `json:"realm,omitempty"`
	Description *Description `json:"description,omitempty"`
}

// Schema contains the URI to the schema
type Schema struct {
	URI string `json:"uri,omitempty"`
}

// Author contains the email address of the author
type Author struct {
	Email string `json:"email,omitempty"`
}
