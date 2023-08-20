package persistance

import (
	"context"
	"fmt"
	"os"

	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/Jordation/jsonl/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholder with your Atlas connection string

type Persistance struct {
	Client *mongo.Client
}

func (p *Persistance) StoreGameConfig(cfg *riotTypes.GameConfig, ID string) error {
	collection := p.Client.Database("val-feed").Collection("game-configs")
	if _, err := collection.InsertOne(
		context.Background(),
		bson.M{"_id": ID, "config": cfg}); err != nil {
		return err
	}

	return nil
}

func (p *Persistance) GetGameConfig(ID string) (*riotTypes.GameConfig, error) {
	collection := p.Client.Database("val-feed").Collection("game-configs")
	var result bson.M

	response := &riotTypes.GameConfig{}
	if err := collection.FindOne(context.Background(),
		bson.M{"_id": ID}).Decode(&result); err != nil {
		return nil, err
	}
	bytes, err := bson.Marshal(result["config"])
	if err != nil {
		return nil, err
	}

	if err := bson.Unmarshal(bytes, response); err != nil {
		return nil, err
	}

	return response, nil
}

func New() (*Persistance, error) {
	var uri = os.Getenv("MONGO_CONN_STR")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return &Persistance{
		Client: client,
	}, nil
}

func init() {
	godotenv.Load(utils.GetRelativePath("../../local.env"))
}

func Ping(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Ping Ack'd.")
}
