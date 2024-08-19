package models

// Item adalah struktur data untuk item
type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price string `json:"price"`
}

// GetItems mendapatkan semua item dari database
func GetItems() ([]Item, error) {
    rows, err := db.Query("SELECT id, name, price FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

// GetItem mendapatkan item berdasarkan ID dari database
func GetItem(id string) (Item, error) {
    var item Item
    row := db.QueryRow("SELECT id, name, price FROM items WHERE id = ?", id)
    err := row.Scan(&item.ID, &item.Name, &item.Price)
    if err != nil {
        return item, err
    }
    return item, nil
}

// CreateItem menambahkan item baru ke database
func CreateItem(item Item) error {
    _, err := db.Exec("INSERT INTO items (id, name, price) VALUES (?, ?, ?)", item.ID, item.Name, item.Price)
    return err
}

// UpdateItem memperbarui item yang sudah ada di database
func UpdateItem(item Item) error {
    _, err := db.Exec("UPDATE items SET name = ?, price = ? WHERE id = ?", item.Name, item.Price, item.ID)
    return err
}

// DeleteItem menghapus item berdasarkan ID dari database
func DeleteItem(id string) error {
    _, err := db.Exec("DELETE FROM items WHERE id = ?", id)
    return err
}
