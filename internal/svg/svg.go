package svg

import (
	"bytes"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"image/color"
	"io"
	"strconv"

	"github.com/Fremenkiel/gophant/v2/internal/col"
)

func Colorize(src []byte, clr color.Color) ([]byte, error) {
	rdr := bytes.NewReader(src)
	s, err := svgFromXML(rdr)
	if err != nil {
		return src, fmt.Errorf("could not load SVG, falling back to static content: %v", err)
	}
	if err := s.replaceFillColor(clr); err != nil {
		return src, fmt.Errorf("could not replace fill color, falling back to static content: %v", err)
	}
	colorized, err := xml.Marshal(s)
	if err != nil {
		return src, fmt.Errorf("could not marshal svg, falling back to static content: %v", err)
	}
	return colorized, nil
}

func ColorizeStroke(src []byte, clr color.Color) ([]byte, error) {
	rdr := bytes.NewReader(src)
	s, err := svgFromXML(rdr)
	if err != nil {
		return src, fmt.Errorf("could not load SVG, falling back to static content: %v", err)
	}
	if err := s.replaceStrokeColor(clr); err != nil {
		return src, fmt.Errorf("could not replace fill color, falling back to static content: %v", err)
	}
	colorized, err := xml.Marshal(s)
	if err != nil {
		return src, fmt.Errorf("could not marshal svg, falling back to static content: %v", err)
	}
	return colorized, nil
}

type svg struct {
	XMLName  xml.Name      `xml:"svg"`
	XMLNS    string        `xml:"xmlns,attr"`
	Width    string        `xml:"width,attr,omitempty"`
	Height   string        `xml:"height,attr,omitempty"`
	ViewBox  string        `xml:"viewBox,attr,omitempty"`
	Paths    []*pathObj    `xml:"path"`
	Rects    []*rectObj    `xml:"rect"`
	Circles  []*circleObj  `xml:"circle"`
	Ellipses []*ellipseObj `xml:"ellipse"`
	Polygons []*polygonObj `xml:"polygon"`
	Groups   []*objGroup   `xml:"g"`
}

type pathObj struct {
	XMLName         xml.Name `xml:"path"`
	Fill            string   `xml:"fill,attr,omitempty"`
	FillOpacity     string   `xml:"fill-opacity,attr,omitempty"`
	Stroke          string   `xml:"stroke,attr,omitempty"`
	StrokeWidth     string   `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string   `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string   `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string   `xml:"stroke-dasharray,attr,omitempty"`
	D               string   `xml:"d,attr"`
	Transform       string   `xml:"transform,attr,omitempty"`
}

type rectObj struct {
	XMLName         xml.Name `xml:"rect"`
	Fill            string   `xml:"fill,attr,omitempty"`
	FillOpacity     string   `xml:"fill-opacity,attr,omitempty"`
	Stroke          string   `xml:"stroke,attr,omitempty"`
	StrokeWidth     string   `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string   `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string   `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string   `xml:"stroke-dasharray,attr,omitempty"`
	X               string   `xml:"x,attr,omitempty"`
	Y               string   `xml:"y,attr,omitempty"`
	Width           string   `xml:"width,attr,omitempty"`
	Height          string   `xml:"height,attr,omitempty"`
	Transform       string   `xml:"transform,attr,omitempty"`
}

type circleObj struct {
	XMLName         xml.Name `xml:"circle"`
	Fill            string   `xml:"fill,attr,omitempty"`
	FillOpacity     string   `xml:"fill-opacity,attr,omitempty"`
	Stroke          string   `xml:"stroke,attr,omitempty"`
	StrokeWidth     string   `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string   `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string   `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string   `xml:"stroke-dasharray,attr,omitempty"`
	CX              string   `xml:"cx,attr,omitempty"`
	CY              string   `xml:"cy,attr,omitempty"`
	R               string   `xml:"r,attr,omitempty"`
	Transform       string   `xml:"transform,attr,omitempty"`
}

type ellipseObj struct {
	XMLName         xml.Name `xml:"ellipse"`
	Fill            string   `xml:"fill,attr,omitempty"`
	FillOpacity     string   `xml:"fill-opacity,attr,omitempty"`
	Stroke          string   `xml:"stroke,attr,omitempty"`
	StrokeWidth     string   `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string   `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string   `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string   `xml:"stroke-dasharray,attr,omitempty"`
	CX              string   `xml:"cx,attr,omitempty"`
	CY              string   `xml:"cy,attr,omitempty"`
	RX              string   `xml:"rx,attr,omitempty"`
	RY              string   `xml:"ry,attr,omitempty"`
	Transform       string   `xml:"transform,attr,omitempty"`
}

type polygonObj struct {
	XMLName         xml.Name `xml:"polygon"`
	Fill            string   `xml:"fill,attr,omitempty"`
	FillOpacity     string   `xml:"fill-opacity,attr,omitempty"`
	Stroke          string   `xml:"stroke,attr,omitempty"`
	StrokeWidth     string   `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string   `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string   `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string   `xml:"stroke-dasharray,attr,omitempty"`
	Points          string   `xml:"points,attr"`
	Transform       string   `xml:"transform,attr,omitempty"`
}

type objGroup struct {
	XMLName         xml.Name      `xml:"g"`
	ID              string        `xml:"id,attr,omitempty"`
	Fill            string        `xml:"fill,attr,omitempty"`
	Stroke          string        `xml:"stroke,attr,omitempty"`
	StrokeWidth     string        `xml:"stroke-width,attr,omitempty"`
	StrokeLineCap   string        `xml:"stroke-linecap,attr,omitempty"`
	StrokeLineJoin  string        `xml:"stroke-linejoin,attr,omitempty"`
	StrokeDashArray string        `xml:"stroke-dasharray,attr,omitempty"`
	Transform       string        `xml:"transform,attr,omitempty"`
	Paths           []*pathObj    `xml:"path"`
	Circles         []*circleObj  `xml:"circle"`
	Ellipses        []*ellipseObj `xml:"ellipse"`
	Rects           []*rectObj    `xml:"rect"`
	Polygons        []*polygonObj `xml:"polygon"`
	Groups          []*objGroup   `xml:"g"`
}

func replacePathsFill(paths []*pathObj, hexColor string, opacity string) {
	for _, path := range paths {
		if path.Fill != "none" {
			path.Fill = hexColor
			path.FillOpacity = opacity
		}
	}
}

func replacePathsStroke(paths []*pathObj, hexColor string) {
	for _, path := range paths {
		if path.Stroke != "none" {
			path.Stroke = hexColor
		}
	}
}

func replaceRectsFill(rects []*rectObj, hexColor string, opacity string) {
	for _, rect := range rects {
		if rect.Fill != "none" {
			rect.Fill = hexColor
			rect.FillOpacity = opacity
		}
	}
}

func replaceRectsStroke(rects []*rectObj, hexColor string) {
	for _, rect := range rects {
		if rect.Stroke != "none" {
			rect.Stroke = hexColor
		}
	}
}

func replaceCirclesFill(circles []*circleObj, hexColor string, opacity string) {
	for _, circle := range circles {
		if circle.Fill != "none" {
			circle.Fill = hexColor
			circle.FillOpacity = opacity
		}
	}
}

func replaceCirclesStroke(circles []*circleObj, hexColor string) {
	for _, circle := range circles {
		if circle.Stroke != "none" {
			circle.Stroke = hexColor
		}
	}
}

func replaceEllipsesFill(ellipses []*ellipseObj, hexColor string, opacity string) {
	for _, ellipse := range ellipses {
		if ellipse.Fill != "none" {
			ellipse.Fill = hexColor
			ellipse.FillOpacity = opacity
		}
	}
}

func replaceEllipsesStroke(ellipses []*ellipseObj, hexColor string) {
	for _, ellipse := range ellipses {
		if ellipse.Stroke != "none" {
			ellipse.Stroke = hexColor
		}
	}
}

func replacePolygonsFill(polys []*polygonObj, hexColor string, opacity string) {
	for _, poly := range polys {
		if poly.Fill != "none" {
			poly.Fill = hexColor
			poly.FillOpacity = opacity
		}
	}
}

func replacePolygonsStroke(polys []*polygonObj, hexColor string) {
	for _, poly := range polys {
		if poly.Stroke != "none" {
			poly.Stroke = hexColor
		}
	}
}

func replaceGroupObjectFill(groups []*objGroup, hexColor string, opacity string) {
	for _, grp := range groups {
		replaceCirclesFill(grp.Circles, hexColor, opacity)
		replaceEllipsesFill(grp.Ellipses, hexColor, opacity)
		replacePathsFill(grp.Paths, hexColor, opacity)
		replaceRectsFill(grp.Rects, hexColor, opacity)
		replacePolygonsFill(grp.Polygons, hexColor, opacity)
		replaceGroupObjectFill(grp.Groups, hexColor, opacity)
	}
}

func replaceGroupObjectStroke(groups []*objGroup, hexColor string) {
	for _, grp := range groups {
		replaceCirclesStroke(grp.Circles, hexColor)
		replaceEllipsesStroke(grp.Ellipses, hexColor)
		replacePathsStroke(grp.Paths, hexColor)
		replaceRectsStroke(grp.Rects, hexColor)
		replacePolygonsStroke(grp.Polygons, hexColor)
		replaceGroupObjectStroke(grp.Groups, hexColor)
	}
}

// replaceFillColor alters an svg objects fill color.  Note that if an svg with multiple fill
// colors is being operated upon, all fills will be converted to a single color.  Mostly used
// to recolor Icons to match the theme's IconColor.
func (s *svg) replaceFillColor(color color.Color) error {
	hexColor, opacity := colorToHexAndOpacity(color)
	replacePathsFill(s.Paths, hexColor, opacity)
	replaceRectsFill(s.Rects, hexColor, opacity)
	replaceCirclesFill(s.Circles, hexColor, opacity)
	replaceEllipsesFill(s.Ellipses, hexColor, opacity)
	replacePolygonsFill(s.Polygons, hexColor, opacity)
	replaceGroupObjectFill(s.Groups, hexColor, opacity)
	return nil
}

func (s *svg) replaceStrokeColor(color color.Color) error {
	hexColor, _ := colorToHexAndOpacity(color)
	replacePathsStroke(s.Paths, hexColor)
	replaceRectsStroke(s.Rects, hexColor)
	replaceCirclesStroke(s.Circles, hexColor)
	replaceEllipsesStroke(s.Ellipses, hexColor)
	replacePolygonsStroke(s.Polygons, hexColor)
	replaceGroupObjectStroke(s.Groups, hexColor)
	return nil
}

func svgFromXML(reader io.Reader) (*svg, error) {
	var s svg
	bSlice, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(bSlice, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func colorToHexAndOpacity(color color.Color) (hexStr, aStr string) {
	r, g, b, a := col.ToNRGBA(color)
	cBytes := []byte{byte(r), byte(g), byte(b)}
	hexStr, aStr = "#"+hex.EncodeToString(cBytes), strconv.FormatFloat(float64(a)/0xff, 'f', 6, 64)
	return hexStr, aStr
}

