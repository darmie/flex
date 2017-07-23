package flex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measure3(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount, ok := YGNodeGetContext(node).(*int)
	if ok {
		(*measureCount)++
	}

	return Size{Width: 10, Height: 10}
}

func _simulate_wrapping_text(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if widthMode == MeasureModeUndefined || width >= 68 {
		return Size{Width: 68, Height: 16}
	}

	return Size{Width: 50, Height: 32}
}

func _measure_assert_negative(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if width < 0 {
		panic(fmt.Sprintf("width is %.2f and should be >= 0", width))
	}
	if height < 0 {
		panic(fmt.Sprintf("height is %.2f should be >= 0, height", height))
	}
	// EXPECT_GE(width, 0);
	//EXPECT_GE(height, 0);

	return Size{
		Width: 0, Height: 0,
	}
}

func TestDont_measure_single_grow_shrink_child(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure)
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
}

func TestMeasure_absolute_child_with_no_constraints(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	YGNodeInsertChild(root, rootChild0, 0)

	measureCount := 0

	rootChild0Child0 := NewNode()
	NodeStyleSetPositionType(rootChild0Child0, PositionTypeAbsolute)
	rootChild0Child0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0Child0, _measure3)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, measureCount)
}

func TestDont_measure_when_min_equals_max(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidth(rootChild0, 10)
	YGNodeStyleSetMaxWidth(rootChild0, 10)
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_percentages(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidthPercent(rootChild0, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 10)
	YGNodeStyleSetMinHeightPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_width_percent(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidthPercent(rootChild0, 10)
	YGNodeStyleSetMaxWidthPercent(rootChild0, 10)
	YGNodeStyleSetMinHeight(rootChild0, 10)
	YGNodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_height_percent(t *testing.T) {
	root := NewNode()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	YGNodeStyleSetMinWidth(rootChild0, 10)
	YGNodeStyleSetMaxWidth(rootChild0, 10)
	YGNodeStyleSetMinHeightPercent(rootChild0, 10)
	YGNodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_enough_size_should_be_in_single_line(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 68, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 16, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_not_enough_size_should_wrap(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 55)

	rootChild0 := NewNode()
	YGNodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))
}

func TestMeasure_zero_space_should_grow(t *testing.T) {
	root := NewNode()
	NodeStyleSetHeight(root, 200)
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	YGNodeStyleSetFlexGrow(root, 0)

	measureCount := 0

	rootChild0 := NewNode()
	YGNodeStyleSetFlexDirection(rootChild0, FlexDirectionColumn)
	YGNodeStyleSetPadding(rootChild0, EdgeAll, 100)
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 282, Undefined, DirectionLTR)

	assertFloatEqual(t, 282, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
}

func TestMeasure_flex_direction_row_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetPadding(root, EdgeLeft, 25)
	YGNodeStyleSetPadding(root, EdgeTop, 25)
	YGNodeStyleSetPadding(root, EdgeRight, 25)
	YGNodeStyleSetPadding(root, EdgeBottom, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_column_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 57, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding_align_items_flexstart(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 35, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_flex_shrink(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	YGNodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

func TestMeasure_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 32, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetHeight(rootChild1))
}

/*
#if GTEST_HAS_DEATH_TEST
TEST(YogaDeathTest, cannot_add_child_to_node_with_measure_func) {
  root := YGNodeNew();
  YGNodeSetMeasureFunc(root, _measure3);

  rootChild0 := YGNodeNew();
  ASSERT_DEATH(YGNodeInsertChild(root, rootChild0, 0), "Cannot add child.*");
  YGNodeFree(rootChild0);
  ;
}

TEST(YogaDeathTest, cannot_add_nonnull_measure_func_to_non_leaf_node) {
  root := YGNodeNew();
  rootChild0 := YGNodeNew();
  YGNodeInsertChild(root, rootChild0, 0);

  ASSERT_DEATH(YGNodeSetMeasureFunc(root, _measure3), "Cannot set measure function.*");
  ;
}
#endif
*/

func TestCan_nullify_measure_func_on_any_node(t *testing.T) {
	root := NewNode()
	YGNodeInsertChild(root, NewNode(), 0)

	NodeSetMeasureFunc(root, nil)
	assert.True(t, root.Measure == nil)
}

func TestCant_call_negative_measure(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 10)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	YGNodeStyleSetMargin(rootChild0, EdgeTop, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}

func TestCant_call_negative_measure_horizontal(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	YGNodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 10)
	NodeStyleSetHeight(root, 20)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	YGNodeStyleSetMargin(rootChild0, EdgeStart, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}
