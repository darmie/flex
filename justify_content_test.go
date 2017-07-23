package flex

import "testing"

func TestJustify_content_row_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 92, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 82, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetJustifyContent(root, JustifyFlexEnd)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 72, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 82, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 92, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 20, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_center(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 36, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 56, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 56, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 36, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_space_between(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetJustifyContent(root, JustifySpaceBetween)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 92, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 92, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_row_space_around(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetJustifyContent(root, JustifySpaceAround)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 12, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 46, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 12, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_flex_start(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_flex_end(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyFlexEnd)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 82, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 92, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 72, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 82, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 92, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_center(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifyCenter)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 36, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 56, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 36, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 56, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_space_between(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifySpaceBetween)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 92, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 92, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}

func TestJustify_content_column_space_around(t *testing.T) {
	config := YGConfigNew()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetJustifyContent(root, JustifySpaceAround)
	YGNodeStyleSetWidth(root, 102)
	YGNodeStyleSetHeight(root, 102)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild1, 10)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	YGNodeStyleSetHeight(rootChild2, 10)
	YGNodeInsertChild(root, rootChild2, 2)
	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 12, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))

	YGNodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 102, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 12, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 46, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 102, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild2))
}
