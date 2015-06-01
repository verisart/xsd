package lido

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

var marshalTests = []struct {
	Value         interface{}
	ExpectXML     string
	MarshalOnly   bool
	UnmarshalOnly bool
}{
	{
		Value: &Lido{
			XMLName: xml.Name{Space: "http://www.lido-schema.org", Local: "lido"},
			LidoRecIDs: []*Identifier{
				&Identifier{
					Source:     "Deutsches Dokumentationszentrum für Kunstgeschichte - Bildarchiv Foto Marburg",
					Type:       "local",
					XsdtString: "DE-Mb112/lido-obj00154983",
				},
			},
			Category: &Concept{
				ConceptIDs: []*Identifier{
					&Identifier{
						Type:       "URI",
						XsdtString: "http://www.cidoc-crm.org/crm-concepts/E22",
					},
				},
				Terms: []*Term{
					&Term{
						XsdtString: "Man-Made Object",
						Lang:       "en",
					},
				},
			},
		},
		ExpectXML: `<lido xmlns="http://www.lido-schema.org">` +
			`<lidoRecID` +
			` source="Deutsches Dokumentationszentrum für Kunstgeschichte - Bildarchiv Foto Marburg"` +
			` type="local">DE-Mb112/lido-obj00154983</lidoRecID>` +
			`<category>` +
			`<conceptID type="URI">http://www.cidoc-crm.org/crm-concepts/E22</conceptID>` +
			`<term xml:lang="en">Man-Made Object</term>` +
			`</category>` +
			`</lido>`,
	},
	{
		MarshalOnly: false,
		Value: &MaterialsTech{
			XMLName: xml.Name{Space: "", Local: "materialsTech"},
			TermMaterialsTechs: []*ClassificationElement{
				&ClassificationElement{
					Concept: Concept{
						Terms: []*Term{
							&Term{
								XsdtString: "poplar",
							},
							&Term{
								XsdtString:      "wood",
								AddedSearchTerm: "yes",
							},
						},
					},
					Type: "material",
				},
			},
		},
		ExpectXML: `<materialsTech>` +
			`<termMaterialsTech type="material">` +
			`<term>poplar</term>` +
			`<term addedSearchTerm="yes">wood</term>` +
			`</termMaterialsTech>` +
			`</materialsTech>`,
	},
	{
		MarshalOnly: false,
		Value: &Actor{
			XMLName: xml.Name{Space: "", Local: "actor"},
			Type:    "person",
			ActorIDs: []*Identifier{
				&Identifier{
					XsdtString: "kue 02553338",
					Source:     "Bildindex-KUE-Datei",
					Type:       "local",
				},
			},
			NameActorSets: []*Appellation{
				&Appellation{
					AppellationValues: []*AppellationValue{
						&AppellationValue{
							XsdtString: "Botticelli, Sandro",
							Pref:       "preferred",
						},
					},
				},
				&Appellation{
					AppellationValues: []*AppellationValue{
						&AppellationValue{
							XsdtString: "Filipepi, Alessandro",
							Pref:       "alternate",
						},
					},
				},
				&Appellation{
					AppellationValues: []*AppellationValue{
						&AppellationValue{
							XsdtString: "Filipepi, Sandro",
							Pref:       "alternate",
						},
					},
				},
			},
			NationalityActors: []*ConceptElement{
				&ConceptElement{
					Concept: Concept{
						Terms: []*Term{
							&Term{
								XsdtString: "Italien",
							},
						},
					},
				},
			},
			VitalDatesActor: &DateSpan{
				EarliestDate: &Date{
					XsdtString: "1445",
					Type:       "estimatedDate",
				},
				LatestDate: &Date{
					XsdtString: "1510-05-17",
					Type:       "estimatedDate",
				},
			},
			GenderActors: []*Text{
				&Text{
					XsdtString: "male",
				},
			},
		},
		ExpectXML: `<actor type="person">` +
			`<actorID source="Bildindex-KUE-Datei" type="local">kue 02553338</actorID>` +
			`<nameActorSet>` +
			`<appellationValue pref="preferred">Botticelli, Sandro</appellationValue>` +
			`</nameActorSet>` +
			`<nameActorSet>` +
			`<appellationValue pref="alternate">Filipepi, Alessandro</appellationValue>` +
			`</nameActorSet>` +
			`<nameActorSet>` +
			`<appellationValue pref="alternate">Filipepi, Sandro</appellationValue>` +
			`</nameActorSet>` +
			`<nationalityActor>` +
			`<term>Italien</term>` +
			`</nationalityActor>` +
			`<vitalDatesActor>` +
			`<earliestDate type="estimatedDate">1445</earliestDate>` +
			`<latestDate type="estimatedDate">1510-05-17</latestDate>` +
			`</vitalDatesActor>` +
			`<genderActor>male</genderActor>` +
			`</actor>`,
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

		vt := reflect.TypeOf(test.Value)
		dest := reflect.New(vt.Elem()).Interface()
		err := xml.Unmarshal([]byte(test.ExpectXML), dest)

		if err != nil {
			t.Errorf("#%d: unexpected error: %#v", i, err)
		} else if got, want := dest, test.Value; !reflect.DeepEqual(got, want) {
			t.Errorf("#%d: unmarshal(%q):\nhave %#v\nwant %#v", i, test.ExpectXML, got, want)
		}
	}
}
