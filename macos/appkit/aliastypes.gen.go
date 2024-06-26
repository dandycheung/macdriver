// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A closure that configures and returns a collection view’s supplementary view, such as a header or footer, from a diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasourcesupplementaryviewprovider?language=objc
type CollectionViewDiffableDataSourceSupplementaryViewProvider = func(arg0 CollectionView, arg1 string, arg2 foundation.IndexPath) View

// A closure that configures and returns a row view for a table view from its diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasourcerowprovider?language=objc
type TableViewDiffableDataSourceRowProvider = func(tableView TableView, row int, identifier objc.Object) TableRowView

// A block that you use to handle the custom creation of controller objects from your storyboard file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstoryboardcontrollercreator?language=objc
type StoryboardControllerCreator = func(coder foundation.Coder) objc.Object

// A closure that creates and returns each of the layout's sections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutsectionprovider?language=objc
type CollectionViewCompositionalLayoutSectionProvider = func(section int, arg1 CollectionLayoutEnvironmentObject) CollectionLayoutSection

// A closure that configures and returns a cell for a table view from its diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasourcecellprovider?language=objc
type TableViewDiffableDataSourceCellProvider = func(tableView TableView, column TableColumn, row int, itemId objc.Object) View

// A closure that configures and returns a section header view for a table view from its diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasourcesectionheaderviewprovider?language=objc
type TableViewDiffableDataSourceSectionHeaderViewProvider = func(tableView TableView, row int, sectionId objc.Object) View

// A closure that configures and returns an item for a collection view from its diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasourceitemprovider?language=objc
type CollectionViewDiffableDataSourceItemProvider = func(arg0 CollectionView, arg1 foundation.IndexPath, arg2 objc.Object) CollectionViewItem

// A closure called before each layout cycle to allow modification of items in a section immediately before they’re displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsectionvisibleitemsinvalidationhandler?language=objc
type CollectionLayoutSectionVisibleItemsInvalidationHandler = func(visibleItems []CollectionLayoutVisibleItemObject, contentOffset foundation.Point, layoutEnvironment CollectionLayoutEnvironmentObject)

// A closure that creates and returns each of the custom group’s items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroupcustomitemprovider?language=objc
type CollectionLayoutGroupCustomItemProvider = func(layoutEnvironment CollectionLayoutEnvironmentObject) []CollectionLayoutGroupCustomItem
