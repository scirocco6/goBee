package main

// NickList holds a list of nicks always in most recently used order
type NickList struct {
	list []string
	hash map[string]int
}
