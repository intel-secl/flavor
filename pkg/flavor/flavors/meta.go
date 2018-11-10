package flavors

import (
)

/**
 *
 * @author purvades
 */

type Meta struct {  
    Schema      Schema      `json:"schema"`
    ID          string      `json:"id"`
    Author      Author      `json:"author"`
    Realm       string      `json:"realm"`
    Description Description `json:"description"`
}

type Schema struct {  
    URI         string      `json:"uri"`
}

type Author struct {  
    Email       string      `json:"email"`
}
