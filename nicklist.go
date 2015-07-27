package main

import (
	"container/list"
	"fmt"
)

// nicks contains the list of recently used nicknames
var nicks list.List

// PushNick adds a nick to the list
func PushNick(nick string) {
	//delete the nick if it exists
	//push the nick to the head of the list
	nicks.PushBack(nick)
}

// NickCompleter will be used by readline to get possible nicks
func NickCompleter(query, ctx string) []string {
	var compls []string

	fmt.Printf("\nquery: %q\nctx: %q\n", query, ctx)

	for e := nicks.Front(); e != nil; e = e.Next() {
		fmt.Printf("\nadding %s\n", "/m "+e.Value.(string))
		compls = append(compls, "/m "+e.Value.(string))
	}
	fmt.Printf("\nList is:\n%q\n", compls)
	return compls
}
