package flavor

/**
 *
 * @author purvades
 */

type Meta struct {
	Schema      *Schema      `json:"schema,omitempty"`
	ID          string       `json:"id,omitempty"`
	Author      *Author      `json:"author,omitempty"`
	Realm       string       `json:"realm,omitempty"`
	Description *Description `json:"description,omitempty"`
}

type Schema struct {
	URI string `json:"uri,omitempty"`
}

type Author struct {
	Email string `json:"email,omitempty"`
}
