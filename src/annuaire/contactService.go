package annuaire

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func create(id int, firstname, lastname, phone string) Contact {
	return Contact{
		id,
		firstname,
		lastname,
		phone,
	}
}

func list() []Contact {
	return []Contact{}
}

func loadContacts(filename string) (interface{}, interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func delete(id int) {
	filename := "contacts.json"

	contactsInterface, err := loadContacts(filename)
	if err != nil {
		fmt.Println("Erreur de chargement des contacts:", err)
		return
	}

	contacts := contactsInterface.([]Contact)
	for index, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
}

func (c *Contact) Update(firstName, lastName, phone string) {
	if firstName != "" {
		c.firstname = firstName
	}
	if lastName != "" {
		c.lastname = lastName
	}
	if phone != "" {
		c.phone = phone
	}
}
