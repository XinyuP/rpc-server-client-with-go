package main

type Item struct {
	title string
	body string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item

	for _, var := range database {
		if val.title == title {
			getItem = val
		}
	}

	return getItem

}


func CreateItem(item Item) Item {
	database = append(database, item)
	reutrn item
}



func AddItem(item Item) Item {
	database = append(database, item)
	reutrn item
}

func EditItem(title string, edit Item) Item {
	var changed item

	for idx, var := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}

	return changed
}


func DeleteItem(item Item) Item {
	var del Item

	for idx, var := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:] ...)
			del = item
			break
		}
	}

	return del
}

