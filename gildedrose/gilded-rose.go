package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	GlidedRose()
}

func GlidedRose() {
	for i := 0; i < len(items); i++ {

		if items[i].name == "Sulfuras, Hand of Ragnaros" {
			continue
		}
		items[i].sellIn = items[i].sellIn - 1

		if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			items[i].quality = increaseQuality(items[i])
			if items[i].sellIn < 10 {
				items[i].quality = increaseQuality(items[i])
			}
			if items[i].sellIn < 5 {
				items[i].quality = increaseQuality(items[i])
			}
			if items[i].sellIn < 0 {
				items[i].quality = 0
			}
		} else if items[i].name == "Aged Brie" {
			items[i].quality = increaseQuality(items[i])
			if items[i].sellIn < 0 {
				items[i].quality = increaseQuality(items[i])
			}
		} else {
			items[i].quality = decreaseQuality(items[i])
			if items[i].sellIn < 0 {
				items[i].quality = decreaseQuality(items[i])
			}
		}
	}
}

func increaseQuality(item Item) int {
	if item.quality < 50 {
		return item.quality + 1
	}
	return item.quality
}

func decreaseQuality(item Item) int {
	if item.quality > 0 {
		return item.quality - 1
	}
	return 0
}
