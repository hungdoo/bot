package telecommands

import (
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/bson"
)

func StoreCmd(cmd command.ICommand) error {
	if err := db.GetDb().Insert("commands", cmd); err != nil {
		log.GeneralLogger.Printf("Job [%s] insert db failed: [%s]", cmd.GetName(), err)
		return err
	}
	return nil
}

func UpdateCmd(cmd command.ICommand) error {
	filter := bson.M{"_id": cmd.GetName()}
	update := bson.M{"$set": cmd}
	if err := db.GetDb().Update("commands", filter, update); err != nil {
		log.GeneralLogger.Printf("Job [%s] update db failed: [%s]", cmd.GetName(), err)
		return err
	}
	return nil
}

func UpdateMultiCmd(cmds []command.ICommand) error {
	filter := bson.M{}
	update := bson.M{"$set": cmds}
	if err := db.GetDb().Update("commands", filter, update); err != nil {
		log.GeneralLogger.Printf("Job update many db failed: [%s]", err)
		return err
	}
	return nil
}
