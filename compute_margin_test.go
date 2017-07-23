package flex

import "testing"

func TestComputed_layout_margin(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)
	YGNodeStyleSetMarginPercent(root, EdgeStart, 10)

	NodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 10, NodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeRight))

	NodeCalculateLayout(root, 100, 100, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 10, NodeLayoutGetMargin(root, EdgeRight))
}
