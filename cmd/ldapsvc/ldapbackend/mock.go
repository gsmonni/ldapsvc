package ldapbackend

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"log"
	"math/rand"
	"os"
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
		log.Printf("id %d %s, %s", i, name, lastname)
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

func SaveResult(r *Results, fn string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error while saving %s (%v)", fn, err.Error())
		}
	}()

	if r == nil {
		return fmt.Errorf("empty ldap data")
	}
	if common.Datapath == "" {
		common.Datapath = DefDataDir
	}
	if !common.IsDir(common.Datapath) {
		if err := os.MkdirAll(common.Datapath, os.ModePerm); err != nil {
			return fmt.Errorf("error while creating directory %s (%v)", common.Datapath, err.Error())
		}
	}

	if err := common.SaveJson(fn, *r); err != nil {
		return fmt.Errorf("error saving JSON (%v)", err.Error())
	} else {
		log.Printf("saved ldap-mock-data to %s", fn)
	}
	return err
}

func QueryMockData(q string, r *Results) (*Results, error) {
	if r == nil {
		return nil, fmt.Errorf("ldap-data is empty")
	}
	s := strings.Split(q, "=")
	k := s[0]
	l := make(Results, 0)
	switch len(s) {
	case 0:
		return nil, fmt.Errorf("invalid query")
	case 1:
		return &l, nil
	default:
		v := s[1]
		for _, d := range *(r) {
			if d.GetPropertyValue(k) == v {
				l = append(l, d)
			}
		}
		return &l, nil
	}
}
