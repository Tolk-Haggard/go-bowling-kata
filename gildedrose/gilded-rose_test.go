package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GildedRose_BasicItemDecrementsSellIn(t *testing.T) {
	items = []Item{{"+5 Dexterity Vest", 10, 20}}

	GlidedRose()

	assert.Equal(t, 9, items[0].sellIn)
}

func Test_GildedRose_BasicItemDecrementsQuality(t *testing.T) {
	items = []Item{{"+5 Dexterity Vest", 10, 20}}

	GlidedRose()

	assert.Equal(t, 19, items[0].quality)
}

func Test_GildedRose_BasicQualityDoesNotGoBelowZero(t *testing.T) {
	items = []Item{{"+5 Dexterity Vest", 10, 0}}

	GlidedRose()

	assert.Equal(t, 0, items[0].quality)
}

func Test_GildedRose_BasicQualityDecrementsTwiceAsFastAfterSellIn(t *testing.T) {
	items = []Item{{"+5 Dexterity Vest", 0, 20}}

	GlidedRose()

	assert.Equal(t, 18, items[0].quality)
}

func Test_GildedRose_AgedBrieIncreasesQuality(t *testing.T) {
	items = []Item{{"Aged Brie", 10, 20}}

	GlidedRose()

	assert.Equal(t, 21, items[0].quality)
}

func Test_GildedRose_AgedBrieIncreasesQualityTwiceAsFastAfterSellIn(t *testing.T) {
	items = []Item{{"Aged Brie", 0, 20}}

	GlidedRose()

	assert.Equal(t, 22, items[0].quality)
}

func Test_GildedRose_AgedBrieQualityDoesNotGoAbove50(t *testing.T) {
	items = []Item{{"Aged Brie", 0, 50}}

	GlidedRose()

	assert.Equal(t, 50, items[0].quality)
}

func Test_GildedRose_SulfurasDoesNotChange(t *testing.T) {
	items = []Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}

	GlidedRose()

	assert.Equal(t, 80, items[0].quality)
	assert.Equal(t, 0, items[0].sellIn)
}

func Test_GildedRose_TicketsIncreaseInQuality(t *testing.T) {
	items = []Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 20}}

	GlidedRose()

	assert.Equal(t, 21, items[0].quality)
}

func Test_GildedRose_TicketsIncreaseInQualityTwiceAsFastWhen10DaysLeft(t *testing.T) {
	items = []Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 20}}

	GlidedRose()

	assert.Equal(t, 22, items[0].quality)
}

func Test_GildedRose_TicketsIncreaseInQualityThriceAsFastWhen5DaysLeft(t *testing.T) {
	items = []Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 20}}

	GlidedRose()

	assert.Equal(t, 23, items[0].quality)
}

func Test_GildedRose_TicketsDropToZeroAfterConcert(t *testing.T) {
	items = []Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 20}}

	GlidedRose()

	assert.Equal(t, 0, items[0].quality)
}

func Test_GildedRose_TicketsQualityDoesNotGoAbove50(t *testing.T) {
	items = []Item{{"Backstage passes to a TAFKAL80ETC concert", 2, 50}}

	GlidedRose()

	assert.Equal(t, 50, items[0].quality)
}
