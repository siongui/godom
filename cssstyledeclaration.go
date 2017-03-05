package godom

// This file implements CSSStyleDeclaration interface
// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration

import (
	"github.com/gopherjs/gopherjs/js"
)

type CSSStyleDeclaration struct {
	*js.Object
}

// Attributes

func (s *CSSStyleDeclaration) CssText() string {
	return s.Get("cssText").String()
}

func (s *CSSStyleDeclaration) Length() int {
	return s.Get("length").Int()
}

// CSS Properties

func (s *CSSStyleDeclaration) AlignContent() string {
	return s.Get("alignContent").String()
}

func (s *CSSStyleDeclaration) SetAlignContent(v string) {
	s.Set("alignContent", v)
}

func (s *CSSStyleDeclaration) AlignItems() string {
	return s.Get("alignItems").String()
}

func (s *CSSStyleDeclaration) SetAlignItems(v string) {
	s.Set("alignItems", v)
}

func (s *CSSStyleDeclaration) AlignSelf() string {
	return s.Get("alignSelf").String()
}

func (s *CSSStyleDeclaration) SetAlignSelf(v string) {
	s.Set("alignSelf", v)
}

func (s *CSSStyleDeclaration) Animation() string {
	return s.Get("animation").String()
}

func (s *CSSStyleDeclaration) SetAnimation(v string) {
	s.Set("animation", v)
}

func (s *CSSStyleDeclaration) AnimationDelay() string {
	return s.Get("animationDelay").String()
}

func (s *CSSStyleDeclaration) SetAnimationDelay(v string) {
	s.Set("animationDelay", v)
}

func (s *CSSStyleDeclaration) AnimationDirection() string {
	return s.Get("animationDirection").String()
}

func (s *CSSStyleDeclaration) SetAnimationDirection(v string) {
	s.Set("animationDirection", v)
}

func (s *CSSStyleDeclaration) AnimationDuration() string {
	return s.Get("animationDuration").String()
}

func (s *CSSStyleDeclaration) SetAnimationDuration(v string) {
	s.Set("animationDuration", v)
}

func (s *CSSStyleDeclaration) AnimationFillMode() string {
	return s.Get("animationFillMode").String()
}

func (s *CSSStyleDeclaration) SetAnimationFillMode(v string) {
	s.Set("animationFillMode", v)
}

func (s *CSSStyleDeclaration) AnimationIterationCount() string {
	return s.Get("animationIterationCount").String()
}

func (s *CSSStyleDeclaration) SetAnimationIterationCount(v string) {
	s.Set("animationIterationCount", v)
}

func (s *CSSStyleDeclaration) AnimationName() string {
	return s.Get("animationName").String()
}

func (s *CSSStyleDeclaration) SetAnimationName(v string) {
	s.Set("animationName", v)
}

func (s *CSSStyleDeclaration) AnimationTimingFunction() string {
	return s.Get("animationTimingFunction").String()
}

func (s *CSSStyleDeclaration) SetAnimationTimingFunction(v string) {
	s.Set("animationTimingFunction", v)
}

func (s *CSSStyleDeclaration) AnimationPlayState() string {
	return s.Get("animationPlayState").String()
}

func (s *CSSStyleDeclaration) SetAnimationPlayState(v string) {
	s.Set("animationPlayState", v)
}

func (s *CSSStyleDeclaration) Background() string {
	return s.Get("background").String()
}

func (s *CSSStyleDeclaration) SetBackground(v string) {
	s.Set("background", v)
}

func (s *CSSStyleDeclaration) BackgroundAttachment() string {
	return s.Get("backgroundAttachment").String()
}

func (s *CSSStyleDeclaration) SetBackgroundAttachment(v string) {
	s.Set("backgroundAttachment", v)
}

func (s *CSSStyleDeclaration) BackgroundColor() string {
	return s.Get("backgroundColor").String()
}

func (s *CSSStyleDeclaration) SetBackgroundColor(v string) {
	s.Set("backgroundColor", v)
}

func (s *CSSStyleDeclaration) BackgroundImage() string {
	return s.Get("backgroundImage").String()
}

func (s *CSSStyleDeclaration) SetBackgroundImage(v string) {
	s.Set("backgroundImage", v)
}

func (s *CSSStyleDeclaration) BackgroundPosition() string {
	return s.Get("backgroundPosition").String()
}

func (s *CSSStyleDeclaration) SetBackgroundPosition(v string) {
	s.Set("backgroundPosition", v)
}

func (s *CSSStyleDeclaration) BackgroundRepeat() string {
	return s.Get("backgroundRepeat").String()
}

func (s *CSSStyleDeclaration) SetBackgroundRepeat(v string) {
	s.Set("backgroundRepeat", v)
}

func (s *CSSStyleDeclaration) BackgroundClip() string {
	return s.Get("backgroundClip").String()
}

func (s *CSSStyleDeclaration) SetBackgroundClip(v string) {
	s.Set("backgroundClip", v)
}

func (s *CSSStyleDeclaration) BackgroundOrigin() string {
	return s.Get("backgroundOrigin").String()
}

func (s *CSSStyleDeclaration) SetBackgroundOrigin(v string) {
	s.Set("backgroundOrigin", v)
}

func (s *CSSStyleDeclaration) BackgroundSize() string {
	return s.Get("backgroundSize").String()
}

func (s *CSSStyleDeclaration) SetBackgroundSize(v string) {
	s.Set("backgroundSize", v)
}

func (s *CSSStyleDeclaration) BackfaceVisibility() string {
	return s.Get("backfaceVisibility").String()
}

func (s *CSSStyleDeclaration) SetBackfaceVisibility(v string) {
	s.Set("backfaceVisibility", v)
}

func (s *CSSStyleDeclaration) Border() string {
	return s.Get("border").String()
}

func (s *CSSStyleDeclaration) SetBorder(v string) {
	s.Set("border", v)
}

func (s *CSSStyleDeclaration) BorderBottom() string {
	return s.Get("borderBottom").String()
}

func (s *CSSStyleDeclaration) SetBorderBottom(v string) {
	s.Set("borderBottom", v)
}

func (s *CSSStyleDeclaration) BorderBottomColor() string {
	return s.Get("borderBottomColor").String()
}

func (s *CSSStyleDeclaration) SetBorderBottomColor(v string) {
	s.Set("borderBottomColor", v)
}

func (s *CSSStyleDeclaration) BorderBottomLeftRadius() string {
	return s.Get("borderBottomLeftRadius").String()
}

func (s *CSSStyleDeclaration) SetBorderBottomLeftRadius(v string) {
	s.Set("borderBottomLeftRadius", v)
}

func (s *CSSStyleDeclaration) BorderBottomRightRadius() string {
	return s.Get("borderBottomRightRadius").String()
}

func (s *CSSStyleDeclaration) SetBorderBottomRightRadius(v string) {
	s.Set("borderBottomRightRadius", v)
}

func (s *CSSStyleDeclaration) BorderBottomStyle() string {
	return s.Get("borderBottomStyle").String()
}

func (s *CSSStyleDeclaration) SetBorderBottomStyle(v string) {
	s.Set("borderBottomStyle", v)
}

func (s *CSSStyleDeclaration) BorderBottomWidth() string {
	return s.Get("borderBottomWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderBottomWidth(v string) {
	s.Set("borderBottomWidth", v)
}

func (s *CSSStyleDeclaration) BorderCollapse() string {
	return s.Get("borderCollapse").String()
}

func (s *CSSStyleDeclaration) SetBorderCollapse(v string) {
	s.Set("borderCollapse", v)
}

func (s *CSSStyleDeclaration) BorderColor() string {
	return s.Get("borderColor").String()
}

func (s *CSSStyleDeclaration) SetBorderColor(v string) {
	s.Set("borderColor", v)
}

func (s *CSSStyleDeclaration) BorderImage() string {
	return s.Get("borderImage").String()
}

func (s *CSSStyleDeclaration) SetBorderImage(v string) {
	s.Set("borderImage", v)
}

func (s *CSSStyleDeclaration) BorderImageOutset() string {
	return s.Get("borderImageOutset").String()
}

func (s *CSSStyleDeclaration) SetBorderImageOutset(v string) {
	s.Set("borderImageOutset", v)
}

func (s *CSSStyleDeclaration) BorderImageRepeat() string {
	return s.Get("borderImageRepeat").String()
}

func (s *CSSStyleDeclaration) SetBorderImageRepeat(v string) {
	s.Set("borderImageRepeat", v)
}

func (s *CSSStyleDeclaration) BorderImageSlice() string {
	return s.Get("borderImageSlice").String()
}

func (s *CSSStyleDeclaration) SetBorderImageSlice(v string) {
	s.Set("borderImageSlice", v)
}

func (s *CSSStyleDeclaration) BorderImageSource() string {
	return s.Get("borderImageSource").String()
}

func (s *CSSStyleDeclaration) SetBorderImageSource(v string) {
	s.Set("borderImageSource", v)
}

func (s *CSSStyleDeclaration) BorderImageWidth() string {
	return s.Get("borderImageWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderImageWidth(v string) {
	s.Set("borderImageWidth", v)
}

func (s *CSSStyleDeclaration) BorderLeft() string {
	return s.Get("borderLeft").String()
}

func (s *CSSStyleDeclaration) SetBorderLeft(v string) {
	s.Set("borderLeft", v)
}

func (s *CSSStyleDeclaration) BorderLeftColor() string {
	return s.Get("borderLeftColor").String()
}

func (s *CSSStyleDeclaration) SetBorderLeftColor(v string) {
	s.Set("borderLeftColor", v)
}

func (s *CSSStyleDeclaration) BorderLeftStyle() string {
	return s.Get("borderLeftStyle").String()
}

func (s *CSSStyleDeclaration) SetBorderLeftStyle(v string) {
	s.Set("borderLeftStyle", v)
}

func (s *CSSStyleDeclaration) BorderLeftWidth() string {
	return s.Get("borderLeftWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderLeftWidth(v string) {
	s.Set("borderLeftWidth", v)
}

func (s *CSSStyleDeclaration) BorderRadius() string {
	return s.Get("borderRadius").String()
}

func (s *CSSStyleDeclaration) SetBorderRadius(v string) {
	s.Set("borderRadius", v)
}

func (s *CSSStyleDeclaration) BorderRight() string {
	return s.Get("borderRight").String()
}

func (s *CSSStyleDeclaration) SetBorderRight(v string) {
	s.Set("borderRight", v)
}

func (s *CSSStyleDeclaration) BorderRightColor() string {
	return s.Get("borderRightColor").String()
}

func (s *CSSStyleDeclaration) SetBorderRightColor(v string) {
	s.Set("borderRightColor", v)
}

func (s *CSSStyleDeclaration) BorderRightStyle() string {
	return s.Get("borderRightStyle").String()
}

func (s *CSSStyleDeclaration) SetBorderRightStyle(v string) {
	s.Set("borderRightStyle", v)
}

func (s *CSSStyleDeclaration) BorderRightWidth() string {
	return s.Get("borderRightWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderRightWidth(v string) {
	s.Set("borderRightWidth", v)
}

func (s *CSSStyleDeclaration) BorderSpacing() string {
	return s.Get("borderSpacing").String()
}

func (s *CSSStyleDeclaration) SetBorderSpacing(v string) {
	s.Set("borderSpacing", v)
}

func (s *CSSStyleDeclaration) BorderStyle() string {
	return s.Get("borderStyle").String()
}

func (s *CSSStyleDeclaration) SetBorderStyle(v string) {
	s.Set("borderStyle", v)
}

func (s *CSSStyleDeclaration) BorderTop() string {
	return s.Get("borderTop").String()
}

func (s *CSSStyleDeclaration) SetBorderTop(v string) {
	s.Set("borderTop", v)
}

func (s *CSSStyleDeclaration) BorderTopColor() string {
	return s.Get("borderTopColor").String()
}

func (s *CSSStyleDeclaration) SetBorderTopColor(v string) {
	s.Set("borderTopColor", v)
}

func (s *CSSStyleDeclaration) BorderTopLeftRadius() string {
	return s.Get("borderTopLeftRadius").String()
}

func (s *CSSStyleDeclaration) SetBorderTopLeftRadius(v string) {
	s.Set("borderTopLeftRadius", v)
}

func (s *CSSStyleDeclaration) BorderTopRightRadius() string {
	return s.Get("borderTopRightRadius").String()
}

func (s *CSSStyleDeclaration) SetBorderTopRightRadius(v string) {
	s.Set("borderTopRightRadius", v)
}

func (s *CSSStyleDeclaration) BorderTopStyle() string {
	return s.Get("borderTopStyle").String()
}

func (s *CSSStyleDeclaration) SetBorderTopStyle(v string) {
	s.Set("borderTopStyle", v)
}

func (s *CSSStyleDeclaration) BorderTopWidth() string {
	return s.Get("borderTopWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderTopWidth(v string) {
	s.Set("borderTopWidth", v)
}

func (s *CSSStyleDeclaration) BorderWidth() string {
	return s.Get("borderWidth").String()
}

func (s *CSSStyleDeclaration) SetBorderWidth(v string) {
	s.Set("borderWidth", v)
}

func (s *CSSStyleDeclaration) Bottom() string {
	return s.Get("bottom").String()
}

func (s *CSSStyleDeclaration) SetBottom(v string) {
	s.Set("bottom", v)
}

func (s *CSSStyleDeclaration) BoxShadow() string {
	return s.Get("boxShadow").String()
}

func (s *CSSStyleDeclaration) SetBoxShadow(v string) {
	s.Set("boxShadow", v)
}

func (s *CSSStyleDeclaration) BoxSizing() string {
	return s.Get("boxSizing").String()
}

func (s *CSSStyleDeclaration) SetBoxSizing(v string) {
	s.Set("boxSizing", v)
}

func (s *CSSStyleDeclaration) CaptionSide() string {
	return s.Get("captionSide").String()
}

func (s *CSSStyleDeclaration) SetCaptionSide(v string) {
	s.Set("captionSide", v)
}

func (s *CSSStyleDeclaration) Clear() string {
	return s.Get("clear").String()
}

func (s *CSSStyleDeclaration) SetClear(v string) {
	s.Set("clear", v)
}

func (s *CSSStyleDeclaration) Clip() string {
	return s.Get("clip").String()
}

func (s *CSSStyleDeclaration) SetClip(v string) {
	s.Set("clip", v)
}

func (s *CSSStyleDeclaration) Color() string {
	return s.Get("color").String()
}

func (s *CSSStyleDeclaration) SetColor(v string) {
	s.Set("color", v)
}

func (s *CSSStyleDeclaration) ColumnCount() string {
	return s.Get("columnCount").String()
}

func (s *CSSStyleDeclaration) SetColumnCount(v string) {
	s.Set("columnCount", v)
}

func (s *CSSStyleDeclaration) ColumnFill() string {
	return s.Get("columnFill").String()
}

func (s *CSSStyleDeclaration) SetColumnFill(v string) {
	s.Set("columnFill", v)
}

func (s *CSSStyleDeclaration) ColumnGap() string {
	return s.Get("columnGap").String()
}

func (s *CSSStyleDeclaration) SetColumnGap(v string) {
	s.Set("columnGap", v)
}

func (s *CSSStyleDeclaration) ColumnRule() string {
	return s.Get("columnRule").String()
}

func (s *CSSStyleDeclaration) SetColumnRule(v string) {
	s.Set("columnRule", v)
}

func (s *CSSStyleDeclaration) ColumnRuleColor() string {
	return s.Get("columnRuleColor").String()
}

func (s *CSSStyleDeclaration) SetColumnRuleColor(v string) {
	s.Set("columnRuleColor", v)
}

func (s *CSSStyleDeclaration) ColumnRuleStyle() string {
	return s.Get("columnRuleStyle").String()
}

func (s *CSSStyleDeclaration) SetColumnRuleStyle(v string) {
	s.Set("columnRuleStyle", v)
}

func (s *CSSStyleDeclaration) ColumnRuleWidth() string {
	return s.Get("columnRuleWidth").String()
}

func (s *CSSStyleDeclaration) SetColumnRuleWidth(v string) {
	s.Set("columnRuleWidth", v)
}

func (s *CSSStyleDeclaration) Columns() string {
	return s.Get("columns").String()
}

func (s *CSSStyleDeclaration) SetColumns(v string) {
	s.Set("columns", v)
}

func (s *CSSStyleDeclaration) ColumnSpan() string {
	return s.Get("columnSpan").String()
}

func (s *CSSStyleDeclaration) SetColumnSpan(v string) {
	s.Set("columnSpan", v)
}

func (s *CSSStyleDeclaration) ColumnWidth() string {
	return s.Get("columnWidth").String()
}

func (s *CSSStyleDeclaration) SetColumnWidth(v string) {
	s.Set("columnWidth", v)
}

func (s *CSSStyleDeclaration) CounterIncrement() string {
	return s.Get("counterIncrement").String()
}

func (s *CSSStyleDeclaration) SetCounterIncrement(v string) {
	s.Set("counterIncrement", v)
}

func (s *CSSStyleDeclaration) CounterReset() string {
	return s.Get("counterReset").String()
}

func (s *CSSStyleDeclaration) SetCounterReset(v string) {
	s.Set("counterReset", v)
}

func (s *CSSStyleDeclaration) Cursor() string {
	return s.Get("cursor").String()
}

func (s *CSSStyleDeclaration) SetCursor(v string) {
	s.Set("cursor", v)
}

func (s *CSSStyleDeclaration) Direction() string {
	return s.Get("direction").String()
}

func (s *CSSStyleDeclaration) SetDirection(v string) {
	s.Set("direction", v)
}

func (s *CSSStyleDeclaration) Display() string {
	return s.Get("display").String()
}

func (s *CSSStyleDeclaration) SetDisplay(v string) {
	s.Set("display", v)
}

func (s *CSSStyleDeclaration) EmptyCells() string {
	return s.Get("emptyCells").String()
}

func (s *CSSStyleDeclaration) SetEmptyCells(v string) {
	s.Set("emptyCells", v)
}

func (s *CSSStyleDeclaration) Filter() string {
	return s.Get("filter").String()
}

func (s *CSSStyleDeclaration) SetFilter(v string) {
	s.Set("filter", v)
}

func (s *CSSStyleDeclaration) Flex() string {
	return s.Get("flex").String()
}

func (s *CSSStyleDeclaration) SetFlex(v string) {
	s.Set("flex", v)
}

func (s *CSSStyleDeclaration) FlexBasis() string {
	return s.Get("flexBasis").String()
}

func (s *CSSStyleDeclaration) SetFlexBasis(v string) {
	s.Set("flexBasis", v)
}

func (s *CSSStyleDeclaration) FlexDirection() string {
	return s.Get("flexDirection").String()
}

func (s *CSSStyleDeclaration) SetFlexDirection(v string) {
	s.Set("flexDirection", v)
}

func (s *CSSStyleDeclaration) FlexFlow() string {
	return s.Get("flexFlow").String()
}

func (s *CSSStyleDeclaration) SetFlexFlow(v string) {
	s.Set("flexFlow", v)
}

func (s *CSSStyleDeclaration) FlexGrow() string {
	return s.Get("flexGrow").String()
}

func (s *CSSStyleDeclaration) SetFlexGrow(v string) {
	s.Set("flexGrow", v)
}

func (s *CSSStyleDeclaration) FlexShrink() string {
	return s.Get("flexShrink").String()
}

func (s *CSSStyleDeclaration) SetFlexShrink(v string) {
	s.Set("flexShrink", v)
}

func (s *CSSStyleDeclaration) FlexWrap() string {
	return s.Get("flexWrap").String()
}

func (s *CSSStyleDeclaration) SetFlexWrap(v string) {
	s.Set("flexWrap", v)
}

func (s *CSSStyleDeclaration) CssFloat() string {
	return s.Get("cssFloat").String()
}

func (s *CSSStyleDeclaration) SetCssFloat(v string) {
	s.Set("cssFloat", v)
}

func (s *CSSStyleDeclaration) Font() string {
	return s.Get("font").String()
}

func (s *CSSStyleDeclaration) SetFont(v string) {
	s.Set("font", v)
}

func (s *CSSStyleDeclaration) FontFamily() string {
	return s.Get("fontFamily").String()
}

func (s *CSSStyleDeclaration) SetFontFamily(v string) {
	s.Set("fontFamily", v)
}

func (s *CSSStyleDeclaration) FontSize() string {
	return s.Get("fontSize").String()
}

func (s *CSSStyleDeclaration) SetFontSize(v string) {
	s.Set("fontSize", v)
}

func (s *CSSStyleDeclaration) FontStyle() string {
	return s.Get("fontStyle").String()
}

func (s *CSSStyleDeclaration) SetFontStyle(v string) {
	s.Set("fontStyle", v)
}

func (s *CSSStyleDeclaration) FontVariant() string {
	return s.Get("fontVariant").String()
}

func (s *CSSStyleDeclaration) SetFontVariant(v string) {
	s.Set("fontVariant", v)
}

func (s *CSSStyleDeclaration) FontWeight() string {
	return s.Get("fontWeight").String()
}

func (s *CSSStyleDeclaration) SetFontWeight(v string) {
	s.Set("fontWeight", v)
}

func (s *CSSStyleDeclaration) FontSizeAdjust() string {
	return s.Get("fontSizeAdjust").String()
}

func (s *CSSStyleDeclaration) SetFontSizeAdjust(v string) {
	s.Set("fontSizeAdjust", v)
}

func (s *CSSStyleDeclaration) Height() string {
	return s.Get("height").String()
}

func (s *CSSStyleDeclaration) SetHeight(v string) {
	s.Set("height", v)
}

func (s *CSSStyleDeclaration) JustifyContent() string {
	return s.Get("justifyContent").String()
}

func (s *CSSStyleDeclaration) SetJustifyContent(v string) {
	s.Set("justifyContent", v)
}

func (s *CSSStyleDeclaration) Left() string {
	return s.Get("left").String()
}

func (s *CSSStyleDeclaration) SetLeft(v string) {
	s.Set("left", v)
}

func (s *CSSStyleDeclaration) LetterSpacing() string {
	return s.Get("letterSpacing").String()
}

func (s *CSSStyleDeclaration) SetLetterSpacing(v string) {
	s.Set("letterSpacing", v)
}

func (s *CSSStyleDeclaration) LineHeight() string {
	return s.Get("lineHeight").String()
}

func (s *CSSStyleDeclaration) SetLineHeight(v string) {
	s.Set("lineHeight", v)
}

func (s *CSSStyleDeclaration) ListStyle() string {
	return s.Get("listStyle").String()
}

func (s *CSSStyleDeclaration) SetListStyle(v string) {
	s.Set("listStyle", v)
}

func (s *CSSStyleDeclaration) ListStyleImage() string {
	return s.Get("listStyleImage").String()
}

func (s *CSSStyleDeclaration) SetListStyleImage(v string) {
	s.Set("listStyleImage", v)
}

func (s *CSSStyleDeclaration) ListStylePosition() string {
	return s.Get("listStylePosition").String()
}

func (s *CSSStyleDeclaration) SetListStylePosition(v string) {
	s.Set("listStylePosition", v)
}

func (s *CSSStyleDeclaration) ListStyleType() string {
	return s.Get("listStyleType").String()
}

func (s *CSSStyleDeclaration) SetListStyleType(v string) {
	s.Set("listStyleType", v)
}

func (s *CSSStyleDeclaration) Margin() string {
	return s.Get("margin").String()
}

func (s *CSSStyleDeclaration) SetMargin(v string) {
	s.Set("margin", v)
}

func (s *CSSStyleDeclaration) MarginBottom() string {
	return s.Get("marginBottom").String()
}

func (s *CSSStyleDeclaration) SetMarginBottom(v string) {
	s.Set("marginBottom", v)
}

func (s *CSSStyleDeclaration) MarginLeft() string {
	return s.Get("marginLeft").String()
}

func (s *CSSStyleDeclaration) SetMarginLeft(v string) {
	s.Set("marginLeft", v)
}

func (s *CSSStyleDeclaration) MarginRight() string {
	return s.Get("marginRight").String()
}

func (s *CSSStyleDeclaration) SetMarginRight(v string) {
	s.Set("marginRight", v)
}

func (s *CSSStyleDeclaration) MarginTop() string {
	return s.Get("marginTop").String()
}

func (s *CSSStyleDeclaration) SetMarginTop(v string) {
	s.Set("marginTop", v)
}

func (s *CSSStyleDeclaration) MaxHeight() string {
	return s.Get("maxHeight").String()
}

func (s *CSSStyleDeclaration) SetMaxHeight(v string) {
	s.Set("maxHeight", v)
}

func (s *CSSStyleDeclaration) MaxWidth() string {
	return s.Get("maxWidth").String()
}

func (s *CSSStyleDeclaration) SetMaxWidth(v string) {
	s.Set("maxWidth", v)
}

func (s *CSSStyleDeclaration) MinHeight() string {
	return s.Get("minHeight").String()
}

func (s *CSSStyleDeclaration) SetMinHeight(v string) {
	s.Set("minHeight", v)
}

func (s *CSSStyleDeclaration) MinWidth() string {
	return s.Get("minWidth").String()
}

func (s *CSSStyleDeclaration) SetMinWidth(v string) {
	s.Set("minWidth", v)
}

func (s *CSSStyleDeclaration) Opacity() string {
	return s.Get("opacity").String()
}

func (s *CSSStyleDeclaration) SetOpacity(v string) {
	s.Set("opacity", v)
}

func (s *CSSStyleDeclaration) Order() string {
	return s.Get("order").String()
}

func (s *CSSStyleDeclaration) SetOrder(v string) {
	s.Set("order", v)
}

func (s *CSSStyleDeclaration) Orphans() string {
	return s.Get("orphans").String()
}

func (s *CSSStyleDeclaration) SetOrphans(v string) {
	s.Set("orphans", v)
}

func (s *CSSStyleDeclaration) Outline() string {
	return s.Get("outline").String()
}

func (s *CSSStyleDeclaration) SetOutline(v string) {
	s.Set("outline", v)
}

func (s *CSSStyleDeclaration) OutlineColor() string {
	return s.Get("outlineColor").String()
}

func (s *CSSStyleDeclaration) SetOutlineColor(v string) {
	s.Set("outlineColor", v)
}

func (s *CSSStyleDeclaration) OutlineOffset() string {
	return s.Get("outlineOffset").String()
}

func (s *CSSStyleDeclaration) SetOutlineOffset(v string) {
	s.Set("outlineOffset", v)
}

func (s *CSSStyleDeclaration) OutlineStyle() string {
	return s.Get("outlineStyle").String()
}

func (s *CSSStyleDeclaration) SetOutlineStyle(v string) {
	s.Set("outlineStyle", v)
}

func (s *CSSStyleDeclaration) OutlineWidth() string {
	return s.Get("outlineWidth").String()
}

func (s *CSSStyleDeclaration) SetOutlineWidth(v string) {
	s.Set("outlineWidth", v)
}

func (s *CSSStyleDeclaration) Overflow() string {
	return s.Get("overflow").String()
}

func (s *CSSStyleDeclaration) SetOverflow(v string) {
	s.Set("overflow", v)
}

func (s *CSSStyleDeclaration) OverflowX() string {
	return s.Get("overflowX").String()
}

func (s *CSSStyleDeclaration) SetOverflowX(v string) {
	s.Set("overflowX", v)
}

func (s *CSSStyleDeclaration) OverflowY() string {
	return s.Get("overflowY").String()
}

func (s *CSSStyleDeclaration) SetOverflowY(v string) {
	s.Set("overflowY", v)
}

func (s *CSSStyleDeclaration) Padding() string {
	return s.Get("padding").String()
}

func (s *CSSStyleDeclaration) SetPadding(v string) {
	s.Set("padding", v)
}

func (s *CSSStyleDeclaration) PaddingBottom() string {
	return s.Get("paddingBottom").String()
}

func (s *CSSStyleDeclaration) SetPaddingBottom(v string) {
	s.Set("paddingBottom", v)
}

func (s *CSSStyleDeclaration) PaddingLeft() string {
	return s.Get("paddingLeft").String()
}

func (s *CSSStyleDeclaration) SetPaddingLeft(v string) {
	s.Set("paddingLeft", v)
}

func (s *CSSStyleDeclaration) PaddingRight() string {
	return s.Get("paddingRight").String()
}

func (s *CSSStyleDeclaration) SetPaddingRight(v string) {
	s.Set("paddingRight", v)
}

func (s *CSSStyleDeclaration) PaddingTop() string {
	return s.Get("paddingTop").String()
}

func (s *CSSStyleDeclaration) SetPaddingTop(v string) {
	s.Set("paddingTop", v)
}

func (s *CSSStyleDeclaration) PageBreakAfter() string {
	return s.Get("pageBreakAfter").String()
}

func (s *CSSStyleDeclaration) SetPageBreakAfter(v string) {
	s.Set("pageBreakAfter", v)
}

func (s *CSSStyleDeclaration) PageBreakBefore() string {
	return s.Get("pageBreakBefore").String()
}

func (s *CSSStyleDeclaration) SetPageBreakBefore(v string) {
	s.Set("pageBreakBefore", v)
}

func (s *CSSStyleDeclaration) PageBreakInside() string {
	return s.Get("pageBreakInside").String()
}

func (s *CSSStyleDeclaration) SetPageBreakInside(v string) {
	s.Set("pageBreakInside", v)
}

func (s *CSSStyleDeclaration) Perspective() string {
	return s.Get("perspective").String()
}

func (s *CSSStyleDeclaration) SetPerspective(v string) {
	s.Set("perspective", v)
}

func (s *CSSStyleDeclaration) PerspectiveOrigin() string {
	return s.Get("perspectiveOrigin").String()
}

func (s *CSSStyleDeclaration) SetPerspectiveOrigin(v string) {
	s.Set("perspectiveOrigin", v)
}

func (s *CSSStyleDeclaration) Position() string {
	return s.Get("position").String()
}

func (s *CSSStyleDeclaration) SetPosition(v string) {
	s.Set("position", v)
}

func (s *CSSStyleDeclaration) Quotes() string {
	return s.Get("quotes").String()
}

func (s *CSSStyleDeclaration) SetQuotes(v string) {
	s.Set("quotes", v)
}

func (s *CSSStyleDeclaration) Resize() string {
	return s.Get("resize").String()
}

func (s *CSSStyleDeclaration) SetResize(v string) {
	s.Set("resize", v)
}

func (s *CSSStyleDeclaration) Right() string {
	return s.Get("right").String()
}

func (s *CSSStyleDeclaration) SetRight(v string) {
	s.Set("right", v)
}

func (s *CSSStyleDeclaration) TableLayout() string {
	return s.Get("tableLayout").String()
}

func (s *CSSStyleDeclaration) SetTableLayout(v string) {
	s.Set("tableLayout", v)
}

func (s *CSSStyleDeclaration) TabSize() string {
	return s.Get("tabSize").String()
}

func (s *CSSStyleDeclaration) SetTabSize(v string) {
	s.Set("tabSize", v)
}

func (s *CSSStyleDeclaration) TextAlign() string {
	return s.Get("textAlign").String()
}

func (s *CSSStyleDeclaration) SetTextAlign(v string) {
	s.Set("textAlign", v)
}

func (s *CSSStyleDeclaration) TextAlignLast() string {
	return s.Get("textAlignLast").String()
}

func (s *CSSStyleDeclaration) SetTextAlignLast(v string) {
	s.Set("textAlignLast", v)
}

func (s *CSSStyleDeclaration) TextDecoration() string {
	return s.Get("textDecoration").String()
}

func (s *CSSStyleDeclaration) SetTextDecoration(v string) {
	s.Set("textDecoration", v)
}

func (s *CSSStyleDeclaration) TextDecorationColor() string {
	return s.Get("textDecorationColor").String()
}

func (s *CSSStyleDeclaration) SetTextDecorationColor(v string) {
	s.Set("textDecorationColor", v)
}

func (s *CSSStyleDeclaration) TextDecorationLine() string {
	return s.Get("textDecorationLine").String()
}

func (s *CSSStyleDeclaration) SetTextDecorationLine(v string) {
	s.Set("textDecorationLine", v)
}

func (s *CSSStyleDeclaration) TextDecorationStyle() string {
	return s.Get("textDecorationStyle").String()
}

func (s *CSSStyleDeclaration) SetTextDecorationStyle(v string) {
	s.Set("textDecorationStyle", v)
}

func (s *CSSStyleDeclaration) TextIndent() string {
	return s.Get("textIndent").String()
}

func (s *CSSStyleDeclaration) SetTextIndent(v string) {
	s.Set("textIndent", v)
}

func (s *CSSStyleDeclaration) TextOverflow() string {
	return s.Get("textOverflow").String()
}

func (s *CSSStyleDeclaration) SetTextOverflow(v string) {
	s.Set("textOverflow", v)
}

func (s *CSSStyleDeclaration) TextShadow() string {
	return s.Get("textShadow").String()
}

func (s *CSSStyleDeclaration) SetTextShadow(v string) {
	s.Set("textShadow", v)
}

func (s *CSSStyleDeclaration) TextTransform() string {
	return s.Get("textTransform").String()
}

func (s *CSSStyleDeclaration) SetTextTransform(v string) {
	s.Set("textTransform", v)
}

func (s *CSSStyleDeclaration) Top() string {
	return s.Get("top").String()
}

func (s *CSSStyleDeclaration) SetTop(v string) {
	s.Set("top", v)
}

func (s *CSSStyleDeclaration) Transform() string {
	return s.Get("transform").String()
}

func (s *CSSStyleDeclaration) SetTransform(v string) {
	s.Set("transform", v)
}

func (s *CSSStyleDeclaration) TransformOrigin() string {
	return s.Get("transformOrigin").String()
}

func (s *CSSStyleDeclaration) SetTransformOrigin(v string) {
	s.Set("transformOrigin", v)
}

func (s *CSSStyleDeclaration) TransformStyle() string {
	return s.Get("transformStyle").String()
}

func (s *CSSStyleDeclaration) SetTransformStyle(v string) {
	s.Set("transformStyle", v)
}

func (s *CSSStyleDeclaration) Transition() string {
	return s.Get("transition").String()
}

func (s *CSSStyleDeclaration) SetTransition(v string) {
	s.Set("transition", v)
}

func (s *CSSStyleDeclaration) TransitionProperty() string {
	return s.Get("transitionProperty").String()
}

func (s *CSSStyleDeclaration) SetTransitionProperty(v string) {
	s.Set("transitionProperty", v)
}

func (s *CSSStyleDeclaration) TransitionDuration() string {
	return s.Get("transitionDuration").String()
}

func (s *CSSStyleDeclaration) SetTransitionDuration(v string) {
	s.Set("transitionDuration", v)
}

func (s *CSSStyleDeclaration) TransitionTimingFunction() string {
	return s.Get("transitionTimingFunction").String()
}

func (s *CSSStyleDeclaration) SetTransitionTimingFunction(v string) {
	s.Set("transitionTimingFunction", v)
}

func (s *CSSStyleDeclaration) TransitionDelay() string {
	return s.Get("transitionDelay").String()
}

func (s *CSSStyleDeclaration) SetTransitionDelay(v string) {
	s.Set("transitionDelay", v)
}

func (s *CSSStyleDeclaration) UnicodeBidi() string {
	return s.Get("unicodeBidi").String()
}

func (s *CSSStyleDeclaration) SetUnicodeBidi(v string) {
	s.Set("unicodeBidi", v)
}

func (s *CSSStyleDeclaration) UserSelect() string {
	return s.Get("userSelect").String()
}

func (s *CSSStyleDeclaration) SetUserSelect(v string) {
	s.Set("userSelect", v)
}

func (s *CSSStyleDeclaration) VerticalAlign() string {
	return s.Get("verticalAlign").String()
}

func (s *CSSStyleDeclaration) SetVerticalAlign(v string) {
	s.Set("verticalAlign", v)
}

func (s *CSSStyleDeclaration) Visibility() string {
	return s.Get("visibility").String()
}

func (s *CSSStyleDeclaration) SetVisibility(v string) {
	s.Set("visibility", v)
}

func (s *CSSStyleDeclaration) WhiteSpace() string {
	return s.Get("whiteSpace").String()
}

func (s *CSSStyleDeclaration) SetWhiteSpace(v string) {
	s.Set("whiteSpace", v)
}

func (s *CSSStyleDeclaration) Width() string {
	return s.Get("width").String()
}

func (s *CSSStyleDeclaration) SetWidth(v string) {
	s.Set("width", v)
}

func (s *CSSStyleDeclaration) WordBreak() string {
	return s.Get("wordBreak").String()
}

func (s *CSSStyleDeclaration) SetWordBreak(v string) {
	s.Set("wordBreak", v)
}

func (s *CSSStyleDeclaration) WordSpacing() string {
	return s.Get("wordSpacing").String()
}

func (s *CSSStyleDeclaration) SetWordSpacing(v string) {
	s.Set("wordSpacing", v)
}

func (s *CSSStyleDeclaration) WordWrap() string {
	return s.Get("wordWrap").String()
}

func (s *CSSStyleDeclaration) SetWordWrap(v string) {
	s.Set("wordWrap", v)
}

func (s *CSSStyleDeclaration) Widows() string {
	return s.Get("widows").String()
}

func (s *CSSStyleDeclaration) SetWidows(v string) {
	s.Set("widows", v)
}

func (s *CSSStyleDeclaration) ZIndex() string {
	return s.Get("zIndex").String()
}

func (s *CSSStyleDeclaration) SetZIndex(v string) {
	s.Set("zIndex", v)
}
