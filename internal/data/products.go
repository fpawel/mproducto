package data

import (
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type ProductsCategoryNode struct {
	ID       int                     `json:"id"`
	Name     string                  `json:"name"`
	Level    int                     `json:"level"`
	Nodes    []*ProductsCategoryNode `json:"nodes,omitempty"`
	Products []*Product              `json:"products,omitempty"`
}

type Product struct {
	ID              int    `json:"id" db:"-"`
	Name2           string `json:"name2" db:"name2"`
	HierarchyLevel4 string `json:"level4" db:"hierarchy_level4"`
	HierarchyLevel5 string `json:"level5" db:"hierarchy_level5"`
	HierarchyLevel6 string `json:"level6" db:"hierarchy_level6"`
	HierarchyLevel7 string `json:"level7" db:"hierarchy_level7"`
}

func GetProductsCategoryNodes(db *sqlx.DB, ) (result [] *ProductsCategoryNode) {
	var level1Names []string
	if err := db.Select(&level1Names,
		`
SELECT DISTINCT lower(hierarchy_level1) 
FROM product 
ORDER BY 1`, ); err != nil {
		log.Fatal(err)
	}
	var id int
	for _, level1Name := range level1Names {
		level1Name = strings.Title(level1Name)
		id++
		level1 := &ProductsCategoryNode{
			Name: level1Name,
			ID:   id,
		}

		result = append(result, level1)

		var level2Names []string
		if err := db.Select(&level2Names,
			`
SELECT DISTINCT lower(hierarchy_level2) 
FROM product 
WHERE lower(hierarchy_level1) = lower($1) 
ORDER BY 1`, level1Name); err != nil {
			log.Fatal(err)
		}
		for _, level2Name := range level2Names {
			level2Name := strings.Title(level2Name)
			id++
			level2 := &ProductsCategoryNode{
				Name:  level2Name,
				ID:    id,
				Level: 1,
			}

			level1.Nodes = append(level1.Nodes, level2)

			var level3Names []string
			if err := db.Select(&level3Names,
				`
SELECT DISTINCT lower(subgroup) 
FROM product 
WHERE lower(hierarchy_level1) = lower($1) AND lower(hierarchy_level2) = lower($2) 
ORDER BY 1`, level1Name, level2Name); err != nil {
				log.Fatal(err)
			}
			for _, level3Name := range level3Names {
				level3Name = strings.Title(level3Name)
				id++
				level3 := &ProductsCategoryNode{
					Name:  level3Name,
					ID:    id,
					Level: 2,
				}
				level2.Nodes = append(level2.Nodes, level3)

				if err := db.Select(&level3.Products,
					`
SELECT name2, hierarchy_level4, hierarchy_level5, hierarchy_level6, hierarchy_level7 
FROM product 
WHERE lower(hierarchy_level1) = lower($1) 
  AND lower(hierarchy_level2) = lower($2)
  AND lower(subgroup) = lower($3)
ORDER BY 1`, level1Name, level2Name, level3Name); err != nil {
					log.Fatal(err)
				}
				for _,x := range level3.Products{
					id++
					x.ID = id
				}
			}
		}
	}
	return
}
