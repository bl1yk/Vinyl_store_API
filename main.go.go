package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product представляет продукт
type Product struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Artist    string   `json:"artist"`
	Cover     string   `json:"cover"`
	Price     float64  `json:"price"`
	Info      string   `json:"info"`
	Tracklist []string `json:"tracklist"`
	Listen    string   `json:"listen"`
	Favourite bool     `json:"favourite"`
	Quantity  int      `json:"quantity"`
}

// Пример списка продуктов
var products = []Product{
	{ID: 1, Name: "The Piper at the Gates of Dawn", Artist: "Pink Floyd", Cover: "assets/images/PinkFloyd_PiperAtTheGatesOfDawn.jpg", Price: 2590, Info: "Дебютный студийный альбом группы Pink Floyd, выпущенный в 1967 году и единственный, записанный под руководством Сида Барретта, который был основным автором песен и творческим лидером группы в тот период. По общему признанию, альбом оказал огромное влияние на становление и развитие психоделического рока.", Tracklist: []string{"1. Astronomy Domine	04:12",
		"2. Lucifer Sam	03:07",
		"3. Matilda Mother	03:08",
		"4. Flaming	02:46",
		"5. Pow R. Toc H.	04:26",
		"6. Take Up Thy Stethoscope and Walk	03:05",
		"7. Interstellar Overdrive	09:41",
		"8. The Gnome	02:13",
		"9. Chapter 24	03:42",
		"10. Scarecrow	02:11",
		"11. Bike	03:21"}, Listen: "https://open.spotify.com/album/0Fke5eiQ6lszQHlwiFygqn?si=V3O3O88oQWOomT_jpR3Chw", Favourite: false, Quantity: 0},
	{ID: 2, Name: "A Saucerful of Secrets", Artist: "Pink Floyd", Cover: "assets/images/saucer.jpg", Price: 3100, Info: "Второй студийный альбом британской рок-группы Pink Floyd, выпущенный 29 июня 1968 года лейблом Columbia, принадлежащим EMI, в Великобритании и 27 июля того же года лейблом Capitol Records в США.", Tracklist: []string{"1. Let There Be More Light	05:38",
		"2. Remember a Day	04:33",
		"3. Set the Controls for the Heart of the Sun	05:28",
		"4. Corporal Clegg	04:13",
		"5. A Saucerful of Secrets	11:52",
		"6. See-Saw	04:36",
		"7. Jugband Blues	03:00"}, Listen: "https://open.spotify.com/album/5rwuexO7oiRJKqzZrd1upQ?si=g-sZD5kORtmHbgcRA1NGKA", Favourite: false, Quantity: 0},
	{ID: 3, Name: "More", Artist: "Pink Floyd", Cover: "assets/images/More_Pink_Floyd_Cover.jpg", Price: 2999, Info: "Первый саундтрек и третий студийный альбом британской прогрессивной рок группы Pink Floyd, выпущенный 13 июня 1969 года на лейбле EMI Columbia в Великобритании и 9 августа 1969 года на лейбле Tower Records в США.", Tracklist: []string{"1. Cirrus Minor	 	05:18",
		"2. The Nile Song	 	03:26",
		"3. Crying Song	 	03:33",
		"4. Up the Khyber	 	02:12",
		"5. Green Is the Colour	 	02:58",
		"6. Cymbaline	 	04:50",
		"7. Party Sequence	 	01:07",
		"8. Main Theme	 	05:27",
		"9. Ibiza Bar	 	03:19",
		"10. More Blues	 	02:12",
		"11. Quicksilver	 	07:13",
		"12. A Spanish Piece	 	01:05",
		"13. Dramatic Theme	 	02:15"}, Listen: "https://open.spotify.com/album/6xWRCFsaqoYj3ZwxMkCE85?si=2kQa10opQQOTqwi4bq85Vw", Favourite: false, Quantity: 0},
	{ID: 4, Name: "Ummagumma", Artist: "Pink Floyd", Cover: "assets/images/Ummagumma.jpeg", Price: 4990, Info: "Четвёртый альбом британской рок-группы Pink Floyd, выпущенный 7 ноября 1969 года лейблом Harvest Records. Это двойной альбом, содержащий концертные и студийные треки.", Tracklist: []string{"1. Astronomy Domine	 	08:29",
		"2. Careful with That Axe, Eugene	 	08:50",
		"3. Set the Controls for the Heart of the Sun	 	09:12",
		"4. Saucerful of Secrets	 	12:48",
		"5. Sysyphus	 	12:59",
		"6. Grantchester Meadows	 	07:19",
		"7. Several Species of Small Furry Animals Gathered Together in a Cave and Grooving With a Pict	 	05:01",
		"8. The Narrow Way	 	12:17",
		"9. The Grand Vizier's Garden Party	 	08:44"}, Listen: "https://open.spotify.com/album/4uuyGVEZpYaGB1HtewAogW?si=7vweUI7USs-RG9KOcRoeww", Favourite: false, Quantity: 0},
	{ID: 5, Name: "Atom Heart Mother", Artist: "Pink Floyd", Cover: "Atom Heart Mother", Price: 3199, Info: "Пятый студийный альбом британской рок-группы Pink Floyd, выпущенный 10 октября 1970 года. «Atom Heart Mother» достиг первого места в хит-парадах в Великобритании и 55-го места в США.", Tracklist: []string{"1. Atom Heart Mother	 	23:44",
		"2. If	 	04:31",
		"3. Summer '68	 	05:29",
		"4. Fat Old Sun	 	05:22",
		"5. Alan's Psychedelic Breakfast	 	13:00"}, Listen: "https://open.spotify.com/album/5c1ZTzT4oSkiiFS4wmEuOe?si=vvR4T3SPTHucS8ejh4YyJw", Favourite: false, Quantity: 0},
	{ID: 6, Name: "Meddle", Artist: "Pink Floyd", Cover: "assets/images/Meddle_album_cover.jpg", Price: 3690, Info: "Шестой студийный альбом британской рок-группы Pink Floyd, выпущенный 13 ноября 1971 года в Великобритании. В США альбом был издан лейблом Harvest Records немного раньше, 30 октября того же года. Meddle достиг 3-го места в хит-парадах в Великобритании и 70-го места в США. По словам музыкального критика (Daryl Easlea), этот альбом «представляет собой рождение Pink Floyd в том виде, в каком мы их знаем сегодня».", Tracklist: []string{"1. One of These Days	05:57",
		"2. A Pillow of Winds	05:10",
		"3. Fearless	06:08",
		"4. San Tropez	03:43",
		"5. Seamus	02:16",
		"6. Echoes	23:29"}, Listen: "https://open.spotify.com/album/7yKRvvqspSxfLkr7C7RaAI?si=Q1M6RIUOSIWPeE5xTeku5w", Favourite: false, Quantity: 0},
	{ID: 7, Name: "Obscured by Clouds", Artist: "Pink Floyd", Cover: "assets/images/Obscured_by_Clouds_Pink_Floyd.jpg", Price: 3199, Info: "Седьмой студийный альбом британской прогрессив-рок-группы Pink Floyd, выпущенный 3 июня 1972 года, саундтрек к французскому фильму «Долина» Барбе Шрёдера. Он же был режиссёром фильма «More», к которому Pink Floyd также записывали звуковую дорожку.", Tracklist: []string{"1. Obscured by Clouds	 	03:03",
		"2. When You're in	 	02:30",
		"3. Burning Bridges	 	03:29",
		"4. The Gold It's in the...	 	03:07",
		"5. Wot's... Uh the Deal	 	05:08",
		"6. Mudmen	 	04:20",
		"7. Childhood's End	 	04:31",
		"8. Free Four	 	04:15",
		"9. Stay	 	04:05",
		"10. Absolutely Curtains	 	05:52"}, Listen: "https://open.spotify.com/album/3HZKOTmpigZcWHxACjENyh?si=c7TSBz49Sw6c7LiZA61BIQ", Favourite: false, Quantity: 0},
	{ID: 8, Name: "The Dark Side of the Moon", Artist: "Pink Floyd", Cover: "assets/images/The_Dark_Side_of_the_Moon.png", Price: 3499, Info: "Восьмой студийный альбом британской рок-группы Pink Floyd, выпущенный 1 марта 1973 года. Один из самых продаваемых альбомов в истории звукозаписи — общее число проданных экземпляров превышает 45 миллионов. Является одним из наиболее известных концептуальных альбомов прогрессивного рока.", Tracklist: []string{"1. Speak to Me / Breathe	 	03:57",
		"2. On the Run	 	03:35",
		"3. Time	 	07:04",
		"4. The Great Gig in the Sky	 	04:47",
		"5. Money	 	06:22",
		"6. Us and Them	 	07:50",
		"7. Any Colour You Like	 	03:25",
		"8. Brain Damage	 	03:50",
		"9. Eclipse	 	02:01"}, Listen: "https://open.spotify.com/album/2WT1pbYjLJciAR26yMebkH?si=tIBlKzrvRpqzABT7-Z35mg", Favourite: false, Quantity: 0},
	{ID: 9, Name: "Wish You Were Here", Artist: "Pink Floyd", Cover: "assets/images/WishYouWereHere-300.jpg", Price: 3499, Info: "Девятый студийный альбом английской рок-группы Pink Floyd, выпущенный в сентябре 1975 года. Материал для пластинки собирался в ходе концертного тура по Европе, а запись после множественных сессий была произведена в лондонской студии Эбби Роуд.", Tracklist: []string{"1. Shine on You Crazy Diamond (Part One)	 	13:30",
		"2. Welcome to the Machine	 	07:26",
		"3. Have a Cigar	 	05:08",
		"4. Wish You Were Here	 	05:40",
		"5. Shine on You Crazy Diamond (Part Two)	 	12:22"}, Listen: "https://open.spotify.com/album/6uvBKDGlJAYLH5Vy6RQVsc?si=RG0sXoOCT26pCWaQ-ZQCWA", Favourite: false, Quantity: 0},
	{ID: 10, Name: "Animals", Artist: "Pink Floyd", Cover: "assets/images/Pink_Floyd-Animals-Frontal.jpg", Price: 3899, Info: "Десятый студийный альбом британской прогрессив-рок-группы Pink Floyd, выпущенный 23 января 1977 года. Записан в студии «Британиа роу». Достиг второго места в хит-параде Великобритании и третьего в США.", Tracklist: []string{"1. Pigs on the Wings (Part 1)	 	01:25",
		"2. Dogs	 	17:04",
		"3. Pigs (Three Different Ones)	 	11:21",
		"4. Sheep	 	10:23",
		"5. Pigs on the Wings (Part 2)	 	01:24"}, Listen: "https://open.spotify.com/album/21jUB9RqplD6OqtsTjKBnO?si=Da79dCN-SMqj2sidTOzoEQ", Favourite: false, Quantity: 0},
	{ID: 11, Name: "The Wall", Artist: "Pink Floyd", Cover: "assets/images/PinkFloydWallCoverOriginalNoText.jpg", Price: 4999, Info: "Одиннадцатый студийный альбом британской прогрессив-рок-группы Pink Floyd. Был выпущен 30 ноября 1979 года на двух дисках. Это последний релиз группы, записанный в классическом составе: Дэвид Гилмор, Роджер Уотерс, Ник Мейсон  и Ричард Райт. В поддержку альбома был организован помпезный гастрольный тур со сложными театрализованными постановками.", Tracklist: []string{"Side A:",
		" 1. In the Flesh?	 	03:20",
		"2. The Thin Ice	 	02:26",
		"3. Another Brick in the Wall, Part 1	 	03:11",
		"4. The Happiest Days of Our Lives	 	01:50",
		"5. Another Brick in the Wall, Part 2	 	03:58",
		"6. Mother	 	05:34",
		"7. Goodbye Blue Sky	 	02:47",
		"8. Empty Spaces	 	02:07",
		"9. Young Lust	 	03:29",
		"10. One of My Turns	 	03:36",
		"11. Don't Leave Me Now	 	04:15",
		"12. Another Brick in the Wall, Part 3	 	01:14",
		"13. Goodbye Cruel World	 	01:17",
		"Side B:",
		"1. Hey You	 	04:40",
		"2. Is There Anybody Out There?	 	02:41",
		"3. Nobody Home	 	03:22",
		"4. Vera	 	01:33",
		"5. Bring the Boys Back Home	 	01:27",
		"6. Comfortably Numb	 	06:22",
		"7. The Show Must Go on	 	01:36",
		"8. In the Flesh	 	04:15",
		"9. Run Like Hell	 	04:23",
		"10. Waiting for the Worms	 	03:57",
		"11. Stop	 	00:30",
		"12. The Trial	 	05:18",
		"13. Outside the Wall	 	01:46"}, Listen: "https://open.spotify.com/album/6WaIQHxEHtZL0RZ62AuY0g?si=ErE9XXiEReGCEZnl6CbWgA", Favourite: false, Quantity: 0},
	{ID: 12, Name: "The Final Cut", Artist: "Pink Floyd", Cover: "assets/images/411px-Thefinalcutcover.jpg", Price: 2999, Info: "Двенадцатый студийный альбом британской рок-группы Pink Floyd, выпущенный 21 марта 1983 года на лейбле Harvest компании EMI Records. Подзаголовок: Реквием по послевоенной мечте Роджера Уотерса, исполненный Пинк Флойд. Памяти Эрика Флетчера Уотерса 1913−1944. Альбом посвящён Эрику Флетчеру Уотерсу, отцу Роджера Уотерса, погибшему во Второй мировой войне.", Tracklist: []string{"1. The Post War Dream	 	02:59",
		"2. Your Possible Pasts	 	04:19",
		"3. One of the Few	 	01:27",
		"4. The Hero's Return	 	02:57",
		"5. The Gunner's Dream	 	05:03",
		"6. Paranoid Eyes	 	03:42",
		"7. Get Your Filthy Hands off My Desert	 	01:17",
		"8. The Fletcher Memorial Home	 	04:10",
		"9. Southampton Dock	 	02:13",
		"10. The Final Cut	 	04:43",
		"11. Not Now John	 	05:01",
		"12. Two Suns in the Sunset	 	05:16"}, Listen: "https://open.spotify.com/album/5ChHkKb5VhZe0pgQRsvpek?si=kbpzG8xsRsKC1Wc2DCQsXg", Favourite: false, Quantity: 0},
	{ID: 13, Name: "A Momentary Lapse of Reason", Artist: "Pink Floyd", Cover: "assets/images/Lapse-l.jpg", Price: 3799, Info: "Тринадцатый студийный альбом британской рок-группы Pink Floyd, выпущенный звукозаписывающими компаниями EMI Records (7 сентября 1987 года) и Columbia Records (8 сентября 1987 года). Это первый диск, записанный группой после ухода из Pink Floyd Роджера Уотерса.", Tracklist: []string{"1. Signs of Life	 	04:24",
		"2. Learning to Fly	 	04:53",
		"3. The Dogs of War	 	06:05",
		"4. One Slip	 	05:10",
		"5. On the Turning Away	 	05:42",
		"6. Yet Another Movie - a) Round and Around	 	07:28",
		"7. A New Machine (Part 1)	 	01:46",
		"8. Terminal Frost	 	06:17",
		"9. A New Machine (Part 2)	 	00:38",
		"10. Sorrow	 	08:46"}, Listen: "https://open.spotify.com/album/1tWgv9v78StWukBRBVNyxA?si=aJagnUMrQl68z5dFE_xCCQ", Favourite: false, Quantity: 0},
	{ID: 14, Name: "The Division Bell", Artist: "Pink Floyd", Cover: "assets/images/Pink_floyd_—_The_Division_Bell_front.jpg", Price: 4499, Info: "Четырнадцатый студийный альбом британской прогрессив-рок-группы Pink Floyd. На территории Великобритании диск вышел 28 марта 1994 года на лейбле EMI, в США релиз состоялся 4 апреля того же года на Columbia Records.", Tracklist: []string{"1. Cluster One	 	05:58",
		"2. What Do You Want from Me	 	04:22",
		"3. Poles Apart	 	07:05",
		"4. Marooned	 	05:28",
		"5. A Great Day for Freedom	 	04:18",
		"6. Wearing the Inside Out	 	06:49",
		"7. Take It Back	 	06:13",
		"8. Coming Back to Life	 	06:19",
		"9. Keep Talking	 	06:11",
		"10. Lost for Words	 	05:15",
		"11. High Hopes	 	08:32"}, Listen: "https://open.spotify.com/album/7wzStEU2keGohEu8jpVMZW?si=JheLf2ThQTOsMQazfGq0rw", Favourite: false, Quantity: 0},
	{ID: 15, Name: "The Endless River", Artist: "Pink Floyd", Cover: "assets/images/Pink_Floyd_-_The_Endless_River_(Artwork).jpg", Price: 5690, Info: "Пятнадцатый и последний студийный альбом британской рок-группы Pink Floyd, вышедший 10 ноября 2014 года. Продюсером альбома выступил Дэвид Гилмор. Выпуском пластинки в Великобритании занимается Parlophone Records, а в США — Warner Bros. Records. Альбом основан на студийных записях, сделанных при подготовке альбома The Division Bell.", Tracklist: []string{"1. Things Left Unsaid	04:26",
		"2. It's What We Do	06:15",
		"3. Ebb and Flow	01:52",
		"4. Sum	04:50",
		"5. Skins	02:32",
		"6. Unsung	01:06",
		"7. Anisina	03:10",
		"8. The Lost Art of Conversation	01:43",
		"9. On Noodle Street	01:42",
		"10. Night Light	01:42",
		"11. Allons-y (Part 1)	01:56",
		"12. Autumn '68	01:37",
		"13. Allons-y (Part 2)	01:35",
		"14. Talkin Hawkin'	03:26",
		"15. Calling	03:38",
		"16. Eyes to Pearls	01:52",
		"17. Surfacing	02:47",
		"18. Louder Than Words	06:25"}, Listen: "https://open.spotify.com/album/0fXAlQ9wTG2glNJvZEkBZc?si=K4GlkZO3SV21WXCtaymZjg", Favourite: false, Quantity: 0},
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)
	var lastID int = len(products)

	for _, productItem := range products {
		if productItem.ID > lastID {
			lastID = productItem.ID
		}
	}
	newProduct.ID = lastID + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

//Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {

			products[i].Name = updatedProduct.Name
			products[i].Artist = updatedProduct.Artist
			products[i].Cover = updatedProduct.Cover
			products[i].Price = updatedProduct.Price
			products[i].Info = updatedProduct.Info
			products[i].Tracklist = updatedProduct.Tracklist
			products[i].Listen = updatedProduct.Listen
			products[i].Quantity = updatedProduct.Quantity

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	fmt.Println("Server is running on http://localhost:8080 !")
	http.ListenAndServe(":8080", nil)
}
