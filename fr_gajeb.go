package nameinfo

import "fmt"

type FrGajeb struct{}

func (_ FrGajeb) Name() string { return "identité de gajeb" }
func (_ FrGajeb) Generate(firstname, lastname string) string {
	return fmt.Sprintf("%s %s",
		[]string{
			"une pomme de terre",
			"une tasse de café",
			"une fraise",
			"un tiramitsu",
			"un caisson de basse",
			"un coquelicot",
			"une cuillère",
			"un soda",
			"une cravate",
			"des lacets",
			"des lunettes",
			"une casquette",
			"un bracelet",
			"une langue",
			"une lampe",
			"une lentille de contact",
			"une touche de piano",
			"un crayon",
			"une fermeture éclaire",
			"un pantalon",
			"une paquerette",
			"un micro",
			"un kazoo",
			"un doigt de pied",
			"des chaussons",
			"une braise",
		}[firstLetterIdx(firstname)],
		[]string{
			"soyeux/soyeuse",
			"élegant/élegante",
			"qui danse le mia",
			"au sel",
			"sur la cheminée",
			"qui appartenait à Louis",
			"pour ramasser des esgargots",
			"avec une truite",
			"touchant/touchante",
			"doux/douce",
			"poli/polie",
			"à absynthe",
			"bien rempli/remplie",
			"un peu emrobé/emrobée",
			"qui a le gout de mousse",
			"sucré/sucrée",
			"gras/grasse",
			"acide",
			"fade",
			"rose",
			"qui sonne bien",
			"très très lourd/lourde",
			"de ta soeur",
			"qui siffle dans les oreilles",
			"qui déraille",
			"indélibile",
		}[firstLetterIdx(lastname)],
	)
}

func init() {
	Generators = append(Generators, FrGajeb{})
}
