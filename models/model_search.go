package models
import (
	"log"
	"os"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	. "github.com/ro4tub/gamedb/util"
)


// game在es中的文档
type GameDocument struct {
	Id int64		`json:"id"`
	Name string		`json:"name"`
	Genre string	`json:"genre"`
	Platform string	`json:"platform"`
	Tags []string	`json:"tags"`
	SimpleDesc string	`json:"simple_desc"`
	Detail string	`json:"detail"`
}

var client *elastic.Client

func init() {
	var tracelogger *log.Logger
	if estracelog, _ := beego.AppConfig.Bool("estracelog"); estracelog == true {
		tracelogger = log.New(os.Stdout, "ES Trace: ", log.Lshortfile)
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(beego.AppConfig.String("esurl")), elastic.SetTraceLog(tracelogger))
	if err != nil {
		panic(fmt.Sprintf("elastic.NewClient failed: %v", err))
		return
	}
}


func CreateSearchIndex(game GameDocument) error {
	Log.Debug("CreateSearchIndex")
	// 创建索引
	exists, err := client.IndexExists("iedb").Do()
	if err != nil {
	    Log.Error("client.IndexExists failed: %v", err)
		return err
	}
	if !exists {
		_, err = client.CreateIndex("iedb").Do()
		if err != nil {
			Log.Error("create index iedb failed: %v", err)
			return err
		}
	}
	
	// 增加文档
	_, err = client.Index().Index("iedb").Type("game").Id(string(game.Id)).BodyJson(game).Do()
	if err != nil {
		Log.Error("add document failed: %v", err)
		return err
	}
	return nil
}

func Search(text string, from int ,size int) ([]GameDocument, error) {
	Log.Debug("Search")
	query := elastic.NewMultiMatchQuery(text, "name", "genre", "platform", "tags", "simple_desc", "detail")
	searchResult, err := client.Search().Index("iedb").Query(&query).From(from).Size(size).Do()
	if err != nil {
		Log.Error("client search failed: %v", err)
		return nil, err
	}
	var result []GameDocument
	if searchResult.Hits != nil {
		Log.Debug("Found a total of %d games", searchResult.Hits.TotalHits)
		result = make([]GameDocument, searchResult.Hits.TotalHits)
	    // Iterate through results
	    for i, hit := range searchResult.Hits.Hits {
	        // hit.Index contains the name of the index

	        // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	        var t GameDocument
	        err := json.Unmarshal(*hit.Source, &t)
	        if err != nil {
	            Log.Error("json.Unmarshal failed: %v", err)
				continue
	        }
			result[i] = t
	    }
	} else {
	    Log.Warn("no search result: %s, %d, %d", text, from , size)
	}
	return result, nil
}