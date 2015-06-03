package lido

import (
	"github.com/verisart/xsd/xsdt"
)

type MeasurementsWrap struct {
	MeasurementsSets []*MeasurementsSet `xml:"http://www.lido-schema.org objectMeasurementsSet"`
}

type MeasurementsSet struct {
	//  Definition: Display element for one object measurement, corresponding to the following objectMeasurement element.
	//  How to record: Repeat this element only for language variants.
	DisplayMeasurements []*Text `xml:"http://www.lido-schema.org displayObjectMeasurements"`

	// Structured measurement information about the dimensions, size, or scale of
	// the object / work. it may also include the parts of a complex object /
	// work, series, or collection.
	Measurements *Measurements `xml:"http://www.lido-schema.org objectMeasurements"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

type Measurements struct {
	//  Definition: The configuration of an object / work, including technical formats. Used as necessary.
	//  How to record: Example values: Vignette, VHS, IMAX, and DOS
	FormatMeasurements []*ExtentMeasurement `xml:"http://www.lido-schema.org formatMeasurements"`

	//  Definition: The shape of an object / work. Used for unusual shapes (e.g., an oval painting).
	//  How to record: Example values: oval, round, square, rectangular, and irregular.
	ShapeMeasurements []*ExtentMeasurement `xml:"http://www.lido-schema.org shapeMeasurements"`

	//  Definition: An expression of the ratio between the size of the representation of something and that thing (e.g., the size of the drawn structure and the actual built work).
	//  How to record: Example values for scale: numeric (e.g., 1 inch = 1 foot), full-size, life-size, half size,monumental. and others as recommended in CCO and CDWA. Combine this tag with Measurement Sets for numeric scales. For measurementsSet type for Scale, use "base" for the left side of the equation, and "target" for the right side of the equation).
	//  Notes: Used for studies, record drawings, models, and other representations drawn or constructed to scale.
	ScaleMeasurements []*ExtentMeasurement `xml:"http://www.lido-schema.org scaleMeasurements"`

	// The dimensions or other measurements for one aspect of an object / work
	// (e.g., width). May be combined with extent, qualifier, and other
	// sub-elements as necessary.The subelements "measurementUnit",
	// "measurementValue" and "measurementType" are mandatory.
	MeasurementsSets []*AspectMeasurements `xml:"http://www.lido-schema.org measurementsSet"`

	//  Definition: An explanation of the part of the object / work being measured included, when necessary, for clarity.
	//  How to record: Example values: overall, components, sheet, plate mark, chain lines, pattern repeat, lid, base, laid lines, folios, leaves, columns per page, lines per page, tessera, footprint, panel, interior, mat, window of mat, secondary support, frame, and mount
	ExtentMeasurements []*ExtentMeasurement `xml:"http://www.lido-schema.org extentMeasurements"`

	//  Definition: A word or phrase that elaborates on the nature of the measurements of the object / work when necessary, e.g. when the measurements are approximate.
	//  How to record: Example values: approximate, sight, maximum, larges, smallest, average, variable, assembled, before restoration, before restoration, at corners, rounded, framed, and with base.
	QualifierMeasurements []*ExtentMeasurement `xml:"http://www.lido-schema.org qualifierMeasurements"`
}

type ExtentMeasurement struct {
	Text

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}

//  Definition: Structured measurement information about the dimensions, size, or scale of the object / work.
//  Notes: It may also include the parts of a complex object / work, series, or collection.
//  Definition: The dimensions, size, shape, scale, format, or storage configuration of the object / work, including volume, weight, area or running time.
//  How to record: Measurements are formatted to allow retrieval; preferably in metric units where applicable.
//  Definition: The dimensions or other measurements for one aspect of an object / work (e.g., width).
//  How to record: May be combined with extent, qualifier, and other sub-elements as necessary.The subelements "measurementUnit", "measurementValue" and
//  "measurementType" are mandatory.

type AspectMeasurements struct {
	//  Definition: Indicates what kind of measurement is taken.
	//  How to record: Data values for type: height, width, depth, length, diameter, circumference, stories, count, area, volume, running time, size.Repeat this element only for language variants.
	Types []*Text `xml:"http://www.lido-schema.org measurementType"`

	//  Definition: The unit of the measurement.
	//  How to record: E.g. cm, mm, m, g, kg, kb, Mb or Gb.Repeat this element only for language variants.
	Units []*Text `xml:"http://www.lido-schema.org measurementUnit"`

	//  Definition: The value of the measurement.
	//  How to record: Whole numbers or decimal fractions.
	Value Text `xml:"http://www.lido-schema.org measurementValue"`

	// Assigns a priority order for online presentation of the element. Has to be
	// a positive integer, with descending priority from 1 to x.
	SortOrder xsdt.Integer `xml:"http://www.lido-schema.org sortorder,attr,omitempty"`
}
