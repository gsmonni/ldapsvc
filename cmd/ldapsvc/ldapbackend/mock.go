package ldapbackend

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"log"
	"math/rand"
	"strings"
)

func GenerateMockData(N uint16) *Results {
	if N == 0 {
		return nil
	}
	l := make(Results, 0)
	log.Printf("generating %d entries for ldap-mock-data", N)
	for i := 0; uint16(i) < N; i++ {
		name := mockFirstNames[rand.Intn(len(mockFirstNames))]
		lastname := mockLastNames[rand.Intn(len(mockLastNames))]
		d := QueryResult{
			CN:       fmt.Sprintf("%s.%s@example.com", name, lastname),
			UID:      fmt.Sprintf("LDAP-User-%d", i),
			ClientId: uuid.NewString(),
			Country:  mockCountries[rand.Intn(len(mockCountries))],
			Groups: []string{
				mockGroups[rand.Intn(len(mockGroups))], mockGroups[rand.Intn(len(mockGroups))],
			},
			Roles: []string{
				mockRoles[rand.Intn(len(mockRoles))], mockGroups[rand.Intn(len(mockRoles))],
			},
			FirstName: name,
			LastName:  lastname,
		}
		l = append(l, d)
	}
	return &l
}

func SaveResult(r *Results, fn string) error {
	if r == nil {
		return fmt.Errorf("empty result")
	}
	if err := common.SaveJson(fn, *r); err != nil {
		return fmt.Errorf("error saving data (%v)", err.Error())
	}
	log.Printf("saved ldap-mock-data to %s", fn)

	return nil
}

func QueryMockData(q string, r *Results) (*Results, error) {
	if r == nil {
		return nil, fmt.Errorf("ldap-data is empty")
	}
	id := strings.Split(q, "client-id=")
	l := make(Results, 0)
	switch len(id) {
	case 0:
		return nil, fmt.Errorf("invalid query")
	case 1:
		return &l, nil
	default:
		k := id[1]
		for _, d := range *(r) {
			if d.ClientId == k {
				l = append(l, d)
			}
		}
		return &l, nil
	}
}
