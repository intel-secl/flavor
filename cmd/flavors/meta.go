package flavors

import (
)

/**
 *
 * @author purvades
 */

type Meta struct {  
    Schema      Schema
    ID          string
    Author      Author
    Realm       string
    Description Description
}

type Schema struct {  
    URI         string
}

type Author struct {  
    Email       string
}
