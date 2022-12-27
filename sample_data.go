package main

import (
	"github.com/jmcerezo/corporate-america/models"
	"github.com/uadmin/uadmin"
)

func SampleData() {
	tag_type := []models.TagType{}
	uadmin.All(&tag_type)

	default_tagType := []models.TagType{
		{
			Name: "Log in",
		},
		{
			Name: "Maintenance",
		},
	}

	for i := range default_tagType {
		isExisting := false

		for j := range tag_type {
			if default_tagType[i].Name == tag_type[j].Name {
				isExisting = true
				break
			}
		}

		if !isExisting {
			uadmin.Save(&default_tagType[i])
		}
	}
}
