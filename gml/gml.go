package gml

import (
	xlink "github.com/verisart/xsd/xlink"
	xsdt "github.com/verisart/xsd/xsdt"
)

//  A LineString is a special curve that consists of a single segment with linear interpolation. It is defined by two or more coordinate
//  tuples, with linear interpolation between them. It is backwards compatible with the LineString of GML 2, GM_LineString of ISO 19107 is
//  implemented by LineStringSegment.
type LineString struct {
	AbstractGeometry

	//  GML supports two different ways to specify the control points of a line string. 1. A sequence of "pos"
	//  (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part
	//  of this curve, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference
	//  another point defined outside of this curve (reuse of existing points). 2. The "posList" element allows for a compact way to
	//  specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong
	//  to this curve only. The number of direct positions in the list must be at least two.
	PosList *DirectPosition `xml:"http://www.opengis.net/gml posList"`

	//  GML supports two different ways to specify the control points of a line string. 1. A sequence of "pos"
	//  (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part
	//  of this curve, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference
	//  another point defined outside of this curve (reuse of existing points). 2. The "posList" element allows for a compact way to
	//  specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong
	//  to this curve only. The number of direct positions in the list must be at least two.
	//  Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *Coordinates `xml:"http://www.opengis.net/gml coordinates"`

	Poses []*DirectPosition `xml:"http://www.opengis.net/gml pos"`

	PointProperties []*PointProperty `xml:"http://www.opengis.net/gml pointProperty"`

	//  Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility
	//  with GML 3.0.0.
	PointReps []*PointProperty `xml:"http://www.opengis.net/gml pointRep"`

	//  Deprecated with GML version 3.0. Use "pos" instead. The "coord" element is included for backwards
	//  compatibility with GML 2.
	Coord *Coord `xml:"http://www.opengis.net/gml coord"`
}

// A Polygon is a special surface that is defined by a single surface patch. The
// boundary of this patch is coplanar and the polygon uses planar interpolation
// in its interior. It is backwards compatible with the Polygon of GML 2,
// GM_Polygon of ISO 19107 is implemented by PolygonPatch.
type Polygon struct {
	Exterior

	Interior

	AbstractGeometry
}

type Point struct {
	//  GML supports two different ways to specify the direct poisiton of a point.
	// 1. The "pos" element is of type DirectPositionType.
	Pos *DirectPosition `xml:"http://www.opengis.net/gml pos"`

	//  GML supports two different ways to specify the direct poisiton of a point. 1. The "pos" element is of type
	//  DirectPositionType.
	//  Deprecated with GML version 3.1.0 for coordinates with ordinate values that are numbers. Use "pos"
	//  instead. The "coordinates" element shall only be used for coordinates with ordinates that require a string
	//  representation, e.g. DMS representations.
	Coordinates *Coordinates `xml:"http://www.opengis.net/gml coordinates"`

	// GML supports two different ways to specify the direct poisiton of a point.
	// 1. The "pos" element is of type DirectPositionType.
	// Deprecated with GML version 3.0. Use "pos" instead. The "coord" element is
	// included for backwards compatibility with GML 2.
	Coord *Coord `xml:"http://www.opengis.net/gml coord"`

	AbstractGeometry
}

type AbstractGeometry struct {
	AbstractGML

	SRSReferenceGroup

	//  This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	//  This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	//  anymore and may be deleted in future versions of GML without further notice.
	GID xsdt.String `xml:"http://www.opengis.net/gml gid,attr"`
}

type AbstractGML struct {
	ID xsdt.String `xml:"http://www.opengis.net/gml id,attr"`

	StandardObjectProperties
}

type StandardObjectProperties struct {
	MetaDataProperties []*MetaDataProperty `xml:"http://www.opengis.net/gml metaDataProperty"`

	Description *StringOrRef `xml:"http://www.opengis.net/gml description"`

	// Label for the object, normally a descriptive name. An object may have
	// several names, typically assigned by different authorities.  The authority
	// for a name is indicated by the value of its (optional) codeSpace attribute.
	// The name may or may not be unique, as determined by the rules of the
	// organization responsible for the codeSpace.
	Names []*Code `xml:"http://www.opengis.net/gml name"`

	//  The name by which this coordinate system is identified.
	CsNames []*Code `xml:"http://www.opengis.net/gml csName"`

	//  The name by which this reference system is identified.
	SrsNames []*Code `xml:"http://www.opengis.net/gml srsName"`

	//  The name by which this operation parameter group is identified.
	GroupNames []*Code `xml:"http://www.opengis.net/gml groupName"`

	//  The name by which this datum is identified.
	DatumNames []*Code `xml:"http://www.opengis.net/gml datumName"`

	//  The name by which this prime meridian is identified. The meridianName most common value is Greenwich, and that value shall be used when the greenwichLongitude value is zero.
	MeridianNames []*Code `xml:"http://www.opengis.net/gml meridianName"`

	//  The name by which this ellipsoid is identified.
	EllipsoidNames []*Code `xml:"http://www.opengis.net/gml ellipsoidName"`

	//  The name by which this coordinate operation is identified.
	CoordinateOperationNames []*Code `xml:"http://www.opengis.net/gml coordinateOperationName"`

	//  The name by which this operation method is identified.
	MethodNames []*Code `xml:"http://www.opengis.net/gml methodName"`

	//  The name by which this operation parameter is identified.
	ParameterNames []*Code `xml:"http://www.opengis.net/gml parameterName"`
}

// DirectPositions, as data types, will often be included in larger objects
// (such as geometry elements) that have references to CRS, the "srsName"
// attribute will in general be missing, if this particular DirectPosition is
// included in a larger element with such a reference to a CRS. In this case,
// the CRS is implicitly assumed to take on the value of the containing object's
// CRS. XML List based on XML Schema double type.  An element of this type
// contains a space-separated list of double values
type DoubleList xsdt.String

type DirectPosition struct {
	SRSReferenceGroup
	DoubleList
}

type SRSReferenceGroup struct {
	//  Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	//  labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	//  When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels TNCNameList `xml:"http://www.opengis.net/gml axisLabels,attr"`

	//  Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	//  gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	//  axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	//  shall also be omitted.
	UomLabels TNCNameList `xml:"http://www.opengis.net/gml uomLabels,attr"`

	//  In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	//  (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	//  location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	//  geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	//  at the direct position level only in rare cases.
	SrsName xsdt.AnyURI `xml:"http://www.opengis.net/gml srsName,attr"`

	//  The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	//  specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension xsdt.PositiveInteger `xml:"http://www.opengis.net/gml srsDimension,attr"`
}

//  Optional reference to the CRS used by this geometry, with optional additional information to simplify use when
//  a more complete definition of the CRS is not needed.
//  Optional additional and redundant information for a CRS to simplify use when a more complete definition of the
//  CRS is not needed. This information shall be the same as included in the more complete definition of the CRS, referenced by the
//  srsName attribute. When the srsName attribute is included, either both or neither of the axisLabels and uomLabels attributes
//  shall be included. When the srsName attribute is omitted, both of these attributes shall be omitted.
//  Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
//  labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
//  When the srsName attribute is omitted, this attribute shall also be omitted.
//  A set of values, representing a list of token with the lexical value space of NCName. The tokens are seperated by whitespace.
type TNCNameList xsdt.String

type SRSInformationGroup struct {
	//  Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	//  labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	//  When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels TNCNameList `xml:"http://www.opengis.net/gml axisLabels,attr"`

	//  Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	//  gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	//  axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	//  shall also be omitted.
	UomLabels TNCNameList `xml:"http://www.opengis.net/gml uomLabels,attr"`
}

type Coordinates struct {
	XsdtString xsdt.String `xml:",chardata"`

	Decimal xsdt.String `xml:"http://www.opengis.net/gml decimal,attr"`

	Cs xsdt.String `xml:"http://www.opengis.net/gml cs,attr"`

	Ts xsdt.String `xml:"http://www.opengis.net/gml ts,attr"`
}

type Coord struct {
	X xsdt.Decimal `xml:"http://www.opengis.net/gml X"`

	Y xsdt.Decimal `xml:"http://www.opengis.net/gml Y"`

	Z xsdt.Decimal `xml:"http://www.opengis.net/gml Z"`
}

//  This type encapsulates various dynamic properties of moving objects
//  (points, lines, regions). It is useful for dealing with features whose
//  geometry or topology changes over time.
//  A timeslice encapsulates the time-varying properties of a dynamic feature--it
//  must be extended to represent a timestamped projection of a feature. The dataSource
//  property describes how the temporal data was acquired.
//  All complexContent GML elements are directly or indirectly derived from this abstract supertype
//  to establish a hierarchy of GML types that may be distinguished from other XML types by their ancestry.
//  Elements in this hierarchy may have an ID and are thus referenceable.
//  This content model group makes it easier to construct types that
//  derive from AbstractGMLType and its descendents "by restriction".
//  A reference to the group saves having to enumerate the standard object properties.
//  Label for the object, normally a descriptive name. An object may have several names, typically assigned by different authorities.  The authority for a name is indicated by the value of its (optional) codeSpace attribute.  The name may or may not be unique, as determined by the rules of the organization responsible for the codeSpace.
//  The name by which this operation method is identified.
//  Name or code with an (optional) authority.  Text token.
//  If the codeSpace attribute is present, then its value should identify a dictionary, thesaurus
//  or authority for the term, such as the organisation who assigned the value,
//  or the dictionary from which it is taken.
//  A text string with an optional codeSpace attribute.
type Code struct {
	XsdtString xsdt.String `xml:",chardata"`

	CodeSpace xsdt.AnyURI `xml:"http://www.opengis.net/gml codeSpace,attr"`
}

//  This type is available wherever there is a need for a "text" type property. It is of string type, so the text can be included inline, but the value can also be referenced remotely via xlinks from the AssociationAttributeGroup. If the remote reference is present, then the value obtained by traversing the link should be used, and the string content of the element can be used for an annotation.
type StringOrRef struct {
	XsdtString xsdt.String `xml:",chardata"`

	xlink.SimpleLink

	//  Reference to an XML Schema fragment that specifies the content model of
	// the propertys value. This is in conformance with the XML Schema Section
	// 4.14 Referencing Schemas from Elsewhere.
	RemoteSchema xsdt.AnyURI `xml:"http://www.opengis.net/gml remoteSchema,attr"`
}

type MetaDataProperty struct {
	xlink.SimpleLink

	About xsdt.AnyURI `xml:"http://www.opengis.net/gml about,attr"`

	//  Reference to an XML Schema fragment that specifies the content model of
	// the propertys value. This is in conformance with the XML Schema Section
	// 4.14 Referencing Schemas from Elsewhere.
	RemoteSchema xsdt.AnyURI `xml:"http://www.opengis.net/gml remoteSchema,attr"`
}

type Interior struct {
	// A boundary of a surface consists of a number of rings. The "interior" rings
	// seperate the surface / surface patch from the area enclosed by the rings.
	Interiors []*AbstractRingProperty `xml:"http://www.opengis.net/gml interior"`

	// Deprecated with GML 3.0, included only for backwards compatibility with
	// GML 2. Use "interior" instead.
	InnerBoundaryIses []*AbstractRingProperty `xml:"http://www.opengis.net/gml innerBoundaryIs"`
}

type Exterior struct {
	//  A boundary of a surface consists of a number of rings. In the normal 2D case, one of these rings is distinguished as being the exterior boundary. In a general manifold this is not always possible, in which case all boundaries shall be listed as interior boundaries, and the exterior will be empty.
	Exterior *AbstractRingProperty `xml:"http://www.opengis.net/gml exterior"`

	//  Deprecated with GML 3.0, included only for backwards compatibility with GML 2. Use "exterior" instead.
	OuterBoundaryIs *AbstractRingProperty `xml:"http://www.opengis.net/gml outerBoundaryIs"`
}

type AbstractRingProperty struct {
	//  The "_Ring" element is the abstract head of the substituition group for all closed boundaries of a surface patch.
	_Ring *AbstractGeometry `xml:"http://www.opengis.net/gml _Ring"`

	LinearRing *LinearRing `xml:"http://www.opengis.net/gml LinearRing"`

	Ring *Ring `xml:"http://www.opengis.net/gml Ring"`
}

type LinearRing struct {
	//  GML supports two different ways to specify the control points of a linear ring.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this ring, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this ring (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this ring only. The number of direct positions in the list must be at least four.
	//  Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *Coordinates `xml:"http://www.opengis.net/gml coordinates"`

	//  GML supports two different ways to specify the control points of a linear ring.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this ring, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this ring (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this ring only. The number of direct positions in the list must be at least four.
	//  Deprecated with GML version 3.0 and included for backwards compatibility with GML 2. Use "pos" elements instead.
	Coord *Coord `xml:"http://www.opengis.net/gml coord"`

	Poses []*DirectPosition `xml:"http://www.opengis.net/gml pos"`

	PointProperties []*PointProperty `xml:"http://www.opengis.net/gml pointProperty"`

	//  Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointReps []*PointProperty `xml:"http://www.opengis.net/gml pointRep"`

	AbstractGeometry

	//  GML supports two different ways to specify the control points of a linear ring.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this ring, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this ring (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this ring only. The number of direct positions in the list must be at least four.
	PosList *DirectPosition `xml:"http://www.opengis.net/gml posList"`
}

type Ring struct {
	AbstractGeometry

	//  This element references or contains one curve in the composite curve. The curves are contiguous, the collection of curves is ordered.
	//  NOTE: This definition allows for a nested structure, i.e. a CompositeCurve may use, for example, another CompositeCurve as a curve member.
	CurveMembers []*CurveProperty `xml:"http://www.opengis.net/gml curveMember"`
}

type CurveProperty struct {
	//  The "_Curve" element is the abstract head of the substituition group for all (continuous) curve elements.
	_Curve *AbstractGeometry `xml:"http://www.opengis.net/gml _Curve"`

	LineString *LineString `xml:"http://www.opengis.net/gml LineString"`

	CompositeCurve *CompositeCurve `xml:"http://www.opengis.net/gml CompositeCurve"`

	Curve *Curve `xml:"http://www.opengis.net/gml Curve"`

	OrientableCurve *OrientableCurve `xml:"http://www.opengis.net/gml OrientableCurve"`

	AssociationGroup
}

type AssociationGroup struct {
	//  This attribute group includes the XLink attributes (see xlinks.xsd). XLink is used in GML to reference remote
	//  resources (including those elsewhere in the same document). A simple link element can be constructed by including a specific
	//  set of XLink attributes. The XML Linking Language (XLink) is currently a Proposed Recommendation of the World Wide Web Consortium.
	//  XLink allows elements to be inserted into XML documents so as to create sophisticated links between resources; such links can be used
	//  to reference remote properties. A simple link element can be used to implement pointer functionality, and this functionality has been built
	//  into various GML 3 elements by including the gml:AssociationAttributeGroup.
	xlink.SimpleLink

	//  Reference to an XML Schema fragment that specifies the content model of
	// the propertys value. This is in conformance with the XML Schema Section
	// 4.14 Referencing Schemas from Elsewhere.
	RemoteSchema xsdt.AnyURI `xml:"http://www.opengis.net/gml remoteSchema,attr"`
}

type PointProperty struct {
	Point *Point `xml:"http://www.opengis.net/gml Point"`

	AssociationGroup
}

//  A CompositeCurve is defined by a sequence of (orientable) curves such that the each curve in the sequence terminates at the start point of the subsequent curve in the list.
type CompositeCurve struct {
	AbstractGeometry

	//  This element references or contains one curve in the composite curve. The curves are contiguous, the collection of curves is ordered.
	//  NOTE: This definition allows for a nested structure, i.e. a CompositeCurve may use, for example, another CompositeCurve as a curve member.
	CurveMembers []*CurveProperty `xml:"http://www.opengis.net/gml curveMember"`
}

//  Utility type used in various places
//  - e.g. to indicate the direction of topological objects;
//  "+" for forwards, or "-" for backwards.
type Sign xsdt.String

type OrientableCurve struct {
	//  If the orientation is "+", then the OrientableCurve is identical to the baseCurve. If the orientation is "-", then the OrientableCurve is related to another _Curve with a parameterization that reverses the sense of the curve traversal. "+" is the default value.
	Orientation Sign `xml:"http://www.opengis.net/gml orientation,attr"`

	AbstractGeometry

	//  References or contains the base curve (positive orientation).
	//  NOTE: This definition allows for a nested structure, i.e. an OrientableCurve may use another OrientableCurve as its base curve.
	BaseCurve *CurveProperty `xml:"http://www.opengis.net/gml baseCurve"`
}

type Curve struct {
	AbstractGeometry

	//  This element encapsulates the segments of the curve.
	Segments *CurveSegments `xml:"http://www.opengis.net/gml segments"`
}

type CurveSegments struct {
	//  The "_CurveSegment" element is the abstract head of the substituition group for all curve segment elements, i.e. continuous segments of the same interpolation mechanism.
	_CurveSegments []*CurveSegment `xml:"http://www.opengis.net/gml _CurveSegment"`

	ArcStrings []*ArcString `xml:"http://www.opengis.net/gml ArcString"`

	// TODO:
	//XsdGoPkgHasElems_ArcStringByBulge

	// TODO:
	//XsdGoPkgHasElems_CubicSpline

	// TODO:
	//XsdGoPkgHasElems_BSpline

	// TODO:
	//XsdGoPkgHasElems_GeodesicString

	// TODO:
	//XsdGoPkgHasElems_LineStringSegment

	// TODO:
	//XsdGoPkgHasElems_ArcByCenterPoint

	// TODO:
	//XsdGoPkgHasElems_OffsetCurve

	// TODO:
	//XsdGoPkgHasElems_Clothoid
}

type ArcStrings struct {
	ArcStrings []*ArcString `xml:"http://www.opengis.net/gml ArcString"`

	Arc
}

type Arcs struct {
	Arcs []*Arc `xml:"http://www.opengis.net/gml Arc"`

	Circles []*Arc `xml:"http://www.opengis.net/gml Circle"`
}

type Arc struct {
	ArcString

	//  GML supports two different ways to specify the control points of a curve segment.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this curve segment, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this curve segment (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this curve segment only. The number of direct positions in the list must be three.
	PosList *DirectPosition `xml:"http://www.opengis.net/gml posList"`

	//  GML supports two different ways to specify the control points of a curve segment.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this curve segment, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this curve segment (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this curve segment only. The number of direct positions in the list must be three.
	//  Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *Coordinates `xml:"http://www.opengis.net/gml coordinates"`

	Poses []*DirectPosition `xml:"http://www.opengis.net/gml pos"`

	PointProperties []*PointProperty `xml:"http://www.opengis.net/gml pointProperty"`

	//  Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointReps []*PointProperty `xml:"http://www.opengis.net/gml pointRep"`

	//  An arc is an arc string consiting of a single arc, the attribute is fixed to "1".
	NumArc xsdt.Integer `xml:"http://www.opengis.net/gml numArc,attr"`
}

//  A Ring is used to represent a single connected component of a surface boundary. It consists of a sequence of curves connected in a cycle (an object whose boundary is empty).
//  A Ring is structurally similar to a composite curve in that the endPoint of each curve in the sequence is the startPoint of the next curve in the Sequence. Since the sequence is circular, there is no exception to this rule. Each ring, like all boundaries, is a cycle and each ring is simple.
//  NOTE: Even though each Ring is simple, the boundary need not be simple. The easiest case of this is where one of the interior rings of a surface is tangent to its exterior ring.
//  This property element either references a curve via the XLink-attributes or contains the curve element. A curve element is any element which is substitutable for "_Curve".
//  A property that has a curve as its value domain can either be an appropriate geometry element encapsulated in an
//  element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere
//  in the same document). Either the reference or the contained element must be given, but neither both nor none.
//  The "_Curve" element is the abstract head of the substituition group for all (continuous) curve elements.
//  Curve is a 1-dimensional primitive. Curves are continuous, connected, and have a measurable length in terms of the coordinate system.
//  A curve is composed of one or more curve segments. Each curve segment within a curve may be defined using a different interpolation method. The curve segments are connected to one another, with the end point of each segment except the last being the start point of the next segment in the segment list.
//  The orientation of the curve is positive.
//  This property element contains a list of curve segments. The order of the elements is significant and shall be preserved when processing the array.
//  A container for an array of curve segments.
//  The "_CurveSegment" element is the abstract head of the substituition group for all curve segment elements, i.e. continuous segments of the same interpolation mechanism.
//  A LineStringSegment is a curve segment that is defined by two or more coordinate tuples, with linear interpolation between them.
//  Note: LineStringSegment implements GM_LineString of ISO 19107.
//  The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
//  uses the control points and control parameters to determine the position of this curve segment. For a LineStringSegment the interpolation is fixed as "linear".
//  CurveInterpolationType is a list of codes that may be used to identify the interpolation mechanisms specified by an
//  application schema.
type CurveInterpolation xsdt.String

type ArcString struct {
	//  The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	//  uses the control points and control parameters to determine the position of this curve segment. For an ArcString the interpolation is fixed as "circularArc3Points".
	Interpolation CurveInterpolation `xml:"http://www.opengis.net/gml interpolation,attr"`

	//  The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be 2 * numArc + 1.
	NumArc xsdt.Integer `xml:"http://www.opengis.net/gml numArc,attr"`

	CurveSegment

	//  GML supports two different ways to specify the control points of a curve segment.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this curve segment, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this curve segment (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this curve segment only. The number of direct positions in the list must be at least three.
	PosList *DirectPosition `xml:"http://www.opengis.net/gml posList"`

	//  GML supports two different ways to specify the control points of a curve segment.
	//  1. A sequence of "pos" (DirectPositionType) or "pointProperty" (PointPropertyType) elements. "pos" elements are control points that are only part of this curve segment, "pointProperty" elements contain a point that may be referenced from other geometry elements or reference another point defined outside of this curve segment (reuse of existing points).
	//  2. The "posList" element allows for a compact way to specifiy the coordinates of the control points, if all control points are in the same coordinate reference systems and belong to this curve segment only. The number of direct positions in the list must be at least three.
	//  Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *Coordinates `xml:"http://www.opengis.net/gml coordinates"`

	Poses []*DirectPosition `xml:"http://www.opengis.net/gml pos"`

	PointProperties []*PointProperty `xml:"http://www.opengis.net/gml pointProperty"`

	//  Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointReps []*PointProperty `xml:"http://www.opengis.net/gml pointRep"`
}

type CurveSegment struct {
	//  The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	//  NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart xsdt.Integer `xml:"http://www.opengis.net/gml numDerivativesAtStart,attr"`

	//  The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	//  NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd xsdt.Integer `xml:"http://www.opengis.net/gml numDerivativesAtEnd,attr"`

	//  The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	//  NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior xsdt.Integer `xml:"http://www.opengis.net/gml numDerivativeInterior,attr"`
}
