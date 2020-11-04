package scraper

import (
	"context"
	"log"
	"strconv"
	"strings"

	"simplescraper/data"

	"github.com/gocolly/colly"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func saveData(newData []data.Entry) error {

	ctx := context.TODO()

	var cred options.Credential
	// cred.AuthSource = YourAuthSource
	cred.Username = "admin"
	cred.Password = "asd123"
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred)

	client, e := mongo.Connect(ctx, clientOpt)
	if nil != e {
		return e
	}
	defer client.Disconnect(ctx)

	collection := client.Database("scraper-db").Collection("dolarhoy")

	// simple := data.Entry{Title: "Prueba", Buy: 999.9, Sell: 999.9}
	for _, elem := range newData {
		insertResult, err := collection.InsertOne(ctx, elem)
		if nil != err {
			return err
		}
		log.Println("INSERT", insertResult)
	}
	return nil
}

//Scrap do the scrap stuff
func Scrap() {

	URL := "https://dolarhoy.com/"

	if URL == "" {
		log.Println("missing URL argument")
		return
	}

	log.Println("visiting", URL)

	c := colly.NewCollector()
	var response []data.Entry

	c.OnHTML(".pill", func(e *colly.HTMLElement) {
		newEntry := data.Entry{}
		newEntry.Title = e.ChildText("a")
		r := strings.NewReplacer(",", ".", "\n", "", " ", "", "\t", "", "$", " ")
		prices := strings.Split(strings.Trim(r.Replace(e.ChildText(".price")), " "), " ")

		if 0 < len(prices) {
			if s, err := strconv.ParseFloat(prices[0], 8); err == nil {
				newEntry.Buy = s
			}
			if 1 < len(prices) {
				if s, err := strconv.ParseFloat(prices[1], 8); err == nil {
					newEntry.Sell = s
				}
			} else {
				newEntry.Sell = newEntry.Buy
			}
		}
		response = append(response, newEntry)
	})
	c.Visit(URL)

	// b, err := json.Marshal(response)

	// if err != nil {
	// 	log.Println("failed to serialize response:", err)
	// 	return
	// }
	dbErr := saveData(response)
	if nil != dbErr {
		log.Println(dbErr)
	}
}
