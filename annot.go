package poppler

// #cgo pkg-config: poppler-glib
// #include <poppler.h>
// #include <glib.h>
// #include <cairo.h>
import "C"
import "unsafe"
//import "github.com/ungerik/go-cairo"

//import "fmt"

type PopplerAnnotType int

const (
    POPPLER_ANNOT_UNKNOWN         PopplerAnnotType = C.POPPLER_ANNOT_UNKNOWN
    POPPLER_ANNOT_TEXT            PopplerAnnotType = C.POPPLER_ANNOT_TEXT
    POPPLER_ANNOT_LINK            PopplerAnnotType = C.POPPLER_ANNOT_LINK
    POPPLER_ANNOT_FREE_TEXT       PopplerAnnotType = C.POPPLER_ANNOT_FREE_TEXT
    POPPLER_ANNOT_LINE            PopplerAnnotType = C.POPPLER_ANNOT_LINE
    POPPLER_ANNOT_SQUARE          PopplerAnnotType = C.POPPLER_ANNOT_SQUARE
    POPPLER_ANNOT_CIRCLE          PopplerAnnotType = C.POPPLER_ANNOT_CIRCLE
    POPPLER_ANNOT_POLYGON         PopplerAnnotType = C.POPPLER_ANNOT_POLYGON
    POPPLER_ANNOT_POLY_LINE       PopplerAnnotType = C.POPPLER_ANNOT_POLY_LINE
    POPPLER_ANNOT_HIGHLIGHT       PopplerAnnotType = C.POPPLER_ANNOT_HIGHLIGHT
    POPPLER_ANNOT_UNDERLINE       PopplerAnnotType = C.POPPLER_ANNOT_UNDERLINE
    POPPLER_ANNOT_SQUIGGLY        PopplerAnnotType = C.POPPLER_ANNOT_SQUIGGLY
    POPPLER_ANNOT_STRIKE_OUT      PopplerAnnotType = C.POPPLER_ANNOT_STRIKE_OUT
    POPPLER_ANNOT_STAMP           PopplerAnnotType = C.POPPLER_ANNOT_STAMP
    POPPLER_ANNOT_CARET           PopplerAnnotType = C.POPPLER_ANNOT_CARET
    POPPLER_ANNOT_INK             PopplerAnnotType = C.POPPLER_ANNOT_INK
    POPPLER_ANNOT_POPUP           PopplerAnnotType = C.POPPLER_ANNOT_POPUP
    POPPLER_ANNOT_FILE_ATTACHMENT PopplerAnnotType = C.POPPLER_ANNOT_FILE_ATTACHMENT
    POPPLER_ANNOT_SOUND           PopplerAnnotType = C.POPPLER_ANNOT_SOUND
    POPPLER_ANNOT_MOVIE           PopplerAnnotType = C.POPPLER_ANNOT_MOVIE
    POPPLER_ANNOT_WIDGET          PopplerAnnotType = C.POPPLER_ANNOT_WIDGET
    POPPLER_ANNOT_SCREEN          PopplerAnnotType = C.POPPLER_ANNOT_SCREEN
    POPPLER_ANNOT_PRINTER_MARK    PopplerAnnotType = C.POPPLER_ANNOT_PRINTER_MARK
    POPPLER_ANNOT_TRAP_NET        PopplerAnnotType = C.POPPLER_ANNOT_TRAP_NET
    POPPLER_ANNOT_WATERMARK       PopplerAnnotType = C.POPPLER_ANNOT_WATERMARK
    POPPLER_ANNOT_3D              PopplerAnnotType = C.POPPLER_ANNOT_3D
)


type Annot struct {
	a *C.struct__PopplerAnnot
}

type AnnotMapping struct {
	a *C.struct__PopplerAnnotMapping
	Page
}


//func (am *AnnotMapping) ExtractPopplerAnnot() Annot {
	//popplerAnnot := am.a.annot
	//return Annot{p: popplerAnnot}
//}

func (am *AnnotMapping) Type() PopplerAnnotType{
	// method getAnnot
	popplerAnnot := am.a.annot
	return PopplerAnnotType(C.poppler_annot_get_annot_type(popplerAnnot))
}

func (am *AnnotMapping) Text() string {
	var rect C.struct__PopplerRectangle
	popplerAnnot := am.a.annot

	C.poppler_annot_get_rectangle(popplerAnnot, &rect)
	cText := C.poppler_page_get_text_for_area(am.p, &rect)

	if cText == nil {
		return ""
	}
	defer C.g_free(C.gpointer(unsafe.Pointer(cText)))

	return C.GoString(cText)
}

