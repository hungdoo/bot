package telecommands

import (
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/bson"
)

func StoreDb(cmd command.ICommand) {
	filter := bson.M{"_id": cmd.GetName()}
	update := bson.M{"$set": cmd}
	if err := db.GetDb().Update("commands", filter, update); err != nil {
		log.GeneralLogger.Printf("Job [%s] update db failed: [%s]", cmd.GetName(), err)
	}
}
