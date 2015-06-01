package aat

import (
	"encoding/xml"
	"github.com/davecgh/go-spew/spew"
	"github.com/verisart/xsd/rdf"
	"github.com/verisart/xsd/rdfs"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var marshalTests = []struct {
	Value         interface{}
	ExpectXML     string
	MarshalOnly   bool
	UnmarshalOnly bool
	IsFile        bool
}{
	{
		UnmarshalOnly: true,
		IsFile:        true,
		ExpectXML:     "testdata/surrealism.rdf",
		Value: &Term{
			//XMLName: xml.Name{Space: "http://www.w3.org/1999/02/22-rdf-syntax-ns#", Local: "rdf"},
			Subject: &GVPSubject{
				About: "http://vocab.getty.edu/aat/300021512",
				Types: []*rdf.Type{
					&rdf.Type{
						Resource: "http://www.w3.org/2004/02/skos/core#Concept",
					},
					&rdf.Type{
						Resource: "http://vocab.getty.edu/ontology#Concept",
					},
				},
				Labels: []*rdfs.Label{
					&rdfs.Label{
						XsdtString: "Surrealist",
						Lang:       "en",
					},
					&rdfs.Label{
						XsdtString: "Surrealism",
						Lang:       "en",
					},
					&rdfs.Label{
						XsdtString: "Supperrealism",
						Lang:       "en",
					},
					&rdfs.Label{
						XsdtString: "超現實主義",
						Lang:       "zh-hant",
					},
					&rdfs.Label{
						XsdtString: "chāo xiàn shí zhǔ yì",
						Lang:       "zh-latn-pinyin-x-hanyu",
					},
					&rdfs.Label{
						XsdtString: "chao xian shi zhu yi",
						Lang:       "zh-latn-pinyin-x-notone",
					},
					&rdfs.Label{
						XsdtString: "ch'ao hsien shih chu i",
						Lang:       "zh-latn-wadegile",
					},
					&rdfs.Label{
						XsdtString: "surrealistisch",
						Lang:       "nl",
					},
					&rdfs.Label{
						XsdtString: "surrealisme",
						Lang:       "nl",
					},
					&rdfs.Label{
						XsdtString: "Surrealista",
						Lang:       "es",
					},
					&rdfs.Label{
						XsdtString: "surrealismo",
						Lang:       "es",
					},
				},
				BroaderTerms: []*rdf.ResourceAttr{
					&rdf.ResourceAttr{
						Resource: "http://vocab.getty.edu/aat/300021494",
					},
				},
				BroaderPreferredTerms: []*rdf.ResourceAttr{
					&rdf.ResourceAttr{
						Resource: "http://vocab.getty.edu/aat/300021494",
					},
				},
			},
			Statements: []*rdf.Statement{
				&rdf.Statement{
					&rdf.Subject{
						Resource: "http://vocab.getty.edu/aat/300021512",
					},
					&rdf.Predicate{
						Resource: "http://vocab.getty.edu/ontology#aat2812_followed",
					},
					&rdf.Object{
						Resource: "http://vocab.getty.edu/aat/300021514",
					},
				},
				&rdf.Statement{
					&rdf.Subject{
						Resource: "http://vocab.getty.edu/aat/300021512",
					},
					&rdf.Predicate{
						Resource: "http://vocab.getty.edu/ontology#aat2811_preceded",
					},
					&rdf.Object{
						Resource: "http://vocab.getty.edu/aat/300022099",
					},
				},
				&rdf.Statement{
					&rdf.Subject{
						Resource: "http://vocab.getty.edu/aat/300021512",
					},
					&rdf.Predicate{
						Resource: "http://vocab.getty.edu/ontology#aat2552_reflected_in-produces",
					},
					&rdf.Object{
						Resource: "http://vocab.getty.edu/aat/300182745",
					},
				},
			},
		},
	},
	{
		UnmarshalOnly: true,
		IsFile:        true,
		ExpectXML:     "testdata/materials_by_function.rdf",
		Value: &Term{
			//XMLName: xml.Name{Space: "http://www.w3.org/1999/02/22-rdf-syntax-ns#", Local: "rdf"},
			Subject: &GVPSubject{
				About: "http://vocab.getty.edu/aat/300014692",
				Types: []*rdf.Type{
					&rdf.Type{
						Resource: "http://vocab.getty.edu/ontology#GuideTerm",
					},
					&rdf.Type{
						Resource: "http://purl.org/iso25964/skos-thes#ThesaurusArray",
					},
					&rdf.Type{
						Resource: "http://www.w3.org/2004/02/skos/core#Collection",
					},
				},
				Labels: []*rdfs.Label{
					&rdfs.Label{
						XsdtString: "<materials by function>",
						Lang:       "en",
					},
					&rdfs.Label{
						XsdtString: "<materialen naar functie>",
						Lang:       "nl",
					},
					&rdfs.Label{
						XsdtString: "<materiaal naar functie>",
						Lang:       "nl",
					},
					&rdfs.Label{
						XsdtString: "<matériaux selon la fonction>",
						Lang:       "fr",
					},
					&rdfs.Label{
						XsdtString: "<materiales por función>",
						Lang:       "es",
					},
				},
				BroaderTerms: []*rdf.ResourceAttr{
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300010358"},
				},
				BroaderPreferredTerms: []*rdf.ResourceAttr{
					&rdf.ResourceAttr{Resource: "http://vocab.getty.edu/aat/300010358"},
				},
				MemberTerms: []*rdf.ResourceAttr{
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014693"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014701"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014801"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300389869"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265590"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300235576"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379033"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300080663"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300236306"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014842"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300254871"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300257037"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014720"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014846"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265585"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379608"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300011146"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014855"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300263384"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014857"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014902"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379664"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265211"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014904"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300014907"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300264870"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300013026"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015114"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265588"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015119"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265586"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379869"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300183758"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300234632"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300386688"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015131"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015132"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300080665"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015133"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015135"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015134"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015137"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300254496"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015139"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265591"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300246920"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300251577"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300227846"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015152"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300266340"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015153"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300391379"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379584"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300254740"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300227860"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015166"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300127358"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300264512"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300212646"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300183876"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015172"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015186"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300379616"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300194439"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015187"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015158"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300389687"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015192"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015217"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300311593"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300015295"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300265584"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300250441"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300183798"},
					&rdf.ResourceAttr{"http://vocab.getty.edu/aat/300204912"},
				},
			},
		},
	},
}

func TestMarshal(t *testing.T) {
	for idx, test := range marshalTests {
		if test.UnmarshalOnly {
			continue
		}
		data, err := xml.Marshal(test.Value)
		if err != nil {
			t.Errorf("#%d: Error: %s", idx, err)
			continue
		}

		if got, want := string(data), test.ExpectXML; got != want {
			if strings.Contains(want, "\n") {
				t.Errorf("#%d: marshal(%#v):\nHAVE:\n%s\nWANT:\n%s", idx, test.Value, got, want)
			} else {

				t.Errorf("#%d: marshal(%#v):\nhave %#q\nwant %#q", idx, test.Value, got, want)
			}
		}
	}
}

func TestUnmarshal(t *testing.T) {
	for i, test := range marshalTests {
		if test.MarshalOnly {
			continue
		}

		expectXML := test.ExpectXML

		if test.IsFile {
			bytes, err := ioutil.ReadFile(expectXML)

			if err != nil {
				t.Errorf("Failed to load test data: %s", err)
			}

			expectXML = string(bytes[:])
		}

		vt := reflect.TypeOf(test.Value)
		dest := reflect.New(vt.Elem()).Interface()
		err := xml.Unmarshal([]byte(expectXML), dest)

		if err != nil {
			t.Errorf("#%d: unexpected error: %#v", i, err)
		} else if got, want := dest, test.Value; !reflect.DeepEqual(got, want) {
			t.Error(spew.Errorf("#%d: unmarshal(%q):\nhave %#v\nwant %#v", i, test.ExpectXML, got, want))
		}
	}
}
