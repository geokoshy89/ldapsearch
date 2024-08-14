package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	fmt.Print("New golang project for ldap access")
	ldapURL := "ldap://localhost:389"
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("cn=admin,dc=example,dc=org", "admin")
	if err != nil {
		log.Fatal(err)
	}
	// addReq := ldap.NewAddRequest("CN=testgroup,ou=Groups,dc=example,dc=org", []ldap.Control{})

	// addReq.Attribute("objectClass", []string{"top", "groupOfNames"})
	// addReq.Attribute("cn", []string{"testgroup"})
	// addReq.Attribute("member", []string{"cn=testuser,dc=example,dc=com"})

	// if err := l.Add(addReq); err != nil {
	// 	log.Fatal("error adding group:", addReq, err)
	// }

	// addReq := ldap.NewAddRequest("CN=fooUser,OU=People,dc=example,dc=org", []ldap.Control{})
	// addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "person"})
	// addReq.Attribute("cn", []string{"fooUser"})
	// addReq.Attribute("sn", []string{"fooSurname"})

	// if err := l.Add(addReq); err != nil {
	// 	log.Fatal("error adding service:", addReq, err)
	// }

	// connect code comes here

	user := "fooUser"
	baseDN := "DC=example,DC=org"
	filter := fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(user))

	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{"cn"}, []ldap.Control{})

	result, err := l.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Got", len(result.Entries), "search results")
	for _, entry := range result.Entries {
		fmt.Printf("DN: %s\n", entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf("%s: %s\n", attr.Name, attr.Values)
		}
		fmt.Println()
	}
}
