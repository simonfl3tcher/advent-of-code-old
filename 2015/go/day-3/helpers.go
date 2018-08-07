package main

func switcher(char string, c *counter, store *[]counter) {
	switch char {
	case "^":
		c.y++
	case "v":
		c.y--
	case "<":
		c.x--
	case ">":
		c.x++
	}
	appendIfNotInList(store, *c)
}

func appendIfNotInList(store *[]counter, c counter) {
	var exists bool
	for _, a := range *store {
		if a == c {
			exists = true
		}
	}

	if !exists {
		*store = append(*store, c)
	}
}
