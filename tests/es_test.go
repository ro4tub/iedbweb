package test
// ElasticSearch的单元测试
import (
	"testing"
	"log"
	"os"
	"github.com/olivere/elastic"
)

type GameDocument struct {
	Name string		`json:"name"`
	Genre string	`json:"genre"`
	Platform string	`json:"platform"`
	Tags []string	`json:"tags"`
	SimpleDesc string	`json:"simple_desc"`
	Detail string	`json:"detail"`
}


func TestSearch(t *testing.T) {
	// 建立链接
	tracelogger := log.New(os.Stdout, "ES Trace: ", log.Lshortfile)
	client, err := elastic.NewClient(elastic.SetURL("http://backend1:9200"), elastic.SetTraceLog(tracelogger))
	if err != nil {
		t.Fatal("elastic.NewClient failed: %v", err)
		return
	}
	
	// 创建索引
	exists, err := client.IndexExists("iedbtest").Do()
	if err != nil {
	    t.Fatal("client.IndexExists failed: %v", err)
		return
	}
	if !exists {
		_, err = client.CreateIndex("iedbtest").Do()
		if err != nil {
			t.Fatal("create index iedb failed: %v", err)
			return
		}
	}

	
	// 增加文档
	// game := GameDocument{Name:"天天魔斗士", Genre:"角色扮演", Platform:"Android", Tags:[]string{"Unity3D", "腾讯"}, SimpleDesc:"腾讯出品的好游戏", Detail:"哇哈哈哈哈"}
	game := GameDocument{Name:"天天魔斗士", Genre:"角色扮演", Platform:"Android", Tags:[]string{"Unity3D", "腾讯"}, SimpleDesc:"这款游戏由腾讯天美工作室出品", Detail:"哇哈哈哈哈"}
	_, err = client.Index().Index("iedbtest").Type("game").Id("1").BodyJson(game).Do()
	if err != nil {
		t.Fatal("add document failed: %v", err)
		return
	}
	
	// 查询
	// 通过名字查询 term query
	// termQuery := elastic.NewTermQuery("Name", "天天魔斗士")
	query1 := elastic.NewMatchQuery("name", "天天")
	//.From(0).Size(10)
	searchResult1, err := client.Search().Index("iedbtest").Query(&query1).From(0).Size(10).Do()
	if err != nil {
		t.Fatal("client search failed: %v", err)
		return
	}
	
	if searchResult1.Hits == nil || searchResult1.Hits.TotalHits != 1 {
		t.Error("not found result", err)
	}
	
	// 好游戏 可以
	// 游戏 可以
	// Unity3D 可以
	// unity3d 可以
	// 腾讯  可以
	// 角色扮演 可以
	// android 可以
	// 天天 可以
	// 魔斗士 可以
	query2 := elastic.NewMultiMatchQuery("游戏", "name", "genre", "platform", "tags", "simple_desc", "detail")
	searchResult2, err := client.Search().Index("iedbtest").Query(&query2).From(0).Size(10).Do()
	if err != nil {
		t.Fatal("client search failed: %v", err)
		return
	}
	if searchResult2.Hits == nil || searchResult2.Hits.TotalHits != 1 {
		t.Error("not found result", err)
	}
}

